package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i := range len(ss) {
		res[i], _ = strconv.Atoi(ss[i])
	}
	return res
}

func drive(reader *bufio.Reader) []int {
	nums := readNums(reader)
	m := nums[1]
	a := readNums(reader)

	queries := make([][]int, m)
	for i := range m {
		s := readString(reader)
		ss := strings.Split(s, " ")
		if s[0] == 'C' {
			queries[i] = make([]int, 4)
			queries[i][0] = 1
		} else {
			queries[i] = make([]int, 3)
			queries[i][0] = 2
		}
		for j := 1; j < len(queries[i]); j++ {
			queries[i][j], _ = strconv.Atoi(ss[j])
		}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	tr1 := make(Tree1, 2*n)
	tr2 := make(Tree2, 2*n)
	for i := range n {
		tr1.Update(i, i+1, a[i])
		if i+1 < n {
			diff := a[i+1] - a[i]
			tr2.Update(i, diff)
		}
	}

	update := func(l int, r int, d int) {
		tr1.Update(l, r+1, d)
		if l > 0 {
			v1 := tr1.Get(l - 1)
			v2 := tr1.Get(l)
			diff := v2 - v1
			tr2.Update(l-1, diff)
		}
		if r+1 < n {
			v1 := tr1.Get(r)
			v2 := tr1.Get(r + 1)
			diff := v2 - v1
			tr2.Update(r, diff)
		}
	}

	get := func(l int, r int) int {
		tmp := tr2.Get(l+1, r)
		v1 := tr1.Get(l)
		return gcd(tmp, v1)
	}

	var res []int
	for _, cur := range queries {
		if cur[0] == 1 {
			update(cur[1]-1, cur[2]-1, cur[3])
		} else {
			tmp := get(cur[1]-1, cur[2]-1)
			res = append(res, tmp)
		}
	}
	return res
}

type Tree1 []int

func (tr Tree1) Update(l int, r int, v int) {
	n := len(tr) / 2
	l += n
	r += n
	for l < r {
		if l&1 == 1 {
			tr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			tr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (tr Tree1) Get(p int) int {
	n := len(tr) / 2
	p += n
	var res int
	for p > 0 {
		res += tr[p]
		p >>= 1
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}

func gcd(a, b int) int {
	a = abs(a)
	b = abs(b)
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Tree2 []int

func (tr Tree2) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v
	for p > 1 {
		tr[p>>1] = gcd(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr Tree2) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = gcd(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = gcd(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

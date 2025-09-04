package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) int {
	n := len(a)
	m := len(a[0])
	L := make([][]int, n)
	R := make([][]int, n)
	for i := range n {
		L[i] = make([]int, m)
		R[i] = make([]int, m)
	}

	for i := range n {
		for j := range m {
			if a[i][j] == 'z' {
				L[i][j] = 1
				if j > 0 {
					L[i][j] += L[i][j-1]
				}
			}
			// else 0
		}
		for j := m - 1; j >= 0; j-- {
			if a[i][j] == 'z' {
				R[i][j] = 1
				if j+1 < m {
					R[i][j] += R[i][j+1]
				}
			}
		}
	}

	bit := make([]int, m+10)
	time := make([]int, m+10)
	var now int

	add := func(i int, v int) {
		i++
		for i < len(bit) {
			if time[i] < now {
				bit[i] = 0
				time[i] = now
			}
			bit[i] += v
			i += i & -i
		}
	}

	pre := func(i int) int {
		i++
		var res int
		for i > 0 {
			if time[i] == now {
				res += bit[i]
			}
			i -= i & -i
		}
		return res
	}

	todo := make([][]int, m)
	var keys []int
	clear := func() {
		now++
		for _, k := range keys {
			todo[k] = todo[k][:0]
		}
		keys = keys[:0]
	}

	var res int

	for d := 0; d < n+m-1; d++ {
		clear()
		// i + j = d
		// i <= n - 1 i >= 0
		// j <= m - 1, j >= 0
		start := max(0, d-(m-1))
		end := min(n-1, d)
		for i := start; i <= end; i++ {
			j := d - i
			if a[i][j] == '.' {
				clear()
				continue
			}
			add(j, 1)
			res += pre(j + R[i][j] - 1)
			l := j - L[i][j] + 1
			todo[l] = append(todo[l], j)
			keys = append(keys, l)
			for _, k := range todo[j] {
				add(k, -1)
			}
		}

	}

	return res
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func drive(reader *bufio.Reader) []int {
	s := readString(reader)
	n := readNum(reader)
	queries := make([]string, n)
	for i := range n {
		queries[i] = readString(reader)
	}
	return solve(s, queries)
}

func getId(x byte) int {
	if x == 'A' {
		return 0
	}
	if x == 'T' {
		return 1
	}
	if x == 'G' {
		return 2
	}
	return 3
}

func solve(s string, queries []string) []int {
	n := len(s)
	trs := make([][][]BIT, 4)
	for x := range 4 {
		trs[x] = make([][]BIT, 11)
		for k := 1; k <= 10; k++ {
			trs[x][k] = make([]BIT, k)
			for r := range k {
				trs[x][k][r] = make(BIT, n+1)
			}
		}
	}

	buf := []byte(s)

	for i := range n {
		x := getId(buf[i])
		for k := 1; k <= 10; k++ {
			trs[x][k][i%k].update(i, 1)
		}
	}

	change := func(pos int, x byte) {
		if buf[pos] == x {
			return
		}
		old := getId(buf[pos])
		for k := 1; k <= 10; k++ {
			trs[old][k][pos%k].update(pos, -1)
		}
		cur := getId(x)
		for k := 1; k <= 10; k++ {
			trs[cur][k][pos%k].update(pos, 1)
		}
		buf[pos] = x
	}

	find := func(l int, r int, e string) int {
		k := len(e)
		var res int
		// l % k
		w := l % k
		for i := range k {
			x := getId(e[i])
			cnt := trs[x][k][(w+i)%k].query(l, r)
			res += cnt
		}
		return res
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == '1' {
			var i int
			pos := readInt([]byte(cur), 2, &i)
			change(i-1, cur[pos+1])
		} else {
			var l, r int
			pos := readInt([]byte(cur), 2, &l) + 1
			pos = readInt([]byte(cur), pos, &r) + 1
			e := cur[pos:]
			ans = append(ans, find(l-1, r-1, e))
		}
	}

	return ans
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) pre(i int) int {
	var res int
	i++
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return bit.pre(r) - bit.pre(l-1)
}

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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	s := readString(reader)
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(s, queries)
}

func solve(s string, queries [][]int) []int {
	n := len(s)

	at := make([][]int, n)
	for i, cur := range queries {
		r := cur[1] - 1
		at[r] = append(at[r], i)
	}

	bit := make(BIT, n+2)

	ans := make([]int, len(queries))

	stack := make([]int, n)
	var top int
	for i := range n {
		if s[i] == '(' {
			stack[top] = i
			top++
		} else {
			if top > 0 {
				j := stack[top-1]
				bit.Set(j, 1)
				top--
			}
		}
		for _, j := range at[i] {
			l := queries[j][0] - 1
			ans[j] = bit.Query(l, i) * 2
		}
	}
	return ans
}

type BIT []int

func (bit BIT) Set(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}
func (bit BIT) Get(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) Query(l int, r int) int {
	res := bit.Get(r)
	if l > 0 {
		res -= bit.Get(l - 1)
	}
	return res
}

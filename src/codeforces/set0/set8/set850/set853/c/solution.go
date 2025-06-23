package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

func process(reader *bufio.Reader) []int {
	n, q := readTwoNums(reader)
	p := readNNums(reader, n)
	qs := make([][]int, q)
	for i := range q {
		qs[i] = readNNums(reader, 4)
	}
	return solve(n, p, qs)
}

func ways(n int) int {
	return n * (n - 1) / 2
}

func solve(n int, p []int, qs [][]int) []int {
	// n * (n - 1) / 2 - 上下左右 + 四个角落
	ans := make([]int, len(qs))
	at_l := make([][]int, n+1)
	at_r := make([][]int, n+1)
	for i := range qs {
		l, d, r, u := qs[i][0], qs[i][1], qs[i][2], qs[i][3]
		ans[i] = ways(n) - ways(l-1) - ways(d-1) - ways(n-r) - ways(n-u)
		at_l[l] = append(at_l[l], i)
		at_r[r] = append(at_r[r], i)
	}

	cnt := make(BIT, n+3)

	for l := 1; l <= n; l++ {
		for _, i := range at_l[l] {
			d, u := qs[i][1], qs[i][3]
			w := cnt.Get(d - 1)
			ans[i] += ways(w)
			w = cnt.GetRange(u+1, n)
			ans[i] += ways(w)
		}
		cnt.Add(p[l-1], 1)
	}

	clear(cnt)

	for r := n; r > 0; r-- {
		for _, i := range at_r[r] {
			d, u := qs[i][1], qs[i][3]
			w := cnt.Get(d - 1)
			ans[i] += ways(w)
			w = cnt.GetRange(u+1, n)
			ans[i] += ways(w)
		}
		cnt.Add(p[r-1], 1)
	}

	return ans
}

type BIT []int

func (bit BIT) Add(i int, v int) {
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) Get(i int) int {
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) GetRange(l int, r int) int {
	return bit.Get(r) - bit.Get(l-1)
}

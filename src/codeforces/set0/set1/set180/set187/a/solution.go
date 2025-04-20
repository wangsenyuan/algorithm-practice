package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	s := readNNums(reader, n)
	t := readNNums(reader, n)
	return solve(s, t)
}

func solve(s []int, t []int) int {
	n := len(s)

	pos := make([]int, n)
	for i := range n {
		s[i]--
		t[i]--
		pos[t[i]] = i
	}

	// 要对s重新排序
	for i := 0; i < n; i++ {
		s[i] = pos[s[i]]
	}

	var res int
	marked := make([]bool, n)
	// 然后将s变成 0,1,.... n
	// 似乎不对的, 0, 1, 4, 2, 3, 5
	// 答案是3

	for i, j := n-1, n-1; i >= 0; i-- {
		for marked[j] {
			j--
		}
		marked[s[i]] = true
		if s[i] == j {
			j--
			continue
		}
		res = n - i
	}

	return res
}

type BIT []int

func (bit BIT) update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) query(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	segs := make([][]int, n)
	for i := range n {
		segs[i] = readNNums(reader, 4)
	}
	return solve(segs, k)
}

func solve(segs [][]int, k int) int {
	// n := len(segs)
	// ti, Ti, xi, costi

	var res int

	for _, cur := range segs {
		t, T, x, c := cur[0], cur[1], cur[2], cur[3]
		if t+k <= T {
			// 只需要一辆bus
			res += c
			continue
		}
		// need to add bus or x
		// 假设每辆车有m个人, m + t <= T => m = T - t
		m := T - t
		if m <= 0 {
			// 肯定超了，只能赔钱
			res += k*x + c
			continue
		}
		// m <= k
		cnt := (k + m - 1) / m
		res += min(c+k*x, cnt*c)
	}

	return res
}

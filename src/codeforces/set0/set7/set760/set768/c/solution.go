package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1])
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
	n, k, x := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(x, k, a)
}

const X = 1000

func solve(x int, k int, a []int) []int {
	dp := make([]int, 2048)
	ndp := make([]int, 2048)
	for _, v := range a {
		dp[v]++
	}

	for range k {
		// 这个地方不是这样弄的
		var cnt int
		for i, c := range dp {
			if cnt%2 == 0 {
				// 一半迁移到 i ^ x
				ndp[i^x] += (c + 1) / 2
				ndp[i] += c / 2
			} else {
				ndp[i] += (c + 1) / 2
				ndp[i^x] += c / 2
			}
			cnt += c
		}
		copy(dp, ndp)
		clear(ndp)
	}

	res := []int{-1, 0}

	for i, c := range dp {
		if c > 0 {
			if res[0] == -1 {
				res[0] = i
			}
			res[1] = i
		}
	}
	res[0], res[1] = res[1], res[0]
	return res
}

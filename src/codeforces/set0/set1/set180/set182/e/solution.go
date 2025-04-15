package main

import (
	"bufio"
	"fmt"
	"os"
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

func process(reader *bufio.Reader) int {
	n, l := readTwoNums(reader)
	boards := make([][]int, n)
	for i := range n {
		boards[i] = readNNums(reader, 2)
	}
	return solve(l, boards)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(l int, boards [][]int) int {
	// dp[x][i][0/1] 表示长度为x，且最后一个是是类型i，使用的是它的宽0, 或者是长(1)的计数
	n := len(boards)
	dp := make([][][]int, l+1)
	for i := range l + 1 {
		dp[i] = make([][]int, n)
		for j := range n {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := range n {
		w, h := boards[i][0], boards[i][1]
		if w <= l {
			dp[w][i][0] = 1
		}
		if w != h && h <= l {
			dp[h][i][1] = 1
		}
	}

	for x := 1; x <= l; x++ {
		for i := range n {
			for j := range 2 {
				if dp[x][i][j] == 0 {
					continue
				}
				// j = 0 的时候， cur = 宽度
				cur := boards[i][j]

				for k := range n {
					if i == k {
						continue
					}
					w, h := boards[k][0], boards[k][1]
					if w+x <= l && h == cur {
						dp[w+x][k][0] = add(dp[w+x][k][0], dp[x][i][j])
					}
					if w != h && h+x <= l && w == cur {
						dp[h+x][k][1] = add(dp[h+x][k][1], dp[x][i][j])
					}
				}
			}
		}
	}

	var res int
	for i := range n {
		res = add(res, dp[l][i][0])
		res = add(res, dp[l][i][1])
	}
	return res
}

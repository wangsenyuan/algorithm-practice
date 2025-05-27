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
	n := readNum(reader)
	arr := make([][]int, 3)
	for i := range 3 {
		arr[i] = readNNums(reader, n)
	}

	return solve(n, arr)
}

const inf = 1 << 60

func solve(n int, arr [][]int) int {
	a := arr[0]
	b := arr[1]
	c := arr[2]
	dp := make([]int, 2)
	dp[0] = a[n-1]
	dp[1] = b[n-1]

	for i := n - 2; i >= 0; i-- {
		// 要饿先i，那么对于后面来说就是dp[1], 要么后i
		x := max(a[i]+dp[1], b[i]+dp[0])
		// 要么先i，那么对于i+1，来说，就是要计算dp[1]
		// 要么后i，那么对于i+1来说，就是dp[0]
		y := max(b[i]+dp[1], c[i]+dp[0])
		dp[0] = x
		dp[1] = y
	}

	return dp[0]
}

func solve1(n int, arr [][]int) int {
	dp := make([][]int, n)
	ndp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, 4)
		ndp[i] = make([]int, 4)
	}
	cnt := []int{0, 1, 1, 2}
	fp := make([]int, 4)
	for r := 0; r < n; r++ {
		clear(fp)
		for l := r; l >= 0; l-- {
			for s := range 4 {
				if l == r {
					fp[s] = arr[cnt[s]][l]
					ndp[l][s] = arr[cnt[s]][l]
				} else {
					// l < r
					// 先处理l+1...r
					sr := s & 1
					var tmp int
					if s&2 == 2 {
						tmp++
					}
					// 先处理l
					ndp[l][s] = max(ndp[l][s], fp[2|sr]+arr[tmp][l])
					tmp++
					ndp[l][s] = max(ndp[l][s], fp[sr]+arr[tmp][l])
					sl := s & 2
					tmp = 0
					if s&1 == 1 {
						tmp++
					}
					ndp[l][s] = max(ndp[l][s], dp[l][sl|1]+arr[tmp][r])
					tmp++
					ndp[l][s] = max(ndp[l][s], dp[l][sl]+arr[tmp][r])
				}
			}
			for s := range 4 {
				fp[s] = ndp[l][s]
			}
		}
		for l := range r + 1 {
			for s := range 4 {
				dp[l][s] = ndp[l][s]
				ndp[l][s] = 0
			}
		}
	}

	return dp[0][0]
}

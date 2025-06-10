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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func process(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)
	}
	return solve(grid)
}

func solve(grid []string) int {
	n := len(grid)
	m := len(grid[0])

	var res int

	dp := make([][]int, n)
	ndp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
		ndp[i] = make([]int, m)
		for j := range m {
			dp[i][j] = -1
			ndp[i][j] = -1
		}
	}

	for i := 0; i+2 <= n; i++ {
		for j := 0; j+2 <= m; j++ {
			dp[i][j] = 0
			for di := 0; di < 2; di++ {
				for dj := 0; dj < 2; dj++ {
					dp[i][j] *= 2
					if grid[i+di][j+dj] == '*' {
						dp[i][j]++
					}
				}
			}
		}
	}

	for sz := 2; sz*2 <= min(n, m); sz *= 2 {
		for i := 0; i+2*sz <= n; i++ {
			for j := 0; j+2*sz <= m; j++ {
				var mask int
				for di := range 2 {
					for dj := range 2 {
						mask *= 2
						if dp[i+di*sz][j+dj*sz] == 15 {
							mask++
						}
					}
				}
				ok := true
				for di := range 2 {
					for dj := range 2 {
						if dp[i+di*sz][j+dj*sz] != 15 && dp[i+di*sz][j+dj*sz] != mask {
							ok = false
						}
					}
				}
				if ok {
					ndp[i][j] = mask
				}
			}
		}
		for i := range n {
			for j := range m {
				dp[i][j] = ndp[i][j]
				ndp[i][j] = -1
				if dp[i][j] != -1 {
					res++
				}
			}
		}
	}

	return res
}

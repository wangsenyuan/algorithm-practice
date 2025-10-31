package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
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

func drive(reader *bufio.Reader) int {
	_, _, k := readThreeNums(reader)
	s := readString(reader)
	t := readString(reader)
	return solve(k, s, t)
}

func solve(k int, s string, t string) int {
	n := len(s)
	m := len(t)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
	}
	for i := range n {
		for j := range m {
			if s[i] == t[j] {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] += dp[i-1][j-1]
				}
			}
		}
	}

	fp := make([][][]int, n)
	for i := range n {
		fp[i] = make([][]int, m)
		for j := range m {
			fp[i][j] = make([]int, k+1)
		}
	}
	for i := range n {
		for j := range m {
			for x := range k + 1 {
				if i > 0 {
					fp[i][j][x] = max(fp[i][j][x], fp[i-1][j][x])
				}
				if j > 0 {
					fp[i][j][x] = max(fp[i][j][x], fp[i][j-1][x])
				}
				if x > 0 && dp[i][j] > 0 {
					i1 := i - dp[i][j]
					j1 := j - dp[i][j]
					if i1 < 0 || j1 < 0 {
						fp[i][j][x] = max(fp[i][j][x], dp[i][j])
					} else {
						fp[i][j][x] = max(fp[i][j][x], fp[i1][j1][x-1]+dp[i][j])
					}
				}
			}
			for x := range k {
				fp[i][j][x+1] = max(fp[i][j][x+1], fp[i][j][x])
			}
		}
	}
	return fp[n-1][m-1][k]
}

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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
	readNum(reader)
	s := readString(reader)
	return solve(s)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a int, b int) int {
	return a * b % mod
}

func solve(s string) int {
	n := len(s)
	X := 1 << 10

	dp := make([][]int, X)
	for i := 0; i < X; i++ {
		dp[i] = make([]int, 10)
	}
	fp := make([]int, X)

	for i := 0; i < n; i++ {
		x := int(s[i] - 'A')
		for state := 0; state < X; state++ {
			if (state>>x)&1 == 0 {
				// 前面还没有出现x
				next := state | (1 << x)
				if state == 0 {
					fp[next] = add(fp[next], 1)
				} else {
					for y := 0; y < 10; y++ {
						if (state>>y)&1 == 1 {
							fp[next] = add(fp[next], dp[state][y])
						}
					}
				}
			} else {
				// 前面已经出现了state, 必须选择以x结尾的状态
				fp[state] = add(fp[state], dp[state][x])
			}
		}

		for state := 0; state < X; state++ {
			dp[state][x] = add(dp[state][x], fp[state])
			fp[state] = 0
		}
	}

	var ans int

	for i := 0; i < 10; i++ {
		for state := 0; state < X; state++ {
			if (state>>i)&1 == 1 {
				ans = add(ans, dp[state][i])
			}
		}
	}

	return ans
}

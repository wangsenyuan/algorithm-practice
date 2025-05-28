package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0])
	fmt.Println(res[1])
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
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(k, a)
}

type state struct {
	a int
	b int
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(k int, weights []int) []int {
	n := len(weights)

	C := make([][]int, n+1)
	for i := range n + 1 {
		C[i] = make([]int, i+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j-1], C[i-1][j])
		}
	}

	cnt := make([]int, 2)
	for _, w := range weights {
		if w == 50 {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}
	dp := make([][]int, cnt[0]+1)
	ways := make([][]int, cnt[0]+1)
	for i := range cnt[0] + 1 {
		dp[i] = make([]int, cnt[1]+1)
		ways[i] = make([]int, cnt[1]+1)
		for j := range cnt[1] + 1 {
			dp[i][j] = -1
		}
	}

	dp[cnt[0]][cnt[1]] = 0
	ways[cnt[0]][cnt[1]] = 1

	var que []state
	que = append(que, state{cnt[0], cnt[1]})

	update := func(v state, d int, w int) {
		if dp[v.a][v.b] == -1 {
			dp[v.a][v.b] = d
			ways[v.a][v.b] = w
			que = append(que, v)
		} else if dp[v.a][v.b] == d {
			ways[v.a][v.b] = add(ways[v.a][v.b], w)
		}
		// else ignore
	}

	var tail int
	for tail < len(que) {
		u := que[tail]
		tail++

		// 然后产生新的状态
		for i := 0; i <= u.a; i++ {
			if i*50 > k {
				break
			}
			for j := 0; j <= u.b; j++ {
				if i*50+j*100 > k {
					break
				}
				if i+j == 0 {
					// 不送人
					continue
				}
				// 对岸有这么多人
				x, y := cnt[0]-u.a, cnt[1]-u.b
				// 送到对岸后这么多人
				x += i
				y += j
				w := mul(C[u.a][i], C[u.b][j])
				w = mul(w, ways[u.a][u.b])
				if x+y == n {
					// 不用返回
					next := state{0, 0}
					update(next, dp[u.a][u.b]+1, w)
				} else {
					for c := 0; c <= x; c++ {
						for d := 0; d <= y; d++ {
							if c+d == 0 {
								continue
							}
							if c*50+d*100 > k {
								break
							}
							next := state{u.a - i + c, u.b - j + d}
							update(next, dp[u.a][u.b]+1, mul(w, mul(C[x][c], C[y][d])))
						}
					}
				}
			}
		}
	}

	if dp[0][0] != -1 {
		return []int{dp[0][0]*2 - 1, ways[0][0]}
	}

	return []int{-1, 0}
}

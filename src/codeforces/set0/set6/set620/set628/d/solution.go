package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
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

func drive(reader *bufio.Reader) int {
	m, d := readTwoNums(reader)
	a := readString(reader)
	b := readString(reader)
	return solve(m, d, a, b)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(m int, d int, a string, b string) int {
	n := len(a)
	dp := make([][][]int, m)
	ndp := make([][][]int, m)
	for i := range m {
		dp[i] = make([][]int, 2)
		ndp[i] = make([][]int, 2)
		for j := range 2 {
			dp[i][j] = make([]int, 2)
			ndp[i][j] = make([]int, 2)
		}
	}

	dp[0][1][1] = 1

	update := func(w1 int, w2 int, x int, i int, pos int) {
		nw1 := w1
		if i > int(a[pos]-'0') {
			nw1 = 0
		}
		nw2 := w2
		if i < int(b[pos]-'0') {
			nw2 = 0
		}
		nx := (x*10 + i) % m
		ndp[nx][nw1][nw2] = add(ndp[nx][nw1][nw2], dp[x][w1][w2])
	}

	for pos := range n {
		for x := range m {
			for w1 := range 2 {
				for w2 := range 2 {
					if dp[x][w1][w2] == 0 {
						continue
					}
					if pos&1 == 0 {
						// 不能是d
						s, e := 0, 9
						if pos == 0 {
							s++
						}
						for i := s; i <= e; i++ {
							if i == d || w1 == 1 && i < int(a[pos]-'0') || w2 == 1 && i > int(b[pos]-'0') {
								continue
							}
							update(w1, w2, x, i, pos)
						}
					} else {
						// 只能是d
						if (w1 == 0 || d >= int(a[pos]-'0')) && (w2 == 0 || d <= int(b[pos]-'0')) {
							update(w1, w2, x, d, pos)
						}
					}
				}
			}
		}
		for x := range m {
			for w1 := range 2 {
				for w2 := range 2 {
					dp[x][w1][w2] = ndp[x][w1][w2]
					ndp[x][w1][w2] = 0
				}
			}
		}
	}

	var ans int
	for w1 := range 2 {
		for w2 := range 2 {
			ans = add(ans, dp[0][w1][w2])
		}
	}
	return ans
}

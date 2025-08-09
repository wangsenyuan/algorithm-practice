package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	fmt.Println(ans)
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
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) int {
	ans1 := solve1(a)
	ans2 := solve2(a)
	var ans3 int
	var ans4 int
	for range 4 {
		tmp := solve3(a)
		ans3 += tmp[0]
		ans4 += tmp[1]
		a = rotate(a)
	}
	return ans1 + ans2 + ans3/2 + ans4
}

func rotate(a []string) []string {
	n := len(a)
	m := len(a[0])
	buf := make([][]byte, m)
	for i := range m {
		buf[i] = make([]byte, n)
	}
	for i := range n {
		for j := range m {
			buf[j][n-1-i] = a[i][j]
		}
	}
	res := make([]string, m)
	for i := range m {
		res[i] = string(buf[i])
	}
	return res
}

func getPrefixSum(a []string) [][]int {
	n := len(a)
	m := len(a[0])
	sum := make([][]int, n+1)
	for i := range n + 1 {
		sum[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j]
			if a[i][j] == '#' {
				sum[i+1][j+1]++
			}
		}
	}
	return sum
}

func get(sum [][]int, x1, y1, x2, y2 int) int {
	return sum[x2+1][y2+1] - sum[x1][y2+1] - sum[x2+1][y1] + sum[x1][y1]
}

func solve1(a []string) int {
	sum := getPrefixSum(a)
	n := len(a)
	m := len(a[0])

	var res int

	for i := 1; i < n-1; i++ {
		if get(sum, i, 0, i, m-1) == 0 {
			res++
		}
	}
	for j := 1; j < m-1; j++ {
		if get(sum, 0, j, n-1, j) == 0 {
			res++
		}
	}
	return res
}

func solve2(a []string) int {
	sum := getPrefixSum(a)
	n := len(a)
	m := len(a[0])

	var res int

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if a[i][j] == '.' {
				if get(sum, i, 0, i, j) == 0 {
					// 从左边过来是空的
					if get(sum, 0, j, i, j) == 0 {
						res++
					}
					if get(sum, i, j, n-1, j) == 0 {
						res++
					}
				}
				if get(sum, i, j, i, m-1) == 0 {
					if get(sum, 0, j, i, j) == 0 {
						res++
					}
					if get(sum, i, j, n-1, j) == 0 {
						res++
					}
				}
			}
		}
	}

	return res
}

func solve3(a []string) []int {
	sum := getPrefixSum(a)
	n := len(a)
	m := len(a[0])

	res := make([]int, 4)

	for i := 1; i < n-1; i++ {
		var cnt int
		for j := 1; j < m-1; j++ {
			if a[i][j] == '.' {
				// 上右下结构
				if get(sum, i, j, n-1, j) == 0 {
					res[0] += cnt
				}
				if get(sum, 0, j, i, j) == 0 {
					res[1] += cnt
					if j > 1 && get(sum, 0, j-1, i, j-1) == 0 {
						// 靠近的U型，不符合条件
						res[1]--
					}
					cnt++
				}
			} else {
				cnt = 0
			}
		}
		cnt = 0
		for j := m - 2; j > 0; j-- {
			if a[i][j] == '.' {
				if get(sum, i, j, n-1, j) == 0 {
					if get(sum, i, j, n-1, j) == 0 {
						res[0] += cnt
					}
				}
				// 不需要考虑u型，上面已经包括了
				if get(sum, 0, j, i, j) == 0 {
					cnt++
				}
			} else {
				cnt = 0
			}
		}
	}

	return res
}

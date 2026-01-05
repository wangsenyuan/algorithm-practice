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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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

func drive(reader *bufio.Reader) int {
	n, _, k := readThreeNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a, k)
}

func solve(a []string, k int) int {
	n := len(a)
	m := len(a[0])
	sum := make([][]int, n+1)
	for i := range n + 1 {
		sum[i] = make([]int, m+1)
	}

	for i := range n {
		for j := range m {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j]
			if a[i][j] == '1' {
				sum[i+1][j+1]++
			}
		}
	}

	get := func(x1, y1, x2, y2 int) int {
		return sum[x2+1][y2+1] - sum[x1][y2+1] - sum[x2+1][y1] + sum[x1][y1]
	}

	if sum[n][m] < k {
		return 0
	}
	res := n * m

	for r1 := range n {
		for r2 := r1; r2 < n; r2++ {
			var c1 int
			for c2 := range m {
				for c1 < c2 && get(r1, c1, r2, c2) >= k {
					c1++
					if get(r1, c1, r2, c2) < k {
						c1--
						break
					}
				}
				if get(r1, c1, r2, c2) >= k {
					res = min(res, (r2-r1+1)*(c2-c1+1))
				}
			}
		}
	}
	return res
}

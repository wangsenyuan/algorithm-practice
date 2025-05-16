package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	res := process(reader)
	fmt.Fprintln(w, res)
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
	n, _, k := readThreeNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)
	}
	return solve(k, grid)
}

func solve(k int, grid []string) int {
	n := len(grid)
	m := len(grid[0])
	col := make([]int, m)
	cnt := make([]int, 26)
	var res int
	for r1 := 0; r1 < n; r1++ {
		clear(col)
		for r2 := r1; r2 < n; r2++ {
			var sum int
			clear(cnt)
			for c1, c2 := 0, 0; c2 < m; c2++ {
				if grid[r2][c2] == 'a' {
					col[c2]++
				}
				sum += col[c2]
				for sum > k {
					sum -= col[c1]
					if grid[r1][c1] == grid[r2][c1] {
						y := int(grid[r2][c1] - 'a')
						cnt[y]--
					}
					c1++
				}
				// sum <= k
				if grid[r2][c2] == grid[r1][c2] {
					y := int(grid[r2][c2] - 'a')
					if r1 < r2 {
						// 必须要有max(0)的检查。因为有可能同一列，在上面被减掉了
						res += max(cnt[y], 0)
					}
					cnt[y]++
				}
			}
		}
	}

	return res
}

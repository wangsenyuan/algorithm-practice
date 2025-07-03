package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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
	n, m := readTwoNums(reader)
	grid := make([][]int, n)
	for i := range n {
		grid[i] = readNNums(reader, m)
	}
	return solve(grid)
}

func solve(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	found := make([][]bool, n)
	for i := range n {
		found[i] = make([]bool, n)
	}

	cols := make([][]int, m)

	// 理解了，其实是n * n 的
	check := func(x int) bool {
		for i := range n {
			clear(found[i])
		}

		clear(cols)

		for r1 := 0; r1 < n; r1++ {
			for c := 0; c < m; c++ {
				if grid[r1][c] >= x {
					for _, r2 := range cols[c] {
						if found[r2][r1] {
							return true
						}
						found[r2][r1] = true
					}
					cols[c] = append(cols[c], r1)
				}
			}
		}

		return false
	}

	var nums []int
	for i := range n {
		for j := range m {
			nums = append(nums, grid[i][j])
		}
	}
	sort.Ints(nums)
	nums = slices.Compact(nums)

	pos := sort.Search(len(nums), func(i int) bool {
		return !check(nums[i])
	})

	return nums[pos-1]
}

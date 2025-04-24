package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%d", res)
	fmt.Println(s)
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

	readGrid := func(h int, w int) [][]int {
		grid := make([][]int, h)
		for i := range grid {
			grid[i] = readNNums(reader, w)
		}
		return grid
	}

	a := readGrid(n, m)
	b := readGrid(n, m)

	return solve(a, b)

}

func solve(a [][]int, b [][]int) int {
	n := len(a)
	m := len(a[0])

	rows := make([]int, n)
	cols := make([]int, m)
	for i := range rows {
		rows[i] = i
	}

	check := func(rows []int, cols []int) int {
		for i := range rows {
			for j := range cols {
				if a[rows[i]][cols[j]] != b[i][j] {
					return -1
				}
			}
		}

		return countSwaps(rows) + countSwaps(cols)
	}

	res := -1

	for {
		for i := range cols {
			cols[i] = i
		}

		for {
			cnt := check(rows, cols)

			if cnt >= 0 && (res < 0 || cnt < res) {
				res = cnt
			}

			if !nextPermutation(cols) {
				break
			}
		}

		if !nextPermutation(rows) {
			break
		}
	}

	return res
}

func nextPermutation(arr []int) bool {
	// Find longest decreasing suffix
	i := len(arr) - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}

	if i < 0 {
		return false // No next permutation
	}

	// Find successor to pivot in suffix
	j := len(arr) - 1
	for arr[j] <= arr[i] {
		j--
	}

	// Swap pivot with successor
	arr[i], arr[j] = arr[j], arr[i]

	// Reverse the suffix
	for k := 1; i+k < len(arr)-k; k++ {
		arr[i+k], arr[len(arr)-k] = arr[len(arr)-k], arr[i+k]
	}

	return true
}

func countSwaps(arr []int) int {
	n := len(arr)
	var res int

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[i] {
				res++
			}
		}
	}

	return res
}

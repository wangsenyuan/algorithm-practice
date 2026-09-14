package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	fmt.Fprintln(w, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	c := make([][]int, n)
	for i := range n {
		c[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &c[i][j])
		}
	}
	var a, b, cc int
	fmt.Fscan(reader, &a, &b, &cc)
	return solve(c, a, b, cc)
}

func solve(g [][]int, a, b, c int) int {
	n := len(g)
	m := len(g[0])
	row := make([]int, n)
	col := make([]int, m)
	var tot int
	for i, cur := range g {
		for j, v := range cur {
			row[i] += v
			col[j] += v
			tot += v
		}
	}

	if tot != a+b+c {
		return 0
	}

	arr := []int{a, b, c}
	slices.Sort(arr)

	var res int

	play := func(row []int) {
		// 这个0还比较麻烦
		var pref []int
		var sum int
		for i, v := range row {
			sum += v
			if sum > arr[0] {
				break
			}
			if sum == arr[0] {
				pref = append(pref, i)
			}
		}

		sum = 0
		for i := len(row) - 1; i > 0 && len(pref) > 0; i-- {
			sum += row[i]
			if sum > arr[2] {
				break
			}
			for len(pref) > 0 && pref[len(pref)-1] >= i-1 {
				pref = pref[:len(pref)-1]
			}

			if len(pref) == 0 {
				break
			}
			if sum == arr[2] {
				res += len(pref)
			}
		}
	}

	for {
		play(row)
		play(col)
		if !nextPermutation(arr) {
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

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	nums := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &nums[i])
	}
	return solve(k, nums)
}

const inf = 1 << 60

func solve(k int, nums []string) int {
	pos := make([]int, k)
	for i := range k {
		pos[i] = i
	}

	check := func(pos []int, num string) int {
		var res int
		for _, i := range pos {
			res = res*10 + int(num[i]-'0')
		}
		return res
	}

	play := func(pos []int) int {
		x, y := inf, -inf

		for _, num := range nums {
			tmp := check(pos, num)
			x = min(x, tmp)
			y = max(y, tmp)
		}
		return y - x
	}

	best := inf
	for {
		best = min(best, play(pos))
		if !nextPermutation(pos) {
			break
		}
	}
	return best
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

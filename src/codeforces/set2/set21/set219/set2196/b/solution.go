package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(reader, &nums[i])
	}

	return solve(nums)
}

const SQ = 500

func solve(nums []int) int64 {
	n := len(nums)
	var ans int
	for i, v := range nums {
		if v >= SQ {
			// w = i + w * v
			for w := 1; i+w*v < n; w++ {
				if nums[i+w*v] == w {
					ans++
				}
			}
			for w := 1; i-w*v >= 0; w++ {
				if nums[i-w*v] == w {
					ans++
				}
			}
			continue
		}

		for w := 1; i+w*v < n && w < SQ; w++ {
			if nums[i+w*v] == w {
				ans++
			}
		}
	}

	return int64(ans)
}

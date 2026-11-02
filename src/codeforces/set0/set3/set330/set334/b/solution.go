package main

import (
	"bufio"
	"fmt"
	"maps"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	pts := make([][]int, 8)
	for i := range pts {
		pts[i] = make([]int, 2)
		fmt.Fscan(reader, &pts[i][0], &pts[i][1])
	}
	return solve(pts)
}

type point struct {
	x int
	y int
}

func solve(pts [][]int) string {
	// x1 < x2 < x3
	// y1 < y2 < y3
	// x1 很好判断, 就是最小的那个点, 它对应的就是y1
	var nums []int
	freq := make(map[point]int)
	for _, pt := range pts {
		x, y := pt[0], pt[1]
		freq[point{x, y}]++
		nums = append(nums, pt[0], pt[1])
	}

	slices.Sort(nums)
	nums = slices.Compact(nums)
	if len(nums) < 3 {
		return "ugly"
	}
	n := len(nums)

	check := func(s1 int, s2 int) bool {
		var x []int
		var y []int
		for i := range n {
			if (s1>>i)&1 == 1 {
				x = append(x, nums[i])
			}
			if (s2>>i)&1 == 1 {
				y = append(y, nums[i])
			}
		}

		f2 := make(map[point]int)

		for i := range 3 {
			for j := range 3 {
				if i == j && i == 1 {
					continue
				}
				x, y := x[i], y[j]
				f2[point{x, y}]++
			}
		}

		return maps.Equal(freq, f2)
	}

	for s1 := 1; s1 < 1<<n; s1++ {
		w := bits.OnesCount(uint(s1))
		if w == 3 {
			for s2 := 1; s2 < 1<<n; s2++ {
				v := bits.OnesCount(uint(s2))
				if v == 3 && s1|s2 == (1<<n)-1 && check(s1, s2) {
					return "respectable"
				}
			}
		}
	}

	return "ugly"
}

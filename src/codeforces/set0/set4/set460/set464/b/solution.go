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
	res := drive(reader)
	if !res.ok {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for _, p := range res.pts {
		fmt.Println(p[0], p[1], p[2])
	}
}

type result struct {
	ok  bool
	pts [][]int
}

func drive(reader *bufio.Reader) result {
	lines := make([][]int, 8)
	for i := range 8 {
		lines[i] = make([]int, 3)
		fmt.Fscan(reader, &lines[i][0], &lines[i][1], &lines[i][2])
	}
	return solve(lines)
}

func solve(lines [][]int) result {

	var rows [][]int

	checkConcide := func(cur []int) bool {
		for _, row := range rows {
			if slices.Equal(row, cur) {
				return true
			}
		}
		return false
	}

	var play func(i int) bool

	play = func(i int) bool {
		if i == len(lines) {
			return check(rows)
		}

		cur := slices.Clone(lines[i])
		slices.Sort(cur)

		for cur != nil {
			if !checkConcide(cur) {
				rows = append(rows, cur)
				if play(i + 1) {
					return true
				}
				rows = rows[:len(rows)-1]
			}
			cur = nextPermutation(cur)
		}
		return false
	}

	ok := play(0)
	if !ok {
		return result{ok: false}
	}
	return result{ok: true, pts: rows}
}

type node struct {
	id   int
	dist int
}

func check(pts [][]int) bool {
	// 它们是否是一个cube的8个顶点
	// dist from first
	n := len(pts)
	d1 := 1 << 60
	dists := make([]int, n)
	for i := 1; i < n; i++ {
		dx := pts[i][0] - pts[0][0]
		dy := pts[i][1] - pts[0][1]
		dz := pts[i][2] - pts[0][2]
		dists[i] = dx*dx + dy*dy + dz*dz
		d1 = min(d1, dists[i])
	}

	cnt := make([]int, 3)
	for i := 1; i < n; i++ {
		switch dists[i] {
		case d1:
			cnt[0]++
		case 2 * d1:
			cnt[1]++
		case 3 * d1:
			cnt[2]++
		default:
			return false
		}
	}

	return cnt[0] == 3 && cnt[1] == 3 && cnt[2] == 1
}

func nextPermutation(nums []int) []int {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return nil
	}
	j := n - 1
	for j >= 0 && nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	sort.Ints(nums[i+1:])
	return nums
}

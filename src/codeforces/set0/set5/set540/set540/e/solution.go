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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	swaps := make([][]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		swaps[i] = []int{x, y}
	}
	return solve(swaps)
}

type pair struct {
	first  int
	second int
}

func solve(swaps [][]int) int {
	var nums []int
	for _, cur := range swaps {
		nums = append(nums, cur[0], cur[1])
	}
	sort.Ints(nums)
	nums = slices.Compact(nums)
	n := len(nums)
	arr := make([]pair, n)
	where := make([]int, n)
	for i, v := range nums {
		arr[i] = pair{v, v}
		where[i] = i
	}

	play := func(x int, y int) {
		if x > y {
			x, y = y, x
		}
		l := sort.SearchInts(nums, x)
		r := sort.SearchInts(nums, y)
		// l是x的排序后的位置
		i, j := where[l], where[r]
		arr[i].first = y
		arr[j].first = x
		where[r] = i
		where[l] = j
	}

	find := func(i int) int {
		// arr is sorted by second always
		pos := sort.Search(n, func(j int) bool {
			return arr[j].second >= i
		})
		return arr[pos].first
	}

	for _, cur := range swaps {
		i, j := cur[0], cur[1]
		// i, j 是两个位置
		x := find(i)
		y := find(j)
		play(x, y)
	}
	cnt := NewBit(n)

	var res int

	for _, cur := range arr {
		x, p := cur.first, cur.second
		// p是x现在所在的位置
		j := sort.SearchInts(nums, x)
		res += cnt.Get(n) - cnt.Get(j)

		if x != p {
			// 要知道中间没有出现的数
			l, r := min(x, p), max(x, p)
			pl := sort.SearchInts(nums, l)
			pr := sort.SearchInts(nums, r)
			if pr == n || nums[pr] > r {
				pr--
			}

			res += r - l + 1 - (pr - pl + 1)
		}

		cnt.Update(j, 1)
	}

	return res
}

type BIT []int

func NewBit(n int) BIT {
	arr := make([]int, n+3)
	return BIT(arr)
}

func (bit BIT) Update(pos int, val int) {
	pos++
	n := len(bit) - 1
	for pos <= n {
		bit[pos] += val
		pos += pos & (-pos)
	}
}

func (bit BIT) Get(pos int) int {
	pos++
	var res int
	for pos > 0 {
		res += bit[pos]
		pos -= pos & (-pos)
	}
	return res
}

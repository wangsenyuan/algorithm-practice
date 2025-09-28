package p3695

import "sort"

func maxAlternatingSum(nums []int, swaps [][]int) int64 {
	n := len(nums)

	set := NewDSU(n)
	for _, cur := range swaps {
		set.Union(cur[0], cur[1])
	}

	arr := make([][]int, n)

	for i := range n {
		j := set.Find(i)
		arr[j] = append(arr[j], nums[i])
	}

	for i := range n {
		sort.Ints(arr[i])
	}

	var sum int64
	for i := range n {
		j := set.Find(i)
		if i&1 == 1 {
			sum -= int64(arr[j][0])
			arr[j] = arr[j][1:]
		} else {
			sum += int64(arr[j][len(arr[j])-1])
			arr[j] = arr[j][:len(arr[j])-1]
		}
	}

	return sum
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

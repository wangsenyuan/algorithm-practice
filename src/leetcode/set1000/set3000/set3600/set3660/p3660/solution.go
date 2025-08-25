package p3660

import (
	"cmp"
	"slices"
)

type pair struct {
	first  int
	second int
}

func maxValue(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	pref := make([]int, n)
	pref[0] = nums[0]
	for i := 1; i < n; i++ {
		pref[i] = max(pref[i-1], nums[i])
	}
	ans[n-1] = pref[n-1]
	suf := 1 << 60
	for i := n - 1; i >= 0; i-- {
		if pref[i] <= suf {
			ans[i] = pref[i]
		} else {
			ans[i] = ans[i+1]
		}
		suf = min(suf, nums[i])
	}
	return ans
}

func maxValue1(nums []int) []int {
	n := len(nums)

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{nums[i], i}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	best := slices.Clone(nums)

	tr := NewTree(nums)
	// 感觉是个树么
	marked := make([]bool, n)
	for p := n - 1; p >= 0; p-- {
		cur := arr[p]
		i := cur.second
		if marked[i] {
			continue
		}
		u := i
		tmp := nums[i]
		for {
			for j := u; j < n && !marked[j]; j++ {
				tmp = min(tmp, nums[j])
				tr.Update(j, -1)
				best[j] = cur.first
				marked[j] = true
			}
			if tr.val[0] <= tmp {
				break
			}
			u = tr.QueryFirst(tmp)
		}
	}

	return best
}

type Tree struct {
	val []int
	sz  int
}

func NewTree(arr []int) *Tree {
	tr := new(Tree)
	n := len(arr)
	tr.sz = n
	tr.val = make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
	}
	build(0, 0, n-1)
	return tr
}

func (tr *Tree) QueryFirst(x int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.val[i*2+1] > x {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	return f(0, 0, tr.sz-1)
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
	}
	f(0, 0, tr.sz-1)
}

package p3962

import (
	"slices"
	"sort"
)

func maxSum(nums []int, k int) int64 {
	// 肯定可以把k个最大的数放在一起的,但是这个不一定是最优解
	arr := slices.Clone(nums)
	slices.Sort(arr)
	arr = slices.Compact(arr)
	// 如果在区间 l...r中间,少于0的数 x <= k, 那么就把它们都替换成外面最大的x个数
	// 如果在区间l...r 中间 ... > k, 那么就把它们替换成外面最大的k个数
	// 但是这样就能得到最优解吗?
	// 假设k次替换后,最优解是区间[l...r], 那么只能是把区间内和区间外进行交换,区间里面相互不需要交换
	// 用区间内,最小的数和区间外最大的数,依次匹配,在前k个里面选择最大值
	// 如果区间内第k大的数 > 区间外第k小的数,那么就不能交换这么多.
	// 否则交换k个数,是最优的.

	m := len(arr)
	tr1 := NewTree(m)
	for _, v := range nums {
		i := sort.SearchInts(arr, v)
		tr1.Add(i, v)
	}

	best := -inf
	if k > 0 {
		// 始终可以把最大的k个数放在一起
		best = tr1.GetMaxKSum(k)
	}

	tr2 := NewTree(m)

	n := len(nums)
	for r := range n {
		var sum int
		var w int
		delta := -inf
		for l := r; l >= 0; l-- {
			sum += nums[l]
			i := sort.SearchInts(arr, nums[l])
			tr1.Remove(i, nums[l])
			tr2.Add(i, nums[l])
			w2 := w
			for w1 := max(0, w-1); w1 <= min(k, w+1); w1++ {
				if w1 > min(tr1.cnt[0], tr2.cnt[0]) {
					break
				}
				tmp := sum + tr1.GetMaxKSum(w1) - tr2.GetMinKSum(w1)
				if tmp > delta {
					delta = tmp
					w2 = w1
				}
			}

			best = max(best, delta)
			w = w2
		}
		for l := range r + 1 {
			i := sort.SearchInts(arr, nums[l])
			tr1.Add(i, nums[l])
			tr2.Remove(i, nums[l])
		}
	}

	return int64(best)
}

type Tree struct {
	arr []int
	cnt []int
}

const inf = 1 << 60

func NewTree(n int) *Tree {
	arr := make([]int, 4*n)
	cnt := make([]int, 4*n)
	return &Tree{arr, cnt}
}

func (t *Tree) Add(pos int, v int) {
	t.update(pos, 0, 0, len(t.arr)/4-1, v, true)
}

func (t *Tree) Remove(pos int, v int) {
	t.update(pos, 0, 0, len(t.arr)/4-1, v, false)
}

func (t *Tree) update(pos int, i int, l int, r int, v int, flag bool) {
	if l == r {
		if flag {
			t.arr[i] += v
			t.cnt[i]++
		} else {
			t.arr[i] -= v
			t.cnt[i]--
		}
	} else {
		mid := (l + r) >> 1
		if pos <= mid {
			t.update(pos, 2*i+1, l, mid, v, flag)
		} else {
			t.update(pos, 2*i+2, mid+1, r, v, flag)
		}
		t.arr[i] = t.arr[i*2+1] + t.arr[i*2+2]
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
	}
}

func (t *Tree) GetMaxKSum(k int) int {
	if k == 0 {
		return 0
	}
	if t.cnt[0] <= k {
		return t.arr[0]
	}
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r {
			return t.arr[i] / t.cnt[i] * k
		}
		mid := (l + r) >> 1
		if t.cnt[i*2+2] >= k {
			return f(i*2+2, mid+1, r, k)
		}
		return t.arr[i*2+2] + f(i*2+1, l, mid, k-t.cnt[i*2+2])
	}
	return f(0, 0, len(t.arr)/4-1, k)
}

func (t *Tree) GetMinKSum(k int) int {
	if k == 0 {
		return 0
	}
	if t.cnt[0] <= k {
		return t.arr[0]
	}
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r {
			return t.arr[i] / t.cnt[i] * k
		}
		mid := (l + r) >> 1
		if t.cnt[i*2+1] >= k {
			return f(i*2+1, l, mid, k)
		}
		return t.arr[i*2+1] + f(i*2+2, mid+1, r, k-t.cnt[i*2+1])
	}
	return f(0, 0, len(t.arr)/4-1, k)
}

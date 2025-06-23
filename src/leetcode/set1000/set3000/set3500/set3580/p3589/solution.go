package p3589

import "slices"

func primeSubarray(nums []int, k int) int {
	lpf := getLpf(nums)
	n := len(nums)
	var pos []int
	for i, v := range nums {
		if lpf[v] == v {
			pos = append(pos, i)
		}
	}

	q1 := make([]int, 0, n)
	q2 := make([]int, 0, n)

	var ans int
	var it int
	var l int
	for r, v := range nums {
		if lpf[v] == v {
			for len(q1) > 0 && nums[last(q1)] < v {
				q1 = q1[:len(q1)-1]
			}
			q1 = append(q1, r)

			for len(q2) > 0 && nums[last(q2)] > v {
				q2 = q2[:len(q2)-1]
			}
			q2 = append(q2, r)
			it++

			for nums[first(q1)]-nums[first(q2)] > k {
				l++
				if first(q1) < l {
					q1 = q1[1:]
				}
				if first(q2) < l {
					q2 = q2[1:]
				}
			}
		}

		if it >= 2 && nums[first(q1)]-nums[first(q2)] <= k {
			ans += pos[it-2] - l + 1
		}
	}

	return ans
}

func last(q []int) int {
	return q[len(q)-1]
}

func first(q []int) int {
	return q[0]
}

func primeSubarray1(nums []int, k int) int {
	lpf := getLpf(nums)
	n := len(nums)
	var pos []int
	for i, v := range nums {
		if lpf[v] == v {
			pos = append(pos, i)
		}
	}
	tr := NewTree(n)

	var ans int
	var it int
	var l int
	for r, v := range nums {
		if lpf[v] == v {
			tr.Update(r, v)
			it++
		}
		for l < r && tr.GetMax(l, r)-tr.GetMin(l, r) > k {
			l++
		}
		if tr.GetMax(l, r)-tr.GetMin(l, r) <= k && it > 0 && pos[it-1] >= l {
			if it >= 2 {
				j := pos[it-2]
				if j >= l {
					ans += j - l + 1
				}
			}
		}
	}

	return ans
}

func getLpf(nums []int) []int {
	x := slices.Max(nums)
	var primes []int
	lpf := make([]int, x+1)
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > x {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}
	return lpf
}

type Tree struct {
	vals [][2]int
	sz   int
}

const inf = 1 << 30

func NewTree(n int) *Tree {
	vals := make([][2]int, 4*n)
	for i := range vals {
		vals[i] = [2]int{-inf, inf}
	}
	return &Tree{vals, n}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if l == r {
			tr.vals[i] = [2]int{v, v}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		tr.vals[i][0] = max(tr.vals[i*2+1][0], tr.vals[i*2+2][0])
		tr.vals[i][1] = min(tr.vals[i*2+1][1], tr.vals[i*2+2][1])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) GetMax(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) int
	loop = func(i int, l int, r int, L int, R int) int {
		if L == l && R == r {
			return tr.vals[i][0]
		}
		mid := (l + r) / 2
		if R <= mid {
			return loop(2*i+1, l, mid, L, R)
		}
		if mid < L {
			return loop(2*i+2, mid+1, r, L, R)
		}
		return max(loop(2*i+1, l, mid, L, mid), loop(2*i+2, mid+1, r, mid+1, R))
	}
	return loop(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) GetMin(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) int
	loop = func(i int, l int, r int, L int, R int) int {
		if L == l && R == r {
			return tr.vals[i][1]
		}
		mid := (l + r) / 2
		if R <= mid {
			return loop(2*i+1, l, mid, L, R)
		}
		if mid < L {
			return loop(2*i+2, mid+1, r, L, R)
		}
		return min(loop(2*i+1, l, mid, L, mid), loop(2*i+2, mid+1, r, mid+1, R))
	}
	return loop(0, 0, tr.sz-1, L, R)
}

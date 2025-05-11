package p3542

func minOperations(nums []int) int {
	var ans int
	st := nums[:0] // 原地
	for _, x := range nums {
		for len(st) > 0 && x < st[len(st)-1] {
			st = st[:len(st)-1]
			ans++
		}
		// 如果 x 与栈顶相同，那么 x 与栈顶可以在同一次操作中都变成 0，x 无需入栈
		if len(st) == 0 || x != st[len(st)-1] {
			st = append(st, x)
		}
	}
	if st[0] == 0 { // 0 不需要操作
		ans--
	}
	return ans + len(st)
}

func minOperations1(nums []int) int {

	n := len(nums)
	next := make([]int, n)
	pos := make(map[int]int)

	set := NewSegTree(n)

	for i := n - 1; i >= 0; i-- {
		if j, ok := pos[nums[i]]; ok {
			next[i] = j
		} else {
			next[i] = n
		}
		pos[nums[i]] = i
		set.Update(i, nums[i])
	}

	var loop func(l int, r int) int

	loop = func(l int, r int) int {
		x := set.Get(l, r+1)
		var ans int
		if x.first > 0 {
			ans++
		}
		if l < x.second {
			ans += loop(l, x.second-1)
		}
		i := x.second
		for i < r {
			j := min(next[i]-1, r)
			if i+1 <= j {
				ans += loop(i+1, j)
			}

			i = next[i]
		}
		return ans
	}
	return loop(0, len(nums)-1)
}

type pair struct {
	first  int
	second int
}

func min_pair(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

const inf = 1 << 60

type SegTree []pair

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = pair{inf, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[i*2], arr[i*2+1])
	}
	return arr
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p].first = v

	for p > 1 {
		tr[p>>1] = min_pair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) pair {
	n := len(tr) / 2
	l += n
	r += n
	res := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = min_pair(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min_pair(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

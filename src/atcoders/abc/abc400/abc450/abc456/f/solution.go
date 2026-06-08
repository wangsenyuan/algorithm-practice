package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintf(&buf, "%d\n", res)
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

const inf = 1 << 60

type FoldableQueue[T any] struct {
	op    func(T, T) T
	front []T
	back  []T
	last  T
}

func (q *FoldableQueue[T]) Push(v T) {
	if len(q.back) != 0 {
		q.last = q.op(q.last, v)
	} else {
		q.last = v
	}
	q.back = append(q.back, v)
}

func (q *FoldableQueue[T]) Pop() T {
	if len(q.front) == 0 {
		for i := len(q.back) - 1; i > 0; i-- {
			q.back[i-1] = q.op(q.back[i-1], q.back[i])
		}
		for i := 0; i < len(q.back); i++ {
			q.front = append(q.front, q.back[i])
		}
		q.back = q.back[:0]
	}

	res := q.front[0]
	q.front = q.front[1:]
	return res
}

func (q *FoldableQueue[T]) AllProd() T {
	if len(q.front) == 0 {
		return q.last
	}
	if len(q.back) == 0 {
		return q.front[0]
	}

	return q.op(q.front[0], q.last)
}

func (q *FoldableQueue[T]) Size() int {
	return len(q.front) + len(q.back)
}

type Mat [4]int

func matMul(b Mat, a Mat) Mat {
	var res Mat
	for i := range 2 {
		for j := range 2 {
			res[i*2+j] = min(a[i*2+0]+b[2*0+j], a[i*2+1]+b[2*1+j])
		}
	}
	return res
}

func solve(k int, nums []int) int {
	n := len(nums)
	arr := make([]Mat, n)
	for i := range n {
		arr[i] = Mat{inf, 0, nums[i], nums[i]}
	}

	que := &FoldableQueue[Mat]{op: matMul}
	var r int
	ans := inf
	for l := 0; l+k <= n; l++ {
		for r < l+k {
			que.Push(arr[r])
			r++
		}

		for que.Size() > k {
			que.Pop()
		}

		tmp := que.AllProd()
		ans = min(ans, tmp[2])
		if l > 0 {
			ans = min(ans, nums[l-1]+tmp[3])
		}
	}

	return ans
}

func solve1(k int, nums []int) int {
	if k == 1 {
		return slices.Min(nums)
	}
	n := len(nums)

	// val 表示这个节点上，没有限制（相邻不能取，还是要保证）的最优解
	// d[0/1][0/1] 第一个0/1表示左端点，第二个0/1表示右端点
	// 0表示不限制，1不选择
	type node struct {
		d00 int
		d01 int
		d10 int
		d11 int
	}

	merge := func(a node, b node) node {
		d00 := max(a.d00+b.d10, a.d01+b.d00)
		d01 := max(a.d00+b.d11, a.d01+b.d01)
		d10 := max(a.d10+b.d10, a.d11+b.d00)
		d11 := max(a.d11+b.d01, a.d10+b.d11)
		return node{d00, d01, d10, d11}
	}

	arr := make([]node, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			arr[i] = node{d00: nums[l]}
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		arr[i] = merge(arr[i*2+1], arr[i*2+2])
	}

	build(0, 0, n-1)

	var get func(i int, l int, r int, L int, R int) node

	get = func(i int, l int, r int, L int, R int) node {
		if l == L && r == R {
			return arr[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return get(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return get(i*2+2, mid+1, r, L, R)
		}
		a := get(i*2+1, l, mid, L, mid)
		b := get(i*2+2, mid+1, r, mid+1, R)
		return merge(a, b)
	}

	best := inf

	sum := make([]int, n+1)
	for i := range nums {
		sum[i+1] = sum[i] + nums[i]
	}

	check := func(l int, r int) int {
		if l < 0 {
			return inf
		}

		tmp := sum[r+1] - sum[l]
		if l+1 == r {
			return tmp
		}
		tmp -= get(0, 0, n-1, l+1, r-1).d00
		return tmp
	}

	for r := k - 1; r < n; r++ {
		l := r - k + 1
		best = min(best, check(l, r))
		best = min(best, check(l-1, r))
	}

	return best
}

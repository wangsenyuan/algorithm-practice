package p3578

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func countPartitions(nums []int, k int) int {

	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 1

	type pair struct {
		first  int
		second int
	}

	q1 := make([]pair, n)
	var h1, t1 int
	q2 := make([]pair, n)
	var h2, t2 int

	sum := 1

	for l, r := 0, 0; r < n; r++ {
		for h1 > t1 && q1[h1-1].first < nums[r] {
			h1--
		}
		q1[h1] = pair{nums[r], r}
		h1++
		for h2 > t2 && q2[h2-1].first > nums[r] {
			h2--
		}
		q2[h2] = pair{nums[r], r}
		h2++

		for q1[t1].first-q2[t2].first > k {
			sum = sub(sum, dp[l])
			if l == q1[t1].second {
				t1++
			}
			if l == q2[t2].second {
				t2++
			}
			l++
		}

		dp[r+1] = sum
		sum = add(sum, dp[r+1])
	}

	return dp[n]
}

func countPartitions1(nums []int, k int) int {
	// 对于一个r来说，假设区间l...r不满足条件,
	// 那么区间l-1...r也不满足条件
	// 假设dp[i]表示以i为右节点的区间的数量
	// dp[r] = sum(dp[l].... dp[r-1])
	// 貌似可以用双端队列来优化，最大值，最小值
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 1

	tr1 := NewSegTree(n, func(a, b int) int {
		return min(a, b)
	}, 1<<30)

	tr2 := NewSegTree(n, func(a, b int) int {
		return max(a, b)
	}, -1)

	sum := 1

	for l, r := 0, 1; r <= n; r++ {
		tr1.Update(r-1, nums[r-1])
		tr2.Update(r-1, nums[r-1])
		for tr2.Find(l, r)-tr1.Find(l, r) > k {
			sum = sub(sum, dp[l])
			l++
		}
		dp[r] = sum
		sum = add(sum, dp[r])
	}

	return dp[n]
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewSegTree(n int, fn func(int, int) int, iv int) *SegTree {
	size := n
	arr := make([]int, 2*size)
	for i := range arr {
		arr[i] = iv
	}
	return &SegTree{arr, size, iv, fn}
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}

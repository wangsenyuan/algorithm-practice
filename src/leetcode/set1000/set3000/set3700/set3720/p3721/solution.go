package p3721

func longestBalanced(nums []int) int {
	n := len(nums)
	// 0 is for sum, 1 is for max
	t := NewTree(n + 1)
	var sum int

	pos := make(map[int]int)

	var ans int

	for i := 1; i <= n; i++ {
		v := 1
		if nums[i-1]&1 == 0 {
			v = -1
		}

		if j, ok := pos[nums[i-1]]; ok {
			t.Update(j, n, -v)
			sum -= v
		}
		pos[nums[i-1]] = i
		t.Update(i, n, v)
		sum += v
		j := t.Find(sum)
		ans = max(ans, i-j)
	}

	return ans
}

type Tree struct {
	val  [][2]int
	lazy []int
}

func NewTree(n int) *Tree {
	return &Tree{
		val:  make([][2]int, n*4),
		lazy: make([]int, n*4),
	}
}

func (t *Tree) update(i int, v int) {
	t.val[i][0] += v
	t.val[i][1] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.update(i*2+1, t.lazy[i])
		t.update(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i][0] = min(t.val[i*2+1][0], t.val[i*2+2][0])
	t.val[i][1] = max(t.val[i*2+1][1], t.val[i*2+2][1])
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.update(i, v)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if R <= mid {
			f(i*2+1, l, mid, L, R)
		} else if mid < L {
			f(i*2+2, mid+1, r, L, R)
		} else {
			f(i*2+1, l, mid, L, mid)
			f(i*2+2, mid+1, r, mid+1, R)
		}
		t.pull(i)
	}
	n := len(t.val) / 4
	f(0, 0, n-1, L, R)
}

func (t *Tree) Find(v int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		t.push(i)
		mid := (l + r) / 2
		if t.val[i*2+1][0] <= v && v <= t.val[i*2+1][1] {
			return f(2*i+1, l, mid)
		}
		return f(2*i+2, mid+1, r)
	}
	n := len(t.val) / 4
	return f(0, 0, n-1)
}

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
	fmt.Println(process(reader))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 2)
	}
	return solve(m, segs)
}

type Seg struct {
	l  int
	r  int
	sz int
}

const inf = 1 << 60

func solve(m int, segs [][]int) int {
	// 区间压缩
	var nums []int
	for _, seg := range segs {
		nums = append(nums, seg[0], seg[1])
	}
	nums = sortAndUnique(nums)
	n := len(nums)

	arr := make([]Seg, len(segs))
	for i, cur := range segs {
		l := sort.SearchInts(nums, cur[0])
		r := sort.SearchInts(nums, cur[1])
		sz := cur[1] - cur[0]
		arr[i] = Seg{l, r, sz}
	}

	slices.SortFunc(arr, func(i Seg, j Seg) int {
		if i.sz != j.sz {
			return i.sz - j.sz
		}
		return i.l - j.l
	})

	tr := NewTree(n)

	ans := inf

	for i, j := 0, 0; j < len(arr); j++ {
		tr.Update(arr[j].l, arr[j].r, 1)
		for tr.Get(0, n-1) >= m {
			ans = min(ans, arr[j].sz-arr[i].sz)
			tr.Update(arr[i].l, arr[i].r, -1)
			if tr.Get(0, n-1) < m {
				tr.Update(arr[i].l, arr[i].r, 1)
				break
			}
			i++
		}
	}

	if ans == inf {
		return -1
	}

	return ans
}

func sortAndUnique(nums []int) []int {
	sort.Ints(nums)
	var n int
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[i-1] {
			nums[n] = nums[i-1]
			n++
		}
	}
	return nums[:n]
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]int, n*4)
	lazy := make([]int, n*4)
	return &Tree{val, lazy, n}
}

func (t *Tree) update(i int, v int) {
	t.val[i] += v
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
	t.val[i] = max(t.val[i*2+1], t.val[i*2+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int, L int, R int)
	loop = func(i int, l int, r int, L int, R int) {
		if L == l && R == r {
			t.update(i, v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			loop(i*2+1, l, mid, L, R)
		} else if mid < L {
			loop(i*2+2, mid+1, r, L, R)
		} else {
			loop(i*2+1, l, mid, L, mid)
			loop(i*2+2, mid+1, r, mid+1, R)
		}
		t.pull(i)
	}

	loop(0, 0, t.sz-1, L, R)
}

func (t *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) int
	loop = func(i int, l int, r int, L int, R int) int {
		if L == l && R == r {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			return loop(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return loop(i*2+2, mid+1, r, L, R)
		}
		return max(loop(i*2+1, l, mid, L, mid), loop(i*2+2, mid+1, r, mid+1, R))
	}
	return loop(0, 0, t.sz-1, L, R)
}

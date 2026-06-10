package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) (res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	qs := make([]int, q)
	for i := range qs {
		fmt.Fscan(reader, &qs[i])
	}
	res = solve(m, a, qs)
	return
}

type pair struct {
	first  int
	second int
}

func solve(m int, a []int, qs []int) []int {
	// TODO: solve by hand first.
	n := len(a)
	freq := make([]int, m)
	for _, v := range a {
		freq[v-1]++
	}
	arr := make([]pair, m)
	for i := range m {
		arr[i] = pair{freq[i], i}
	}
	slices.SortFunc(arr, func(x pair, y pair) int {
		return cmp.Or(x.first-y.first, x.second-y.second)
	})

	pref := make([]int, m+1)
	for i, cur := range arr {
		pref[i+1] = pref[i] + cur.first
	}

	w := arr[m-1].first

	find := func(k int) pair {
		if k > w*m-n {
			return pair{m + 1, k - (w*m - n)}
		}
		// k < w * m - n
		// 这个地方感觉要二分.但是我不知道呐
		// 假设i是第一次被添加进来, 那么它前面的部分,都要和它有相同的freq
		// (freq[i]) * i - sum[i] 这个是操作的次数, 那么这样子可以找到最后一个要被添加进来的字符
		// 对于例子 [1, 1, 2], arr = (3, 0), (2, 1), (1, 2)
		// 2第一次被添加,需要经过1次操作后, 1第一次被添加,需要经过3次后(2次3,1次2)
		i := sort.Search(m, func(i int) bool {
			return i*arr[i].first-pref[i] >= k
		})
		i--
		tot := i*arr[i].first - pref[i]
		k -= tot
		// 然后还需要知道前i+1个数字,具体是什么, 比如(2, 3), 那么在这个后面开始,就是2,3,2,3这样循环,直到下一个数字
		k = (k - 1) % (i + 1)
		return pair{i + 1, k + 1}
	}

	ans := make([]int, len(qs))

	todo := make([][]pair, m+1)

	for i, k := range qs {
		if k <= n {
			ans[i] = a[k-1]
			continue
		}
		tmp := find(k - n)
		if tmp.first == m+1 {
			ans[i] = (tmp.second-1)%m + 1
		} else {
			todo[tmp.first] = append(todo[tmp.first], pair{i, tmp.second})
		}
	}

	tr := NewTree(m)

	for i := 1; i <= m; i++ {
		tr.Set(arr[i-1].second)
		for _, tmp := range todo[i] {
			id, k := tmp.first, tmp.second
			ans[id] = tr.GetKth(k) + 1
		}
	}

	return ans
}

type Tree struct {
	cnt []int
}

func NewTree(n int) *Tree {
	return &Tree{
		cnt: make([]int, 4*n),
	}
}

func (tr *Tree) Set(p int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i]++
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	}
	f(0, 0, len(tr.cnt)/4-1)
}

func (tr *Tree) GetKth(k int) int {
	var f func(i int, l int, r int, k int) int
	f = func(i int, l int, r int, k int) int {
		if l == r {
			return l
		}
		mid := (l + r) >> 1
		if tr.cnt[2*i+1] >= k {
			return f(i*2+1, l, mid, k)
		}
		return f(i*2+2, mid+1, r, k-tr.cnt[i*2+1])
	}

	return f(0, 0, len(tr.cnt)/4-1, k)
}

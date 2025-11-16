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
	_, _, best, res := drive(reader)
	fmt.Println(best)
	fmt.Println(len(res))
	for _, cur := range res {
		fmt.Println(cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (a []int, b []int, best int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	b = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	best, res = solve(a, b)
	return
}

func solve(a []int, b []int) (int, [][]int) {
	var s1, s2 int
	for _, v := range a {
		s1 += v
	}

	for _, v := range b {
		s2 += v
	}

	best := abs(s1 - s2)
	var res [][]int
	tmp1, res1 := solve1(a, s1, b, s2)
	if tmp1 < best {
		res = res1
		best = tmp1
	}

	tmp2, res2 := solve2(a, s1, b, s2)
	if tmp2 < best {
		res = res2
		best = tmp2
	}

	return best, res
}

func abs(num int) int {
	return max(num, -num)
}

type pair struct {
	first  int
	second int
}

func solve1(a []int, s1 int, b []int, s2 int) (best int, res [][]int) {
	// 只交换一个位置能够得到的最优解
	// x1 = s1 - a[i] + b[j]
	// y1 = s2 + a[i] - b[j]
	// y1 - x1 = s2 - s1 + 2 * a[i] - 2 * b[j]
	// y1 - x1 = s2 - 2 * b[j] - (s1 - 2 * a[i])
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{s1 - 2*a[i], i}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	best = inf

	res = make([][]int, 1)

	for i, v := range b {
		tmp := s2 - 2*v
		j := sort.Search(n, func(j int) bool {
			return arr[j].first >= tmp
		})
		if j < n && arr[j].first-tmp < best {
			best = arr[j].first - tmp
			res[0] = []int{arr[j].second + 1, i + 1}
		}
		j--
		if j >= 0 && tmp-arr[j].first < best {
			best = tmp - arr[j].first
			res[0] = []int{arr[j].second + 1, i + 1}
		}
	}

	return
}

const inf = 1 << 60

func solve2(a []int, s1 int, b []int, s2 int) (best int, res [][]int) {
	if len(a) == 1 || len(b) == 1 {
		return inf, nil
	}

	best = inf
	res = make([][]int, 2)

	n := len(a)
	var arr []pair

	for i := range n {
		for j := range i {
			arr = append(arr, pair{s1 - 2*a[i] - 2*a[j], i*n + j})
		}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	update := func(p int, val int, idx1 int, idx2 int) {
		if p < 0 || p == len(arr) {
			return
		}
		cur := abs(arr[p].first - val)
		if cur < best {
			best = cur
			u, v := arr[p].second/n, arr[p].second%n
			res[0] = []int{u + 1, idx1 + 1}
			res[1] = []int{v + 1, idx2 + 1}
		}
	}

	m := len(b)
	for i := range m {
		for j := range i {
			tmp := s2 - 2*b[i] - 2*b[j]

			k := sort.Search(len(arr), func(k int) bool {
				return arr[k].first >= tmp
			})
			update(k, tmp, i, j)
			update(k-1, tmp, i, j)
		}
	}

	return
}

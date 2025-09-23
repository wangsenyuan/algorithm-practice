package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, _, best, ans := drive(reader)
	fmt.Println(best)
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

type pair struct {
	first  int
	second int
}

func drive(reader *bufio.Reader) (A int, cf int, cm int, m int, a []int, best int, ans []int) {
	var n int
	fmt.Fscan(reader, &n, &A, &cf, &cm, &m)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	best, ans = solve(A, cf, cm, m, a)
	return
}

func solve(A int, cf int, cm int, m int, a []int) (best int, ans []int) {
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return b.first - a.first
	})

	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + arr[i].first
	}

	find := func(l int, tmp int) int {
		if l == n {
			return A
		}

		r := n
		for l < r {
			mid := (l + r) / 2
			// 能否得到arr[mid].first
			s1 := arr[mid].first*(n-mid) - suf[mid]
			if s1 <= tmp {
				r = mid
			} else {
				l = mid + 1
			}
		}

		if l == n {
			// 这里tmp太小了？
			return tmp
		}
		return min(A, (tmp+suf[l])/(n-l))
	}

	bestAt := 0
	bestW := -1
	var sum int
	for i := 0; i <= n; i++ {
		// 这些应该放到最后面几个上面
		tmp := m - sum
		// 假设最小值 w,  w * cnt - suf <= tmp
		w := find(i, tmp)
		if i*cf+w*cm > best {
			best = i*cf + w*cm
			bestAt = i
			bestW = w
		}
		if i == n || sum+A-arr[i].first > m {
			break
		}
		sum += A - arr[i].first
	}

	ans = slices.Clone(a)

	for i := range bestAt {
		ans[arr[i].second] = A
	}

	for i := bestAt; i < n; i++ {
		ans[arr[i].second] = max(arr[i].first, bestW)
	}

	return best, ans
}

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
	_, _, best, changes, res := drive(reader)
	fmt.Println(best, changes)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (m int, a []int, best int, changes int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &m)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	best, changes, res = solve(m, a)
	return
}

func solve(m int, a []int) (best int, changes int, res []int) {
	// 超过m的，都应该修改掉？好像也不一定
	// 只有当超过的部分，能够让喜欢的band增加，才需要修改

	// n := len(a)
	cnt := make([]int, m+1)

	var bad []int
	for i, v := range a {
		if v <= m {
			cnt[v]++
		} else {
			bad = append(bad, i)
		}
	}

	check := func(expect int) bool {
		// 1...m都需要这么多
		var sum int
		for i := 1; i <= m; i++ {
			sum += max(0, cnt[i]-expect)
		}
		sum += len(bad)
		for i := 1; i <= m; i++ {
			if cnt[i] < expect {
				need := expect - cnt[i]
				if need > sum {
					return false
				}
				sum -= need
			}
		}
		return true
	}

	best = sort.Search(len(a)+1, func(expect int) bool {
		return !check(expect)
	})
	// 可以到达这么多
	best--
	// 凡是超过的都可以被修改，但优先修改bad
	for i, v := range a {
		if v <= m && cnt[v] > best {
			cnt[v]--
			bad = append(bad, i)
		}
	}

	res = slices.Clone(a)

	for v := 1; v <= m; v++ {
		for cnt[v] < best {
			changes++
			i := bad[0]
			bad = bad[1:]
			res[i] = v
			cnt[v]++
		}
	}

	return
}

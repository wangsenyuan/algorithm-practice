package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, k, m, a int
	fmt.Fscan(reader, &n, &k, &m, &a)
	g := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(reader, &g[i])
	}
	return solve(n, k, m, a, g)
}

type data struct {
	cnt  int
	last int
	id   int
}

func solve(n int, k int, m int, a int, g []int) []int {
	arr := make([]data, n)
	for i := range n {
		arr[i].id = i
	}

	reset := func() {
		for i := range n {
			arr[i].id = i
			arr[i].cnt = 0
			arr[i].last = 0
		}
	}

	check1 := func(i int) bool {
		reset()
		for j, v := range g {
			v--
			arr[v].cnt++
			arr[v].last = j
		}
		if arr[i].cnt == 0 {
			return false
		}

		slices.SortFunc(arr, func(x, y data) int {
			return cmp.Or(y.cnt-x.cnt, x.last-y.last)
		})
		// 已经排在i前面的人，不需要分配选票
		var pos int
		for pos < n && arr[pos].id != i {
			pos++
		}
		if pos >= k {
			return false
		}
		// 还有w张选票
		w := m - a
		v := k - pos
		for j := pos + 1; j < n; j++ {
			w -= (arr[pos].cnt + 1 - arr[j].cnt)
			if w < 0 {
				return true
			}
			// w >= 0
			v--
			if v == 0 {
				return false
			}
		}
		return true
	}

	// 第i个人是否肯定会落选
	check2 := func(i int) bool {
		reset()
		for j, v := range g {
			v--
			arr[v].cnt++
			arr[v].last = j
		}
		arr[i].cnt += m - a
		if m-a > 0 {
			arr[i].last = n
		}
		if arr[i].cnt == 0 {
			return true
		}
		slices.SortFunc(arr, func(x, y data) int {
			return cmp.Or(y.cnt-x.cnt, x.last-y.last)
		})
		for j := range n {
			if arr[j].id == i {
				return j >= k
			}
		}
		return false
	}

	ans := make([]int, n)

	for i := range n {
		tmp := check1(i)
		if tmp {
			ans[i] = 1
		} else {
			tmp = check2(i)
			if tmp {
				ans[i] = 3
			} else {
				ans[i] = 2
			}
		}
	}
	return ans
}

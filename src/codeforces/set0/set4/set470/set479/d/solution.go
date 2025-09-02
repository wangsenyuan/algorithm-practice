package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	if len(res) > 0 {
		for _, v := range res {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (a []int, x int, y int, res []int) {
	var n, l int
	fmt.Fscan(reader, &n, &l, &x, &y)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a, x, y)
	return
}

func solve(a []int, x int, y int) []int {
	var flag int
	n := len(a)

	search := func(v int, x int) bool {
		j := sort.SearchInts(a, v+x)
		if j < n && a[j] == v+x {
			return true
		}
		j = sort.SearchInts(a, v-x)
		if j < n && a[j] == v-x {
			return true
		}
		return false
	}

	for i := 0; i < n; i++ {
		if search(a[i], x) {
			flag |= 1
		}
		if search(a[i], y) {
			flag |= 2
		}
	}
	if flag == 3 {
		return nil
	}
	if flag == 1 {
		return []int{y}
	}
	if flag == 2 {
		return []int{x}
	}
	// flag == 0
	for i := 0; i < n; i++ {
		// 先添加一个
		if a[i]+y <= a[n-1] && search(a[i]+y, x) {
			return []int{a[i] + y}
		}
		if a[i]-y >= 0 && search(a[i]-y, x) {
			return []int{a[i] - y}
		}
		if a[i]+x <= a[n-1] && search(a[i]+x, y) {
			return []int{a[i] + x}
		}
		if a[i]-x >= 0 && search(a[i]-x, y) {
			return []int{a[i] - x}
		}
	}
	return []int{x, y}
}

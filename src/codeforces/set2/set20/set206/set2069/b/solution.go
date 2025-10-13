package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a [][]int) int {
	n := len(a)
	m := len(a[0])
	var arr []pair
	for i := range n {
		for j := range m {
			cnt := 1
			if i > 0 && a[i-1][j] == a[i][j] || j > 0 && a[i][j-1] == a[i][j] {
				cnt++
			}
			arr = append(arr, pair{a[i][j], cnt})
		}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	k := len(arr)
	var sum int
	var mx int
	for i := 0; i < k; {
		j := i
		var cnt int
		for i < k && arr[i].first == arr[j].first {
			cnt = max(cnt, arr[i].second)
			i++
		}
		sum += cnt
		mx = max(mx, cnt)
	}
	return sum - mx
}

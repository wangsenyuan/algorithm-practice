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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)

	for range tc {
		res := drive(reader)
		if res {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func drive(reader *bufio.Reader) bool {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) bool {
	arr := slices.Clone(a)
	slices.Sort(arr)
	w := arr[k-1]
	// < w 的数删除不了
	l := sort.SearchInts(arr, w)
	// k - l 至少保留这么多的w
	var a1 []int
	for _, v := range a {
		if v <= w {
			a1 = append(a1, v)
		}
	}
	// xyw..  wyxw
	var cnt int
	for i, j := 0, len(a1)-1; i <= j; {
		if a1[i] == a1[j] {
			if a1[i] == w {
				if i == j {
					cnt++
				} else {
					cnt += 2
				}
			}
			i++
			j--
			continue
		} else {
			// a1[i] != a1[j], must delete one
			if a1[i] < w && a1[j] < w {
				return false
			}
			// 删掉一个
			if a1[i] == w {
				i++
			} else {
				j--
			}
		}
	}

	return cnt >= k-l-1
}

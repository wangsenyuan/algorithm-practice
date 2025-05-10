package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	var k int
	fmt.Scanf("%d", &k)
	res := solve(k)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, s := range res {
		buf.WriteString(s)
		buf.WriteByte('\n')
	}

	buf.WriteTo(os.Stdout)
}

func solve(k int) []string {
	n := sort.Search(100, func(n int) bool {
		return n*(n-1)*(n-2)/6 > k
	})
	n--
	// 共n个点

	arr := make([][]int, n)
	for i := range n {
		for j := range n {
			if i != j {
				arr[i] = append(arr[i], j)
			}
		}
	}
	tot := n * (n - 1) * (n - 2) / 6

	for k > tot {
		diff := k - tot
		arr = append(arr, make([]int, 0, 1))
		// 然后增加一个节点去连接任意的两个节点
		m := sort.Search(n, func(m int) bool {
			return m*(m-1)/2 > diff
		})
		m--
		for i := range m {
			arr[i] = append(arr[i], n)
			arr[n] = append(arr[n], i)
		}
		tot += m * (m - 1) / 2
		n++
	}

	ans := make([]string, len(arr))
	for i, v := range arr {
		buf := make([]byte, len(arr))
		for j := range buf {
			buf[j] = '0'
		}
		for _, j := range v {
			buf[j] = '1'
		}
		ans[i] = string(buf)
	}
	return ans
}

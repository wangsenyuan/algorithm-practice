package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inf, res := drive(reader)
	if inf {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d ", v))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) (inf bool, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	inf, res = solve(a)
	return
}

func solve(a []int) (inf bool, res []int) {
	n := len(a)
	if n == 1 {
		return true, nil
	}
	sort.Ints(a)
	// 一类是，多出来的这张牌是第一张，第2张，或者是之后的某一张
	// 先考虑它是第一张的情况
	if check(a, a[1]-a[0]) {
		// 放在两头
		res = append(res, a[0]-(a[1]-a[0]))
		res = append(res, a[n-1]+(a[1]-a[0]))
	} else {
		// 在某个地方缺失了
		diff := a[1] - a[0]
		prev := a[1]
		var cnt int
		for i := 2; i < n; i++ {
			if a[i]-prev != diff {
				// 把 a[i] - diff 补上
				if a[i]-prev == 2*diff {
					cnt++
				} else {
					// 如果出现这种情况，说明这种情况下，没有解
					cnt = -n
				}
			}
			prev = a[i]
		}
		if cnt == 1 {
			prev = a[1]
			for i := 2; i < n; i++ {
				if a[i]-prev != diff {
					res = append(res, a[i]-diff)
					break
				}
				prev = a[i]
			}
		}
	}
	// 放在a[0],a[1]中间
	if (a[1]-a[0])%2 == 0 && check(a[1:], (a[1]-a[0])/2) {
		res = append(res, a[0]+(a[1]-a[0])/2)
	}

	slices.Sort(res)
	res = slices.Compact(res)
	return false, res
}

func check(a []int, diff int) bool {
	for i := 0; i+1 < len(a); i++ {
		if a[i+1]-a[i] != diff {
			return false
		}
	}
	return true
}

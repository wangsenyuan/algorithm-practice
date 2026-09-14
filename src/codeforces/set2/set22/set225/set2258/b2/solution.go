package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return formatInts(solve(m, a))
}

func formatInts(a []int) string {
	s := fmt.Sprintf("%v", a)
	return s[1 : len(s)-1]
}

func solve(m int, a []int) []int {
	ans := make([]int, m)

	freq := make([]int, m+2)
	var sum int

	for _, v := range a {
		freq[v]++
		sum += v
	}

	for i := m - 1; i >= 0; i-- {
		freq[i] += freq[i+1]
	}

	for k := 1; k <= m; k++ {
		if k > 18 || 1<<k > m {
			ans[k-1] = sum
		} else {
			for x := 1; x <= m/(1<<k)+3; x++ {
				var cnt int

				for j := 1; j < 1<<k && x*j <= m; j++ {
					cnt += freq[x*j]
				}
				if x*(1<<k) <= m {
					cnt += freq[x*(1<<k)] - freq[x*(1<<k)+1]
				}

				ans[k-1] = max(ans[k-1], cnt)
			}
		}
	}

	return ans
}

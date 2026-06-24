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
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	// 两边是连续的,
	n := len(a)
	// 从i, 开始, 是否有长度为1,2,3...的递增序列, 不能有重复的元素
	// fp[j][x] = w
	// dp[l][r] = true 这段是连续的, 然后对于这样的一段, 找到它的最大值, w
	// 那么另外一段, 必须是包含w+1的,且长度正好是r-l+1, 且在l...r的外面(只要是以w+1为最小值的, 肯定在外面)
	todo := make([][]int, n+1)

	freq := make([]int, n+1)

	for i := range n {
		j := i
		var sum int
		w := n + 1
		for j < n {
			if freq[a[j]] > 0 {
				// duplicates
				break
			}
			freq[a[j]]++
			sum += a[j]
			w = min(w, a[j])
			// 如果这段是一个连续的, w, w + 1, w + 2, ... w + j - i
			// w * (j-i+1) + (0 + j - i) * (j - i + 1) / 2
			if w*(j-i+1)+(j-i)*(j-i+1)/2 == sum {
				// 连续的一段
				todo[j-i+1] = append(todo[j-i+1], w)
			}
			j++
		}
		for j1 := i; j1 < j; j1++ {
			freq[a[j1]]--
		}
	}

	var best int
	for d := 1; d <= n/2; d++ {
		slices.Sort(todo[d])
		for _, v := range todo[d] {
			i := sort.SearchInts(todo[d], v+d)
			if i < len(todo[d]) && todo[d][i] == v+d {
				best = d
			}
		}
	}

	return best
}

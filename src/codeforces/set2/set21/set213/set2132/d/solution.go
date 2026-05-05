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
		var k int
		fmt.Fscan(reader, &k)
		res := solve(k)
		fmt.Fprintln(writer, res)
	}
}

func solve(k int) int {

	count := func(n int) int {
		var sum int
		cur := 1
		cnt := 1
		for cur <= n {
			next := min(10*cur-1, n)
			sum += (next - cur + 1) * cnt
			cur = next + 1
			cnt++
		}
		return sum
	}
	check := func(n int) bool {
		// 1...n sum count >= k
		return count(n) > k
	}

	n := sort.Search(k+1, check)
	n--
	// count(n) <= k
	// 计算digits sum
	// 123...n

	var digits []int
	for i := n; i > 0; i /= 10 {
		digits = append(digits, i%10)
	}
	slices.Reverse(digits)
	h := len(digits)

	bases := make([]int, h+1)
	bases[0] = 1
	for i := 1; i <= h; i++ {
		bases[i] = bases[i-1] * 10
	}
	var res int
	var hi int
	for i, v := range digits {
		// 有多少个0...v-1
		// 0...hi
		for x := range v {
			res += (hi + 1) * x * bases[h-i-1]
		}
		// 0...hi-1
		for x := v; x < 10; x++ {
			res += hi * x * bases[h-i-1]
		}
		hi = hi*10 + v
		lo := n - hi*bases[h-i-1]
		// 如果当前位是v，且高位一致的情况下，
		// 0....lo
		res += v * (lo + 1)
	}

	w := count(n)
	if w < k {
		s := fmt.Sprintf("%d", n+1)
		s = s[:k-w]
		for _, c := range s {
			res += int(c - '0')
		}
	}

	return res
}

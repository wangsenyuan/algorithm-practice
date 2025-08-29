package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(solve(s))
}

func solve(s string) int {
	n := len(s)
	B := int(math.Ceil(math.Sqrt(float64(n))))

	cnt := make([]int32, B*n+1)

	ans := 0
	for k := 1; k < B; k++ {
		pre := 0
		for _, b := range s {
			cnt[pre+n]++
			pre += k*int(b&1) - 1
			ans += int(cnt[pre+n])
		}
		pre = 0
		for _, b := range s {
			cnt[pre+n]--
			pre += k*int(b&1) - 1
		}
	}

	pos1 := []int{-1}
	for i, b := range s {
		if b == '1' {
			pos1 = append(pos1, i)
		}
	}

	tot1 := 0
	for i, b := range s {
		tot1 += int(b & 1)
		mn := min(tot1, n/B)
		for cnt1 := 1; cnt1 <= mn; cnt1++ {
			j := tot1 - cnt1
			maxK := (i - pos1[j]) / cnt1
			minK := max((i-pos1[j+1])/cnt1+1, B)
			ans += max(maxK-minK+1, 0)
		}
	}
	return ans
}

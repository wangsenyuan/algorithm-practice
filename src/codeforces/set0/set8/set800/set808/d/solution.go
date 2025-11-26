package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) bool {
	// n := len(a)
	var sum int
	for _, v := range a {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	half := sum / 2

	s1, s2 := 0, sum
	var pos int
	n := len(a)
	for pos < n {
		s1 += a[pos]
		s2 -= a[pos]
		if s1 == s2 {
			return true
		}
		if s1 > half {
			s1 -= a[pos]
			s2 += a[pos]
			break
		}
		pos++
	}
	// a[:pos]是前半部分，且s1 < half, s2 > half
	// 如果将某个pos前面的数移动到后面
	// s1 - x + a[pos] = half
	x := s1 + a[pos] - half
	for i := range pos {
		if a[i] == x {
			return true
		}
	}

	// 或者将后面的某个数，移动到第一位
	// s1 + x = half
	for i := pos; i < n; i++ {
		if a[i] == half-s1 {
			return true
		}
	}

	for i := pos + 1; i < n && s1 < half; i++ {
		s1 += a[i]
		if s1 == half {
			return true
		}
	}

	return false
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	if res < 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
		fmt.Println(res)
	}
}

func solve(s string) int {
	n := len(s)
	mex := make([]bool, n+2)
	grundy := make([]int, n+1)
	for i := 1; i <= n; i++ {
		clear(mex)
		for j := range i {
			l := max(0, j-1)
			r := max(0, i-j-2)
			x := grundy[l] ^ grundy[r]
			if x <= n {
				mex[x] = true
			}
		}
		var j int
		for mex[j] {
			j++
		}
		grundy[i] = j
	}
	var sum int
	for i := 1; i+1 < n; i++ {
		if s[i-1] == s[i+1] {
			j := i
			for j+2 < n && s[j] == s[j+2] {
				j++
			}
			sum ^= grundy[j-i+1]
			i = j + 1
		}
	}
	if sum == 0 {
		return -1
	}
	// first

	for i := 1; i < n-1; i++ {
		if s[i-1] == s[i+1] {
			j := i
			for j+2 < n && s[j] == s[j+2] {
				j++
			}
			tmp := sum ^ grundy[j-i+1]
			for k := 0; k < j-i+1; k++ {
				l := max(0, k-1)
				r := max(0, j-i+1-k-2)
				tmp2 := tmp ^ grundy[l] ^ grundy[r]
				if tmp2 == 0 {
					return i + k + 1
				}
			}
			i = j + 1
		}
	}
	return 1
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var w, m int
	fmt.Fscan(reader, &w, &m)
	res := solve(w, m)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(w int, m int) bool {
	if w == 2 {
		return true
	}
	var digits []int
	for m > 0 {
		digits = append(digits, m%w)
		m /= w
	}
	for i := 0; i < len(digits); i++ {
		if digits[i] <= 1 {
			continue
		}
		if i == 101 {
			return false
		}
		v := digits[i]
		// v > 1 and v < w
		// 这里要减去 pow(w, i)
		v -= w
		if v != -1 {
			return false
		}
		j := i + 1
		for j < len(digits) && digits[j] == w-1 {
			digits[j] = 0
			j++
		}
		if j < len(digits) {
			digits[j]++
		} else {
			digits = append(digits, 1)
		}
	}
	return true
}

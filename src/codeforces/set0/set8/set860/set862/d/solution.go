package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	ask := func(s string) int {
		fmt.Println("? " + s)
		var res int
		fmt.Fscan(reader, &res)
		return res
	}
	res := solve(n, ask)
	fmt.Printf("! %d %d\n", res[0], res[1])
}

func solve(n int, ask func(s string) int) []int {
	buf := make([]byte, n)
	for i := range n {
		buf[i] = '0'
	}

	ans0 := ask(string(buf))
	buf[0] = '1'
	ans1 := ask(string(buf))

	res := []int{-1, -1}

	// ans0 是1的个数
	if ans0 > ans1 {
		res[1] = 1
	} else {
		res[0] = 1
	}

	buf[0] = '0'

	if res[0] == -1 {
		res[0] = find0(n-ans0, buf, ask)
	} else {
		res[1] = find1(ans0, buf, ask)
	}

	return res
}

// binarySearch finds the position of the target character using binary search
func binarySearch(cnt int, buf []byte, ask func(s string) int, targetChar, resetChar byte) int {
	n := len(buf)

	// Initialize buffer with reset character
	for i := range n {
		buf[i] = resetChar
	}

	l, r := 0, n
	for l < r {
		mid := (l + r) >> 1

		// Set target character from mid to end
		for i := mid; i < n; i++ {
			buf[i] = targetChar
		}

		ans := ask(string(buf))

		if ans == cnt+(n-mid) {
			// Target characters are all in the first half
			r = mid
		} else {
			// Some target characters are in the second half
			l = mid + 1
		}

		// Reset to original character
		for i := mid; i < n; i++ {
			buf[i] = resetChar
		}
	}

	return r
}

func find1(cnt int, buf []byte, ask func(s string) int) int {
	// 现在需要查出1的位置, cnt表示1的个数
	return binarySearch(cnt, buf, ask, '1', '0')
}

func find0(cnt int, buf []byte, ask func(s string) int) int {
	// 现在需要查出0的位置, cnt表示0的个数
	return binarySearch(cnt, buf, ask, '0', '1')
}

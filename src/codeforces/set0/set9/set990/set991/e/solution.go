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
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	freq := make([]int, 10)
	for i := n; i > 0; i /= 10 {
		freq[i%10]++
	}
	var nums []int
	for i := range 10 {
		if freq[i] > 0 {
			nums = append(nums, i)
		}
	}

	buf := make([]int, 10)

	var res int

	C := make([][]int, 20)
	for i := range 20 {
		C[i] = make([]int, 20)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	calc := func(m int) int {
		ans := 1
		for _, i := range nums {
			ans *= C[m][buf[i]]
			m -= buf[i]
		}
		return ans
	}

	// 一共m位
	play := func(m int) {
		ans := calc(m)
		if buf[0] > 0 {
			buf[0]--
			ans -= calc(m - 1)
			buf[0]++
		}

		res += ans
	}

	var f func(i int, m int)
	f = func(i int, m int) {
		if i == len(nums) {
			play(m)
			return
		}
		x := nums[i]
		for range freq[x] {
			buf[x]++
			f(i+1, m+buf[x])
		}
		buf[x] = 0
	}

	f(0, 0)

	return res
}

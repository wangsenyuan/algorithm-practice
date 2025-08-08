package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	freq := make([]int, 5)
	var sum int
	for _, v := range a {
		freq[v]++
		sum += v
	}
	if sum <= 2 || sum == 5 {
		return -1
	}
	// 先合并1和2
	var res int

	x := min(freq[1], freq[2])

	res += x

	freq[1] -= x
	freq[2] -= x
	freq[3] += x

	if freq[1] > 0 {
		// 尽量3个一组
		w := freq[1] / 3
		res += w * 2
		freq[1] -= w * 3
		freq[3] += w
		freq[1] %= 3
		if freq[1] == 1 && freq[3] == 0 {
			res += 2
		} else {
			res += freq[1]
		}
	} else if freq[2] > 0 {
		w := freq[2] / 3
		res += w * 2
		freq[2] %= 3
		freq[3] += w * 2
		switch freq[2] {
		case 2:
			res += 2
		case 1:
			if freq[4] > 0 {
				res++
			} else {
				res += 2
			}
		}
	}

	return res
}

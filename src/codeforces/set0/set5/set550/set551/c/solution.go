package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, a)
}

func solve(m int, a []int) int {
	n := len(a)

	arr := make([]int, n)

	check := func(x int) bool {
		var cnt int
		copy(arr, a)
		for i := 0; i < n; {
			// 现在分配一个新人，在x时间内处理
			// 移动到这里需要
			for i < n && arr[i] == 0 {
				i++
			}
			if i == n {
				break
			}

			cnt++

			t := i + 1
			if t >= x {
				// 太远了，到达这里，就花费超过了x
				return false
			}
			for i < n && t+arr[i] <= x {
				t += arr[i]
				if i < n-1 && t < x {
					// 移动到下一个位置
					t++
				}
				i++
			}

			if i < n && t < x {
				// 再remove一部分
				arr[i] -= x - t
			}
		}

		return cnt <= m
	}

	var sum int
	for _, v := range a {
		sum += v
	}
	sum += n

	return sort.Search(sum, check)
}

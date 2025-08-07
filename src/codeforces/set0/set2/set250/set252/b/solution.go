package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	fmt.Println(res[0], res[1])
}

func process(reader *bufio.Reader) (a []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

func solve(a []int) []int {
	n := len(a)

	if n <= 2 {
		return nil
	}

	// n >= 3
	// 连续相同的去掉
	var arr []int
	var cnt int
	for i := 0; i < n; i++ {
		if len(arr) == 0 || a[arr[len(arr)-1]] != a[i] {
			if cnt > 1 {
				return []int{i, i + 1}
			}
			arr = append(arr, i)
			cnt = 1
		} else {
			cnt++
			// cnt > 1
			if len(arr) > 1 {
				j1 := arr[len(arr)-2]
				j2 := arr[len(arr)-1]
				return []int{j1 + 1, j2 + 1}
			}
		}
	}

	if len(arr) <= 2 {
		return nil
	}

	// 1 2 3， 或者 3, 2, 1, 或者 1 3 2, 1, 2, 1, 2
	// 如果 arr是sorted的，那么选择靠近的两个位置，就可以了
	// 因为不是sorted， 必然存在 x.... y.... z满足 y > max(x, z) 或者 y < min(x, z)存在
	// 确定y是一个峰顶，
	for i := 1; i+1 < len(arr); i++ {
		x := arr[i-1]
		y := arr[i]
		z := arr[i+1]
		if a[x] < a[y] && a[y] < a[z] || a[x] > a[y] && a[y] > a[z] {
			return []int{y + 1, z + 1}
		}
		if a[x] < a[y] && a[y] > a[z] || a[x] > a[y] && a[y] < a[z] {
			if a[x] != a[z] {
				return []int{x + 1, z + 1}
			}
			// a[x] == a[z]
			if i >= 2 {
				// a[i-2] != x, 且肯定不满足3个连续增加的条件， a[i-2] > x
				if a[arr[i-2]] != a[y] {
					return []int{arr[i-2] + 1, y + 1}
				} else {
					return []int{arr[i-2] + 1, z + 1}
				}
			}
		}
	}

	return nil
}

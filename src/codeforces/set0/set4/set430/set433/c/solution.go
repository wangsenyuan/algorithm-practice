package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) int {
	if len(a) == 1 {
		return 0
	}

	var sum int

	// 先算一遍
	for i := 0; i+1 < len(a); i++ {
		sum += abs(a[i] - a[i+1])
	}

	res := sum

	m := len(a)

	// 假设 x 变成 y, 那么所有 a[i-1] - x 变成了 a[i-1] - y
	// 一下子就变难了么
	// 假设知道某个x的贡献, 多少个a[?] >= x, 多少a[?] < x
	// 如果要改变x，貌似最好是把它改变到中间去

	arr := make([][]int, n+1)

	for i := 0; i < m; i++ {
		x := a[i]
		if i > 0 && a[i-1] != x {
			arr[x] = append(arr[x], a[i-1])
		}
		if i+1 < m && a[i+1] != x {
			arr[x] = append(arr[x], a[i+1])
		}
	}

	for x := 1; x <= n; x++ {
		if len(arr[x]) == 0 {
			continue
		}
		slices.Sort(arr[x])
		k := len(arr[x]) / 2
		y := arr[x][k]
		var diff int
		for _, v := range arr[x] {
			diff -= abs(v - x)
			diff += abs(v - y)
		}
		res = min(res, sum+diff)
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}

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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ops := make([][]int, m)
	for i := range m {
		ops[i] = make([]int, 2)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1])
	}
	return solve(a, ops)
}

func solve(a []int, ops [][]int) []int {
	m := len(ops)
	stack := make([]int, m)
	var top int
	for i, cur := range ops {
		r := cur[1]
		for top > 0 && ops[stack[top-1]][1] <= r {
			top--
		}

		stack[top] = i
		top++
	}

	// n := len(a)

	// 后面的不处理
	r0 := ops[stack[0]][1]
	arr := slices.Clone(a[:r0])
	slices.Sort(arr)

	for i := 0; i < top; i++ {
		j := stack[i]
		r := ops[j][1] - 1
		var r1 = -1
		if i < top-1 {
			r1 = ops[stack[i+1]][1] - 1
		}

		t := ops[j][0]

		if t == 1 {
			// 升序，所以最后的是最大的数
			for u := r; u > r1; u-- {
				a[u] = arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			}
		} else {
			// 降序,所以最后的数是最小的数
			for u := r; u > r1; u-- {
				a[u] = arr[0]
				arr = arr[1:]
			}
		}
	}

	return a
}

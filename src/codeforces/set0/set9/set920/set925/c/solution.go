package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (b []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	b = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	res = solve(b)
	return
}

type pair struct {
	first  int
	second int
}

func solve(b []int) []int {

	// 已经处理掉的数字
	var res []int
	// 还没有被处理的数字
	nums := b

	n := len(nums)

	for d := 60; d >= 0 && len(res) < n; d-- {
		var xs []int
		var other []int
		for _, v := range nums {
			if (v>>d)&1 == 1 {
				xs = append(xs, v)
			} else {
				other = append(other, v)
			}
		}
		if len(xs) == 0 {
			continue
		}
		var next []int
		next = append(next, xs[0])
		xs = xs[1:]

		for _, v := range res {
			next = append(next, v)
			if (v>>d)&1 == 1 && len(xs) > 0 {
				next = append(next, xs[0])
				xs = xs[1:]
			}
		}

		if len(xs) > 0 {
			return nil
		}
		res = next
		nums = other
	}

	return res
}

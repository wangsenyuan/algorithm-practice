package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	r := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &r[i])
	}
	return solve(x, r)
}

type circle struct {
	x int
	r int
}

func cmp(a, b circle) int {
	la := a.x - a.r
	lb := b.x - b.r
	if la != lb {
		return la - lb
	}
	// 大的排在前面
	return (b.x + b.r) - (a.x + a.r)
}

const inf = 1 << 60

func solve(x []int, r []int) int {
	my := make(map[int]int)
	for i := range x {
		a, b := x[i]-r[i], x[i]+r[i]
		for j := a; j <= b; j++ {
			my[j] = max(my[j], int(math.Sqrt(float64(r[i]*r[i]-(j-x[i])*(j-x[i])))))
		}
	}
	var res int
	for _, v := range my {
		res += 2*v + 1
	}
	return res
}

func solve1(x []int, r []int) int {
	n := len(x)

	// 对于任意一个点x, 计算最大的y， 满足 (x - x[i])^^2 + y^^2 <= r[i]^^2
	// 从左到右，任何一个cycle，如果x[i] + r[i] < x, 那么这个cycle就不再起作用了
	// 同时 x[i] - r[i] >= x 的部分，开始起作用
	// 在这些active的圆中，要怎么找到 (x - x[i])^^2 + y^^2 <= r[i]^^2 的i呢？
	// 怎么还要用到凸优化了？
	// 有一个猜想是，如果排除掉，那些完全被包含的圆，能包含一个点的圆的个数不会很多？

	arr := make([]circle, n)
	for i := range n {
		arr[i] = circle{x[i], r[i]}
	}

	slices.SortFunc(arr, cmp)

	// var active IntHeap

	var p int

	far := -inf

	for _, c := range arr {
		if far < c.x+c.r {
			arr[p] = c
			p++
			far = c.x + c.r
		}
	}

	arr = arr[:p]

	findMaxY := func(i int, x int) int {
		var res int
		for j := i; j < len(arr) && (arr[j].x-x) <= arr[j].r; j++ {
			// dx * dx + y * y <= r * r
			dx := arr[j].x - x
			y := math.Sqrt(float64(arr[j].r*arr[j].r - dx*dx))
			res = max(res, int(y))
		}
		for j := i - 1; j >= 0 && (x-arr[j].x) <= arr[j].r; j-- {
			dx := x - arr[j].x
			y := math.Sqrt(float64(arr[j].r*arr[j].r - dx*dx))
			res = max(res, int(y))
		}
		return res
	}

	var res int

	prev := -inf

	for i, c := range arr {
		x := max(prev+1, c.x-c.r)
		for x <= c.x+c.r {
			y := findMaxY(i, x)
			res += 2*y + 1
			x++
		}
		prev = max(prev, c.x+c.r)
	}

	return res
}

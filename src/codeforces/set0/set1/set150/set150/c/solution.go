package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.11f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, m, c int
	fmt.Fscan(reader, &n, &m, &c)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	p := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p[i])
	}
	passengers := make([][]int, m)
	for i := range m {
		passengers[i] = make([]int, 2)
		fmt.Fscan(reader, &passengers[i][0], &passengers[i][1])

	}

	return solve(c, x, p, passengers)
}

func solve(c int, x []int, p []int, passengers [][]int) float64 {
	n := len(x)

	merge := func(a [4]int, b [4]int) [4]int {
		var res [4]int
		res[0] = a[0] + b[0]
		res[1] = max(a[1], a[0]+b[1])
		res[2] = max(b[2], a[2]+b[0])
		res[3] = max(a[3], b[3], a[2]+b[1])
		return res
	}

	arr := make([][4]int, 4*n)

	var f func(i int, l int, r int)

	f = func(i int, l int, r int) {
		if l == r {
			arr[i][0] = (x[l+1]-x[l])*100 - 2*c*p[l]
			arr[i][1] = arr[i][0]
			arr[i][2] = arr[i][0]
			arr[i][3] = arr[i][0]
			return
		}
		mid := (l + r) / 2
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		arr[i] = merge(arr[i*2+1], arr[i*2+2])
	}

	f(0, 0, n-2)

	var g func(i int, l int, r int, L int, R int) [4]int

	g = func(i int, l int, r int, L int, R int) [4]int {
		if l == L && r == R {
			return arr[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return g(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return g(i*2+2, mid+1, r, L, R)
		}
		a := g(i*2+1, l, mid, L, mid)
		b := g(i*2+2, mid+1, r, mid+1, R)
		return merge(a, b)
	}

	var sum int

	// 每个人不买票的区间是不同的
	for _, cur := range passengers {
		l, r := cur[0]-1, cur[1]-1
		cur := g(0, 0, n-2, l, r-1)
		sum += max(0, cur[3])
	}

	return float64(sum) / 200
}

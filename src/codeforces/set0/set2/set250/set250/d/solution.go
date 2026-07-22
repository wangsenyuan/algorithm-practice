package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans, _, _, _, _, _ := drive(reader)
	fmt.Println(ans[0], ans[1])
}

func drive(reader *bufio.Reader) (res []int, a int, b int, ya []int, yb []int, l []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m, &a, &b)
	ya = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &ya[i])
	}
	yb = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &yb[i])
	}
	l = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &l[i])
	}
	res = solve(a, b, ya, yb, l)
	return
}

func solve(a, b int, A, B, L []int) []int {

	type data struct {
		id  int
		val int
	}

	arr := make([]data, len(A))

	for i, v := range A {
		arr[i] = data{id: i, val: v}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return a.val - b.val
	})

	play := func(y0 int, y1 int, l int) float64 {
		d1 := math.Sqrt(float64(a)*float64(a) + float64(y0)*float64(y0))
		d2 := math.Sqrt(float64(b-a)*float64(b-a) + float64(y1-y0)*float64(y1-y0))

		return d1 + d2 + float64(l)
	}

	var best float64 = math.MaxFloat64

	ans := []int{1, 1}

	for j, yb := range B {
		// 连接 O- (b, yb), 和 x = a的交点
		// y = k * x
		k := float64(yb) / float64(b)
		ya := k * float64(a)
		// 找这个点附近的i

		i := sort.Search(len(arr), func(i int) bool {
			return float64(arr[i].val) >= ya
		})

		if i < len(arr) {
			dist := play(arr[i].val, yb, L[j])
			if dist < best {
				best = dist
				ans = []int{arr[i].id + 1, j + 1}
			}
		}

		if i > 0 {
			dist := play(arr[i-1].val, yb, L[j])
			if dist < best {
				best = dist
				ans = []int{arr[i-1].id + 1, j + 1}
			}
		}
	}

	return ans
}

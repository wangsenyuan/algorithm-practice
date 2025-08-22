package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, z, k int
	fmt.Scan(&x, &y, &z, &k)
	res := solve(x, y, z, k)
	fmt.Println(res)
}

func solve(x int, y int, z int, k int) int {
	arr := []int{x, y, z}
	sort.Ints(arr)
	x, y, z  = arr[0], arr[1], arr[2]
	// 假设固定x平面的切的次数 <= min(x, k)
	// 那么剩余的次数，用于y和z平面的切割
	// 那么为了在一个平面得到足够多的，应该是尽量平分
	// 所以大体是 (a + 1) *( (k - a) / 2 + 1) * ((k - a + 1) + 1)/ 2
	// 其中 a <= min(x, k)
	// a * (k - a) / 2 * (k - a + 1) / 2
	// 差不多就是在 k / 3， 然后在一个范围内去检验
	l := min(max(k/3-10, 0), x-1)
	r := min(k/3+10, k, x-1)
	var best int
	for i := l; i <= r; i++ {
		// k - i - j >= 0, k - i - j < z
		j := min((k-i)/2, y-1)
		// j >= k - i - z
		u := min(k-i-j, z-1)

		tmp := (i + 1) * (j + 1) * (u + 1)
		best = max(best, tmp)
	}
	return best
}

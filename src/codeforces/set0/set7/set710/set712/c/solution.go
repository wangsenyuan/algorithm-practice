package main

import (
	"fmt"
	"slices"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	res := solve(x, y)
	fmt.Println(res)
}

func solve(x int, y int) int {
	arr := []int{y, y, y}

	var res int
	for {
		slices.Sort(arr)
		if arr[0] == x {
			break
		}
		b, c := arr[1], arr[2]
		d := min(x, b+c-1)
		arr[0] = d
		res++
	}
	return res
}

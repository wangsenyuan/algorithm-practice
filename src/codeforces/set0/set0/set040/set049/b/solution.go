package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res := solve(a, b)
	fmt.Println(res)
}

func solve(a int, b int) int {
	da := fmt.Sprintf("%d", a)
	db := fmt.Sprintf("%d", b)
	var x int
	for i := range len(da) {
		x = max(x, int(da[i]-'0'))
	}
	for i := range len(db) {
		x = max(x, int(db[i]-'0'))
	}
	x++
	// base is x
	i := len(da) - 1
	j := len(db) - 1
	var res int
	var carry int
	for i >= 0 || j >= 0 {
		var u, v int
		if i >= 0 {
			u = int(da[i] - '0')
			i--
		}
		if j >= 0 {
			v = int(db[j] - '0')
			j--
		}
		res++
		sum := (u + v + carry)
		carry = sum / x
	}
	if carry > 0 {
		res++
	}
	return res
}

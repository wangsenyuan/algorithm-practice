package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	user := make([]int, 3)
	for i := range 3 {
		fmt.Fscan(reader, &user[i])
	}
	top := make([]int, 3)
	for i := range 3 {
		fmt.Fscan(reader, &top[i])
	}
	a := make([]int, 6)
	for i := range 6 {
		fmt.Fscan(reader, &a[i])
	}
	return solve(user, top, a)
}

func solve(user []int, top []int, a []int) int {
	var res int

	// z方向
	if user[2] > top[2] {
		res += a[3]
	} else if user[2] < 0 {
		res += a[2]
	}
	// y 方向
	if user[1] > top[1] {
		res += a[1]
	} else if user[1] < 0 {
		res += a[0]
	}
	// x 方向
	if user[0] > top[0] {
		res += a[5]
	} else if user[0] < 0 {
		res += a[4]
	}

	return res
}

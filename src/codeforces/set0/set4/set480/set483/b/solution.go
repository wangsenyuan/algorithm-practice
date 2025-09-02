package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var cnt1, cnt2, x, y int
	fmt.Fscan(reader, &cnt1, &cnt2, &x, &y)
	return solve(cnt1, cnt2, x, y)
}

func solve(cnt1 int, cnt2 int, x int, y int) int {

	check := func(v int) bool {
		// 在 1...v中间，
		// 送个A的至少有cnt1个，且这cnt1个数字，不能整除x
		// 送给B的至少有cnt2个，且这个cnt2个数字，不能整除y
		// 所以，能够整除x * y 的要全部删除掉
		// 然后所有能够整除x的数，要全部给B，能够整除y的数，全部给A;
		// 剩下的数，看情况分配
		// 考虑 (x = 2, y = 3), v = 4, cnt1 = 3, cnt2 = 1 的情况
		// num_xy = 0, num_x = 2, num_y = 1, other = 4 - 3 = 1 (也就是1）
		num_xy := v / (x * y)
		num_x := v/x - num_xy
		num_y := v/y - num_xy
		other := v - num_x - num_y - num_xy
		return num_y+other >= cnt1 && num_x+other >= cnt2 && num_x+num_y+other >= cnt1+cnt2
	}

	return sort.Search(1e10, check)
}

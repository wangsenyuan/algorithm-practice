package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_, ok, res := drive(reader)
	if !ok {
		fmt.Fprintln(writer, "NO")
		return
	}
	fmt.Fprintln(writer, "YES")
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) (a []int, ok bool, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ok, res = solve(a)
	return
}
func solve(nums []int) (bool, []int) {
	if len(nums) == 0 {
		return true, []int{1, 1, 3, 3}
	}
	sort.Ints(nums)
	if len(nums) == 4 {
		return check(nums), nil
	}
	if len(nums) == 1 {
		return true, solve1(nums[0])
	}
	if len(nums) == 2 {
		return solve2(nums[0], nums[1])
	}
	return solve3(nums[0], nums[1], nums[2])
}

func solve1(a int) []int {
	// b, c, d
	return []int{a, 3 * a, 3 * a}
}

func solve2(x int, y int) (bool, []int) {
	if y%x == 0 && y/x == 3 {
		return true, []int{2 * x, 2 * x}
	}
	// 如果 x = a
	if y <= 3*x {
		ok, res := solve3(x, y, 3*x)
		if ok {
			return true, append(res, 3*x)
		}
	}
	// x = b
	if x&1 == y&1 {
		a := (x + y) / 4
		if a <= x && 4*a == x+y && y <= 3*a {
			return true, []int{a, 3 * a}
		}
	}
	//  and y == d
	if y%3 == 0 {
		a := y / 3
		b := 4*a - x
		// b 可以 >= x
		if a <= b && b <= y {
			return true, []int{a, b}
		}
	}
	return false, nil
}

func solve3(x int, y int, z int) (bool, []int) {
	// x <= y and y <= z
	if x*3 == z {
		b := 4*x - y
		if b >= x && b <= z {
			return true, []int{b}
		}
	}
	if x&1 == y&1 {
		a := (x + y) / 4
		if a*3 == z {
			return true, []int{a}
		}
	}

	if x*3 >= z && 4*x == y+z {
		return true, []int{3 * x}
	}

	return false, nil
}

func check(nums []int) bool {
	a := nums[0]
	if 3*a != nums[3] {
		return false
	}
	if 4*a != nums[1]+nums[2] {
		return false
	}
	return (nums[1]+nums[2])%2 == 0
}

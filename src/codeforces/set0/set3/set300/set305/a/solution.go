package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var k int
	fmt.Fscan(reader, &k)
	a := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	// 理解错了，每一位上，只能有一个0
	seen := make([]bool, 101)
	for _, v := range a {
		seen[v] = true
	}

	var res []int
	// 这两个和其他的没有冲突
	if seen[100] {
		res = append(res, 100)
	}

	if seen[0] {
		res = append(res, 0)
	}

	for x := 10; x < 100; x += 10 {
		if seen[x] {
			res = append(res, x)
			break
		}
	}

	canAdd := func(x int, y int) bool {
		for x > 0 && y > 0 {
			if x%10 > 0 && y%10 > 0 {
				return false
			}
			x /= 10
			y /= 10
		}
		return true
	}

	// 然后再找其他的，如果和已有的不冲突，就添加进去，否则就不加
	for _, v := range a {
		ok := true
		for _, x := range res {
			if v == x || !canAdd(v, x) {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, v)
		}
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		if len(res) == 0 {
			fmt.Fprintln(writer, "NO")
			continue
		}
		fmt.Fprintln(writer, "YES")
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	slices.Sort(a)
	// 逆时针4个点
	// x1 = x4, y1 <= y4
	// x1 <= x2, y1 = y2
	// x2 = x3, y2 <= y3
	// x4 <= x3, y4 == y3
	// 同一个数只能被用一次, 但是它可能出现了多次
	// 固定x1 和 x2(各有两个)
	// 然后在剩余的数中找出(y1, y3)
	// 假设 w出现了至少两次, 且它是最小的那个数, 那么要么它是x1, 要么它是y1
	// 记录那些出现了2次的
	var nums []int
	for i := 0; i < len(a); {
		j := i
		for i < len(a) && a[i] == a[j] {
			i++
			if (i-j)%2 == 0 {
				nums = append(nums, a[j])
			}
		}
	}
	// 必须要有8个数
	if len(nums) < 4 {
		return nil
	}
	k := len(nums)
	x1 := nums[0]
	y1 := nums[1]
	u := nums[k-1]
	v := nums[k-2]
	var x2, y3 int
	if (u-x1)*(v-y1) >= (v-x1)*(u-y1) {
		x2 = u
		y3 = v
	} else {
		x2 = v
		y3 = u
	}

	return []int{x1, y1, x2, y1, x2, y3, x1, y3}
}

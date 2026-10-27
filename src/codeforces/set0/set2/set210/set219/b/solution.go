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
	var p, d int
	fmt.Fscan(reader, &p, &d)
	return solve(p, d)
}

func solve(p int, d int) int {
	// 使用尽可能多的9
	best := p

	var nums []int

	for i := p; i > 0; i /= 10 {
		nums = append(nums, i%10)
	}

	bases := make([]int, len(nums))
	bases[0] = 1
	for i := 1; i < len(bases); i++ {
		bases[i] = bases[i-1] * 10
	}
	var cur int

	for i := range nums {
		// 把 [0...i]变成9
		cur = cur*10 + 9
		if cur > p || cur < 0 {
			break
		}
		// 然后在前面放置足够的数字, 保证cur < p, 且cur最大
		var tmp int
		for j := len(nums) - 1; j > i; j-- {
			v := nums[j]
			w := min(v, (p-tmp-cur)/bases[j])
			tmp += w * bases[j]
		}
		//9 越多越好
		if p-tmp-cur <= d {
			best = tmp + cur
		}
	}

	return best
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	n := len(a)

	if k == 1 {
		// 这个居然不是, k = 1的时候，
		// 1 2 3
		// 1 2 3
		var res int
		for i := range n {
			prod := 1
			var sum int
			for j := i; j < min(n, i+30); j++ {
				prod *= a[j]
				sum += a[j]
				if prod == sum {
					res++
				}
			}
		}

		return res
	}

	L := make([]int, n)

	for i := range n {
		if a[i] == 1 {
			L[i] = i - 1
			if i > 0 {
				L[i] = L[i-1]
			}
		} else {
			L[i] = i
		}
	}
	// L[i]当前1的上一个位置1的位置
	var res int
	var all int
	for r := range n {
		all += a[r]
		prod := 1
		sum := 0
		var cnt int
		// 如果 all * k < prod 了，那么再增加序列一个数，prod只会更大
		for l := r; l >= 0 && cnt < 100 && all*k >= prod; {
			if a[l] == 1 {
				// 第一个遇到1
				j := L[l]
				if prod%k == 0 {
					x := max(0, prod/k-sum)
					if x > 0 && x <= l-j {
						res++
					}
				}
				// prod no change
				sum += l - j
				l = j
			} else {
				prod *= a[l]
				sum += a[l]
				if prod == sum*k {
					res++
				}
				l--
			}
			cnt++
		}
	}

	return res
}

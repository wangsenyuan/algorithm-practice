package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (p int, x int, y int, a []int, res []int) {
	var n, k int
	fmt.Fscan(reader, &n, &k, &p, &x, &y)
	a = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(n, p, x, y, a)
	return
}

func solve(n int, p int, x int, y int, a []int) []int {
	k := len(a)
	var sum int
	var cnt int
	for _, v := range a {
		if v >= y {
			cnt++
		}
		sum += v
	}
	sum += n - k
	res := make([]int, n-k)
	for i := range res {
		res[i] = 1
	}
	// sum <= x
	// 如果 cnt < (n + 1) / 2
	// 那么需要补充 (n + 1) / 2 - cnt 个人
	if cnt < (n+1)/2 {
		if (n+1)/2-cnt > n-k {
			return nil
		}
		// 这些人需要获得 y - 1 的分数
		sum += (y - 1) * ((n+1)/2 - cnt)
		for i := 0; i < (n+1)/2-cnt; i++ {
			res[i] += y - 1
		}
	}
	if sum > x {
		// no answer
		return nil
	}
	return res
}

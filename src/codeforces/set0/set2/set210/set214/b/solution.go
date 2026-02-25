package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const inf = 1 << 59

func solve(a []int) string {
	// 比如以0结尾，且sum % 3 = 0
	slices.Sort(a)
	if a[0] != 0 {
		return "-1"
	}
	slices.Reverse(a)
	// 最后一个是0，然后把其他的数按照倒序排列
	// 如果不能整除3
	// dp[1] 最小的数是多少，dp[2]最小的数是多少？删掉这几个数就可以了
	dp := make([]int, 3)
	ndp := make([]int, 3)

	for i := range 3 {
		dp[i] = inf
		ndp[i] = inf
	}
	dp[0] = 0
	n := len(a)
	var sum int
	for i := range n - 1 {
		v := a[i]

		for j := range 3 {
			ndp[(j+v)%3] = min(ndp[(j+v)%3], dp[j]*10+v)
		}
		for j := range 3 {
			dp[j] = min(dp[j], ndp[j])
			ndp[j] = inf
		}
		sum = (sum + v) % 3
	}
	var buf bytes.Buffer

	if sum == 0 {
		for _, x := range a {
			fmt.Fprintf(&buf, "%d", x)
		}
	} else {
		if dp[sum] == inf {
			return "-1"
		}
		// 那么需要删掉dp[rem]对应的数字
		marked := make([]bool, n)
		for i := n - 2; i >= 0 && dp[sum] > 0; i-- {
			x := dp[sum] % 10
			if a[i] == x {
				marked[i] = true
				dp[sum] /= 10
			}
		}

		for i := range n {
			if !marked[i] {
				fmt.Fprintf(&buf, "%d", a[i])
			}
		}
	}

	res := buf.String()
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return res
}

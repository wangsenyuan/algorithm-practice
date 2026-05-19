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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		a[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int) int {
	n := len(a)
	var m int
	for _, v := range a {
		m = max(m, len(v))
	}
	// 当前的a[?]的mex的sum
	var sum int
	// dp[i] 表示，如果使的a[i]mex变成i，它的贡献是多少
	dp := make([]int, m+3)
	mex1 := make([]int, n)
	for r, v := range a {
		sort.Ints(v)
		var mex int

		var i int
		for i < len(v) {
			if v[i] > mex {
				// mex will not change
				break
			}
			if v[i] == mex {
				mex++
			}
			i++
		}
		sum += mex
		mex1[r] = mex
		mex++
		for i < len(v) {
			if v[i] > mex {
				break
			}
			// v[i] < mex may holds
			if v[i] == mex {
				mex++
			}
			i++
		}
		dp[mex1[r]] += mex - mex1[r]
	}

	var ans int

	for i := range n {
		// 还必须知道v得freq
		for j, v := range a[i] {
			if v > mex1[i] {
				// 不用增加freq, 但是要减去当前i的额外的贡献
				tmp := sum * (n - 1)
				if v < len(dp) {
					tmp += dp[v]
				}
				ans += tmp
			} else {
				// if v < mex1[i]

				if j > 0 && a[i][j-1] == v || j+1 < len(a[i]) && a[i][j+1] == v {
					// 移除v，不会改变a[i]
					ans += sum*(n-1) + dp[v]
				} else {
					// v变成了新的mex
					tmp := (sum-mex1[i]+v)*(n-1) + dp[v]
					ans += tmp
				}
			}
			// v 不可能等于 mex1[i] 因为 mex1[i]是a[i]不存在的
		}
	}

	return ans
}

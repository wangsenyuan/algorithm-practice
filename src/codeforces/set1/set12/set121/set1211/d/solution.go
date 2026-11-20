package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, a, b, k int
	fmt.Fscan(reader, &n, &a, &b, &k)
	r := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &r[i])
	}
	return solve(a, b, k, r)
}

func solve(a int, b int, k int, r []int) int {
	slices.Sort(r)

	todo := make(map[int][]int)

	n := len(r)
	for i := 0; i < n; {
		j := i
		for i < n && r[i] == r[j] {
			i++
		}

		if _, ok := todo[r[j]/k]; !ok || r[j]%k != 0 {
			todo[r[j]] = []int{i - j}
		} else {
			todo[r[j]] = append(todo[r[j]/k], i-j)
			delete(todo, r[j]/k)
		}
	}

	var ans int

	for _, v := range todo {
		ans += play(v, a, b)
	}

	return ans
}

func play(arr []int, a int, b int) int {
	// 假设w个人, 那么前面分配x个人(和前面配对), y个人(和后面配对)
	// x * b = y + a
	// x + y <= w
	// len(arr) < 20
	// dp[i] 表示剩余i个人可以和当前配对时的最优解
	w := slices.Max(arr)
	dp := make([]int, w+1)
	ndp := make([]int, w+1)

	for i := 1; i < len(arr); i++ {
		v := arr[i]
		for x := 0; x <= v; x += b {
			// 使用x个人和前面配对, 那么前面就需要剩余x1个人
			x1 := x / b * a
			if x1 > arr[i-1] {
				break
			}
			// x1 <= arr[i-1]
			ndp[v-x] = max(ndp[v-x], dp[x1]+x/b)
		}
		for x := w; x >= 0; x-- {
			if x+1 <= w {
				// 如果剩余10个人, 能组成5对, 那么剩余1个人的时候, 最好也是5对?
				ndp[x] = max(ndp[x], ndp[x+1])
			}
			dp[x] = ndp[x]
		}
		clear(ndp)
	}

	return dp[0]
}

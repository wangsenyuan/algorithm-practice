package main

import "fmt"

func main() {
	var k int
	fmt.Scanf("%d\n", &k)
	res := solve(k)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

const X = 1e6
const N = 1000

func solve(k int) []int {
	// (x - w) * 2 - x = k
	// 2 * x - 2 * w - x = k
	// x - 2 * w = k
	// 如果k = 1
	// x = 3, w = 1
	// -1, 3
	// 可以让w = 1
	if k+2 <= X {
		// 2 * (k + 1) - (k + 2) = k
		return []int{-1, k + 2}
	}
	// a[0] = -1
	// a[1] >= 2
	// S - n + 1 = k
	// s - n = k - 1
	n := 2 * N
	S := k - 1 + n
	arr := make([]int, n)
	arr[0] = -1
	arr[1] = 2
	S--
	S -= n - 2
	for i := 2; i < n; i++ {
		arr[i] = 1
		x := min(S, X-1)
		arr[i] += x
		S -= x
	}
	return arr
}

func abs(num int) int {
	return max(num, -num)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

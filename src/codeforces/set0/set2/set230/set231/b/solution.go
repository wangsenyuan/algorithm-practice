package main

import "fmt"

func main() {
	var n, d, l int
	fmt.Scanf("%d %d %d", &n, &d, &l)
	res := solve(n, d, l)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int, d int, l int) []int {
	// 奇数位全部是1，偶数位全部是l，这个时d最小
	if (n+1)/2-n/2*l > d {
		return nil
	}
	// 奇数位全部是l，偶数位全部是1，这个时候d最大
	if (n+1)/2*l-n/2 < d {
		return nil
	}
	// 那么都有ans
	res := make([]int, n)
	sum := make([]int, 2)
	for i := range n {
		// 假设这位上设为x, x >= 1 and x <= l
		if i%2 == 0 {
			// sum[0] + x - sum[1] + (n - i) / 2 * l - (n - i)/ 2 >= d
			// x >= d - sum[0] + sum[1] - (n - i) / 2 * l + (n - i)/ 2
			x := max(1, d-sum[0]+sum[1]-(n-i-1)/2*l+(n-i)/2)
			x = min(x, l)
			res[i] = x
			sum[0] += x
		} else {
			// sum[0] - sum[1] - x + (n - i + 1) / 2 - (n - i) / 2 >= d
			// x <= sum[0] - sum[1] + (n - i + 1) / 2 - (n - i) / 2 - d
			x := min(l, sum[0]-sum[1]+(n-i)/2-(n-i-1)/2-d)
			x = max(1, x)
			res[i] = x
			sum[1] += x
		}
	}

	return res
}

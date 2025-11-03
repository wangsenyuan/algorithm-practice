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
	a := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

const H = 30

func solve(a []int, k int) int {
	n := len(a) - 1
	ds := make([]int, n*2+H)
	copy(ds, a)
	for i := 0; i <= n; i++ {
		if ds[i] > 0 {
			ds[i+1] += ds[i] / 2
			ds[i] = ds[i] & 1
		} else if ds[i] < 0 {
			if ds[i]%2 != 0 {
				ds[i+1] += ds[i]/2 - 1
				ds[i] = 1
			} else {
				ds[i+1] += ds[i] / 2
				ds[i] = 0
			}
		}
	}

	st := n + 1
	for i := 0; i <= n; i++ {
		if ds[i] != 0 {
			st = i
			break
		}
	}

	var now int
	for i := n + 1; i > st; i-- {
		now = now*2 + ds[i]
		if abs(now) > 2*k {
			now = 2*k + 1
			break
		}
	}

	var ans int

	for i := st; i >= 0; i-- {
		now = now*2 + ds[i]
		if abs(now) > 2*k {
			break
		}
		if i <= n {
			diff := a[i] - now
			if abs(diff) <= k && (i != n || diff != 0) {
				ans++
			}
		}
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}

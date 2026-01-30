package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func solve(n int) []int {

	// 考虑1， 如果把1放在 1 p1, p2
	// p1 和 p2 互质的话，那么1就是一个bad的位置
	// 所以最好p1和p2不互质
	// 最好是 1 p1, p2, x, p3, p4, y
	// 但是如果p2和p3互质的话，(x，p3) 互质， 那么 p2就是一个bad的位置
	// 所以最好p2和p3不互质
	// 但是这样子，总会有一个数，比如这里的p4,它会成为一个bad
	// 除非p4和y不互质, 比如如果 p1 = 2, y = 3, 那么p4 = 6
	pos := make([]int, n+1)
	for i := range n + 1 {
		pos[i] = -1
	}
	marked := make([]bool, n+1)

	var special []int
	var res []int

	special = append(special, 1)

	var lastPrime int
	for i := 2; i <= n; i++ {
		if marked[i] {
			continue
		}
		// i是质数
		var arr []int
		for j := i; j <= n; j += i {
			if !marked[j] {
				arr = append(arr, j)
				marked[j] = true
			}
		}

		if len(arr) == 1 {
			special = append(special, i)
		} else {
			// len(arr) >= 2
			if len(res) > 0 {
				x := lastPrime * i
				if x <= n && pos[x] >= 0 && pos[x] < len(res)-1 {
					// 交换过来
					j := pos[x]
					res[j], res[len(res)-1] = res[len(res)-1], res[j]
					pos[x] = len(res) - 1
				}
			}
			for _, j := range arr {
				res = append(res, j)
				pos[j] = len(res) - 1
			}
			lastPrime = i
		}
	}

	var ans []int
	var i, j int
	for i < len(special) && j+1 < len(res) {
		if gcd(res[j], res[j+1]) == 1 {
			ans = append(ans, res[j])
			j++
			continue
		}
		ans = append(ans, special[i])
		i++
		ans = append(ans, res[j], res[j+1])
		j += 2
	}
	for j < len(res) {
		ans = append(ans, res[j])
		j++
	}
	for i < len(special) {
		ans = append(ans, special[i])
		i++
	}
	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

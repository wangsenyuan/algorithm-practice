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

const mod = 998244353

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res *= num
		res %= mod
	}
	return res
}

func solve(k int, a []int) int {
	n := len(a)
	// a -1, -1... a
	same := make([]int, n+1)
	diff := make([]int, n+1)
	same[0] = 0
	diff[0] = 1

	for l := 1; l <= n; l++ {
		if l&1 == 1 {
			// 中间选择两头相同的
			x := mul(same[l/2], same[l/2])
			// 中间选择和两头不相同的
			y := mul(k-1, diff[l/2], diff[l/2])
			same[l] = add(x, y)
			// 中间选择两头中的一个，那么一边相同，一边不同
			x = mul(2, same[l/2], diff[l/2])
			// 中间选择和两头都不一样的
			y = mul(k-2, diff[l/2], diff[l/2])
			diff[l] = add(x, y)
		} else {
			//a..-1.... -1, a
			// 最后一个-1肯定a不同，所以有k-1种选择，中间是diff[l-1]
			same[l] = mul(diff[l-1], k-1)
			// a...-1...-1,b
			// 如果-1和a相同，但肯定不能是b
			diff[l] = add(same[l-1], mul(diff[l-1], k-2))
		}
	}

	// 奇偶位可以分开处理
	play := func(arr []int) int {
		n := len(arr)
		res := 1
		for i := 0; i < n; {
			if arr[i] != -1 {
				i++
				continue
			}
			// arr[i] == -1
			j := i
			for i < n && arr[i] == -1 {
				i++
			}
			d := i - j
			if d == n {
				if n == 1 {
					return k
				}
				// 整个序列都是-1
				return add(mul(k, same[n-2]), mul(k*(k-1), diff[n-2]))
			}
			if j == 0 || i == n {
				// 两端没有数字
				// 和另外一端相同
				cur := add(same[d-1], mul(k-1, diff[d-1]))
				res = mul(res, cur)
			} else {
				a := arr[j-1]
				b := arr[i]
				if a == b {
					res = mul(res, same[d])
				} else {
					res = mul(res, diff[d])
				}
			}
		}

		return res
	}

	var odd []int
	var even []int
	for i := range n {
		if a[i] > -1 && i > 1 && a[i] == a[i-2] {
			return 0
		}
		if i&1 == 1 {
			odd = append(odd, a[i])
		} else {
			even = append(even, a[i])
		}
	}

	res := play(odd)
	res = mul(res, play(even))

	return res
}

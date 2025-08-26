package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

const X = 1000010

var F [X]int

func init() {
	var primes []int
	for i := 2; i < X; i++ {
		if F[i] == 0 {
			F[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= X {
				break
			}
			F[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
}

func drive(reader *bufio.Reader) int {
	var x, p, k int
	fmt.Fscan(reader, &x, &p, &k)
	return solve(x, p, k)
}

func solve(x int, p int, k int) int {
	if p == 1 {
		return x + k
	}

	var arr []int
	for i := p; i > 1; {
		j := F[i]
		arr = append(arr, j)

		for i%j == 0 {
			i /= j
		}
	}

	x++

	n := len(arr)
	// n <= 8
	N := 1 << n

	check := func(num int) bool {
		if num < x {
			return false
		}

		cnt := num - x + 1

		for state := N - 1; state > 0; state-- {
			prod := 1
			var w int
			for i := range n {
				if (state>>i)&1 == 1 {
					prod *= arr[i]
					w++
				}
			}
			// 在 x...num中间能够整除prod的数
			l := (x + prod - 1) / prod
			r := num / prod

			tmp := max(0, r-l+1)

			// dp[2] - dp[2 * 3] + dp[2 * 3 * 5] - dp[2 * 3 * 5 * 7]
			if w&1 == 1 {
				cnt -= tmp
			} else {
				cnt += tmp
			}
		}

		return cnt >= k
	}

	res := sort.Search(1e10, check)

	return res
}

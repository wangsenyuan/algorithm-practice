package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

const N = 1e5 + 10

var lpf [N]int

func init() {
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			for j := i; j < N; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func solve(n int) []int {
	// 如果i是质数, 且不存在 2 * i 的时候, p[i] = i (否则的话, gcd(p[i], i) = 1)
	// 然后其他的, 肯定可以得到合适的数字
	parts := make([][]int, n+1)

	p := make([]int, n+1)
	p[1] = 1

	for i := 2; i <= n; i++ {
		if lpf[i] == i {
			parts[i] = append(parts[i], i)
			if i*2 <= n {
				parts[i] = append(parts[i], i*2)
			}
		}
		// 其他的合数
		if lpf[i] != i && (i%2 != 0 || lpf[i/2] != i/2) {
			parts[lpf[i]] = append(parts[lpf[i]], i)
		}
	}

	for _, cur := range parts {
		if len(cur) > 0 {
			slices.Sort(cur)
			for i := range cur {
				p[cur[i]] = cur[(i+1)%len(cur)]
			}
		}
	}

	return p[1:]
}

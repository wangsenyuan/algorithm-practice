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

	res := drive(reader)
	for _, s := range res {
		fmt.Fprintln(writer, s)
	}
}

func drive(reader *bufio.Reader) []string {
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(x, y, a)
}

func solve(x int, y int, a []int) []string {
	cnt := make([]int, 2)

	var winner []string

	for len(winner) < x+y {
		if (cnt[0]+1)*y > (cnt[1]+1)*x {
			winner = append(winner, "Vova")
			cnt[1]++
		} else if (cnt[0]+1)*y < (cnt[1]+1)*x {
			winner = append(winner, "Vanya")
			cnt[0]++
		} else {
			winner = append(winner, "Both")
			winner = append(winner, "Both")
			cnt[0]++
			cnt[1]++
		}
	}

	ans := make([]string, len(a))

	for i, v := range a {
		ans[i] = winner[(v-1)%(x+y)]
	}

	return ans
}

func gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

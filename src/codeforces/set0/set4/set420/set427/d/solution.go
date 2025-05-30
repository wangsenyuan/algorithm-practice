package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	s1 := readString(reader)
	s2 := readString(reader)
	return solve(s1, s2)
}

func solve(s1 string, s2 string) int {
	n := len(s1)
	m := len(s2)
	dp := make([][]int, max(n, m)+1)
	for i := range dp {
		dp[i] = make([]int, max(n, m)+1)
	}
	compute := func(x string, y string) {
		for i := range dp {
			clear(dp[i])
		}
		for i := len(x) - 1; i >= 0; i-- {
			for j := len(y) - 1; j >= 0; j-- {
				if x[i] == y[j] {
					dp[i][j] = dp[i+1][j+1] + 1
				} else {
					dp[i][j] = 0
				}
			}
		}
	}
	compute(s1, s1)
	fp1 := make([]int, n)
	for i := range n {
		for j := range n {
			if i != j {
				fp1[i] = max(fp1[i], dp[i][j])
			}
		}
		fp1[i]++
	}
	compute(s2, s2)
	fp2 := make([]int, m)
	for i := range m {
		for j := range m {
			if i != j {
				fp2[i] = max(fp2[i], dp[i][j])
			}
		}
		fp2[i]++
	}

	best := -1
	compute(s1, s2)
	for i := range n {
		for j := range m {
			if dp[i][j] >= max(fp1[i], fp2[j]) {
				if best < 0 || best > max(fp1[i], fp2[j]) {
					best = max(fp1[i], fp2[j])
				}
			}
		}
	}
	return best
}

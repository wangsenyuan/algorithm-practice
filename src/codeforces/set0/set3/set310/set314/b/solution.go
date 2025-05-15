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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	b, d := readTwoNums(reader)
	a := readString(reader)
	c := readString(reader)
	return solve(b, d, a, c)
}

type state struct {
	i int
	j int
}

func solve(b int, d int, a string, c string) int {
	n := len(a)
	m := len(c)

	dp := make([]state, n)
	for i := range n {
		dp[i] = state{i: -1, j: -1}
	}

	find := func(i1 int, j1 int, i2 int, j2 int) int {
		before_cycle_c := i1 / m
		cycle_cnt_c := (i2 - i1) / m
		before_cycle_a := j1 / n
		cycle_cnt_a := (j2 - j1) / n

		// 循环了这么多次
		cnt := (b - before_cycle_a) / cycle_cnt_a
		if j2%n != 0 {
			cnt--
		}

		j2 = j1 + cnt*n*cycle_cnt_a

		cycle_cnt_c *= cnt

		i2 = i1 + cnt*m*cycle_cnt_c

		cycle_cnt_c += before_cycle_c

		for i, j := i2, j2; j < b*n; {
			for c[i%m] != a[j%n] {
				j++
			}
			if j/n == b {
				break
			}
			// 肯定能找到
			i++
			if i%m == 0 {
				cycle_cnt_c++
			}
			j++
		}

		return cycle_cnt_c / d
	}

	for i, j := 0, 0; ; i++ {
		k := j
		for c[i%m] != a[j%n] {
			j++
			if j-k >= n {
				return 0
			}
		}
		if dp[j%n].i > 0 && dp[j%n].i%m == i%m {
			return find(dp[j%n].i, dp[j%n].j, i, j)
		}
		if dp[j%n].i == -1 {
			dp[j%n] = state{i, j}
		}
		j++
		if j/n == b {
			return ((i + 1) / m) / d
		}
	}
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

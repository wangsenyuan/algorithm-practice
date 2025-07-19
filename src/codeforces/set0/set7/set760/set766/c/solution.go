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
	fmt.Println(res[0])
	fmt.Println(res[1])
	fmt.Println(res[2])
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	s := readString(reader)[:n]
	a := readNNums(reader, 26)
	return solve(s, a)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

// const inf = 1 << 30

func solve(s string, a []int) []int {
	n := len(s)

	// dp[i] = 到i为止，满足条件的分隔数
	// fp[i] = 到i为止，满足条件的最小分段数
	// gp[i] = 到i为止，满足条件的最长分段数
	dp := make([]int, n+1)
	dp[0] = 1
	fp := make([]int, n+1)
	gp := make([]int, n+1)

	for i := 0; i < n; i++ {
		fp[i+1] = n
		gp[i+1] = 0

		w := i + 1

		for j := i; j >= 0; j-- {
			x := int(s[j] - 'a')
			w = min(w, a[x])
			if w < i-j+1 {
				break
			}
			// 要把x的要求放上去
			dp[i+1] = add(dp[i+1], dp[j])
			fp[i+1] = min(fp[i+1], fp[j]+1)
			gp[i+1] = max(gp[i+1], gp[j], i-j+1)
		}
	}

	return []int{dp[n], gp[n], fp[n]}
}

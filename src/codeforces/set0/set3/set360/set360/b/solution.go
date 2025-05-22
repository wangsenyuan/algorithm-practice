package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, k)
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(a []int, k int) int {
	n := len(a)

	dp := make([]int, n)

	check := func(x int) bool {
		for j := 0; j < n; j++ {
			// 最差的情况是改变前面所有的数字
			dp[j] = j
			for i := j - 1; i >= 0; i-- {
				if abs(a[j]-a[i]) <= (j-i)*x {
					dp[j] = min(dp[j], dp[i]+j-i-1)
				}
			}
			if n-1-j+dp[j] <= k {
				return true
			}
		}
		return false
	}
	var r int
	for i := 0; i+1 < n; i++ {
		r = max(r, abs(a[i+1]-a[i]))
	}
	return sort.Search(r, check)
}

func abs(x int) int {
	return max(x, -x)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {
	n := len(a)
	dp := make([][]pair, n)

	checkAndAdd := func(i int, sum int, cnt int) {
		top := len(dp[i])
		if top == 0 || dp[i][top-1].second > cnt {
			dp[i] = append(dp[i], pair{sum, cnt})
		}
	}

	for i := 0; i < n; i++ {
		sum := a[i]
		// 这里sum是越来越大的
		for j := i - 1; j >= 0; j-- {
			pos := sort.Search(len(dp[j]), func(k int) bool {
				return dp[j][k].first > sum
			})
			pos--
			if pos >= 0 {
				checkAndAdd(i, sum, dp[j][pos].second+i-j-1)
			}
			sum += a[j]
		}
		// 把前面的都放起来
		checkAndAdd(i, sum, i)
	}
	return dp[n-1][len(dp[n-1])-1].second
}

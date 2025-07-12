package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1], res[2])
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

func process(reader *bufio.Reader) []int {
	nums := readNNums(reader, 5)
	return solve(nums[0], nums[1], nums[2], nums[3], nums[4])
}

func solve(n int, m int, k int, x int, y int) []int {
	if n == 1 {
		return solve1(m, k, x, y)
	}
	// 1.2...n-1,n,n-1....2
	round := (2*n - 2) * m
	a, b := k/round, k%round
	cnt := make([][]int, n)
	for i := range n {
		cnt[i] = make([]int, m)
	}
	for j := range m {
		cnt[0][j] = a
		cnt[n-1][j] = a
	}
	for i := 1; i < n-1; i++ {
		for j := range m {
			cnt[i][j] = 2 * a
		}
	}
	for i := 0; i < n && b > 0; i++ {
		for j := 0; j < m && b > 0; j++ {
			cnt[i][j]++
			b--
		}
	}
	for i := n - 2; i >= 0 && b > 0; i-- {
		for j := 0; j < m && b > 0; j++ {
			cnt[i][j]++
			b--
		}
	}
	ans := make([]int, 3)
	ans[1] = cnt[0][0]
	for i := range n {
		for j := range m {
			ans[0] = max(ans[0], cnt[i][j])
			ans[1] = min(ans[1], cnt[i][j])
		}
	}
	ans[2] = cnt[x-1][y-1]
	return ans
}

func solve1(m int, k int, x int, y int) []int {
	a, b := k/m, k%m
	res := []int{a, a, 0}
	if b > 0 {
		res[0]++
	}
	y--
	if y < b {
		res[2] = res[0]
	} else {
		res[2] = res[1]
	}
	return res
}

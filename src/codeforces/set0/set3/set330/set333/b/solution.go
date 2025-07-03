package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, m := readTwoNums(reader)
	banned := make([][]int, m)
	for i := range banned {
		banned[i] = readNNums(reader, 2)
	}
	return solve(n, banned)
}

func solve(n int, banned [][]int) int {
	blocked_row := make([]bool, n)
	blocked_col := make([]bool, n)

	for _, cur := range banned {
		r, c := cur[0]-1, cur[1]-1
		blocked_row[r] = true
		blocked_col[c] = true
	}

	var ans int

	for i, j := 1, n-2; i <= j; i, j = i+1, j-1 {
		if i == j {
			// 最多只能放置一个
			if !blocked_row[i] || !blocked_col[i] {
				ans++
			}

			continue
		}

		tmp := check(!blocked_col[i]) + check(!blocked_col[j]) + check(!blocked_row[i]) + check(!blocked_row[j])

		ans += tmp
	}

	return ans
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}

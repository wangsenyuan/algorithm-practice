package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res < 0 {
		fmt.Println("Poor Inna and pony!")
	} else {
		fmt.Println(res)
	}
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
	nums := readNNums(reader, 6)
	return solve(nums[0], nums[1], nums[2], nums[3], nums[4], nums[5])
}

func solve(n, m, i, j, a, b int) int {

	check := func(x0 int, y0 int, x1 int, y1 int) int {
		dx := abs(x0 - x1)
		dy := abs(y0 - y1)
		if dx%a != 0 || dy%b != 0 {
			return -1
		}
		u := dx / a
		v := dy / b
		if u&1 != v&1 {
			return -1
		}
		if u == v {
			return u
		}

		if n <= a || m <= b {
			return -1
		}

		return max(u, v)
	}

	pos := [][]int{
		{1, 1},
		{1, m},
		{n, 1},
		{n, m},
	}

	ans := -1

	for _, cur := range pos {
		tmp := check(cur[0], cur[1], i, j)
		if tmp >= 0 {
			if ans < 0 || tmp < ans {
				ans = tmp
			}
		}
	}
	return ans
}

func abs(num int) int {
	return max(num, -num)
}

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
	k := readNum(reader)
	s := readNNums(reader, 2)
	t := readNNums(reader, 2)
	return solve(k, s, t)
}

func solve(k int, s []int, t []int) int {
	s[0] += k
	s[1] += k
	t[0] += k
	t[1] += k
	dx := abs(s[0] - t[0])
	dy := abs(s[1] - t[1])
	if k == 1 {
		return dx + dy
	}
	// k > 1
	a := getKTileNeighbors(s, k)
	b := getKTileNeighbors(t, k)

	ans := dx + dy

	for _, u := range a {
		for _, v := range b {
			x, y := u[0], u[1]
			z, w := v[0], v[1]
			tmp := u[2] + v[2]
			if k == 2 {
				x_diff := max(x, z) - min(x, z)
				y_diff := max(y, w) - min(y, w)
				tmp += x_diff + y_diff + (max(x_diff, y_diff)-min(x_diff, y_diff))/2
			} else {
				tmp += max(x+y, z+w) - min(x+y, z+w) + max(x+w, z+y) - min(x+w, z+y)
			}
			ans = min(ans, tmp)
		}
	}

	return ans
}

func getKTileNeighbors(s []int, k int) [][]int {
	x, y := s[0], s[1]
	// 要向左取整
	i := x / k
	j := y / k

	if (i^j)&1 == 1 {
		return [][]int{
			{i, j, 0},
		}
	}

	return [][]int{
		{i - 1, j, x%k + 1},
		{i + 1, j, k - x%k},
		{i, j + 1, k - y%k},
		{i, j - 1, y%k + 1},
	}
}
func abs(num int) int {
	return max(num, -num)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func drive(reader *bufio.Reader) []int {
	nums := readNNums(reader, 4)
	h, w, x0, y0 := nums[0], nums[1], nums[2], nums[3]
	s := readString(reader)
	return solve(h, w, x0, y0, s)
}
func solve(h int, w int, x0 int, y0 int, s string) []int {
	marked := make([][]bool, h)
	for i := range h {
		marked[i] = make([]bool, w)
	}
	x0--
	y0--

	k := len(s)
	ans := make([]int, k+1)
	ans[0] = 1

	cnt := 1
	marked[x0][y0] = true

	for i := range k {
		switch s[i] {
		case 'L':
			y0 = max(0, y0-1)
		case 'R':
			y0 = min(w-1, y0+1)
		case 'U':
			x0 = max(0, x0-1)
		default:
			x0 = min(h-1, x0+1)
		}
		if !marked[x0][y0] {
			marked[x0][y0] = true
			ans[i+1] = 1
			cnt++
		}
	}

	ans[k] += h*w - cnt

	return ans
}

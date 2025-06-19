package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) (a []int, res []int) {
	n := readNum(reader)
	a = readNNums(reader, n)
	res = solve(slices.Clone(a))
	return
}

func solve(a []int) []int {
	n := len(a)

	if n == 1 {
		return []int{1, a[0]}
	}

	for i := range n {
		a[i]--
	}
	// y = 1 是一个special的case
	var y int
	for i := 0; i+1 < n; i++ {
		if a[i] == a[i+1] {
			return nil
		}
		if abs(a[i]-a[i+1]) != 1 {
			y = abs(a[i] - a[i+1])
			break
		}
	}
	if y == 0 {
		return []int{1, slices.Max(a) + 1}
	}
	// y > 0

	r, c := a[0]/y, a[0]%y
	x := r
	for i := 1; i < n; i++ {
		u, v := a[i]/y, a[i]%y
		if abs(u-r)+abs(v-c) != 1 {
			return nil
		}
		r, c = u, v
		x = max(x, r)
	}

	return []int{x + 1, y}
}

func abs(num int) int {
	return max(num, -num)
}

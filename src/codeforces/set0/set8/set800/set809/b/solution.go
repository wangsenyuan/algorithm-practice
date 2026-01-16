package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	ask := func(x int, y int) string {
		fmt.Printf("1 %d %d\n", x, y)
		return readString(reader)
	}
	res := solve(n, k, ask)
	fmt.Printf("2 %d %d\n", res[0], res[1])
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

func solve(n int, k int, ask func(int, int) string) []int {
	// 先找出第一个点
	// 怎么问出第二个呢？
	// first的左边和右边？
	find := func(l int, r int) int {
		for l < r {
			mid := (l + r) / 2
			if ask(mid, mid+1) == "TAK" {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r
	}

	first := find(1, n)
	left := 1
	if first > 1 {
		left = find(1, first-1)
	}
	right := n
	if first+1 < n {
		right = find(first+1, n)
	}
	if left != first && ask(left, right) == "TAK" {
		return []int{left, first}
	}
	return []int{first, right}
}

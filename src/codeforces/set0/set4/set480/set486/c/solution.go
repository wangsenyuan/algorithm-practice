package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, p := readTwoNums(reader)
	s := readString(reader)
	res := solve(p, s)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(p int, s string) int {
	n := len(s)

	get := func(x byte, y byte) int {
		if x > y {
			x, y = y, x
		}
		res := int(y - x)
		return min(res, 26-res)
	}

	var sum int

	var arr []int
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			d := get(s[i], s[j])
			sum += d
			arr = append(arr, i)
		}
	}

	if len(arr) == 0 {
		return 0
	}

	p--

	if p > (n-1)/2 {
		p = n - 1 - p
	}
	//p在前半段
	u := sort.SearchInts(arr, p)
	// u == len(arr) or arr[u] >= p
	if u == 0 {
		// 只能往后边移动
		return arr[len(arr)-1] - p + sum
	}
	if u == len(arr) {
		// 只能往前边移动
		return p - arr[0] + sum
	}
	l := p - arr[0]
	r := arr[len(arr)-1] - p
	return sum + min(2*l+r, 2*r+l)
}

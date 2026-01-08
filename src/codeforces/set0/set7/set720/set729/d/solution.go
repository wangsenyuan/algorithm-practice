package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	fmt.Println(len(res))
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

func drive(reader *bufio.Reader) (a int, b int, s string, res []int) {
	firstLine := readNNums(reader, 4)
	n, a, b, k := firstLine[0], firstLine[1], firstLine[2], firstLine[3]
	s = readString(reader)
	res = solve(n, a, b, k, s)
	return
}

func solve(n int, a int, b int, k int, s string) []int {
	var sum int
	var arr [][]int

	last := -1
	for i := 0; i <= n; i++ {
		if i == n || s[i] == '1' {
			sum += (i - last - 1) / b
			arr = append(arr, []int{last + 1, i - 1})
			last = i
		}
	}

	if sum < a {
		return nil
	}

	var res []int

	for sum >= a {
		cur := arr[0]
		arr = arr[1:]
		l, r := cur[0], cur[1]
		sum -= (r - l + 1) / b
		for l+b-1 <= r && sum+(r-l+1)/b >= a {
			l = l + b
			res = append(res, l)
		}
	}
	return res
}

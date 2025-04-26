package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) []int {
	n, m, c := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(a, b, c)
}

func solve(a []int, b []int, c int) []int {
	// a[1] += b[1]
	// a[2] += b[1] + b[2]
	n := len(a)
	m := len(b)

	add := func(x int, y int) int {
		return (x%c + y%c) % c
	}

	sub := func(x int, y int) int {
		return add(x, c-y)
	}

	diff := make([]int, n+2)
	for i := 0; i < m; i++ {
		diff[i] = add(diff[i], b[i])
		if i+n-m+1 < n {
			diff[i+n-m+1] = sub(diff[i+n-m+1], b[i])
		}
	}
	var sum int
	for i := 0; i < n; i++ {
		sum = add(sum, diff[i])
		a[i] = add(a[i]%c, sum)
	}
	return a
}

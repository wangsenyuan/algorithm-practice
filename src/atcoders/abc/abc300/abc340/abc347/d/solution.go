package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, b, xor := readThreeNums(reader)
	res := solve(a, b, xor)
	if res == nil {
		fmt.Println(-1)
	} else {
		fmt.Println(res[0], res[1])
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

func solve(a int, b int, xor int) []int {
	var arr []int
	for i := 0; i < 60; i++ {
		if (xor>>i)&1 == 1 {
			arr = append(arr, i)
		}
	}

	c := a + b - len(arr)

	if c < 0 || c%2 != 0 {
		return nil
	}
	c /= 2
	a -= c
	b -= c
	if a < 0 || b < 0 {
		return nil
	}
	// 两边同时需要的
	// c := a + b - len(arr)
	// a + b >= len(arr)

	var x, y int
	for i := 0; i < a; i++ {
		x |= 1 << arr[i]
	}

	for i := len(arr) - b; i < len(arr); i++ {
		y |= 1 << arr[i]
	}

	for i := 0; i < 60 && c > 0; i++ {
		if (x>>i)&1 == 1 || (y>>i)&1 == 1 {
			continue
		}
		x |= 1 << i
		y |= 1 << i
		c--
	}
	if c > 0 {
		return nil
	}

	return []int{x, y}
}

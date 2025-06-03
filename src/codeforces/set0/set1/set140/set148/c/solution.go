package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, a, b := readThreeNums(reader)
	arr := solve(n, a, b)
	if len(arr) == 0 {
		fmt.Println(-1)
	} else {
		s := fmt.Sprintf("%v", arr)
		fmt.Println(s[1 : len(s)-1])
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

func solve(n int, a int, b int) []int {
	var arr []int
	if b > 0 {
		arr = append(arr, 1)
		for range b {
			arr = append(arr, arr[len(arr)-1]*2)
		}
		for range a {
			arr = append(arr, arr[len(arr)-1]+1)
		}
	} else if a > 0 {
		// b == 0
		if a+2 > n {
			return nil
		}
		arr = append(arr, 1, 1)
		for range a {
			arr = append(arr, arr[len(arr)-1]+1)
		}
	} else {
		// a = 0, b = 0
		arr = append(arr, 1)
	}

	for len(arr) < n {
		arr = append(arr, arr[len(arr)-1])
	}

	return arr
}

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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	if n <= 2 {
		return 0
	}

	arr := make([]int, n)

	check := func(x int, y int) int {
		copy(arr, a)

		var res int
		if a[0] != x {
			res++
		}
		if a[1] != y {
			res++
		}
		arr[0] = x
		arr[1] = y
		d := y - x
		for i := 2; i < n; i++ {
			if arr[i]-arr[i-1] != d {
				res++
				if arr[i]+1-arr[i-1] == d {
					arr[i]++
				} else if arr[i]-1-arr[i-1] == d {
					arr[i]--
				} else {
					return -1
				}
			}
		}
		return res
	}

	ans := -1
	for _, d0 := range []int{-1, 0, 1} {
		for _, d1 := range []int{-1, 0, 1} {
			tmp := check(a[0]+d0, a[1]+d1)
			if tmp == -1 {
				continue
			}

			if ans == -1 || tmp < ans {
				ans = tmp
			}
		}
	}
	return ans
}

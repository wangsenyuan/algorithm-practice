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
	n, p, m := readThreeNums(reader)
	days := make([][]int, n)
	for i := range n {
		days[i] = readNNums(reader, 2)
	}
	return solve(p, m, days)
}

func solve(p int, m int, days [][]int) int {
	var res int
	var sum int

	days = append(days, []int{m + 1, 0})
	for i, cur := range days {
		d, t := cur[0], cur[1]
		if i == 0 {
			res += d - 1
			sum -= d * p
			sum += t
		} else {
			prev := days[i-1][0]
			if sum < 0 {
				res += d - prev
			} else if sum-(d-prev)*p < 0 {
				res += d - 1 - prev - sum/p
			}
			sum -= (d - prev) * p
			sum += t
		}
	}
	return res
}

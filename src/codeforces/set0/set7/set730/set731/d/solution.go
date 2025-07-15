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
	n, c := readTwoNums(reader)
	w := make([][]int, n)

	for i := range n {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k) + 1
		w[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &w[i][j]) + 1
		}
	}

	return solve(c, w)
}

func solve(c int, w [][]int) int {
	freq := make([]int, c+2)

	for i := 0; i+1 < len(w); i++ {
		a := w[i]
		b := w[i+1]
		var j int
		for j < len(a) && j < len(b) && a[j] == b[j] {
			j++
		}
		if j == len(a) {
			freq[0]++
		} else if j == len(b) {
			// 无法得到一个排好序的结果
			return -1
		} else {
			if a[j] < b[j] {
				freq[0]++
				// 当b[j] = c的时候， a[j] =a[j] + c - b[j]
				freq[c-b[j]+1]--
				// a[j] += c - b[j] + 1
				// c - (a[j] + c - b[j] + 1) + 1 = b[j] - a[j]的时候
				// b[j] = 1 时， a[j] = a[j] + c - b[j] + 1
				// 当 a[j]再一次等于1的时候 c + 1 - (a[j] + c - b[j] + 1) = b[j] - a[j]
				freq[c-a[j]+1]++
			} else {
				// a[j] > b[j]
				// 当 a[j] = c 的时候
				freq[c-a[j]+1]++
				freq[c-b[j]+1]--
			}
		}
	}
	m := len(w)
	for i := 0; i <= c; i++ {
		if i > 0 {
			freq[i] += freq[i-1]
		}
		if freq[i] == m-1 {
			return i
		}
	}
	return -1
}

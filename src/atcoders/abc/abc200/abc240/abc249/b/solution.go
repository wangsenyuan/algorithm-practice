package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve1(a []int) int {
	ma := a[0]
	for _, num := range a {
		if num > ma {
			ma = num
		}
	}

	freq := make([]int, ma+1)

	for _, num := range a {
		freq[num]++
	}
	var res int
	for x := 1; x <= ma; x++ {
		for j := 1; j*x <= ma; j++ {
			res += freq[x] * freq[j] * freq[j*x]
		}
	}

	return res
}

func solve(a []int) int {
	sort.Ints(a)
	n := len(a)
	ma := a[n-1]

	freq := make([]int, ma+1)

	var res int

	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			freq[a[i]]++
			i++
		}

		for x := 1; x <= a[j]/x; x++ {
			if a[j]%x == 0 {
				y := a[j] / x
				if x == y {
					res += freq[x] * freq[y] * freq[a[j]]
				} else {
					res += 2 * freq[x] * freq[y] * freq[a[j]]
				}
			}
		}
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) string {
	// n := len(a)
	x := slices.Max(a)
	if x == 0 {
		return "cslnb"
	}

	freq := make(map[int]int)
	for _, num := range a {
		freq[num]++
	}

	if freq[0] > 1 {
		return "cslnb"
	}

	var cnt int
	for x, v := range freq {
		// 第一个人行动后，马上就输了，不管它做了什么
		if v > 2 {
			return "cslnb"
		}
		if v == 2 {
			cnt++
		}
		// 如果有两个2，第一个人还是输了
		if cnt > 1 {
			return "cslnb"
		}
		// x出现2次，x-1出现一次
		// 那么第一个player，必须移动x才行（否则他马上就输了）
		// 但是因为存在x-1, 所以他不能动
		if x > 0 && v > 1 && freq[x-1] > 0 {
			return "cslnb"
		}
	}
	// cnt <= 1
	var sum int
	sort.Ints(a)
	n := len(a)
	for i := 0; i < n; i++ {
		sum += a[i] - i
	}

	if sum%2 == 1 {
		return "sjfnb"
	}
	return "cslnb"
}

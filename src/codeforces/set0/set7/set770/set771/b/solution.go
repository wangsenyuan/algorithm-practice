package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_, _, _, res := drive(reader)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func drive(reader *bufio.Reader) (n int, k int, s []string, res []string) {
	n, k = readTwoNums(reader)
	x := readString(reader)
	s = strings.Split(x, " ")
	res = solve(n, k, s)
	return
}

func solve(n int, k int, s []string) []string {
	// len(s) = n - k + 1
	var id int

	next := func() string {
		if id > 0 {
			var buf []byte

			for i := id; i > 0; i /= 26 {
				x := i % 26
				buf = append(buf, byte('a'+x))
			}
			slices.Reverse(buf)
			buf[0] = byte(buf[0] - 'a' + 'A')
			id++
			return string(buf)
		}
		id++

		return "A"
	}

	ans := make([]string, n)
	if s[0] == "YES" {
		for i := range k {
			ans[i] = next()
		}
	} else {
		ans[0] = "A"
		for i := 1; i < k; i++ {
			ans[i] = next()
		}
	}
	for i := k; i < n; i++ {
		if s[i-k+1] == "YES" {
			ans[i] = next()
		} else {
			ans[i] = ans[i-k+1]
		}
	}

	return ans
}

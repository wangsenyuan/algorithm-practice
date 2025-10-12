package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, r := range res {
		writer.WriteString(r)
		writer.WriteByte('\n')
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func drive(reader *bufio.Reader) (a []string, res []string) {
	n, _ := readTwoNums(reader)
	a = make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	res = solve(a)
	return
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(a []string) []string {

	n := len(a)
	m := len(a[0])

	res := make([][]byte, n)
	for i := range n {
		res[i] = []byte(a[i])
	}

	get := func(x int, y int, t int) byte {
		x++
		y++
		if (x+y)%2 == 0 {
			return byte('0' + x%3*3 + y%3)
		}
		return byte('0' + (x-t+1)%3*3 + (y+t)%3)
	}

	for i := range n {
		for j := range m {
			if res[i][j] == '.' && j+1 < m && res[i][j+1] == '.' {
				c := get(i, j, 1)
				res[i][j] = c
				res[i][j+1] = c
			}
		}
	}

	for i := range n {
		for j := range m {
			if res[i][j] == '.' && i+1 < n && res[i+1][j] == '.' {
				c := get(i, j, 1)
				res[i][j] = c
				res[i+1][j] = c
			}
		}
	}

	for i := range n {
		for j := range m {
			if res[i][j] == '.' {
				if i > 0 && res[i-1][j] >= '0' && res[i-1][j] <= '9' {
					res[i][j] = res[i-1][j]
				} else if i+1 < n && res[i+1][j] >= '0' && res[i+1][j] <= '9' {
					res[i][j] = res[i+1][j]
				} else if j > 0 && res[i][j-1] >= '0' && res[i][j-1] <= '9' {
					res[i][j] = res[i][j-1]
				} else if j+1 < m && res[i][j+1] >= '0' && res[i][j+1] <= '9' {
					res[i][j] = res[i][j+1]
				} else {
					return nil
				}
			}
		}
	}

	ans := make([]string, n)
	for i := range n {
		ans[i] = string(res[i])
	}

	return ans
}

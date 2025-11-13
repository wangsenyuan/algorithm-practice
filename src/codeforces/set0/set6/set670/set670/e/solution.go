package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) string {
	_, _, p := readThreeNums(reader)
	s := readString(reader)
	cmds := readString(reader)
	return solve(s, p, cmds)
}

func solve(s string, p int, cmds string) string {
	n := len(s)
	pair := make([]int, n)
	stack := make([]int, n)
	L := make([]int, n)
	R := make([]int, n)
	var top int
	for i := range n {
		L[i] = i - 1
		R[i] = i + 1
		if s[i] == '(' {
			stack[top] = i
			top++
		} else {
			j := stack[top-1]
			pair[i] = j
			pair[j] = i
			top--
		}
	}
	p--

	diff := make([]int, n+1)
	for i := range len(cmds) {
		switch cmds[i] {
		case 'L':
			p = L[p]
		case 'R':
			p = R[p]
		default:
			// D
			l, r := p, pair[p]
			if l > r {
				l, r = r, l
			}
			r1 := R[r]
			diff[l]++
			diff[r1]--
			l1 := L[l]
			if l1 >= 0 {
				R[l1] = r1
			}
			if r1 < n {
				L[r1] = l1
			}
			if r1 < n {
				p = r1
			} else {
				p = l1
			}
		}
	}

	var buf []byte
	for i := range n {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		if diff[i] == 0 {
			buf = append(buf, s[i])
		}
	}

	return string(buf)
}

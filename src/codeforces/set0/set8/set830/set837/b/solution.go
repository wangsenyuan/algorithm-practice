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
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
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

func drive(reader *bufio.Reader) bool {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) bool {
	ok := true
	for j := 1; j < len(a[0]); j++ {
		if a[0][j] != a[0][0] {
			ok = false
			break
		}
	}

	if !ok {
		a = rotate(a)
	}

	return check(a)
}

func rotate(a []string) []string {
	n, m := len(a), len(a[0])
	res := make([][]byte, m)
	for i := range m {
		res[i] = make([]byte, n)
	}
	for i := range n {
		for j := range m {
			res[j][i] = a[i][j]
		}
	}
	ans := make([]string, m)
	for i := range m {
		ans[i] = string(res[i])
	}
	return ans
}

func check(a []string) bool {
	id := map[byte]int{
		'R': 0, 'G': 1, 'B': 2,
	}

	n := len(a)
	if n%3 != 0 {
		return false
	}
	m := len(a[0])
	sz := n / 3
	var flag int
	for i := range 3 {
		first := a[i*sz]
		for j := 1; j < m; j++ {
			if first[j] != first[0] {
				return false
			}
		}
		flag |= 1 << id[first[0]]
		for j := i * sz; j < (i+1)*sz; j++ {
			if a[j] != first {
				return false
			}
		}
	}
	return flag == 7
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(writer, "NO")
		return
	}
	fmt.Fprintln(writer, "YES")
	for _, cur := range res {
		fmt.Fprintln(writer, cur)
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

func drive(reader *bufio.Reader) (n int, res []string) {
	n = readNum(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	res = solve(n, a)
	return
}

func solve(n int, a []string) []string {
	var pieces [][]int

	for i := range n {
		for j := range n {
			if a[i][j] == 'o' {
				pieces = append(pieces, []int{i, j})
			}
		}
	}

	marked := make([]bool, n*n)

	check := func(dx, dy int) int {
		for _, cur := range pieces {
			x, y := cur[0], cur[1]
			if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
				if a[x+dx][y+dy] == '.' {
					return -1
				}
			}
		}
		var cnt int
		for _, cur := range pieces {
			x, y := cur[0], cur[1]
			if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
				cnt++
				marked[(x+dx)*n+y+dy] = true
			}
		}
		return cnt
	}

	var ans [][]int

	for dx := -n; dx <= n; dx++ {
		for dy := -n; dy <= n; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			res := check(dx, dy)
			if res > 0 {
				ans = append(ans, []int{dx, dy})
			}
		}
	}

	for i := range n {
		for j := range n {
			if a[i][j] == 'x' && !marked[i*n+j] {
				return nil
			}
		}
	}

	buf := make([][]byte, 2*n-1)
	for i := range 2*n - 1 {
		buf[i] = make([]byte, 2*n-1)
		for j := range 2*n - 1 {
			buf[i][j] = '.'
		}
	}

	buf[n-1][n-1] = 'o'
	for _, cur := range ans {
		dx, dy := cur[0], cur[1]
		x, y := n-1+dx, n-1+dy
		buf[x][y] = 'x'
	}

	s := make([]string, 2*n-1)
	for i := range 2*n - 1 {
		s[i] = string(buf[i])
	}
	return s
}

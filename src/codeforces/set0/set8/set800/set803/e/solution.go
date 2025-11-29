package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
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

func drive(reader *bufio.Reader) (k int, s string, res string) {
	_, k = readTwoNums(reader)
	s = readString(reader)
	res = solve(k, s)
	return
}

func solve(k int, s string) string {
	n := len(s)

	fa := make([][]int, n+1)
	for i := range n + 1 {
		fa[i] = make([]int, 2*k+1)
		for j := range 2*k + 1 {
			fa[i][j] = -1
		}
	}

	k2 := 2*k + 1

	fa[0][k] = 0

	var que []int
	que = append(que, k)

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		i, j := cur/k2, cur%k2
		if i == n {
			continue
		}
		j -= k
		switch s[i] {
		case 'W':
			nj := j + 1
			if (nj < k || nj == k && i+1 == n) && fa[i+1][nj+k] == -1 {
				fa[i+1][nj+k] = cur
				que = append(que, (i+1)*k2+nj+k)
			}
		case 'L':
			nj := j - 1
			if (nj > -k || nj == -k && i+1 == n) && fa[i+1][nj+k] == -1 {
				fa[i+1][nj+k] = cur
				que = append(que, (i+1)*k2+nj+k)
			}
		case 'D':
			nj := j
			if fa[i+1][nj+k] == -1 {
				fa[i+1][nj+k] = cur
				que = append(que, (i+1)*k2+nj+k)
			}
		default:
			// ?
			for d := -1; d <= 1; d++ {
				nj := j + d
				if (abs(nj) < k || abs(nj) == k && i+1 == n) && fa[i+1][nj+k] == -1 {
					fa[i+1][nj+k] = cur
					que = append(que, (i+1)*k2+nj+k)
				}
			}
		}
	}

	if fa[n][2*k] == -1 && fa[n][0] == -1 {
		return "NO"
	}

	j := 2 * k
	if fa[n][2*k] == -1 {
		j = 0
	}

	ans := make([]byte, n)
	for i := n; i > 0; i-- {
		p := fa[i][j]
		pj := p % k2

		if pj+1 == j {
			ans[i-1] = 'W'
		} else if pj-1 == j {
			ans[i-1] = 'L'
		} else {
			ans[i-1] = 'D'
		}

		j = pj
	}

	return string(ans)
}

func abs(num int) int {
	return max(num, -num)
}

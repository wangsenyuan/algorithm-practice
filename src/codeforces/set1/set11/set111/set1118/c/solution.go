package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for i := range res {
		for _, v := range res[i] {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (n int, a []int, res [][]int) {
	n = readNum(reader)
	a = readNNums(reader, n*n)
	res = solve(n, a)
	return
}

func solve(n int, a []int) [][]int {
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, n)
	}
	if n%2 == 0 {
		var x, y int
		for k, v := range freq {
			if v%4 != 0 {
				return nil
			}
			for v > 0 {
				res[x][y] = k
				res[n-1-x][y] = k
				res[x][n-1-y] = k
				res[n-1-x][n-1-y] = k
				y++
				if y == n/2 {
					x++
					y = 0
				}
				v -= 4
			}
		}
		return res
	}
	res[n/2][n/2] = -1
	// 那么中心有一个独立的数字。
	// 且有 (n - 1) * 2个数字的频率%4 = 2
	for k, v := range freq {
		if v%2 == 1 {
			res[n/2][n/2] = k
			freq[k]--
			break
		}
	}
	if res[n/2][n/2] == -1 {
		return nil
	}
	var x, y int
	var arr []int
	for k, v := range freq {
		if v%2 != 0 {
			return nil
		}
		for v >= 4 && x < n/2 {
			res[x][y] = k
			res[n-1-x][y] = k
			res[x][n-1-y] = k
			res[n-1-x][n-1-y] = k
			y++
			if y == n/2 {
				x++
				y = 0
			}
			v -= 4
		}
		freq[k] = v
		if v > 0 {
			arr = append(arr, k)
		}
	}

	if x != n/2 {
		return nil
	}

	for _, k := range arr {
		v := freq[k]
		for v > 0 && y < n/2 {
			res[x][y] = k
			res[x][n-1-y] = k
			v -= 2
			y++
		}
		freq[k] = v
	}
	if y != n/2 {
		return nil
	}
	y = n / 2
	x = 0
	for _, k := range arr {
		v := freq[k]
		for v > 0 && x < n/2 {
			res[x][y] = k
			res[n-1-x][y] = k
			v -= 2
			x++
		}
		freq[k] = v
	}

	if x != n/2 {
		return nil
	}

	return res
}

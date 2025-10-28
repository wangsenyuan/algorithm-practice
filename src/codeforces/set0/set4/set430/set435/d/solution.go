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

func drive(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) int {
	n := len(a)
	m := len(a[0])
	sum := make([][]int, n)
	d1 := make([][]int, n)
	d2 := make([][]int, n)
	for i := range n {
		sum[i] = make([]int, m)
		d1[i] = make([]int, m)
		d2[i] = make([]int, m)
	}

	for i := range n {
		for j := range m {
			x := int(a[i][j] - '0')
			sum[i][j] = x
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
			d1[i][j] = x
			d2[i][j] = x
			if i > 0 && j > 0 {
				d1[i][j] += d1[i-1][j-1]
			}
			if i > 0 && j+1 < m {
				d2[i][j] += d2[i-1][j+1]
			}
		}
	}

	getRect := func(r1 int, c1 int, r2 int, c2 int) int {
		res := sum[r2][c2]
		if r1 > 0 {
			res -= sum[r1-1][c2]
		}
		if c1 > 0 {
			res -= sum[r2][c1-1]
		}
		if r1 > 0 && c1 > 0 {
			res += sum[r1-1][c1-1]
		}
		return res
	}

	getDiag1 := func(r1 int, c1 int, r2 int, c2 int) int {
		// r2 - r1 = c2 - c1
		res := d1[r2][c2]
		if r1 > 0 && c1 > 0 {
			res -= d1[r1-1][c1-1]
		}
		return res
	}

	getDiag2 := func(r1 int, c1 int, r2 int, c2 int) int {
		// r2 - r1 = c1 - c2
		res := d2[r2][c2]
		if r1 > 0 && c1+1 < m {
			res -= d2[r1-1][c1+1]
		}
		return res
	}

	var res int

	for i := range n {
		for j2 := range m {
			if a[i][j2] == '0' {
				for j1 := range j2 {
					if getRect(i, j1, i, j2) == 0 {
						d := j2 - j1 + 1
						if i-d+1 >= 0 && getRect(i-d+1, j1, i, j1) == 0 && getDiag1(i-d+1, j1, i, j2) == 0 {
							res++
						}
						if i+d-1 < n && getRect(i, j1, i+d-1, j1) == 0 && getDiag2(i, j2, i+d-1, j1) == 0 {
							res++
						}
						if i-d+1 >= 0 && getRect(i-d+1, j2, i, j2) == 0 && getDiag2(i-d+1, j2, i, j1) == 0 {
							res++
						}
						if i+d-1 < n && getRect(i, j2, i+d-1, j2) == 0 && getDiag1(i, j1, i+d-1, j2) == 0 {
							res++
						}
						if d&1 == 1 {
							h := d / 2
							if i-h >= 0 && getDiag2(i-h, j1+h, i, j1) == 0 && getDiag1(i-h, j1+h, i, j2) == 0 {
								res++
							}
							if i+h < n && getDiag1(i, j1, i+h, j1+h) == 0 && getDiag2(i, j2, i+h, j1+h) == 0 {
								res++
							}
						}
					}
				}
			}
		}
	}

	for j := range m {
		for i2 := range n {
			for i1 := range i2 {
				if (i2-i1+1)&1 == 1 && getRect(i1, j, i2, j) == 0 {
					h := (i2 - i1 + 1) / 2
					if j-h >= 0 && getDiag2(i1, j, i1+h, j-h) == 0 && getDiag1(i1+h, j-h, i2, j) == 0 {
						res++
					}
					if j+h < m && getDiag1(i1, j, i1+h, j+h) == 0 && getDiag2(i1+h, j+h, i2, j) == 0 {
						res++
					}
				}
			}
		}
	}

	return res
}

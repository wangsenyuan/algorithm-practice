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
	n, m, mod := readThreeNums(reader)
	a := make([]string, m)
	for i := range m {
		a[i] = readString(reader)
	}
	return solve(n, mod, a)
}

func solve(n int, mod int, a []string) int {
	m := len(a)

	add := func(a, b int) int {
		a += b
		if a >= mod {
			a -= mod
		}
		return a
	}

	mul := func(nums ...int) int {
		res := 1
		for _, num := range nums {
			res = res * num % mod
		}
		return res
	}

	nCr := func(n int, r int) int {
		if r == 0 {
			return 1
		}
		if r == 1 {
			return n
		}
		// r == 2
		if n%2 == 0 {
			return mul(n/2, n-1)
		}
		return mul(n, (n-1)/2)
	}

	// dp[i][j] 表示有i行为0，j行为1时的状态
	// i + j <= n
	k := n - m
	dp := make([][]int, k+1)
	ndp := make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, k-i+1)
		ndp[i] = make([]int, k-i+1)
	}

	dp[k][0] = 1

	row := make([][]int, m)
	for i := range m {
		for j := range n {
			if a[i][j] == '1' {
				row[i] = append(row[i], j)
			}
		}
	}

	for c := range n {
		var cnt int
		for i := range m {
			if row[i][0] == c || row[i][1] == c {
				cnt++
			}
		}

		for d0 := range dp {
			for d1, val := range dp[d0] {
				if val > 0 {
					// 这一列不增加1的情况下
					switch cnt {
					case 0:
						// 可以增加最多2个1
						if d0 >= 2 {
							// 将2个0变成1，且这一列只有这两个1
							ndp[d0-2][d1+2] = add(ndp[d0-2][d1+2], mul(nCr(d0, 2), val))
						}
						if d1 >= 2 {
							// 将两个1的变成2，且这一列变成了两个1
							ndp[d0][d1-2] = add(ndp[d0][d1-2], mul(nCr(d1, 2), val))
						}
						// 一个0变成1，一个1变成2
						if d0 > 0 && d1 > 0 {
							ndp[d0-1][d1] = add(ndp[d0-1][d1], mul(d0, d1, val))
						}
					case 1:
						// 可以改变1个
						if d0 > 0 {
							ndp[d0-1][d1+1] = add(ndp[d0-1][d1+1], mul(d0, val))
						}
						if d1 > 0 {
							ndp[d0][d1-1] = add(ndp[d0][d1-1], mul(d1, val))
						}
					default:
						// c0 + c1 == 2
						ndp[d0][d1] = add(ndp[d0][d1], val)
					}
				}
			}
		}

		for d0 := range ndp {
			for d1, v := range ndp[d0] {
				dp[d0][d1] = v
				ndp[d0][d1] = 0
			}
		}
	}

	return dp[0][0]
}

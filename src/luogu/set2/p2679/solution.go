package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, k := readThreeNums(reader)
	s := readString(reader)
	t := readString(reader)
	res := solve(s, t, k)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(s, t string, K int) int {
	// 如果 a[l...r] = b[i...j]  (r - l + 1 == j - i + 1)
	// 那么 a[l...r] 是一个有效的子串
	n := len(s)
	m := len(t)
	f := make([][][][2]int, 2)
	for i := range f {
		f[i] = make([][][2]int, n+1)
		for j := range f[i] {
			f[i][j] = make([][2]int, m+1)
		}
	}
	for i := 0; i <= n; i++ {
		f[0][i][0] = [2]int{1, 1}
	}

	for k := 0; k < K; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				var res int
				if s[i] == t[j] {
					res += f[k%2][i][j][0]     // f(k-1, i-1, j-1, 0)
					res += f[(k+1)%2][i][j][1] // f(k, i-1, j-1, 1)
					res %= mod
				}
				f[(k+1)%2][i+1][j+1][1] = res

				res2 := f[(k+1)%2][i][j+1][0] // f(k, i-1, j, 0)
				if s[i] == t[j] {
					res2 = (res2 + res) % mod
				}
				f[(k+1)%2][i+1][j+1][0] = res2
			}
		}
		for i := 0; i <= n; i++ {
			f[0][i][0][0] = 0
			f[0][i][0][1] = 0
		}
	}

	return f[K%2][n][m][0]
}

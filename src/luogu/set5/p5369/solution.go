package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(a []int) int {
	n := len(a)
	N := 1 << n
	sum := make([]int, N)
	for state := 1; state < N; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				sum[state] += a[i]
			}
		}
	}

	fp := make([][]int, N)
	for i := 0; i < N; i++ {
		fp[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		if a[i] < 0 {
			fp[1<<i][0] = 1
		} else {
			fp[1<<i][1] = 1
		}
	}

	for state := 1; state < N; state++ {
		for j := 0; j < n; j++ {
			if (state>>j)&1 == 0 {
				next := state | (1 << j)
				if sum[next] >= 0 {
					fp[next][1] = add(fp[next][1], fp[state][1])
				} else {
					fp[next][0] = add(fp[next][0], fp[state][1])
				}
			}
		}
	}

	gp := make([]int, N)
	gp[0] = 1

	for state := 0; state < N; state++ {
		for j := 0; j < n; j++ {
			if (state>>j)&1 == 0 {
				next := state | (1 << j)
				if sum[next] < 0 {
					gp[next] = add(gp[next], gp[state])
				}
			}
		}
	}

	var res int

	for state := 1; state < N; state++ {
		tmp := sum[state] % mod
		if tmp < 0 {
			tmp += mod
		}
		tmp = mul(tmp, add(fp[state][0], fp[state][1]))
		tmp = mul(tmp, gp[(N-1)^state])
		res = add(res, tmp)
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, p, d := readThreeNums(reader)
	a := readNNums(reader, n)
	res := solve(n, p, d, a)
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

func solve(n int, p int, d int, a []int) int {
	var best = d
	sum := make([]int64, n+1)

	for i := 0; i < d-1; i++ {
		sum[i+1] = sum[i] + int64(a[i])
	}
	que := make([]int, n)
	var tail, front int

	for l, r := 0, d-1; r < n; r++ {
		sum[r+1] = sum[r] + int64(a[r])
		cur := sum[r+1] - sum[r-d+1]
		for front > tail {
			i := que[front-1]
			tmp := sum[i+d] - sum[i]
			if tmp > cur {
				break
			}
			front--
		}
		que[front] = r - d + 1
		front++
		for r-l+1 >= d {
			for tail < front && que[tail] < l {
				tail++
			}
			mx := sum[que[tail]+d] - sum[que[tail]]
			tmp := sum[r+1] - sum[l]
			if tmp-mx <= int64(p) {
				best = max(best, r-l+1)
				break
			}
			l++
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

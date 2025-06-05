package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, s := readThreeNums(reader)
	res := solve(n, m, s)
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

func count(a int, b int) int {
	return a - b + 1
}

func solve(n int, m int, s int) int {
	if n > m {
		n, m = m, n
	}
	var res int
	for th := 1; th <= n; th += 2 {
		for tw := 1; tw <= m; tw += 2 {
			tmp := count(n, th) * count(m, tw)

			if th*tw <= s {
				if th*tw == s {
					res += tmp * (2*(th/2+1)*(tw/2+1) - 1)
				}
				continue
			}

			corner := th*tw - s
			// corner > 0
			if corner%4 != 0 {
				continue
			}
			corner /= 4

			for h := 1; h*2 < th; h++ {
				if corner%h == 0 {
					w := corner / h
					if 2*w >= tw {
						continue
					}
					res += 2 * tmp
				}
			}
		}
	}
	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) int {
	n, s := readTwoNums(reader)
	cars := make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k) + 1
		cars[i] = make([]int, 2*k)
		for j := range 2 * k {
			pos = readInt(s, pos, &cars[i][j]) + 1
		}
	}
	return solve(s, cars)
}

type segment struct {
	speed int
	time  int
}

func solve(s int, cs [][]int) int {
	n := len(cs)
	cars := make([][]segment, n)
	var stages []int

	merge := func(a, b []int) []int {
		var c []int
		for i, j := 0, 0; i < len(a) || j < len(b); {
			if j == len(b) || i < len(a) && a[i] <= b[j] {
				if j < len(b) && a[i] == b[j] {
					j++
				}
				c = append(c, a[i])
				i++
			} else {
				c = append(c, b[j])
				j++
			}
		}
		return c
	}

	for i := range n {
		var tmp []int
		var cur int
		tmp = append(tmp, 0)
		for j := 0; j < len(cs[i]); j += 2 {
			cur += cs[i][j+1]
			cars[i] = append(cars[i], segment{
				speed: cs[i][j],
				time:  cur,
			})
			tmp = append(tmp, cur)
		}
		stages = merge(stages, tmp)
	}

	pos := make([]int, n)
	id := make([]int, n)
	behind := make([]int, n)
	check := func(x int) int {
		clear(pos)
		clear(id)
		clear(behind)

		var res int
		for i := 0; i+1 < len(stages); i++ {
			// 计算下一个时刻的位置
			diff := stages[i+1] - stages[i]
			// 所有的都运行这么多时间
			for j := range n {
				if id[j] >= len(cars[j]) {
					continue
				}
				tmp := cars[j][id[j]]
				pos[j] += tmp.speed * diff
				if stages[i+1] == tmp.time {
					id[j]++
				}
			}

			for j := range n {
				if j == x {
					continue
				}
				if behind[j] < 0 && pos[j] > pos[x] {
					res++
				}
				if pos[j] < pos[x] {
					behind[j] = -1
				} else if pos[j] > pos[x] {
					behind[j] = 1
				}
			}
		}
		return res
	}

	var res int
	for i := range n {
		res += check(i)
	}
	return res
}

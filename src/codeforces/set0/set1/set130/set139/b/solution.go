package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	m := readNum(reader)
	rooms := make([][]int, m)
	for i := range m {
		rooms[i] = readNNums(reader, 3)
	}

	n := readNum(reader)
	papers := make([][]int, n)
	for i := range n {
		papers[i] = readNNums(reader, 3)
	}
	return solve(rooms, papers)
}

const inf = 1 << 60

func solve(rooms [][]int, papers [][]int) int {

	check := func(room []int) int {
		l, w, h := room[0], room[1], room[2]
		l = 2 * (l + w)

		res := inf

		for _, cur := range papers {
			ch, cw, cp := cur[0], cur[1], cur[2]
			if ch < h {
				// 不够从底贴到顶的
				continue
			}
			// 一卷可以贴出来x次
			x := ch / h
			// 共需要y面
			y := (l + cw - 1) / cw
			tmp := (y + x - 1) / x * cp
			res = min(res, tmp)
		}
		return res
	}

	var ans int

	for _, cur := range rooms {
		ans += check(cur)
	}

	return ans
}

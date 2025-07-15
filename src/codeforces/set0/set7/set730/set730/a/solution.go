package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, best, ans := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n%d\n", best, len(ans)))
	for _, cur := range ans {
		buf.WriteString(cur)
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

func process(reader *bufio.Reader) (r []int, best int, ans []string) {
	n := readNum(reader)
	r = readNNums(reader, n)
	best, ans = solve(n, r)
	return
}

type player struct {
	id     int
	rating int
}

func solve(n int, r []int) (int, []string) {
	players := make([]player, n)
	for i := 0; i < n; i++ {
		players[i] = player{i, r[i]}
	}

	var best int
	var ans []string
	for {
		slices.SortFunc(players, func(a, b player) int {
			return b.rating - a.rating
		})
		if players[0].rating == players[n-1].rating {
			break
		}
		buf := make([]byte, n)
		for i := range n {
			buf[i] = '0'
		}

		j := n - 1
		for j > 0 && players[j].rating == players[n-1].rating {
			j--
		}

		if j < 5 && j > 0 && players[j].rating == players[0].rating {
			for players[0].rating > players[n-1].rating {
				for i := 0; i <= j; i++ {
					buf[players[i].id] = '1'
					players[i].rating--
				}
				ans = append(ans, string(buf))
			}
			break
		}

		buf[players[0].id] = '1'
		players[0].rating--
		buf[players[1].id] = '1'
		if players[1].rating > 0 {
			players[1].rating--
		}

		ans = append(ans, string(buf))
	}

	best = players[0].rating
	return best, ans
}

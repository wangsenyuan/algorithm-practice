package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)

	res := drive(reader)

	for _, cur := range res {
		fmt.Fprintf(w, "%d", len(cur))
		for _, x := range cur {
			fmt.Fprintf(w, " %s", x)
		}
		fmt.Fprintln(w)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func drive(reader *bufio.Reader) [][]string {
	n, t := readTwoNums(reader)
	teams := make([][]int, n)
	for i := range n {
		teams[i] = readNNums(reader, 4)
	}
	m := readNum(reader)
	words := make([]string, m)
	c := make([]int, m)
	for i := range m {
		words[i] = readString(reader)
		c[i] = readNum(reader)
	}
	return solve(t, teams, words, c)
}

func solve(t int, teams [][]int, words []string, c []int) [][]string {
	n := len(teams)
	res := make([][]string, n)

	m := len(words)

	marked := make([]bool, m)

	var team_id, card_id int
	var player int

	d := make([][]int, n)
	for i := range n {
		d[i] = make([]int, m)
	}

	a := make([][]int, n)
	b := make([][]int, n)
	for i := range n {
		a[i] = make([]int, 2)
		b[i] = make([]int, 2)
		a[i][0] = teams[i][0]
		a[i][1] = teams[i][2]
		b[i][0] = teams[i][1]
		b[i][1] = teams[i][3]
	}

	cnt := m

	for cnt > 0 {
		nt := t
		for nt > 0 && cnt > 0 {
			for marked[card_id] {
				card_id = (card_id + 1) % m
			}
			w := max(1, c[card_id]-(a[team_id][player]+b[team_id][1^player])-d[team_id][card_id])
			if w <= nt {
				cnt--
				res[team_id] = append(res[team_id], words[card_id])
				marked[card_id] = true
				nt -= w
			} else {
				// w > nt
				d[team_id][card_id] += nt
				nt = 0
			}
			card_id = (card_id + 1) % m
		}

		team_id = (team_id + 1) % n
		if team_id == 0 {
			player ^= 1
		}
	}

	return res
}

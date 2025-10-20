package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, tot, assign := drive(reader)
	if tot == 0 {
		fmt.Println("No solution")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, tot)
	for i := range assign {
		fmt.Fprintln(writer, assign[i])
	}
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

func drive(reader *bufio.Reader) (M int, T int, requests []string, tot int, assign []int) {
	var n int
	n, M, T = readThreeNums(reader)
	requests = make([]string, n)
	for i := range n {
		requests[i] = readString(reader)
	}
	tot, assign = solve(M, T, requests)
	return
}

const ONE_DAY = 24 * 60 * 60

func solve(M int, T int, requests []string) (tot int, assign []int) {

	n := len(requests)
	arr := make([]int, n)
	for i := range n {
		arr[i] = parseTime(requests[i])
	}

	assign = make([]int, n)

	id := 1

	last := make([]int, n)
	var reach_m bool
	var active []int
	var j int
	for i := range ONE_DAY {
		for len(active) > 0 && last[active[0]] == i {
			//u := active[0]
			active = active[1:]
			//t.Update(arr[u], last[u]-1, -1)
		}
		for j < n && arr[j] == i {
			active = append(active, j)
			assign[j] = id
			id++
			last[j] = i + T
			j++
		}
		m := len(active)

		if m >= M {
			reach_m = true
		}

		if m > M {
			best := active[M-1]
			last[best] = last[active[m-1]]
			for m > M {
				u := active[m-1]
				assign[u] = assign[best]
				m--
			}
			active = active[:M]
		}
	}

	if !reach_m {
		return 0, nil
	}

	seen := make([]bool, id+1)
	for _, v := range assign {
		seen[v] = true
	}

	ord := make([]int, id+1)

	for i := 1; i <= id; i++ {
		if !seen[i] {
			ord[i] = ord[i-1]
		} else {
			tot++
			ord[i] = ord[i-1] + 1
		}
	}

	for i := range n {
		assign[i] = ord[assign[i]]
	}
	return
}

func parseTime(v string) int {
	ss := strings.Split(v, ":")
	h, _ := strconv.Atoi(ss[0])
	m, _ := strconv.Atoi(ss[1])
	s, _ := strconv.Atoi(ss[2])
	return h*60*60 + m*60 + s
}

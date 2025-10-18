package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, u := range res {
		fmt.Fprintf(writer, "%d ", u)
	}
	fmt.Fprintln(writer)
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

func drive(reader *bufio.Reader) (n int, contacts []string, a []int, res []int) {
	n = readNum(reader)
	contacts = make([]string, n)
	for i := range n {
		contacts[i] = readString(reader)
	}
	a = readNNums(reader, n)
	res = solve(n, contacts, a)
	return
}

func solve(n int, contacts []string, a []int) []int {
	deg := make([]int, n)
	for _, cur := range contacts {
		for j := range n {
			if cur[j] == '1' {
				deg[j]++
			}
		}
	}
	marked := make([]bool, n)
	for {
		var arr []int
		for i := range n {
			if !marked[i] && deg[i] == a[i] {
				arr = append(arr, i)
				marked[i] = true
			}
		}
		if len(arr) == 0 {
			break
		}
		for _, u := range arr {
			for j := range n {
				if contacts[u][j] == '1' {
					deg[j]--
				}
			}
		}
	}
	var res []int
	for i := range n {
		if !marked[i] {
			res = append(res, i+1)
		}
	}
	return res
}

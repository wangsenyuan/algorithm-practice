package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0])
	fmt.Println(res[1])
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

func process(reader *bufio.Reader) []string {
	n := readNum(reader)

	readArray := func() []int {
		s, _ := reader.ReadBytes('\n')
		var cnt int
		pos := readInt(s, 0, &cnt) + 1
		arr := make([]int, cnt)
		for j := range cnt {
			pos = readInt(s, pos, &arr[j]) + 1
		}
		return arr
	}

	rick := readArray()
	morty := readArray()

	return solve(n, rick, morty)
}

func solve(n int, rick []int, morty []int) []string {
	deg := make([][2]int, n+1)
	win := make([][2]int, n+1)
	vis := make([][2]bool, n+1)

	arr := [][]int{rick, morty}

	var dfs func(u int, d int)

	dfs = func(u int, d int) {
		if vis[u][d] {
			return
		}
		vis[u][d] = true
		for _, x := range arr[d^1] {
			v := (u-x+n-1)%n + 1
			if v == 1 {
				continue
			}
			// 如果当前状态是个失败状态，那么后一个状态(v, d ^ 1)就是一个胜利状态
			if win[u][d] == 0 {
				win[v][d^1] = 1
				dfs(v, d^1)
			} else {
				// 如果当前是个胜利状态，那么后一个状态(v, d ^ 1)的入度减1
				deg[v][d^1]++
				if deg[v][d^1] == len(arr[d^1]) {
					// 当下一个状态面对的都是胜利状态时，它就是一个失败状态
					win[v][d^1] = 0
					dfs(v, d^1)
				}
			}
		}
	}

	dfs(1, 0)
	dfs(1, 1)

	res := make([][]string, 2)
	for d := range 2 {
		res[d] = make([]string, n-1)
		for i := 2; i <= n; i++ {
			if vis[i][d] {
				if win[i][d] == 1 {
					res[d][i-2] = "Win"
				} else {
					res[d][i-2] = "Lose"
				}
			} else {
				res[d][i-2] = "Loop"
			}
		}
	}

	s1 := strings.Join(res[0], " ")
	s2 := strings.Join(res[1], " ")
	return []string{s1, s2}
}

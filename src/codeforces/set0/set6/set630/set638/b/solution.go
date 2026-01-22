package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (fragments []string, res string) {
	s := readString(reader)
	n, _ := strconv.Atoi(s)
	fragments = make([]string, n)
	for i := range n {
		fragments[i] = readString(reader)
	}
	res = solve(fragments)
	return
}

func solve(fragments []string) string {
	// 知道字符的顺序就可以了
	adj := make([][]int, 26)
	var flag int
	deg := make([]int, 26)
	for _, cur := range fragments {
		for i := 0; i < len(cur); i++ {
			a := int(cur[i] - 'a')
			flag |= 1 << a

			if i+1 < len(cur) {
				b := int(cur[i+1] - 'a')
				// a != b holds
				adj[a] = append(adj[a], b)
				deg[b]++
			}
		}
	}

	// len(adj[?]) <= 1
	for i := range 26 {
		slices.Sort(adj[i])
		adj[i] = slices.Compact(adj[i])
	}

	var res []byte

	for i := range 26 {
		if flag&(1<<i) > 0 && deg[i] == 0 {
			u := i
			res = append(res, byte(u+'a'))
			for len(adj[u]) > 0 {
				u = adj[u][0]
				res = append(res, byte(u+'a'))
			}
		}
	}

	return string(res)
}

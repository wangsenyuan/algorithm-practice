package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (mainCourses []int, deps [][]int, res []int) {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	mainCourses = make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &mainCourses[i])
	}
	deps = make([][]int, n)
	for i := range n {
		var m int
		fmt.Fscan(reader, &m)
		deps[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &deps[i][j])
		}
	}
	res = solve(mainCourses, deps)
	return
}

func solve(mainCourses []int, deps [][]int) []int {
	n := len(deps)

	var que []int
	marked := make([]bool, n)
	for _, i := range mainCourses {
		marked[i-1] = true
		que = append(que, i-1)
	}

	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		for _, v := range deps[u] {
			v--
			if !marked[v] {
				marked[v] = true
				que = append(que, v)
			}
		}
	}

	adj := make([][]int, n)
	deg := make([]int, n)
	for i := range n {
		if marked[i] {
			for _, v := range deps[i] {
				v--
				if marked[v] {
					adj[v] = append(adj[v], i)
					deg[i]++
				}
			}
		}
	}

	inQ := make([]bool, n)
	for i := range n {
		if marked[i] && deg[i] == 0 {
			que = append(que, i)
			inQ[i] = true
		}
	}
	if len(que) == 0 {
		return nil
	}

	var pos int
	for pos < len(que) {
		u := que[pos]
		pos++
		for _, v := range adj[u] {
			deg[v]--
			if deg[v] == 0 {
				inQ[v] = true
				que = append(que, v)
			}
		}
	}

	for i := range n {
		if marked[i] && !inQ[i] {
			return nil
		}
	}

	for i := range que {
		que[i]++
	}
	return que
}

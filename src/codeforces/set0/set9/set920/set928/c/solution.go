package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	fmt.Fprintln(writer, len(res))
	for _, p := range res {
		fmt.Fprintln(writer, p.name, p.ver)
	}
}

type proj struct {
	name string
	ver  int
}

func drive(reader *bufio.Reader) []proj {
	var n int
	fmt.Fscan(reader, &n)
	names := make([]string, n)
	vers := make([]int, n)
	deps := make([][]proj, n)
	for i := range n {
		fmt.Fscan(reader, &names[i], &vers[i])
		var d int
		fmt.Fscan(reader, &d)
		deps[i] = make([]proj, d)
		for j := range d {
			fmt.Fscan(reader, &deps[i][j].name, &deps[i][j].ver)
		}
	}
	return solve(names, vers, deps)
}

func solve(names []string, vers []int, deps [][]proj) []proj {
	projectId := make(map[proj]int)
	uniqueNames := make(map[string]int)
	for i := range len(names) {
		projectId[proj{names[i], vers[i]}] = i
		uniqueNames[names[i]]++
	}

	var arr []string
	// n := len(uniqueNames)
	var dist []int
	for k := range uniqueNames {
		uniqueNames[k] = len(dist)
		dist = append(dist, -1)
		arr = append(arr, k)
	}

	projectVersion := make(map[string]int)
	projectVersion[names[0]] = vers[0]

	var que []proj
	que = append(que, proj{names[0], vers[0]})
	dist[uniqueNames[names[0]]] = 0

	var res []proj

	var head int
	for head < len(que) {
		mark := len(que)
		w := dist[uniqueNames[que[head].name]]
		for head < mark {
			cur := que[head]
			head++
			if projectVersion[cur.name] != cur.ver {
				continue
			}

			if cur.name != names[0] {
				res = append(res, cur)
			}
			u := projectId[cur]
			for _, dep := range deps[u] {
				v := uniqueNames[dep.name]
				if dist[v] == -1 || dist[v] == w+1 && dep.ver > projectVersion[dep.name] {
					dist[v] = w + 1
					projectVersion[dep.name] = dep.ver
					que = append(que, dep)
				}
			}
		}

		head = mark
	}

	slices.SortFunc(res, func(a, b proj) int {
		return strings.Compare(a.name, b.name)
	})

	return res
}

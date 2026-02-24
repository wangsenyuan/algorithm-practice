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
	rd, _ := os.Open("input.txt")
	defer rd.Close()
	wr, _ := os.Create("output.txt")
	defer wr.Close()

	reader := bufio.NewReader(rd)
	res := drive(reader)

	writer := bufio.NewWriter(wr)
	defer writer.Flush()

	for id, cur := range res {
		fmt.Fprintf(writer, "Group %c:\n", byte('A'+id))
		for _, team := range cur {
			fmt.Fprintln(writer, team)
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i, s := range ss {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func drive(reader *bufio.Reader) [][]string {
	n := readNums(reader)[0]
	secondLine := readNums(reader)
	x := secondLine[0]
	a := secondLine[1]
	b := secondLine[2]
	c := secondLine[3]
	teams := make([]string, n)
	for i := range n {
		teams[i] = readString(reader)
	}
	return solve(x, a, b, c, teams)
}

type Team struct {
	name  string
	score int
}

func parse(s string) Team {
	ss := strings.Split(s, " ")
	name := ss[0]
	score, _ := strconv.Atoi(ss[1])
	return Team{name, score}
}

func solve(x int, a int, b int, c int, teams []string) [][]string {
	arr := make([]Team, len(teams))
	for i, team := range teams {
		arr[i] = parse(team)
	}

	slices.SortFunc(arr, func(a Team, b Team) int {
		return b.score - a.score
	})

	var baskets [][]Team

	m := len(arr) / 4

	for i := 0; i < len(arr); i += m {
		baskets = append(baskets, arr[i:i+m])
	}

	marked := make([][]bool, 4)
	for i := range 4 {
		marked[i] = make([]bool, m)
	}

	pick := func(marked []bool, basket []Team, w int) Team {
		var j int
		for j < len(marked) {
			if !marked[j] {
				if w == 0 {
					break
				}
				w--
			}
			j++
		}
		marked[j] = true
		return basket[j]
	}

	var res [][]string

	for i := 0; i < len(arr); i += 4 {
		var cur []string
		for t := range 4 {
			x = (x*a + b) % c
			tmp := pick(marked[t], baskets[t], x%m)
			cur = append(cur, tmp.name)
		}
		m--
		res = append(res, cur)
	}

	return res
}

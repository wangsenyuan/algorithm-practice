package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	players := make([]string, n)
	scores := make([]int, n)
	for i := range n {
		// fmt.Fscanf(reader, "%s %d", &players[i], &scores[i])
		s, _ := reader.ReadBytes('\n')
		var j int
		for j < len(s) && s[j] != ' ' {
			j++
		}
		players[i] = string(s[:j])
		readInt(s, j+1, &scores[i])
	}
	m := readNum(reader)
	b := readNNums(reader, m)
	jim := readString(reader)
	return solve(players, scores, b, jim)
}

func solve(players []string, scores []int, b []int, thePlayer string) []int {
	n := len(players)
	for len(b) < n {
		b = append(b, 0)
	}
	sort.Ints(b)

	var teams []team
	var jim team
	for i := range n {
		cur := team{players[i], scores[i]}
		if players[i] == thePlayer {
			jim = cur
		} else {
			teams = append(teams, cur)
		}
	}

	r1 := getBestRank(teams, b, jim)
	r2 := getLeastRank(teams, b, jim)
	return []int{r1 + 1, r2 + 1}
}

type team struct {
	nick  string
	score int
}

func cmpTeam(a team, b team) bool {
	if a.score > b.score || a.score == b.score && a.nick < b.nick {
		return true
	}
	return false
}

func getLeastRank(teams []team, b []int, jim team) int {
	// jim要排名靠后，使用最低分
	jim.score += b[0]
	b = b[1:]
	n := len(b)
	slices.SortFunc(teams, func(x, y team) int {
		if cmpTeam(x, y) {
			return -1
		}
		return 1
	})
	// 已经比jim高的，应该使用最低分(因为他们的选择不影响结果)
	p := sort.Search(len(teams), func(i int) bool {
		return cmpTeam(jim, teams[i])
	})
	for i := 0; i < p; i++ {
		teams[i].score += b[i]
	}

	for i, j := p, p; i < n && j < n; i++ {
		for j < n {
			teams[i].score += b[j]
			if cmpTeam(teams[i], jim) {
				j++
				break
			}
			teams[i].score -= b[j]
			j++
		}
	}

	slices.SortFunc(teams, func(x, y team) int {
		if cmpTeam(x, y) {
			return -1
		}
		return 1
	})
	return sort.Search(len(teams), func(i int) bool {
		return cmpTeam(jim, teams[i])
	})
}

func getBestRank(ts []team, b []int, jim team) int {
	teams := make([]team, len(ts))
	copy(teams, ts)
	// 为了获取最好的rank，应该把最高分分配给jim
	n := len(b)
	jim.score += b[n-1]
	n--
	b = b[:n]
	slices.SortFunc(teams, func(x, y team) int {
		if cmpTeam(x, y) {
			return -1
		}
		return 1
	})
	// 已经比jim高的，应该使用最高分
	p := sort.Search(len(teams), func(i int) bool {
		return cmpTeam(jim, teams[i])
	})
	// p前面的部分，都比jim的排名高, 给他们分配高分，不影响结果
	for i := n - p; i < n; i++ {
		teams[i-(n-p)].score += b[i]
	}
	b = b[:n-p]
	// 现在分配这些分数
	for i := p; i < len(teams); i++ {
		teams[i].score += b[i-p]
	}

	slices.SortFunc(teams, func(x, y team) int {
		if cmpTeam(x, y) {
			return -1
		}
		return 1
	})
	return sort.Search(len(teams), func(i int) bool {
		return cmpTeam(jim, teams[i])
	})
}

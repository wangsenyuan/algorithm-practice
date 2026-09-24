package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, name := range drive(reader) {
		fmt.Fprintln(writer, name)
	}
}

func drive(reader *bufio.Reader) []string {
	var n int
	fmt.Fscan(reader, &n)
	teams := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &teams[i])
	}
	m := n * (n - 1) / 2
	matches := make([]string, m)
	for i := range m {
		var pair, score string
		fmt.Fscan(reader, &pair, &score)
		matches[i] = pair + " " + score
	}
	return solve(teams, matches)
}

type game struct {
	teams  []string
	scores []int
}

type player struct {
	name     string
	scores   int
	netGoals int
	goals    int
}

func parse(cur string) game {
	ss := strings.Split(cur, " ")
	teams := strings.Split(ss[0], "-")
	_scores := strings.Split(ss[1], ":")
	scores := make([]int, len(_scores))
	for i, x := range _scores {
		scores[i], _ = strconv.Atoi(x)
	}
	if scores[0] < scores[1] {
		scores[0], scores[1] = scores[1], scores[0]
		teams[0], teams[1] = teams[1], teams[0]
	}
	return game{teams, scores}
}

func solve(teams []string, matches []string) []string {
	scores := make(map[string]int)
	netGoals := make(map[string]int)
	goals := make(map[string]int)
	for _, cur := range matches {
		game := parse(cur)
		if game.scores[0] == game.scores[1] {
			scores[game.teams[0]]++
			scores[game.teams[1]]++
		} else {
			scores[game.teams[0]] += 3
		}

		netGoals[game.teams[0]] += game.scores[0] - game.scores[1]
		netGoals[game.teams[1]] += game.scores[1] - game.scores[0]

		goals[game.teams[0]] += game.scores[0]
		goals[game.teams[1]] += game.scores[0]
	}

	arr := make([]player, len(teams))

	for i, cur := range teams {
		arr[i] = player{cur, scores[cur], netGoals[cur], goals[cur]}
	}

	slices.SortFunc(arr, func(a player, b player) int {
		return cmp.Or(b.scores-a.scores, b.netGoals-a.netGoals, b.goals-a.goals)
	})

	res := make([]string, len(arr))
	for i, cur := range arr {
		res[i] = cur.name
	}
	n := len(res)
	res = res[:n/2]
	slices.Sort(res)
	return res
}

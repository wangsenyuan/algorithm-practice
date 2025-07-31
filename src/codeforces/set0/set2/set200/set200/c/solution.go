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
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) string {
	game := make([]string, 5)

	for i := range 5 {
		game[i] = readString(reader)
	}

	return solve(game)
}

const BERLAND = "BERLAND"

func solve(game []string) string {
	scores := make(map[string]int)
	cnt := make(map[string]int)
	goal := make(map[string]int)
	lost := make(map[string]int)

	for _, cur := range game {
		ss := strings.Split(cur, " ")
		a := ss[0]
		b := ss[1]

		c := ss[2]
		ww := strings.Split(c, ":")

		v1, _ := strconv.Atoi(ww[0])
		v2, _ := strconv.Atoi(ww[1])
		goal[a] += v1
		lost[a] += v2

		goal[b] += v2
		lost[b] += v1

		cnt[a]++
		cnt[b]++

		if v1 > v2 {
			scores[a] += 3
			scores[b] += 0
		} else if v1 == v2 {
			scores[a]++
			scores[b]++
		} else {
			scores[a] += 0
			scores[b] += 3
		}
	}

	// 找出BERLAND 和另外一只队伍
	first := BERLAND
	var second string
	for k, v := range cnt {
		if v == 2 && k != BERLAND {
			second = k
			break
		}
	}

	type team struct {
		name  string
		score int
		goal  int
		lost  int
	}

	check := func(x int, y int) bool {
		var arr []team
		for k, v := range scores {
			t := team{name: k, score: v}
			t.goal = goal[k]
			t.lost = lost[k]

			switch k {
			case first:
				t.score += 3
				t.goal += x
				t.lost += y
			case second:
				t.goal += y
				t.lost += x
			}

			arr = append(arr, t)
		}
		slices.SortFunc(arr, func(a team, b team) int {
			if a.score != b.score {
				return b.score - a.score
			}
			if a.goal-a.lost != b.goal-b.lost {
				return (b.goal - b.lost) - (a.goal - a.lost)
			}

			if a.goal != b.goal {
				return b.goal - a.goal
			}
			return cmp.Compare(a.name, b.name)
		})

		return arr[0].name == first || arr[1].name == first
	}

	for diff := 1; diff <= 100; diff++ {
		for y := 0; y < 100; y++ {
			if check(y+diff, y) {
				return fmt.Sprintf("%d:%d", y+diff, y)
			}
		}
	}

	return "IMPOSSIBLE"
}

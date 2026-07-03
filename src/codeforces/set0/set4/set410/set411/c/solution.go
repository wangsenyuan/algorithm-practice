package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	players := make([][]int, 4)
	for i := range 4 {
		players[i] = make([]int, 2)
		fmt.Fscan(reader, &players[i][0], &players[i][1])
	}
	return solve(players)
}

func solve(players [][]int) string {
	type plan struct {
		def []int
		atk []int
	}

	a, b, c, d := players[0], players[1], players[2], players[3]
	t1 := []plan{
		{a, b},
		{b, a},
	}
	t2 := []plan{
		{c, d},
		{d, c},
	}

	beat := func(x, y plan) bool {
		return x.def[0] > y.atk[1] && x.atk[1] > y.def[0]
	}

	for _, first := range t1 {
		win := true
		for _, second := range t2 {
			win = win && beat(first, second)
		}
		if win {
			return "Team 1"
		}
	}

	win := true
	for _, first := range t1 {
		canBeat := false
		for _, second := range t2 {
			canBeat = canBeat || beat(second, first)
		}
		win = win && canBeat
	}
	if win {
		return "Team 2"
	}

	return "Draw"
}

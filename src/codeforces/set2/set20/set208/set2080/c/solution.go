package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	cards := make([]card, n)
	for i := range n {
		var b int
		fmt.Fscan(reader, &b)
		cards[i] = card{a[i], b}
	}

	single := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &single[i])
	}

	if solve(cards, single) {
		return "First"
	}
	return "Second"
}

type card struct {
	a int
	b int
}

func solve(cards []card, single []int) bool {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].a > cards[j].a
	})

	last := 0
	for i := 1; i < len(cards); i++ {
		if cards[i].b < cards[last].a {
			last = i
		}
	}

	var before int
	for _, x := range single {
		if x < cards[last].a {
			before++
		}
	}

	before += len(cards) - 1 - last
	return before%2 == 0
}

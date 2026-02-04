package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, mod int
	fmt.Fscan(reader, &a, &b, &mod)
	player, s1 := solve(a, b, mod)
	if player == 1 {
		fmt.Println(player, s1)
	} else {
		fmt.Println(player)
	}
}

func solve(a int, b int, mod int) (player int, s1 string) {
	for s1 := 1; s1 <= a && s1 < mod; s1++ {
		w := s1 * 1e9 % mod
		v := (mod - w) % mod
		if v > b {
			return 1, fmt.Sprintf("%09d", s1)
		}
	}
	return 2, ""
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	play := func(cmd string) {
		fmt.Println(cmd)
		readString(reader)
	}
	for range tc {
		n := readNum(reader)
		solve(n, play)
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(n int, play func(string)) {
	// x <- x * 9
	play("mul 9")
	// x in [9, 18, 27, 36, 45, 54, 63, 72, 81]
	play("digit")
	// x = 9
	play("digit")

	play(fmt.Sprintf("add %d", n-9))
	play("!")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readInt(reader)
	ask := func(i int, j int) string {
		fmt.Printf("? %d %d\n", i, j)
		return readString(reader)
	}
	for range tc {
		n := readInt(reader)

		res := solve(n, ask)
		fmt.Printf("! %d %d\n", res[0], res[1])
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(reader *bufio.Reader) int {
	s := readString(reader)
	x, _ := strconv.Atoi(s)
	return x
}

func solve(n int, ask func(int, int) string) []int {
	if n == 2 {
		op := ask(1, 2)
		if op == "<" || op == "=" {
			return []int{1, 2}
		}
		return []int{2, 1}
	}
	
	var x []int
	var y []int

	for i := 0; i+1 < n; i += 2 {
		op := ask(i+1, i+2)
		switch op {
		case ">", "=":
			x = append(x, i+1)
			y = append(y, i+2)
		case "<":
			x = append(x, i+2)
			y = append(y, i+1)
		}
	}
	if n&1 == 1 {
		x = append(x, n)
		y = append(y, n)
	}

	var f func(arr []int, op string) int
	f = func(arr []int, op string) int {
		if len(arr) == 1 {
			return arr[0]
		}
		var x []int
		for i := 0; i+1 < len(arr); i += 2 {
			tmp := ask(arr[i], arr[i+1])
			switch tmp {
			case "=", op:
				x = append(x, arr[i])
			default:
				x = append(x, arr[i+1])
			}
		}
		if len(arr)&1 == 1 {
			x = append(x, arr[len(arr)-1])
		}
		return f(x, op)
	}

	i := f(y, "<")
	j := f(x, ">")

	return []int{i, j}
}

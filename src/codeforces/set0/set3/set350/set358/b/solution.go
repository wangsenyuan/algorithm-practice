package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
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

func process(reader *bufio.Reader) bool {
	n := readNum(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	message := readString(reader)
	return solve(words, message)
}

func solve(words []string, message string) bool {

	var i int
	for _, word := range words {
		var state int
		var j int
		for j < len(word) && i < len(message) {
			if message[i] == '<' && state == 0 || message[i] == '3' && state == 1 {
				state++
			} else if message[i] == word[j] && state == 2 {
				j++
			}
			i++
		}

		if j < len(word) || state < 2 {
			return false
		}
	}

	var state int

	for i < len(message) {
		if message[i] == '<' && state == 0 || message[i] == '3' && state == 1 {
			state++
		}
		i++
	}

	return state == 2
}

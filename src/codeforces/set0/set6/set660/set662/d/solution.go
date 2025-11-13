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
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func drive(reader *bufio.Reader) []string {
	n := readNum(reader)
	abbreviations := make([]string, n)
	for i := range n {
		abbreviations[i] = readString(reader)
	}
	return solve(n, abbreviations)
}

func solve(n int, abbreviations []string) []string {
	ans := make([]string, n)
	for j, cur := range abbreviations {
		k := len(cur) - 4
		year, _ := strconv.Atoi(cur[4:])
		var F int
		p10 := 10
		for i := 1; i < k; i++ {
			F += p10
			p10 *= 10
		}
		for year < 1989+F {
			year += p10
		}
		ans[j] = fmt.Sprintf("%d", year)
	}
	return ans
}

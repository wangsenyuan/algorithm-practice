package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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

func drive(reader *bufio.Reader) string {
	n := readNum(reader)
	titles := make([]string, n)
	for i := range n {
		titles[i] = readString(reader)
	}
	return solve(titles)
}

func solve(titles []string) string {

	// 考虑只有一个长度为20的字符串, 它都没法覆盖长度为2的字符串
	// 怀疑，这个长度不会超过4

	check := func(buf []byte) bool {
		s := string(buf)
		for _, title := range titles {
			if strings.Contains(title, s) {
				return false
			}
		}
		return true
	}

	var f func(buf []byte, i int) bool

	f = func(buf []byte, i int) bool {
		if i == len(buf) {
			return check(buf)
		}
		for c := range 26 {
			buf[i] = byte(c + 'a')
			if f(buf, i+1) {
				return true
			}
		}
		return false
	}

	for l := range 4 {
		buf := make([]byte, l+1)
		if f(buf, 0) {
			return string(buf)
		}
	}

	return ""
}

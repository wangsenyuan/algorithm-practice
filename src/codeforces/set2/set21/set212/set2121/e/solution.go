package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		s := readString(reader)
		ss := strings.Split(s, " ")
		l, r := ss[0], ss[1]
		res := solve(l, r)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

const inf = 1 << 60

func solve(l string, r string) int {
	if len(l) < len(r) {
		l = strings.Repeat("0", len(r)-len(l)) + l
	}
	var res int
	var a, b int
	for i := 0; i < len(r); i++ {
		a = a*10 + int(l[i]-'0')
		b = b*10 + int(r[i]-'0')
		if a+1 < b {
			break
		}
		res++
		if a == b {
			res++
		}
	}
	return res
}

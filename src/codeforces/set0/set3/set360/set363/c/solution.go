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
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) string {
	var buf []byte
	for i := range s {
		x := s[i]
		if len(buf) >= 2 && buf[len(buf)-2] == x && buf[len(buf)-1] == x {
			// rule 1
			continue
		}
		if len(buf) >= 3 && buf[len(buf)-1] == x && buf[len(buf)-3] == buf[len(buf)-2] {
			// rule 2
			continue
		}
		buf = append(buf, x)
	}
	return string(buf)
}

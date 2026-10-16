package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) string {
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) string {
	var buf []byte
	for i := range s {
		v := s[i]
		if len(buf) > 0 && buf[len(buf)-1] == v {
			buf = buf[:len(buf)-1]
		} else {
			buf = append(buf, v)
		}
	}
	return string(buf)
}

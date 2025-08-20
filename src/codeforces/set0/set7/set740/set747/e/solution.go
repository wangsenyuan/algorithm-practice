package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	drive(reader, writer)
}

func drive(reader *bufio.Reader, writer *bufio.Writer) {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	solve(s, writer)
}

func solve(s string, writer *bufio.Writer) {
	var roots []*Comment

	var i int

	var max_level int

	var f func(level int) *Comment
	f = func(level int) *Comment {
		max_level = max(max_level, level)
		// i < len(s) always holds
		i1 := i
		for i < len(s) && s[i] != ',' {
			i++
		}
		comment := &Comment{
			p1: i1,
			p2: i,
		}
		if i < len(s) {
			i++
			var cnt int
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				cnt = cnt*10 + int(s[i]-'0')
				i++
			}
			if i < len(s) {
				// s[i] = ,
				i++
			}
			for range cnt {
				child := f(level + 1)
				comment.children = append(comment.children, child)
			}
		}

		if level == 0 {
			roots = append(roots, comment)
		}

		return comment
	}

	for i < len(s) {
		f(0)
	}

	fmt.Fprintf(writer, "%d\n", max_level+1)

	for len(roots) > 0 {
		var next []*Comment

		for i, root := range roots {
			if i == 0 {
				fmt.Fprintf(writer, "%s", s[root.p1:root.p2])
			} else {
				fmt.Fprintf(writer, " %s", s[root.p1:root.p2])
			}
			next = append(next, root.children...)
		}

		fmt.Fprintln(writer)

		roots = next
	}
}

type Comment struct {
	p1       int
	p2       int
	children []*Comment
}

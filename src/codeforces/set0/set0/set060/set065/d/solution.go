package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
	for _, v := range res {
		fmt.Println(v)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

type state [4]int

var houses = []string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"}

var pos = map[byte]int{'G': 0, 'H': 1, 'R': 2, 'S': 3}

func solve(s string) []string {
	n := len(s)

	cur := make(map[state]int)
	cur[state{0, 0, 0, 0}] = 1
	for i := 0; i < n; i++ {
		x := s[i]

		next := make(map[state]int)
		for k := range cur {
			if x != '?' {
				k[pos[x]]++
				next[k]++
			} else {
				// 这里要产生有可能产生新的状态
				v := k[0]
				for i := range 4 {
					v = min(v, k[i])
				}
				for i := range 4 {
					if k[i] == v {
						nk := k
						nk[i]++
						next[nk]++
					}
				}
			}
			cur = next
		}
	}
	res := make(map[string]int)
	for k := range cur {
		v := k[0]
		for i := range 4 {
			v = min(v, k[i])
		}
		for i := range 4 {
			if k[i] == v {
				res[houses[i]]++
			}
		}
	}
	var buf []string
	for k := range res {
		buf = append(buf, k)
	}
	sort.Strings(buf)
	return buf
}

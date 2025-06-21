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
	res := process(reader)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(ans)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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
	bs, _ := reader.ReadBytes('\n')
	return strings.TrimSpace(string(bs))
}

func process(reader *bufio.Reader) []string {
	n := readNum(reader)
	phones := make([]string, n)
	for i := 0; i < n; i++ {
		phones[i] = readString(reader)
	}
	return solve(phones)
}

func solve(phones []string) []string {
	freq := make(map[string]int)

	for _, cur := range phones {
		tmp := make(map[string]bool)
		for i := 0; i < len(cur); i++ {
			for j := i + 1; j <= len(cur); j++ {
				tmp[cur[i:j]] = true
			}
		}
		for k := range tmp {
			freq[k]++
		}
	}

	ans := make([]string, len(phones))
	for id, cur := range phones {
		best := []int{len(cur), 0}
		for i := 0; i < len(cur); i++ {
			for j := i + 1; j <= len(cur); j++ {
				if freq[cur[i:j]] == 1 {
					if j-i < best[0] {
						best[0] = j - i
						best[1] = i
					}
				}
			}
		}
		ans[id] = cur[best[1] : best[1]+best[0]]
	}

	return ans
}

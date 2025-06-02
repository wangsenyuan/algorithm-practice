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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	s := readString(reader)
	k := readNum(reader)
	pairs := make([]string, k)
	for i := 0; i < k; i++ {
		pairs[i] = readString(reader)
	}
	return solve(s, pairs)
}

func solve(s string, pairs []string) int {
	pid := make([]int, 26)
	for i, cur := range pairs {
		pid[cur[0]-'a'] = i + 1
		pid[cur[1]-'a'] = i + 1
	}

	var res int
	n := len(s)

	for i := 0; i < len(s); {
		j := i
		cnt := make([]int, 2)
		for i < n && pid[s[i]-'a'] == pid[s[j]-'a'] {
			if s[i] == s[j] {
				cnt[0]++
			} else {
				cnt[1]++
			}
			i++
		}
		if pid[s[j]-'a'] == 0 {
			continue
		}
		res += min(cnt[0], cnt[1])
	}

	return res
}

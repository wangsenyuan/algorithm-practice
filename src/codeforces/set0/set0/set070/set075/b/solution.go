package main

import (
	"bufio"
	"bytes"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []string {
	player := readString(reader)
	n := readNum(reader)
	actions := make([]string, n)
	for i := range n {
		actions[i] = readString(reader)
	}
	return solve(player, actions)
}

type action struct {
	user1 string
	key   string
	user2 string
}

var values = map[string]int{"posted": 15, "commented": 10, "likes": 5}

func parseAction(s string) action {
	ss := strings.Split(s, " ")
	user1 := ss[0]
	key := ss[1]
	var user2 string
	if key == "likes" {
		user2 = ss[2]
	} else {
		user2 = ss[3]
	}
	i := strings.Index(user2, "'")
	user2 = user2[:i]
	return action{user1, key, user2}
}

func solve(player string, actions []string) []string {
	scores := make(map[string]int)

	for _, s := range actions {
		a := parseAction(s)
		if a.user1 == player {
			scores[a.user2] += values[a.key]
		} else if a.user2 == player {
			scores[a.user1] += values[a.key]
		} else {
			scores[a.user1] += 0
			scores[a.user2] += 0
		}
	}
	var users []string
	for k := range scores {
		users = append(users, k)
	}
	sort.Slice(users, func(i, j int) bool {
		return scores[users[i]] > scores[users[j]] || (scores[users[i]] == scores[users[j]] && users[i] < users[j])
	})

	return users
}

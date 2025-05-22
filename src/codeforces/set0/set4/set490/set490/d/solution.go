package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cnt, size1, size2 := process(reader)
	if cnt < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(cnt)
		fmt.Println(size1[0], size1[1])
		fmt.Println(size2[0], size2[1])
	}
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) (cnt int, size1 []int, size2 []int) {
	size1 = readNNums(reader, 2)
	size2 = readNNums(reader, 2)
	cnt, size1, size2 = solve(size1, size2)
	return
}

type state [2]int

const inf = 1 << 30

func solve(size1 []int, size2 []int) (int, []int, []int) {
	// 最短路径
	var que []state
	checkAndAdd := func(cur state, d int, dist map[state]int, u int, best map[int]state) {
		if d == 0 {
			que = append(que, cur)
			dist[cur] = u
			if best != nil {
				best[cur[0]*cur[1]] = cur
			}
			return
		}

		for i := 0; i < 2; i++ {
			if cur[i]%d == 0 {
				next := cur
				next[i] -= next[i] / d
				if _, ok := dist[next]; !ok {
					dist[next] = u
					que = append(que, next)
					if best != nil {
						if s, ok := best[next[0]*next[1]]; !ok || dist[s] > u {
							best[next[0]*next[1]] = next
						}
					}
				}
			}
		}
	}

	dist := make(map[state]int)
	best := make(map[int]state)

	checkAndAdd(state{size1[0], size1[1]}, 0, dist, 0, best)

	var tail int
	for tail < len(que) {
		cur := que[tail]
		tail++
		checkAndAdd(cur, 2, dist, dist[cur]+1, best)
		checkAndAdd(cur, 3, dist, dist[cur]+1, best)
	}

	que = que[:0]
	tail = 0

	var s1 state
	var s2 state
	cnt := inf

	dist2 := make(map[state]int)

	checkAndAdd(state{size2[0], size2[1]}, 0, dist2, 0, nil)

	for tail < len(que) {
		cur := que[tail]
		tail++
		u := dist2[cur]
		if v, ok := best[cur[0]*cur[1]]; ok {
			if u+dist[v] < cnt {
				cnt = u + dist[v]
				s1 = v
				s2 = cur
			}
		}

		checkAndAdd(cur, 2, dist2, u+1, nil)
		checkAndAdd(cur, 3, dist2, u+1, nil)
	}

	if cnt < inf {
		return cnt, s1[:], s2[:]
	}

	return -1, nil, nil
}

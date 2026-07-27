package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type constraint struct {
	a         int
	b         int
	direction string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	if res == nil {
		fmt.Fprintln(writer, "IMPOSSIBLE")
		return
	}
	output := make([]byte, 0, 1<<20)
	for i, label := range res {
		if i > 0 {
			output = append(output, ' ')
		}
		output = strconv.AppendInt(output, int64(label), 10)
		if len(output) >= 1<<20 {
			_, _ = writer.Write(output)
			output = output[:0]
		}
	}
	output = append(output, '\n')
	_, _ = writer.Write(output)
}

func drive(reader *bufio.Reader) []int {
	data, _ := io.ReadAll(reader)
	pos := 0

	nextToken := func() []byte {
		for pos < len(data) && data[pos] <= ' ' {
			pos++
		}
		start := pos
		for pos < len(data) && data[pos] > ' ' {
			pos++
		}
		return data[start:pos]
	}
	nextInt := func() int {
		token := nextToken()
		value := 0
		for _, digit := range token {
			value = value*10 + int(digit-'0')
		}
		return value
	}

	n, c := nextInt(), nextInt()
	constraints := make([]constraint, c)
	for i := range constraints {
		constraints[i].a = nextInt()
		constraints[i].b = nextInt()
		constraints[i].direction = string(nextToken())
	}
	return solve(n, constraints)
}

func solve(n int, constraints []constraint) []int {
	maxLeft := make([]int, n+1)
	minRight := make([]int, n+1)
	maxRequired := make([]int, n+1)
	invalid := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		minRight[i] = n + 1
		maxRequired[i] = i
	}

	for _, cur := range constraints {
		a, b := cur.a, cur.b
		if b <= a {
			invalid[a] = true
			continue
		}
		maxRequired[a] = max(maxRequired[a], b)
		if cur.direction == "LEFT" {
			maxLeft[a] = max(maxLeft[a], b)
		} else {
			minRight[a] = min(minRight[a], b)
		}
	}

	// split[root] is the last label in root's left subtree.  It equals
	// root when that subtree is empty.
	split := make([]int, n+1)

	type buildFrame struct {
		root  int
		need  int
		stage int
	}

	// Conceptually, build(root, need) constructs the smallest possible
	// subtree starting at root that contains every label through need.
	// The explicit stack avoids O(n) call-stack depth.
	frames := make([]buildFrame, 1, n)
	frames[0] = buildFrame{root: 1, need: n}
	lastEnd := 0

	for len(frames) > 0 {
		top := len(frames) - 1
		root := frames[top].root

		switch frames[top].stage {
		case 0:
			need := max(frames[top].need, maxRequired[root])
			frames[top].need = need
			if invalid[root] || need > n {
				return nil
			}

			if maxLeft[root] == 0 {
				split[root] = root
				if need == root {
					lastEnd = root
					frames = frames[:top]
				} else {
					frames[top].stage = 2
					frames = append(frames, buildFrame{
						root: root + 1,
						need: need,
					})
				}
			} else {
				frames[top].stage = 1
				frames = append(frames, buildFrame{
					root: root + 1,
					need: maxLeft[root],
				})
			}

		case 1:
			// The left subtree was made as short as possible.  If even
			// this subtree contains a required right node, no split works.
			leftEnd := lastEnd
			if leftEnd >= minRight[root] {
				return nil
			}
			split[root] = leftEnd

			if leftEnd >= frames[top].need {
				lastEnd = leftEnd
				frames = frames[:top]
			} else {
				frames[top].stage = 2
				frames = append(frames, buildFrame{
					root: leftEnd + 1,
					need: frames[top].need,
				})
			}

		case 2:
			// The right subtree's end is also this subtree's end.
			frames = frames[:top]
		}
	}

	type interval struct {
		root int
		end  int
	}

	// Recover the inorder traversal without recursion.  For [root, end],
	// the left and right subtree intervals are determined by split[root].
	result := make([]int, 0, n)
	stack := make([]interval, 0, n)
	root, end := 1, n
	for root <= end || len(stack) > 0 {
		for root <= end {
			stack = append(stack, interval{root: root, end: end})
			end = split[root]
			root++
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.root)
		root = split[top.root] + 1
		end = top.end
	}

	return result
}

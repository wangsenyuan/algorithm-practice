package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, cur := range res {
		fmt.Fprintln(writer, cur)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (files []string, res []string) {
	s := readString(reader)
	n, _ := strconv.Atoi(s)
	files = make([]string, n)
	for i := range n {
		files[i] = readString(reader)
	}
	res = solve(files)
	return
}

func solve(files []string) []string {
	examples := make(map[string]int)
	tests := make(map[string]int)

	n := len(files)

	for _, cur := range files {
		k := len(cur)
		if cur[k-1] == '1' {
			// example
			examples[cur[:k-2]] = 2 * n
		} else {
			tests[cur[:k-2]] = 2 * n
		}
	}

	e := len(examples)

	var res []string

	move := func(u string, v string) {
		res = append(res, fmt.Sprintf("move %s %s", u, v))
	}

	var free []int
	var arr1 []string
	var arr2 []string

	for i := 1; i <= n; i++ {
		id := fmt.Sprintf("%d", i)

		if examples[id] == 0 && tests[id] == 0 {
			free = append(free, i)
			continue
		}

		if i <= e {
			if examples[id] > 0 {
				delete(examples, id)
			} else if tests[id] > 0 {
				arr1 = append(arr1, id)
				tests[id] = i
			}
		}
		if i > e {
			if tests[id] > 0 {
				delete(tests, id)
			} else if examples[id] > 0 {
				arr2 = append(arr2, id)
				examples[id] = i
			}
		}
	}

	// 剩余的放在后面的去处理
	for w, v := range tests {
		if v == 2*n {
			arr1 = append(arr1, w)
		}
	}

	for w, v := range examples {
		if v == 2*n {
			arr2 = append(arr2, w)
		}
	}

	// 只要有free, 肯定存在需要移动的文件
	for len(free) > 0 {
		u := free[0]
		free = free[1:]
		if u <= e {
			// 把examples移动到free来
			w := arr2[0]
			arr2 = arr2[1:]
			move(w, fmt.Sprintf("%d", u))
			if pos := examples[w]; pos <= n {
				free = append(free, pos)
			}
		} else {
			w := arr1[0]
			arr1 = arr1[1:]
			move(w, fmt.Sprintf("%d", u))
			if pos := tests[w]; pos <= n {
				free = append(free, pos)
			}
		}
	}

	if len(arr1) > 0 {
		// free < 0, 这种情况只可能是相互占据位置
		last := fmt.Sprintf("%d", n+1)
		for len(arr1) > 0 {
			u := arr1[0]
			v := arr2[0]
			arr1 = arr1[1:]
			arr2 = arr2[1:]
			move(u, last)
			move(v, u)
			last = v
		}
		move(fmt.Sprintf("%d", n+1), last)
	}

	return res
}

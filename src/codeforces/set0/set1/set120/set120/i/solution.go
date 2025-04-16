package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Fprint(w, res)
}

var digits []int

func init() {
	digits = make([]int, 10)
	digits[0] = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5
	digits[1] = 1<<1 | 1<<2
	digits[2] = 1<<0 | 1<<1 | 1<<3 | 1<<4 | 1<<6
	digits[3] = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<6
	digits[4] = 1<<1 | 1<<2 | 1<<5 | 1<<6
	digits[5] = 1<<0 | 1<<2 | 1<<3 | 1<<5 | 1<<6
	digits[6] = 1<<0 | 1<<2 | 1<<3 | 1<<4 | 1<<5 | 1<<6
	digits[7] = 1<<0 | 1<<1 | 1<<2
	digits[8] = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5 | 1<<6
	digits[9] = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<5 | 1<<6
}

func countDigits(i, j int) int {
	v := digits[i] & digits[j]
	return bits.OnesCount(uint(v))
}

func solve(s string) string {
	n := len(s)
	h := n / 2
	arr := make([]int, n)
	for i := range n {
		arr[i] = int(s[i] - '0')
	}
	var update func(i int, pref int, suf int)

	update = func(i int, pref int, suf int) {
		// pref + suf > 0
		if i == n {
			return
		}
		if i < h {
			suf -= 7 - countDigits(arr[i], arr[i])
		} else {
			suf -= countDigits(arr[i-h], arr[i-h]) - countDigits(arr[i], arr[i-h])
		}
		for x := range 10 {
			var tmp int
			if i < h {
				tmp = countDigits(x, x) - countDigits(arr[i], arr[i])
			} else {
				tmp = countDigits(x, arr[i-h]) - countDigits(arr[i], arr[i-h])
			}
			if pref+tmp+suf > 0 {
				arr[i] = x
				update(i+1, pref+tmp, suf)
				return
			}
		}
	}

	var delta int

	ok := false
	for i := n - 1; i >= 0; i-- {
		if arr[i] < 9 {
			for x := arr[i] + 1; x <= 9; x++ {
				var tmp int
				if i < h {
					tmp = countDigits(x, x) - countDigits(arr[i], arr[i])
				} else {
					tmp = countDigits(x, arr[i-h]) - countDigits(arr[i-h], arr[i])
				}

				if tmp+delta > 0 {
					arr[i] = x
					if i < h {
						arr[i+h] = x
					}
					update(i+1, tmp, delta)
					ok = true
					break
				}
			}
		}
		if ok {
			break
		}
		if i < h {
			delta += 7 - countDigits(arr[i], arr[i])
		} else {
			delta += countDigits(arr[i-h], arr[i-h]) - countDigits(arr[i], arr[i-h])
		}
	}

	if ok {
		// create string from arr
		var buf bytes.Buffer
		for _, v := range arr {
			buf.WriteByte(byte(v + '0'))
		}
		return buf.String()
	}

	return "-1"
}

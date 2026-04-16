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
	if drive(reader) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i, cur := range ss {
		nums[i], _ = strconv.Atoi(cur)
	}
	return nums
}

func drive(reader *bufio.Reader) bool {
	k := readNums(reader)[1]
	s := readString(reader)
	return solve(k, s)
}

func solve(k int, s string) bool {
	if k == 0 {
		for i := 0; i < len(s); i++ {
			if s[i] == 'N' {
				return false
			}
		}
		return true
	}

	dp := make([][2]bool, k+1)
	ndp := make([][2]bool, k+1)
	dp[0][0] = true

	for i := 0; i < len(s); i++ {
		for j := 0; j <= k; j++ {
			ndp[j][0] = false
			ndp[j][1] = false
		}

		for run := 0; run <= k; run++ {
			for hasK := range 2 {
				if !dp[run][hasK] {
					continue
				}

				if s[i] != 'Y' && run < k {
					nextHasK := hasK
					if run+1 == k {
						nextHasK = 1
					}
					ndp[run+1][nextHasK] = true
				}

				if s[i] != 'N' {
					ndp[0][hasK] = true
				}
			}
		}

		dp, ndp = ndp, dp
	}

	for run := 0; run <= k; run++ {
		if dp[run][1] {
			return true
		}
	}
	return false
}

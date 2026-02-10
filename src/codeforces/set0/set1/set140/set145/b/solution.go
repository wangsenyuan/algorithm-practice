package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (res string) {
	cnt := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Fscan(reader, &cnt[i])
	}
	res = solve(cnt)
	return
}

func solve(cnt []int) string {
	// cnt[4] = a1, cnt[7] = a2, cnt[47] = a3, cnt[74] = a4
	if cnt[2] > min(cnt[0], cnt[1]) {
		return "-1"
	}
	if cnt[3] > min(cnt[0], cnt[1]) {
		return "-1"
	}

	if abs(cnt[2]-cnt[3]) > 1 {
		return "-1"
	}

	if cnt[0] > 0 && cnt[1] > 0 && cnt[2]+cnt[3] == 0 {
		// 没有分界线
		return "-1"
	}

	if cnt[2] > cnt[3] {
		// 474747
		res := strings.Repeat("4", cnt[0]-cnt[2])
		res += strings.Repeat("47", cnt[2])
		cnt[1] -= cnt[2]
		if cnt[1] < 0 {
			return "-1"
		}
		if cnt[1] > 0 {
			res += strings.Repeat("7", cnt[1])
		}
		return res
	}

	if cnt[2] == cnt[3] {
		if cnt[0] >= cnt[2]+1 {
			// 474
			res := strings.Repeat("4", cnt[0]-cnt[2]-1)
			res += strings.Repeat("47", cnt[2])
			cnt[1] -= cnt[2]
			if cnt[1] > 0 {
				res += strings.Repeat("7", cnt[1])
			}
			res += "4"
			return res
		}
		// cnt[0] < cnt[2] + 1 => cnt[0] = cnt[2]
		// 47474747
		// 7474747474
		res := "7" + strings.Repeat("47", cnt[0])
		cnt[1] -= cnt[0] + 1
		if cnt[1] < 0 {
			return "-1"
		}
		res += strings.Repeat("7", cnt[1])
		return res
	}
	// cnt[2] < cnt[3], 747474
	if cnt[2] == 0 {
		return "-1"
	}
	res := "74" + strings.Repeat("4", cnt[0]-(cnt[2]+1))
	res += strings.Repeat("74", cnt[2]-1)
	// res += "74"
	cnt[1] -= cnt[2]
	if cnt[1] < 1 {
		return "-1"
	}
	res += strings.Repeat("7", cnt[1])
	res += "4"
	return res
}

func abs(num int) int {
	return max(num, -num)
}

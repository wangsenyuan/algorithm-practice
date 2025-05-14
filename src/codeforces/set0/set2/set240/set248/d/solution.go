package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func process(reader *bufio.Reader) int {
	_, k := readTwoNums(reader)
	s := readString(reader)
	return solve(k, s)
}

func solve(t int, s string) int {
	n := len(s)
	var house []int
	var shop []int
	for i := range n {
		if s[i] == 'H' {
			house = append(house, i)
		} else if s[i] == 'S' {
			shop = append(shop, i)
		}
	}

	u := len(house)
	v := len(shop)

	check := func(k int) bool {
		// 提前准备k个礼物时，是否能够在t时间内完成
		if k+v < u {
			return false
		}
		// k + shot[n] >= house[n]
		var ph, ps int
		cur_time := 1
		for i := 0; i < n && cur_time <= t; i++ {

			if ph < u && house[ph] == i {
				if k == 0 {
					// 这个位置要么使用策略2，到达最后的位置，然后返回回来
					// 要么使用策略1, 找到下一个shop, 然后返回
					// 一共有这么多个房子， 所以需要这么多个shop去购买礼物
					cnt := u - ph
					if ps+cnt <= v {
						j := max(shop[ps+cnt-1], house[u-1])
						// 先到达位置j，购买这个区间内所有的礼物，然后在返回的路上分配礼物
						if 2*(j-i)+cur_time <= t {
							// 这个策略可行
							return true
						}
					}
					// 策略2不可行，只能使用策略1
					// 但是如何使用策略1，就比较麻烦了
					w := 1
					for ph+w <= u && ps+w <= v {
						// 到达某个shop的位置 ps + w
						if ph+w == u || shop[ps+w-1] < house[ph+w] {
							break
						}
						// 否则 shop[ps+w] > house[ph+w]
						w++
					}
					if ph+w > u {
						return false
					}
					cur_time += 2 * (shop[ps+w-1] - i)
					ph += w
					ps += w
				} else {
					k--
					ph++
				}
				if ph == u {
					break
				}
			} else if ps < v && shop[ps] == i {
				k++
				ps++
			}
			// 移动到下一个位置
			cur_time++
		}
		return cur_time <= t
	}

	res := sort.Search(n+1, check)
	if res > n {
		return -1
	}
	return res
}

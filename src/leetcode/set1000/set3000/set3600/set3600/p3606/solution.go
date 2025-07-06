package p3606

import (
	"slices"
	"strings"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	order := map[string]int{
		"electronics": 1,
		"grocery":     2,
		"pharmacy":    3,
		"restaurant":  4,
	}

	type Coupon struct {
		id    string
		blOrd int
	}

	var arr []Coupon

	for i := range len(code) {
		if !isActive[i] || !checkCode(code[i]) || order[businessLine[i]] == 0 {
			continue
		}
		arr = append(arr, Coupon{
			id:    code[i],
			blOrd: order[businessLine[i]],
		})
	}

	slices.SortFunc(arr, func(a, b Coupon) int {
		if a.blOrd != b.blOrd {
			return a.blOrd - b.blOrd
		}
		return strings.Compare(a.id, b.id)
	})

	var ans []string
	for _, cur := range arr {
		ans = append(ans, cur.id)
	}

	return ans
}

func checkCode(code string) bool {
	code = strings.Trim(code, " ")
	if code == "" {
		return false
	}
	for i := range len(code) {
		if code[i] == '_' ||
			code[i] >= 'a' && code[i] <= 'z' ||
			code[i] >= 'A' && code[i] <= 'Z' ||
			code[i] >= '0' && code[i] <= '9' {
			continue
		}
		return false
	}
	return true
}

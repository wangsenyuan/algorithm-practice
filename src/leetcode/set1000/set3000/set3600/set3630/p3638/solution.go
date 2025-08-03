package p3638

func maxBalancedShipments(weight []int) int {
	n := len(weight)
	// w[i] 如果是一个[l...i]中的最小值
	// ans[i] = ans[l-1] + 1
	// 如果i是目前最小的那个，那么它就是第一个分组
	var res int
	for i := 0; i+1 < n; {
		if weight[i] > weight[i+1] {
			res++
			i += 2
		} else {
			i++
		}
	}
	return res
}

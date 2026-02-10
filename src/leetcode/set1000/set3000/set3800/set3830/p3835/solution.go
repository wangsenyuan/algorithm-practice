package p3835

func countSubarrays(nums []int, k int64) int64 {
	// 滑动窗口
	var st1 []int
	var st2 []int
	var res int

	var l int
	for r, v := range nums {
		for len(st1) > 0 && nums[st1[len(st1)-1]] < v {
			st1 = st1[:len(st1)-1]
		}
		// 降序队列，
		st1 = append(st1, r)
		for len(st2) > 0 && nums[st2[len(st2)-1]] > v {
			st2 = st2[:len(st2)-1]
		}
		// 升序队列
		st2 = append(st2, r)

		for l < r && (nums[st1[0]]-nums[st2[0]])*(r-l+1) > int(k) {
			if st1[0] == l {
				st1 = st1[1:]
			}
			if st2[0] == l {
				st2 = st2[1:]
			}
			l++
		}

		res += r - l + 1
	}

	return int64(res)
}

package p3834

func mergeAdjacent(nums []int) []int64 {
	var st []int64

	for _, v := range nums {
		u := int64(v)
		for len(st) > 0 && st[len(st)-1] == u {
			u *= 2
			st = st[:len(st)-1]
		}
		st = append(st, u)
	}
	return st
}

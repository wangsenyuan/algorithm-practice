package p3668

import "sort"

func recoverOrder(order []int, friends []int) []int {
	n := len(order)
	pos := make([]int, n)
	for i := range n {
		pos[order[i]-1] = i
	}
	sort.Slice(friends, func(i int, j int) bool {
		return pos[friends[i]-1] < pos[friends[j]-1]
	})
	return friends
}

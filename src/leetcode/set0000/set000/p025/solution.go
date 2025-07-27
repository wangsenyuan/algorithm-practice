package p025

import "learn-go/src/leetcode/util"

type ListNode = util.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	if head == nil {
		return nil
	}

	var cnt int

	for tmp := head; tmp != nil && cnt < k; tmp = tmp.Next {
		cnt++
	}

	if cnt < k {
		return head
	}

	cur := head
	next := head.Next
	head.Next = nil
	for cnt > 1 && next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
		cnt--
	}

	if next != nil {
		head.Next = reverseKGroup(next, k)
	}

	return cur
}

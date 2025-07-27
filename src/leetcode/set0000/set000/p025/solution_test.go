package p025

import (
	"learn-go/src/leetcode/util"
	"testing"
)

func runSample(t *testing.T, nums string, k int, expect string) {
	head := util.ParseAsList(nums)

	head = reverseKGroup(head, k)

	res := head.String()
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "[1,2,3,4,5]", 2, "[2,1,4,3,5]")
}

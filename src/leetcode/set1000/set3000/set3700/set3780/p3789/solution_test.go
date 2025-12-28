package p3789

import "testing"


func runSample(t *testing.T, cost1 int, cost2 int, costBoth int, need1 int, need2 int, expect int) {
	res := minimumCost(cost1, cost2, costBoth, need1, need2)
	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cost1 := 3
	cost2 := 2
	costBoth := 1
	need1 := 3
	need2 := 2
	expect := 3
	runSample(t, cost1, cost2, costBoth, need1, need2, expect)
}

func TestSample2(t *testing.T) {
	cost1 := 5
	cost2 := 4
	costBoth := 15
	need1 := 2
	need2 := 3
	expect :=22
	runSample(t, cost1, cost2, costBoth, need1, need2, expect)
}
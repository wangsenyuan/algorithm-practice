package main

import "testing"

func runSample(t *testing.T, n, m, k int, expect bool) {
	res := solve(n, m, k)
	if res != expect {
		t.Errorf("For n=%d m=%d k=%d: expect %t, got %t", n, m, k, expect, res)
	}
}

func TestSampleK2(t *testing.T) {
	// Examples from problem with k=2
	runSample(t, 1, 1, 2, false) // -
	runSample(t, 1, 2, 2, true)  // +
	runSample(t, 2, 1, 2, true)  // +
	runSample(t, 2, 2, 2, false) // -
	runSample(t, 1, 3, 2, false) // -
	runSample(t, 2, 3, 2, true)  // +
	runSample(t, 3, 1, 2, false) // -
	runSample(t, 3, 2, 2, true)  // +
	runSample(t, 3, 3, 2, true)  // +
	runSample(t, 4, 3, 2, true)  // +
}

func TestLargeK2(t *testing.T) {
	// Large test cases with k=2
	testCases := []struct {
		n, m  int
		expect bool
	}{
		{510468671, 885992046, false}, // -
		{560815578, 528820107, true},  // +
		{345772456, 904487274, true},  // +
		{614145524, 401148549, true},  // +
		{69821620, 761688435, false},  // -
		{555492073, 917117841, false}, // -
		{564072248, 108648678, true},  // +
		{693378609, 918514834, true},  // +
		{845237550, 637668983, false}, // -
		{161488522, 628798147, false}, // -
		{419601502, 805741370, true},  // +
		{481266866, 629833243, true},  // +
		{146688204, 563993768, true},  // +
		{131596024, 467885502, true},  // +
		{76313829, 591791856, true},   // +
		{793172605, 52217114, true},   // +
		{390118935, 365851302, true},  // +
		{46386762, 129675294, true},   // +
		{466969492, 271006705, true},  // +
		{134628830, 847304521, true},  // +
	}

	for i, tc := range testCases {
		runSample(t, tc.n, tc.m, 2, tc.expect)
		if t.Failed() {
			t.Logf("Failed at test case %d: n=%d m=%d", i+1, tc.n, tc.m)
			break
		}
	}
}

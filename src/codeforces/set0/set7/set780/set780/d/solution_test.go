package main

import "testing"

func runSample(t *testing.T, teams []string, expect bool) {
	ok, res := solve(teams)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !expect {
		return
	}

	names := getNameByRules(teams)

	all := make(map[string]int)
	for i, name := range res {
		all[name] = i
	}
	n := len(teams)

	if len(all) != n {
		t.Fatalf("%v, have duplicates", res)
	}

	for i := 0; i < n; i++ {
		if res[i] == names[i][1] {
			// 根据规则2生成的，那么就不能有规则1的名称出现
			if j, ok := all[names[i][0]]; ok && j != i && res[j] != names[j][1] {
				t.Fatalf("Sample result %v, violates rule 1", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	teams := []string{
		"DINAMO BYTECITY",
		"FOOTBALL MOSCOW",
	}
	expect := true
	runSample(t, teams, expect)
}

func TestSample2(t *testing.T) {
	teams := []string{
		"DINAMO BITECITY",
		"DINAMO BITECITY",
	}
	expect := false
	runSample(t, teams, expect)
}

func TestSample3(t *testing.T) {
	teams := []string{
		"PLAYFOOTBALL MOSCOW",
		"PLAYVOLLEYBALL SPB",
		"GOGO TECHNOCUP",
	}
	expect := true
	runSample(t, teams, expect)
}

func TestSample4(t *testing.T) {
	teams := []string{
		"ADAC BABC",
		"ABB DCB",
		"ABB BCDC",
		"DBAC BAC",
		"DBBC DBC",
	}
	expect := true
	runSample(t, teams, expect)
}

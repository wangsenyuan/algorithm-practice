package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFixtureFile(t *testing.T, root, path, content string) {
	t.Helper()
	full := filepath.Join(root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		t.Fatalf("mkdir fixture dir: %v", err)
	}
	if err := os.WriteFile(full, []byte(content), 0644); err != nil {
		t.Fatalf("write fixture file: %v", err)
	}
}

func makeFixtureRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	files := map[string]string{
		"src/codeforces/set1/set18/set185/set1857/g/solution.go":               "package main\n",
		"src/codeforces/set1/set18/set185/set1857/g/problem.md":                "# Problem\n",
		"src/atcoders/arc/arc100/arc127/d/solution.go":                         "package main\n",
		"src/unknown/demo/solution.go":                                         "package main\n",
		"docs/index/README.md":                                                 "old generated content\n",
		"src/codeforces/set1/set18/set185/set1857/g/ignored_solution.go":       "package main\n",
		"src/codeforces/set1/set18/set185/set1857/g/solution_test.go":          "package main\n",
		"src/codeforces/set1/set18/set185/set1857/g/nested/solution.go.backup": "package main\n",
	}
	for path, content := range files {
		writeFixtureFile(t, root, path, content)
	}
	return root
}

func TestCollectRepoCountsSolutionsDocsAndPlatforms(t *testing.T) {
	root := makeFixtureRepo(t)

	index, err := collectRepo(root)
	if err != nil {
		t.Fatalf("collectRepo() error = %v", err)
	}

	if index.TotalSolutions != 3 {
		t.Fatalf("TotalSolutions = %d, want 3", index.TotalSolutions)
	}
	if index.TotalDocs != 1 {
		t.Fatalf("TotalDocs = %d, want 1", index.TotalDocs)
	}

	wantCounts := map[Platform]int{
		platformCodeforces: 1,
		platformAtcoder:    1,
	}
	for platform, want := range wantCounts {
		if got := len(index.Platforms[platform].Entries); got != want {
			t.Fatalf("%s entries = %d, want %d", platform.Name, got, want)
		}
	}
}

func TestRenderIndexHomeLinksLearningTimeline(t *testing.T) {
	root := makeFixtureRepo(t)
	index, err := collectRepo(root)
	if err != nil {
		t.Fatalf("collectRepo() error = %v", err)
	}

	out := renderIndexHome(index)

	if !strings.Contains(out, "[Learning timeline](../timeline/README.md)") {
		t.Fatalf("renderIndexHome() missing learning timeline link in:\n%s", out)
	}
}

func TestRenderPlatformLinksPackageAndDocs(t *testing.T) {
	root := makeFixtureRepo(t)
	index, err := collectRepo(root)
	if err != nil {
		t.Fatalf("collectRepo() error = %v", err)
	}

	out := renderPlatformShard(index.Platforms[platformCodeforces], "set1/set18", "../../../")

	required := []string{
		"# Codeforces set1/set18",
		"[`src/codeforces/set1/set18/set185/set1857/g`](../../../src/codeforces/set1/set18/set185/set1857/g)",
		"[problem.md](../../../src/codeforces/set1/set18/set185/set1857/g/problem.md)",
	}
	for _, text := range required {
		if !strings.Contains(out, text) {
			t.Fatalf("renderPlatformShard() missing %q in:\n%s", text, out)
		}
	}
}

func TestWriteIndexesSplitsLargePlatformIndexesIntoShardFiles(t *testing.T) {
	root := makeFixtureRepo(t)
	index, err := collectRepo(root)
	if err != nil {
		t.Fatalf("collectRepo() error = %v", err)
	}

	if err := writeIndexes(root, index); err != nil {
		t.Fatalf("writeIndexes() error = %v", err)
	}
	landing, err := os.ReadFile(filepath.Join(root, "docs", "index", "codeforces.md"))
	if err != nil {
		t.Fatalf("read Codeforces landing page: %v", err)
	}
	shardPath := filepath.Join(root, "docs", "index", "codeforces", "set1-set18.md")
	first, err := os.ReadFile(shardPath)
	if err != nil {
		t.Fatalf("read Codeforces shard: %v", err)
	}
	if !strings.Contains(string(landing), "[set1/set18](codeforces/set1-set18.md)") {
		t.Fatalf("Codeforces landing page does not link to shard:\n%s", landing)
	}
	if _, err := os.Stat(filepath.Join(root, "docs", "index", "leetcode.md")); !os.IsNotExist(err) {
		t.Fatalf("expected no LeetCode index after writeIndexes, err=%v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "docs", "index", "codechef.md")); !os.IsNotExist(err) {
		t.Fatalf("expected no CodeChef index after writeIndexes, err=%v", err)
	}

	if err := writeIndexes(root, index); err != nil {
		t.Fatalf("second writeIndexes() error = %v", err)
	}
	second, err := os.ReadFile(shardPath)
	if err != nil {
		t.Fatalf("read second generated shard: %v", err)
	}

	if string(first) != string(second) {
		t.Fatalf("generated output changed between runs")
	}
}

package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Platform struct {
	Key     string
	Name    string
	Prefix  string
	OutFile string
}

type Entry struct {
	Path string
	Docs []string
}

type PlatformIndex struct {
	Platform Platform
	Entries  []Entry
}

type RepoIndex struct {
	TotalSolutions int
	TotalDocs      int
	Platforms      map[Platform]PlatformIndex
}

var (
	platformCodeforces = Platform{Key: "codeforces", Name: "Codeforces", Prefix: "src/codeforces/", OutFile: "codeforces.md"}
	platformLeetcode   = Platform{Key: "leetcode", Name: "LeetCode", Prefix: "src/leetcode/", OutFile: "leetcode.md"}
	platformCodechef   = Platform{Key: "codechef", Name: "CodeChef", Prefix: "src/codechef/", OutFile: "codechef.md"}
	platformAtcoder    = Platform{Key: "atcoder", Name: "AtCoder", Prefix: "src/atcoders/", OutFile: "atcoder.md"}
	platforms          = []Platform{platformCodeforces, platformLeetcode, platformCodechef, platformAtcoder}
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	index, err := collectRepo(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "collect repo: %v\n", err)
		os.Exit(1)
	}
	if err := writeIndexes(root, index); err != nil {
		fmt.Fprintf(os.Stderr, "write indexes: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Indexed %d solutions and %d docs\n", index.TotalSolutions, index.TotalDocs)
}

func collectRepo(root string) (RepoIndex, error) {
	index := RepoIndex{
		Platforms: make(map[Platform]PlatformIndex, len(platforms)),
	}
	for _, platform := range platforms {
		index.Platforms[platform] = PlatformIndex{Platform: platform}
	}

	docsByDir := make(map[string][]string)
	var solutionDirs []string

	srcRoot := filepath.Join(root, "src")
	err := filepath.WalkDir(srcRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		name := d.Name()
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		switch {
		case name == "solution.go":
			index.TotalSolutions++
			solutionDirs = append(solutionDirs, filepath.ToSlash(filepath.Dir(rel)))
		case isDocName(name):
			index.TotalDocs++
			dir := filepath.ToSlash(filepath.Dir(rel))
			docsByDir[dir] = append(docsByDir[dir], rel)
		}
		return nil
	})
	if err != nil {
		return RepoIndex{}, err
	}

	sort.Strings(solutionDirs)
	for _, dir := range solutionDirs {
		platform, ok := platformForPath(dir)
		if !ok {
			continue
		}
		docs := append([]string(nil), docsByDir[dir]...)
		sort.Strings(docs)

		platformIndex := index.Platforms[platform]
		platformIndex.Entries = append(platformIndex.Entries, Entry{
			Path: dir,
			Docs: docs,
		})
		index.Platforms[platform] = platformIndex
	}

	return index, nil
}

func isDocName(name string) bool {
	return name == "problem.md" || name == "readme.md" || name == "README.md"
}

func platformForPath(path string) (Platform, bool) {
	for _, platform := range platforms {
		if strings.HasPrefix(path+"/", platform.Prefix) {
			return platform, true
		}
	}
	return Platform{}, false
}

func renderIndexHome(index RepoIndex) string {
	var buf bytes.Buffer
	buf.WriteString("# Solution Index\n\n")
	buf.WriteString("Generated Markdown indexes for the main online judge platforms in this repository.\n\n")
	buf.WriteString("## Stats\n\n")
	buf.WriteString("| Metric | Count |\n")
	buf.WriteString("| --- | ---: |\n")
	fmt.Fprintf(&buf, "| Total solutions | %d |\n", index.TotalSolutions)
	fmt.Fprintf(&buf, "| Local docs | %d |\n", index.TotalDocs)
	for _, platform := range platforms {
		fmt.Fprintf(&buf, "| %s solutions | %d |\n", platform.Name, len(index.Platforms[platform].Entries))
	}
	buf.WriteString("\n## Platform Indexes\n\n")
	for _, platform := range platforms {
		fmt.Fprintf(&buf, "- [%s](%s)\n", platform.Name, platform.OutFile)
	}
	buf.WriteString("\nRegenerate with:\n\n")
	buf.WriteString("```bash\n")
	buf.WriteString("go run ./tools/indexgen\n")
	buf.WriteString("```\n")
	return buf.String()
}

func renderPlatform(index RepoIndex, platform Platform) string {
	platformIndex := index.Platforms[platform]
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "# %s Index\n\n", platform.Name)
	fmt.Fprintf(&buf, "%d packages with `solution.go` under `%s`.\n\n", len(platformIndex.Entries), strings.TrimSuffix(platform.Prefix, "/"))
	buf.WriteString("| Package | Local docs |\n")
	buf.WriteString("| --- | --- |\n")
	for _, entry := range platformIndex.Entries {
		fmt.Fprintf(&buf, "| [`%s`](../../%s) | %s |\n", entry.Path, entry.Path, renderDocLinks(entry.Docs))
	}
	return buf.String()
}

func renderDocLinks(docs []string) string {
	if len(docs) == 0 {
		return ""
	}
	links := make([]string, 0, len(docs))
	for _, doc := range docs {
		links = append(links, fmt.Sprintf("[%s](../../%s)", filepath.Base(doc), doc))
	}
	return strings.Join(links, ", ")
}

func writeIndexes(root string, index RepoIndex) error {
	outDir := filepath.Join(root, "docs", "index")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	files := map[string]string{
		"README.md": renderIndexHome(index),
	}
	for _, platform := range platforms {
		files[platform.OutFile] = renderPlatform(index, platform)
	}

	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		path := filepath.Join(outDir, name)
		if err := os.WriteFile(path, []byte(files[name]), 0644); err != nil {
			return err
		}
	}
	return nil
}

import json
import os
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path


DISCOVER = Path(__file__).with_name("discover.py")


class GitRepo:
    def __init__(self, root):
        self.root = Path(root)
        self.git("init", "-q")
        self.git("config", "user.name", "Timeline Test")
        self.git("config", "user.email", "timeline@example.com")

    def git(self, *args, check=True):
        return subprocess.run(
            ["git", *args],
            cwd=self.root,
            check=check,
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )

    def write(self, relative, contents="package main\n"):
        path = self.root / relative
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(contents, encoding="utf-8")

    def commit_all(self, message):
        self.git("add", ".")
        self.git("commit", "-q", "-m", message)
        return self.git("rev-parse", "HEAD").stdout.strip()

    def discover(self, env=None):
        return subprocess.run(
            [sys.executable, str(DISCOVER), "--repo", str(self.root)],
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            env=env,
        )


class DiscoverTimelineTest(unittest.TestCase):
    def make_repo(self):
        temporary = tempfile.TemporaryDirectory()
        self.addCleanup(temporary.cleanup)
        return GitRepo(temporary.name)

    def assert_success(self, result):
        self.assertEqual(result.returncode, 0, result.stderr)
        return json.loads(result.stdout)

    def test_first_run_uses_parent_of_latest_solution_addition(self):
        repo = self.make_repo()
        repo.write("README.md", "practice\n")
        initial = repo.commit_all("initial")

        repo.write("src/codeforces/new/a/solution.go")
        repo.write("src/codeforces/new/a/solution_test.go", "package main\n")
        repo.commit_all("add solution")
        repo.write("notes.md", "later documentation\n")
        head = repo.commit_all("document it")

        payload = self.assert_success(repo.discover())

        self.assertEqual(payload["baseline"], initial)
        self.assertEqual(payload["head"], head)
        self.assertEqual(
            payload["packages"],
            [
                {
                    "path": "src/codeforces/new/a",
                    "origins": ["committed"],
                    "files": [
                        "src/codeforces/new/a/solution.go",
                        "src/codeforces/new/a/solution_test.go",
                    ],
                }
            ],
        )

    def test_repo_argument_overrides_git_repository_environment(self):
        target = self.make_repo()
        target.write("README.md", "target\n")
        baseline = target.commit_all("target initial")
        target.write("src/target/a/solution.go")
        target_head = target.commit_all("target solution")

        decoy = self.make_repo()
        decoy.write("README.md", "decoy\n")
        decoy.commit_all("decoy initial")
        decoy.write("src/decoy/b/solution.go")
        decoy.commit_all("decoy solution")

        environment = os.environ.copy()
        environment.update(
            {
                "GIT_DIR": str(decoy.root / ".git"),
                "GIT_WORK_TREE": str(decoy.root),
                "GIT_INDEX_FILE": str(decoy.root / ".git" / "index"),
            }
        )

        payload = self.assert_success(target.discover(env=environment))

        self.assertEqual(payload["baseline"], baseline)
        self.assertEqual(payload["head"], target_head)
        self.assertEqual(
            [package["path"] for package in payload["packages"]],
            ["src/target/a"],
        )

    def test_first_run_supports_solution_added_in_root_commit(self):
        repo = self.make_repo()
        repo.write("src/root/a/solution.go")
        head = repo.commit_all("root solution")

        payload = self.assert_success(repo.discover())

        self.assertEqual(
            payload["baseline"],
            "4b825dc642cb6eb9a060e54bf8d69288fbee4904",
        )
        self.assertEqual(payload["head"], head)
        self.assertEqual(
            [package["path"] for package in payload["packages"]],
            ["src/root/a"],
        )

    def test_marker_combines_committed_staged_and_untracked_additions(self):
        repo = self.make_repo()
        repo.write(".gitignore", "src/ignored/\n")
        repo.write("src/existing/a/solution.go")
        marker = repo.commit_all("initial")
        repo.write(
            "docs/timeline/README.md",
            f"# Timeline\n\n<!-- through-commit: {marker} -->\n",
        )
        repo.commit_all("add marker")

        repo.write("src/committed/z/solution.go")
        repo.write("src/committed/z/problem.md", "problem\n")
        repo.write("src/committed/z/README.md", "read me\n")
        repo.commit_all("committed package")

        repo.write("src/existing/a/solution.go", "package main\n// modified\n")
        repo.write("src/staged/b/solution.go")
        repo.write("src/staged/b/readme.md", "lowercase\n")
        repo.git("add", "src/staged/b", "src/existing/a/solution.go")

        repo.write("src/untracked/c/solution.go")
        repo.write("src/untracked/c/solution_test.go", "package main\n")
        repo.write("src/ignored/d/solution.go")
        repo.write("scratch.txt", "unrelated\n")

        payload = self.assert_success(repo.discover())

        self.assertEqual(payload["baseline"], marker)
        self.assertEqual(
            payload["packages"],
            [
                {
                    "path": "src/committed/z",
                    "origins": ["committed"],
                    "files": [
                        "src/committed/z/solution.go",
                        "src/committed/z/problem.md",
                        "src/committed/z/README.md",
                    ],
                },
                {
                    "path": "src/staged/b",
                    "origins": ["staged"],
                    "files": [
                        "src/staged/b/solution.go",
                        "src/staged/b/readme.md",
                    ],
                },
                {
                    "path": "src/untracked/c",
                    "origins": ["untracked"],
                    "files": [
                        "src/untracked/c/solution.go",
                        "src/untracked/c/solution_test.go",
                    ],
                },
            ],
        )

    def test_excludes_package_linked_from_dated_timeline(self):
        repo = self.make_repo()
        repo.write("README.md", "practice\n")
        marker = repo.commit_all("initial")
        repo.write(
            "docs/timeline/README.md",
            f"<!-- through-commit: {marker} -->\n",
        )
        repo.commit_all("marker")
        repo.write("src/codeforces/set2/set20/set2000/a/solution.go")
        repo.write(
            "docs/timeline/2026-07-23.md",
            "[Solved](../../src/codeforces/set2/set20/set2000/a/)\n",
        )
        repo.commit_all("add documented package")

        payload = self.assert_success(repo.discover())

        self.assertEqual(payload["packages"], [])

    def test_dated_timeline_dedupes_only_markdown_links(self):
        repo = self.make_repo()
        repo.write("README.md", "practice\n")
        marker = repo.commit_all("initial")
        repo.write(
            "docs/timeline/README.md",
            f"<!-- through-commit: {marker} -->\n",
        )
        repo.commit_all("marker")
        for package in ("a", "b", "c"):
            repo.write(f"src/x/{package}/solution.go")
        repo.write(
            "docs/timeline/2026-07-24.md",
            "\n".join(
                [
                    "[A](../../src/x/a/)",
                    "[B](../../src/x/b)",
                    "Plain text ../../src/x/c/ is not a link.",
                ]
            ),
        )
        repo.commit_all("add packages and timeline references")

        payload = self.assert_success(repo.discover())

        self.assertEqual(
            [package["path"] for package in payload["packages"]],
            ["src/x/c"],
        )

    def test_invalid_marker_fails_without_mutating_repository(self):
        repo = self.make_repo()
        repo.write(
            "docs/timeline/README.md",
            "<!-- through-commit: 0000000000000000000000000000000000000000 -->\n",
        )
        repo.commit_all("invalid marker")
        before_status = repo.git("status", "--porcelain=v1", "-uall").stdout
        before_head = repo.git("rev-parse", "HEAD").stdout

        result = repo.discover()

        self.assertEqual(result.returncode, 2)
        self.assertIn("discover timeline:", result.stderr)
        self.assertIn("invalid through-commit", result.stderr)
        self.assertEqual(
            repo.git("status", "--porcelain=v1", "-uall").stdout, before_status
        )
        self.assertEqual(repo.git("rev-parse", "HEAD").stdout, before_head)

    def test_duplicate_markers_are_rejected_as_ambiguous(self):
        for conflicting in (False, True):
            with self.subTest(conflicting=conflicting):
                repo = self.make_repo()
                repo.write("README.md", "practice\n")
                first = repo.commit_all("initial")
                repo.write("notes.md", "next\n")
                second = repo.commit_all("second")
                other = second if conflicting else first
                repo.write(
                    "docs/timeline/README.md",
                    "\n".join(
                        [
                            f"<!-- through-commit: {first} -->",
                            f"<!-- through-commit: {other} -->",
                        ]
                    ),
                )
                repo.commit_all("duplicate markers")

                result = repo.discover()

                self.assertEqual(result.returncode, 2)
                self.assertIn("discover timeline:", result.stderr)
                self.assertIn("ambiguous through-commit", result.stderr)

    def test_usage_errors_follow_cli_error_contract(self):
        result = subprocess.run(
            [sys.executable, str(DISCOVER)],
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )

        self.assertEqual(result.returncode, 2)
        self.assertTrue(
            result.stderr.startswith("discover timeline:"),
            result.stderr,
        )


if __name__ == "__main__":
    unittest.main()

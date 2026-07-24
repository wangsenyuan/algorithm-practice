#!/usr/bin/env python3
"""Discover newly added solution packages for a learning timeline."""

import argparse
import json
import os
import re
import subprocess
import sys
from pathlib import Path, PurePosixPath


MARKER = re.compile(r"<!--\s*through-commit:\s*([0-9a-fA-F]{40})\s*-->")
DATED_TIMELINE = re.compile(r"\d{4}-\d{2}-\d{2}\.md")
LINKED_PACKAGE = re.compile(r"\.\./\.\./(src/[^)\s]+?)/(?=[)\s])")
FILE_NAMES = (
    "solution.go",
    "solution_test.go",
    "problem.md",
    "readme.md",
    "README.md",
)
SOLUTION_PATHSPEC = ":(glob)src/**/solution.go"


class DiscoveryError(Exception):
    pass


class TimelineArgumentParser(argparse.ArgumentParser):
    def error(self, message):
        self.exit(2, f"discover timeline: {message}\n")


def run_git(repo, *args, check=True):
    environment = os.environ.copy()
    environment["GIT_OPTIONAL_LOCKS"] = "0"
    result = subprocess.run(
        ["git", "-C", str(repo), *args],
        check=False,
        text=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        env=environment,
    )
    if check and result.returncode:
        detail = result.stderr.strip() or "git command failed"
        raise DiscoveryError(detail)
    return result


def read_marker(repo):
    marker_file = repo / "docs" / "timeline" / "README.md"
    if not marker_file.is_file():
        return None
    contents = marker_file.read_text(encoding="utf-8")
    match = MARKER.search(contents)
    if match:
        return match.group(1).lower()
    if "through-commit" in contents:
        raise DiscoveryError("invalid through-commit marker")
    return None


def validate_marker(repo, marker, head):
    commit = run_git(repo, "cat-file", "-e", f"{marker}^{{commit}}", check=False)
    ancestor = run_git(
        repo, "merge-base", "--is-ancestor", marker, head, check=False
    )
    if commit.returncode or ancestor.returncode:
        raise DiscoveryError(f"invalid through-commit: {marker}")


def first_run_baseline(repo, head):
    result = run_git(
        repo,
        "log",
        "-1",
        "--format=%H",
        "--diff-filter=A",
        head,
        "--",
        SOLUTION_PATHSPEC,
    )
    addition = result.stdout.strip()
    if not addition:
        return head
    parent = run_git(repo, "rev-parse", f"{addition}^", check=False)
    if parent.returncode:
        raise DiscoveryError(
            f"cannot find parent of first solution addition {addition}"
        )
    return parent.stdout.strip()


def added_solution_paths(repo, baseline):
    committed = run_git(
        repo,
        "diff",
        "--name-only",
        "--diff-filter=A",
        f"{baseline}..HEAD",
        "--",
        SOLUTION_PATHSPEC,
    ).stdout.splitlines()
    staged = run_git(
        repo,
        "diff",
        "--cached",
        "--name-only",
        "--diff-filter=A",
        "--",
        SOLUTION_PATHSPEC,
    ).stdout.splitlines()
    untracked = [
        path
        for path in run_git(
            repo, "ls-files", "--others", "--exclude-standard"
        ).stdout.splitlines()
        if is_solution_path(path)
    ]
    return (
        ("committed", committed),
        ("staged", staged),
        ("untracked", untracked),
    )


def is_solution_path(path):
    candidate = PurePosixPath(path)
    return (
        len(candidate.parts) >= 3
        and candidate.parts[0] == "src"
        and candidate.name == "solution.go"
    )


def linked_packages(repo):
    timeline = repo / "docs" / "timeline"
    linked = set()
    if not timeline.is_dir():
        return linked
    for path in timeline.iterdir():
        if path.is_file() and DATED_TIMELINE.fullmatch(path.name):
            contents = path.read_text(encoding="utf-8")
            linked.update(
                match.group(1).rstrip("/")
                for match in LINKED_PACKAGE.finditer(contents)
            )
    return linked


def discover(repo):
    repo = repo.resolve()
    head = run_git(repo, "rev-parse", "HEAD").stdout.strip()
    marker = read_marker(repo)
    if marker:
        validate_marker(repo, marker, head)
        baseline = marker
    else:
        baseline = first_run_baseline(repo, head)

    excluded = linked_packages(repo)
    packages = {}
    for origin, paths in added_solution_paths(repo, baseline):
        for solution in paths:
            if not is_solution_path(solution):
                continue
            package = str(PurePosixPath(solution).parent)
            if package in excluded:
                continue
            packages.setdefault(package, []).append(origin)

    output = []
    for package in sorted(packages):
        package_dir = repo / package
        existing_names = {
            entry.name for entry in package_dir.iterdir() if entry.is_file()
        }
        files = [
            f"{package}/{name}"
            for name in FILE_NAMES
            if name in existing_names
        ]
        output.append(
            {"path": package, "origins": packages[package], "files": files}
        )
    return {"baseline": baseline, "head": head, "packages": output}


def main(argv=None):
    parser = TimelineArgumentParser()
    parser.add_argument("--repo", required=True, type=Path)
    args = parser.parse_args(argv)
    try:
        payload = discover(args.repo)
    except (DiscoveryError, OSError, UnicodeError) as error:
        print(f"discover timeline: {error}", file=sys.stderr)
        return 2
    json.dump(payload, sys.stdout, indent=2)
    sys.stdout.write("\n")
    return 0


if __name__ == "__main__":
    sys.exit(main())

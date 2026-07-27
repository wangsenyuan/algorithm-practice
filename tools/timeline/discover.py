#!/usr/bin/env python3
"""Discover newly added solution packages for a learning timeline.

Symlink containment assumes a stable local filesystem for each invocation.
Stop instead of running discovery when concurrent mutation is possible.
"""

import argparse
import json
import os
import re
import stat
import subprocess
import sys
from pathlib import Path, PurePosixPath


MARKER = re.compile(r"<!--\s*through-commit:\s*([0-9a-fA-F]{40})\s*-->")
MARKER_DECLARATION = re.compile(r"through-commit\s*:")
DATED_TIMELINE = re.compile(r"\d{4}-\d{2}-\d{2}\.md")
LINKED_PACKAGE = re.compile(
    r"\[[^\]\r\n]*\]\(\.\./\.\./(src/[^)\s]+?)/?\)"
)
FILE_NAMES = (
    "solution.go",
    "solution_test.go",
    "problem.md",
    "readme.md",
    "README.md",
)
SOLUTION_PATHSPEC = ":(glob)src/**/solution.go"
EMPTY_TREE = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"
REPOSITORY_ENVIRONMENT = (
    "GIT_DIR",
    "GIT_WORK_TREE",
    "GIT_INDEX_FILE",
    "GIT_COMMON_DIR",
)
MARKDOWN_SAFE_PATH_PUNCTUATION = frozenset("/._-")
STABLE_FILESYSTEM_BOUNDARY = (
    "Symlink containment assumes a stable local filesystem for each "
    "invocation; stop if concurrent mutation is possible."
)


class DiscoveryError(Exception):
    pass


class TimelineArgumentParser(argparse.ArgumentParser):
    def error(self, message):
        self.exit(2, f"discover timeline: {message}\n")


def quote_path(path):
    return json.dumps(os.fspath(path), ensure_ascii=True)


def run_git(repo, *args, check=True, binary=False):
    environment = os.environ.copy()
    for name in REPOSITORY_ENVIRONMENT:
        environment.pop(name, None)
    environment["GIT_OPTIONAL_LOCKS"] = "0"
    result = subprocess.run(
        ["git", "-C", str(repo), *args],
        check=False,
        text=not binary,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        env=environment,
    )
    if check and result.returncode:
        detail = result.stderr.strip()
        if binary:
            detail = os.fsdecode(detail)
        detail = detail or "git command failed"
        raise DiscoveryError(detail)
    return result


def repo_relative(repo, path):
    try:
        return path.relative_to(repo)
    except ValueError as error:
        raise DiscoveryError(
            f"path escapes repository: {quote_path(path)}"
        ) from error


def reject_symlink_components(repo, path):
    relative = repo_relative(repo, path)
    current = repo
    for component in relative.parts:
        current /= component
        try:
            mode = current.lstat().st_mode
        except FileNotFoundError:
            break
        if stat.S_ISLNK(mode):
            raise DiscoveryError(
                "symlink path is not allowed: "
                f"{quote_path(relative.as_posix())}"
            )
    resolved = path.resolve(strict=False)
    try:
        resolved.relative_to(repo)
    except ValueError as error:
        raise DiscoveryError(
            f"path escapes repository: {quote_path(relative.as_posix())}"
        ) from error


def validate_timeline_paths(repo):
    docs = repo / "docs"
    timeline = docs / "timeline"
    marker = timeline / "README.md"
    for path in (docs, timeline, marker):
        reject_symlink_components(repo, path)
    try:
        entries = list(timeline.iterdir())
    except FileNotFoundError:
        return
    except NotADirectoryError as error:
        raise DiscoveryError("timeline path is not a directory") from error
    for path in entries:
        if DATED_TIMELINE.fullmatch(path.name):
            reject_symlink_components(repo, path)


def git_paths(repo, *args):
    result = run_git(repo, *args, binary=True)
    return [
        os.fsdecode(path)
        for path in result.stdout.split(b"\0")
        if path
    ]


def read_marker(repo):
    marker_file = repo / "docs" / "timeline" / "README.md"
    reject_symlink_components(repo, marker_file)
    if not marker_file.is_file():
        return None
    contents = marker_file.read_text(encoding="utf-8")
    declarations = MARKER_DECLARATION.findall(contents)
    if len(declarations) > 1:
        raise DiscoveryError("ambiguous through-commit markers")
    matches = MARKER.findall(contents)
    if matches:
        return matches[0].lower()
    if declarations or "through-commit" in contents:
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
    raw_commit = run_git(repo, "cat-file", "-p", addition).stdout
    parents = []
    for line in raw_commit.splitlines():
        if not line:
            break
        if line.startswith("parent "):
            parents.append(line.removeprefix("parent "))
    if not parents:
        return EMPTY_TREE
    parent = parents[0]
    available = run_git(
        repo, "cat-file", "-e", f"{parent}^{{commit}}", check=False
    )
    if available.returncode:
        raise DiscoveryError(
            f"shallow history: parent {parent} of {addition} is unavailable"
        )
    return parent


def added_solution_paths(repo, baseline):
    committed = git_paths(
        repo,
        "diff",
        "--name-only",
        "--diff-filter=A",
        "-z",
        f"{baseline}..HEAD",
        "--",
        SOLUTION_PATHSPEC,
    )
    staged = git_paths(
        repo,
        "diff",
        "--cached",
        "--name-only",
        "--diff-filter=A",
        "-z",
        "--",
        SOLUTION_PATHSPEC,
    )
    untracked = [
        path
        for path in git_paths(
            repo, "ls-files", "--others", "--exclude-standard", "-z"
        )
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


def validate_package_path(package):
    is_markdown_safe = all(
        character.isalnum()
        or character in MARKDOWN_SAFE_PATH_PUNCTUATION
        for character in package
    )
    if not is_markdown_safe:
        raise DiscoveryError(
            f"invalid package path: {quote_path(package)}"
        )


def linked_packages(repo):
    timeline = repo / "docs" / "timeline"
    linked = set()
    reject_symlink_components(repo, timeline)
    if not timeline.is_dir():
        return linked
    for path in timeline.iterdir():
        if not DATED_TIMELINE.fullmatch(path.name):
            continue
        reject_symlink_components(repo, path)
        if not stat.S_ISREG(path.lstat().st_mode):
            continue
        contents = path.read_text(encoding="utf-8")
        linked.update(
            match.group(1).rstrip("/")
            for match in LINKED_PACKAGE.finditer(contents)
        )
    return linked


def discover(repo):
    repo_argument = repo
    repo = repo.resolve()
    if not repo.is_dir():
        raise DiscoveryError(
            "repository directory does not exist: "
            f"{quote_path(repo_argument)}"
        )
    validate_timeline_paths(repo)
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
            validate_package_path(package)
            solution_path = repo / solution
            reject_symlink_components(repo, solution_path)
            try:
                solution_mode = solution_path.lstat().st_mode
            except FileNotFoundError as error:
                raise DiscoveryError(
                    f"added solution is missing: {quote_path(solution)}"
                ) from error
            if not stat.S_ISREG(solution_mode):
                raise DiscoveryError(
                    "added solution is not a regular file: "
                    f"{quote_path(solution)}"
                )
            if package in excluded:
                continue
            packages.setdefault(package, []).append(origin)

    output = []
    for package in sorted(packages):
        package_dir = repo / package
        reject_symlink_components(repo, package_dir)
        evidence_names = set()
        for evidence in package_dir.iterdir():
            if evidence.name not in FILE_NAMES:
                continue
            reject_symlink_components(repo, evidence)
            mode = evidence.lstat().st_mode
            if stat.S_ISREG(mode):
                evidence_names.add(evidence.name)
        files = [
            f"{package}/{name}" for name in FILE_NAMES if name in evidence_names
        ]
        output.append(
            {"path": package, "origins": packages[package], "files": files}
        )
    return {"baseline": baseline, "head": head, "packages": output}


def main(argv=None):
    parser = TimelineArgumentParser(description=STABLE_FILESYSTEM_BOUNDARY)
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

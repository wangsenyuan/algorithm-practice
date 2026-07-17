# C. Dependency Management

[Problem link](https://codeforces.com/problemset/problem/928/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Polycarp develops a project using a dependency manager (Vamen). Both the
project and libraries are treated as projects. Each project has a unique name
and a version number.

You are given project descriptions. The **first** project is Polycarp's.
Determine all projects that it depends on (directly or through a chain).

If the project depends on two different versions of the same name, collision
resolving applies: choose the version that minimizes the distance (chain
length) from Polycarp's project; on a tie, prefer the **newer** (larger)
version. That version is actual; other versions of the same name and their
dependencies are ignored.

## Constraints

- `1 ≤ n ≤ 1000` — number of projects
- Project names consist of lowercase Latin letters
- Versions are positive integers
- Dependencies refer to projects that appear in the input

## Input

```
n
name version
d
dep_name_1 dep_version_1
...
dep_name_d dep_version_d

name version
...
```

Projects are separated by one empty line. The first project is Polycarp's.
`d` is the number of direct dependencies (`0 ≤ d ≤ n-1`).

## Output

Print the number of distinct dependency projects (excluding Polycarp's
project), then those dependencies in lexicographical order of name, each as
`name version`.

## Samples

### Sample 1

Input:

```
4
a 3
2
b 1
c 1

b 2
0

b 1
1
b 2

c 1
1
b 2
```

Output:

```
2
b 1
c 1
```

### Sample 2

Input:

```
9
codehorses 5
3
webfrmk 6
mashadb 1
mashadb 2

commons 2
0

mashadb 3
0

webfrmk 6
2
mashadb 3
commons 2

extra 4
1
extra 3

extra 3
0

extra 1
0

mashadb 1
1
extra 3

mashadb 2
1
extra 1
```

Output:

```
4
commons 2
extra 1
mashadb 2
webfrmk 6
```

### Sample 3

Input:

```
3
abc 1
2
abc 3
cba 2

abc 3
0

cba 2
0
```

Output:

```
1
cba 2
```

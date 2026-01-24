# Problem C

There are some websites that are accessible through several different addresses. For example, for a long time Codeforces was accessible with two hostnames codeforces.com and codeforces.ru.

You are given a list of page addresses being queried. For simplicity we consider all addresses to have the form `http://<hostname>[/<path>]`, where:

- `<hostname>` — server name (consists of words and maybe some dots separating them),
- `/<path>` — optional part, where `<path>` consists of words separated by slashes.

We consider two `<hostname>` to correspond to one website if for each query to the first `<hostname>` there will be exactly the same query to the second one and vice versa — for each query to the second `<hostname>` there will be the same query to the first one. Take a look at the samples for further clarifications.

Your goal is to determine the groups of server names that correspond to one website. Ignore groups consisting of the only server name.

Please note, that according to the above definition queries `http://<hostname>` and `http://<hostname>/` are different.

## Input

The first line of the input contains a single integer n (1 ≤ n ≤ 100 000) — the number of page queries. Then follow n lines each containing exactly one address. Each address is of the form `http://<hostname>[/<path>]`, where:

- `<hostname>` consists of lowercase English letters and dots, there are no two consecutive dots, `<hostname>` doesn't start or finish with a dot. The length of `<hostname>` is positive and doesn't exceed 20.
- `<path>` consists of lowercase English letters, dots and slashes. There are no two consecutive slashes, `<path>` doesn't start with a slash and its length doesn't exceed 20.
- Addresses are not guaranteed to be distinct.

## Output

First print k — the number of groups of server names that correspond to one website. You should count only groups of size greater than one.

Next k lines should contain the description of groups, one group per line. For each group print all server names separated by a single space. You are allowed to print both groups and names inside any group in arbitrary order.

## Examples

### Example 1

**Input:**
```
10
http://abacaba.ru/test
http://abacaba.ru/
http://abacaba.com
http://abacaba.com/test
http://abacaba.de/
http://abacaba.ru/test
http://abacaba.de/test
http://abacaba.com/
http://abacaba.com/t
http://abacaba.com/test
```

**Output:**
```
1
http://abacaba.de http://abacaba.ru
```

### Example 2

**Input:**
```
14
http://c
http://ccc.bbbb/aba..b
http://cba.com
http://a.c/aba..b/a
http://abc/
http://a.c/
http://ccc.bbbb
http://ab.ac.bc.aa/
http://a.a.a/
http://ccc.bbbb/
http://cba.com/
http://cba.com/aba..b
http://a.a.a/aba..b/a
http://abc/aba..b/a
```

**Output:**
```
2
http://cba.com http://ccc.bbbb
http://a.a.a http://a.c http://abc
```


## Problem Explanation

### Key Concept

Two hostnames correspond to the **same website** if they have the **exact same set of queries (paths)**. This means:
- For every path queried on hostname A, the same path must be queried on hostname B
- For every path queried on hostname B, the same path must be queried on hostname A
- The sets of paths must be **identical** (order doesn't matter)

**Important:** `http://hostname` (no trailing slash) and `http://hostname/` (with trailing slash) are **different queries**.

### Example 2 Detailed Analysis

Let's break down the second example step by step:

**Input queries:**
1. `http://c` → hostname: `c`, path: `""` (empty, no path)
2. `http://ccc.bbbb/aba..b` → hostname: `ccc.bbbb`, path: `aba..b`
3. `http://cba.com` → hostname: `cba.com`, path: `""` (empty)
4. `http://a.c/aba..b/a` → hostname: `a.c`, path: `aba..b/a`
5. `http://abc/` → hostname: `abc`, path: `/` (just the slash)
6. `http://a.c/` → hostname: `a.c`, path: `/`
7. `http://ccc.bbbb` → hostname: `ccc.bbbb`, path: `""` (empty)
8. `http://ab.ac.bc.aa/` → hostname: `ab.ac.bc.aa`, path: `/`
9. `http://a.a.a/` → hostname: `a.a.a`, path: `/`
10. `http://ccc.bbbb/` → hostname: `ccc.bbbb`, path: `/`
11. `http://cba.com/` → hostname: `cba.com`, path: `/`
12. `http://cba.com/aba..b` → hostname: `cba.com`, path: `aba..b`
13. `http://a.a.a/aba..b/a` → hostname: `a.a.a`, path: `aba..b/a`
14. `http://abc/aba..b/a` → hostname: `abc`, path: `aba..b/a`

**Group queries by hostname:**

- **c**: `{""}` (only empty path)
- **ccc.bbbb**: `{"", "aba..b", "/"}` (empty, `aba..b`, and `/`)
- **cba.com**: `{"", "/", "aba..b"}` (empty, `/`, and `aba..b`)
- **a.c**: `{"/", "aba..b/a"}` (`/` and `aba..b/a`)
- **abc**: `{"/", "aba..b/a"}` (`/` and `aba..b/a`)
- **a.a.a**: `{"/", "aba..b/a"}` (`/` and `aba..b/a`)
- **ab.ac.bc.aa**: `{"/"}` (only `/`)

**Find groups with identical path sets:**

1. **Group 1: `cba.com` and `ccc.bbbb`**
   - Both have: `{"", "/", "aba..b"}`
   - These are identical sets! ✓

2. **Group 2: `a.c`, `abc`, and `a.a.a`**
   - All three have: `{"/", "aba..b/a"}`
   - These are identical sets! ✓

3. **Other hostnames:**
   - `c` has only `{""}` - unique, no group
   - `ab.ac.bc.aa` has only `{"/"}` - unique, no group

**Output:**
- 2 groups (only groups with size > 1)
- Group 1: `http://cba.com http://ccc.bbbb`
- Group 2: `http://a.a.a http://a.c http://abc`

### Algorithm Approach

1. Parse all URLs to extract hostname and path
2. Group queries by hostname, collecting all unique paths for each hostname
3. Group hostnames that have identical sets of paths
4. Output only groups with more than one hostname

### ideas
1. Parse each URL to extract hostname and path (note: `http://hostname` has empty path, `http://hostname/` has path `/`)
2. For each hostname, collect all unique paths that were queried
3. Group hostnames by their path sets (use a map/set to compare)
4. Output groups with size > 1
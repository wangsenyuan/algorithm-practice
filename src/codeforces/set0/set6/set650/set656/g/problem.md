# G — Simple recommendation system

## Problem description

You are building a tiny recommendation rule: an item is shown to the user only if enough of their friends “like” it.

You are given:

- how each friend feels about each item (`Y` / `N`), and  
- a threshold **T** — the **minimum** number of friends who must like an item for it to count.

**Task:** print how many items are liked by **at least T** friends.

## Input

- The first line contains three integers **F**, **I**, **T** (1 ≤ F ≤ 10, 1 ≤ I ≤ 10, 1 ≤ T ≤ F): number of **friends**, number of **items**, and the threshold.
- The next **F** lines describe opinions: line **i** has **I** characters with no spaces — the **j**-th character is **`Y`** if friend **i** likes item **j**, and **`N`** otherwise.

## Output

Print one integer — the number of items that at least **T** friends like.

## Examples

### Example 1

**Input:**

```
3 3 2
YYY
NNN
YNY
```

**Output:**

```
2
```

### Example 2

**Input:**

```
4 4 1
NNNY
NNYN
NYNN
YNNN
```

**Output:**

```
4
```

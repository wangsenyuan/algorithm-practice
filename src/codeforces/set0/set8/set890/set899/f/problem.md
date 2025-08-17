# Problem F: String Operations

## Problem Description

Petya has a string of length $n$ consisting of small and large English letters and digits.

He performs $m$ operations. Each operation is described with two integers $l$ and $r$ and a character $c$: Petya removes from the string all characters $c$ on positions between $l$ and $r$, inclusive. It's obvious that the length of the string remains the same or decreases after each operation.

Find how the string will look like after Petya performs all $m$ operations.

## Input

- The first line contains two integers $n$ and $m$ ($1 \leq n, m \leq 2 \cdot 10^5$) — the length of the string and the number of operations.
- The second line contains the string of length $n$, consisting of small and large English letters and digits. Positions in the string are enumerated from 1.
- Each of the next $m$ lines contains two integers $l$ and $r$ ($1 \leq l \leq r$), followed by a character $c$, which is a small or large English letter or a digit. This line describes one operation. It is guaranteed that $r$ doesn't exceed the length of the string $s$ before the current operation.

## Output

Print the string Petya will obtain after performing all $m$ operations. If the string becomes empty after all operations, print an empty line.

## Examples

### Example 1
**Input:**
```
4 2
abac
1 3 a
2 2 c
```

**Output:**
```
b
```

**Explanation:** During the first operation both letters 'a' are removed, so the string becomes "bc". During the second operation the letter 'c' (on the second position) is removed, and the string becomes "b".

### Example 2
**Input:**
```
3 2
A0z
1 3 0
1 1 z
```

**Output:**
```
Az
```

**Explanation:** During the first operation Petya removes '0' from the second position. After that the string becomes "Az". During the second operation the string doesn't change.

### Example 3
**Input:**
```
10 4
agtFrgF4aF
2 5 g
4 9 F
1 5 4
1 7 a
```

**Output:**
```
tFrg4
```

### Example 4
**Input:**
```
9 5
aAAaBBccD
1 4 a
5 6 c
2 3 B
4 4 D
2 3 A
```

**Output:**
```
AB
```



# Problem

Vasya has a sequence of cubes, and exactly one integer is written on each cube. He arranged all cubes in a row, so the sequence from left to right is:

`a1, a2, ..., an`.

While Vasya was walking, his little brother Stepan played with the cubes and changed their order. Now the sequence is:

`b1, b2, ..., bn`.

Stepan says that he swapped only cubes on positions between `l` and `r` (inclusive), and did not remove or add any cubes (that is, he only reordered the segment `[l, r]` in some way).

Determine whether Stepan could be telling the truth, or whether he is guaranteed to be lying.

## Input

The first line contains three integers `n`, `l`, `r` (`1 ≤ n ≤ 10^5`, `1 ≤ l ≤ r ≤ n`) — the number of cubes and the segment described by Stepan.

The second line contains `a1, a2, ..., an` (`1 ≤ ai ≤ n`) — Vasya's original order.

The third line contains `b1, b2, ..., bn` (`1 ≤ bi ≤ n`) — the order after Stepan's rearrangement.

It is guaranteed that Stepan did not remove or add cubes; he only rearranged them.

## Output

Print `"LIE"` if Stepan is guaranteed to have deceived his brother.  
Otherwise print `"TRUTH"`.

## Examples

### Example 1

**Input**

```text
5 2 4
3 4 2 3 1
3 2 3 4 1
```

**Output**

```text
TRUTH
```

### Example 2

**Input**

```text
3 1 2
1 2 3
3 1 2
```

**Output**

```text
LIE
```

### Example 3

**Input**

```text
4 2 4
1 1 1 1
1 1 1 1
```

**Output**

```text
TRUTH
```

## Note

In the first example, Stepan can be truthful. For instance:

- Start: `[3, 4, 2, 3, 1]`
- Swap positions `2` and `3` → `[3, 2, 4, 3, 1]`
- Swap positions `3` and `4` → `[3, 2, 3, 4, 1]`

In the second example, Stepan cannot be truthful: he claims changes only in positions `1..2`, but the element originally at position `3` is moved.

In the third example, for any `l` and `r`, there exists a situation where Stepan's statement can be true.

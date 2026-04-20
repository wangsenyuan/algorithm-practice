# Problem

Vasya the Great Magician and Conjurer loves all kinds of miracles and wizardry. In one wave of a magic wand he can turn an object into something else. But, as you all know, there is no better magic in the Universe than the magic of numbers. That's why Vasya adores math and spends a lot of time turning some numbers into some other ones.

This morning he has `n` cards with integers lined up in front of him. Each integer is not less than `1`, but not greater than `l`. When Vasya waves his magic wand, two rightmost cards vanish from the line and a new card magically appears in their place. It contains the difference between the left and the right numbers on the two vanished cards. Vasya was very interested to know what would happen next, and so he waved with his magic wand on and on, until the table had a single card left.

Suppose that Vasya originally had the following cards: `4, 1, 1, 3` (listed from left to right). Then after the first wave the line would be: `4, 1, -2`, and after the second one: `4, 3`, and after the third one the table would have a single card with number `1`.

Please note that in spite of the fact that initially all the numbers on the cards were not less than `1` and not greater than `l`, the numbers on the appearing cards can be anything, no restrictions are imposed on them.

It is now evening. Vasya is very tired and wants to return everything back, but does not remember which cards he had in the morning. He only remembers that there were `n` cards, they contained integers from `1` to `l`, and after all magical actions he was left with a single card containing number `d`.

Help Vasya recover the initial set of cards with numbers.

## Input

The single line contains three space-separated integers:

- `n` (`2 ≤ n ≤ 100`) — the initial number of cards on the table,
- `d` (`|d| ≤ 10^4`) — the number on the card that was left on the table after all the magical actions,
- `l` (`1 ≤ l ≤ 100`) — the bounds for the initial integers.

## Output

If Vasya is mistaken, that is, if there doesn't exist a set that meets the requirements given in the statement, print a single number `-1`.

Otherwise print the sought set containing `n` integers from `1` to `l`. Separate the integers by spaces. Print the integers in the order in which they were written on the cards from left to right.

If there are several suitable sets of numbers, you may print any of them.

## Examples

### Example 1

**Input**

```text
3 3 2
```

**Output**

```text
2 1 2
```

### Example 2

**Input**

```text
5 -4 3
```

**Output**

```text
-1
```

### Example 3

**Input**

```text
5 -4 4
```

**Output**

```text
2 4 1 4 1
```

### ideas
1. 考虑 n = 2, d[1] = a[1] - a[2], a[1] = d[1] + a[2] => a[1] <= d[1] + l, a[1] >= d[1] + l
2. a[2] = a[1] - d[1] => a[2] >= 1 - d[1], and a[2] <= l - d[1]
3. d = a[1] - (a[2] - (a[3] - (a[4] - ...)))
4.   = a[1] + a[3] + .. - a[2] - a[4] - ...
5. 

## thoughts

Key observation:

`a1 - (a2 - (a3 - ...)) = (sum of odd-indexed ai) - (sum of even-indexed ai)`.

So we only need to construct `a[i] in [1, l]` such that:

`oddSum - evenSum = d`.

### 1) Feasibility range

Let:

- `oddCnt = (n + 1) / 2`
- `evenCnt = n / 2`

Minimum possible value:

`minD = oddCnt * 1 - evenCnt * l`

Maximum possible value:

`maxD = oddCnt * l - evenCnt * 1`

If `d < minD` or `d > maxD`, answer is `-1`.

### 2) Greedy construction (left to right)

Maintain:

- `sum[0]`: current odd-position sum
- `sum[1]`: current even-position sum

For each position `i`:

- if `i` is odd-position (0-based even index), choose the smallest feasible `x` in `[1, l]` that still allows finishing to total `d`;
- if `i` is even-position (0-based odd index), choose the largest feasible `x` in `[1, l]` that still allows finishing to total `d`.

This is exactly what `solve()` does with bound formulas based on remaining counts.

### 3) Complexity

- Time: `O(n)`
- Space: `O(n)` for the output array.
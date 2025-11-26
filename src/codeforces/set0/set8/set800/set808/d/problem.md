# Problem Description

Vasya has an array a consisting of positive integer numbers. Vasya wants to divide this array into two non-empty consecutive parts (the prefix and the suffix) so that the sum of all elements in the first part equals to the sum of elements in the second part. It is not always possible, so Vasya will move some element before dividing the array (Vasya will erase some element and insert it into an arbitrary position).

Inserting an element in the same position he was erased from is also considered moving.

Can Vasya divide the array after choosing the right element to move and its new position?

## Input

The first line contains single integer n (1 ≤ n ≤ 100000) — the size of the array.

The second line contains n integers a1, a2... an (1 ≤ ai ≤ 10^9) — the elements of the array.

## Output

Print YES if Vasya can divide the array after moving one element. Otherwise print NO.

## Examples

### Example 1

**Input:**

```text
3
1 3 2
```

**Output:**

```text
YES
```

### Example 2

**Input:**

```text
5
1 2 3 4 5
```

**Output:**

```text
NO
```

### Example 3

**Input:**

```text
5
2 2 3 4 5
```

**Output:**

```text
YES
```

## Note

In the first example Vasya can move the second element to the end of the array.

In the second example no move can make the division possible.

In the third example Vasya can move the fourth element by one position to the left.


### ideas
1. sum % 2 = 0
2. 只能移动一个数字
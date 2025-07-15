# Problem D: Hieroglyph Lock

## Problem Description

Archeologists have found a secret pass in the dungeon of one of the pyramids of Cycleland. To enter the treasury they have to open an unusual lock on the door. The lock consists of n words, each consisting of some hieroglyphs. The wall near the lock has a round switch. Each rotation of this switch changes the hieroglyphs according to some rules. The instruction nearby says that the door will open only if words written on the lock would be sorted in lexicographical order (the definition of lexicographical comparison in given in notes section).

The rule that changes hieroglyphs is the following. One clockwise rotation of the round switch replaces each hieroglyph with the next hieroglyph in alphabet, i.e. hieroglyph x (1 ≤ x ≤ c - 1) is replaced with hieroglyph (x + 1), and hieroglyph c is replaced with hieroglyph 1.

Help archeologist determine, how many clockwise rotations they should perform in order to open the door, or determine that this is impossible, i.e. no cyclic shift of the alphabet will make the sequence of words sorted lexicographically.

## Input

The first line of the input contains two integers n and c (2 ≤ n ≤ 500,000, 1 ≤ c ≤ 10^6) — the number of words, written on the lock, and the number of different hieroglyphs.

Each of the following n lines contains the description of one word. The i-th of these lines starts with integer li (1 ≤ li ≤ 500,000), that denotes the length of the i-th word, followed by li integers wi,1, wi,2, ..., wi,li (1 ≤ wi,j ≤ c) — the indices of hieroglyphs that make up the i-th word. Hieroglyph with index 1 is the smallest in the alphabet and with index c — the biggest.

It's guaranteed, that the total length of all words doesn't exceed 10^6.

## Output

If it is possible to open the door by rotating the round switch, print integer x (0 ≤ x ≤ c - 1) that defines the required number of clockwise rotations. If there are several valid x, print any of them.

If it is impossible to open the door by this method, print -1.

## Examples

### Example 1
**Input:**
```
4 3
2 3 2
1 1
3 2 3 1
4 2 3 1 2
```

**Output:**
```
1
```

### Example 2
**Input:**
```
2 5
2 4 2
2 4 2
```

**Output:**
```
0
```

### Example 3
**Input:**
```
4 4
1 2
1 3
1 4
1 2
```

**Output:**
```
-1
```

## Notes

Word a1, a2, ..., am of length m is lexicographically not greater than word b1, b2, ..., bk of length k, if one of two conditions hold:

1. At first position i, such that ai ≠ bi, the character ai goes earlier in the alphabet than character bi, i.e. a has smaller character in the first position where they differ;
2. If there is no such position i and m ≤ k, i.e. the first word is a prefix of the second or two words are equal.

The sequence of words is said to be sorted in lexicographical order if each word (except the last one) is lexicographically not greater than the next word.

## Explanation

- **Example 1**: After the round switch is rotated 1 position clockwise the words look as follows:
  ```
  1 3
  2
  3 1 2
  3 1 2 3
  ```

- **Example 2**: Words are already sorted in lexicographical order.

- **Example 3**: One can check that no shift of the alphabet will work.


### ideas
1. 考虑数组x, y， x[i], y[i]是它们第一个不相同的位置
2. x[i] < y[i], 那么[0...c-y[i]）的区间内, x < y, 然后在时刻c-y[i]时， y < x, (此时 x[i] + c - y[i]), 所以经过 c - (x[i] + c - y[i]) = y[i] - c[i]时刻后， x < y
3. 那么可以也就是所有相邻区间，满足a[i] <= a[i+1] 的重叠区间

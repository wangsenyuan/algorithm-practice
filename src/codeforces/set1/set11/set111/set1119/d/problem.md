# Problem Statement

Miyako came to the flea kingdom with a ukulele. She became good friends with local flea residents and played beautiful music for them every day.

In return, the fleas made a bigger ukulele for her: it has $n$ strings, and each string has $(10^{18}+1)$ frets numerated from $0$ to $10^{18}$. The fleas use the array $s_1,s_2,…,s_n$ to describe the ukulele's tuning, that is, the pitch of the $j$-th fret on the $i$-th string is the integer $s_i+j$.

Miyako is about to leave the kingdom, but the fleas hope that Miyako will answer some last questions for them.

Each question is in the form of: "How many different pitches are there, if we consider frets between $l$ and $r$ (inclusive) on all strings?"

Miyako is about to visit the cricket kingdom and has no time to answer all the questions. Please help her with this task!

## Formal Description

You are given a matrix with $n$ rows and $(10^{18}+1)$ columns, where the cell in the $i$-th row and $j$-th column ($0 \leq j \leq 10^{18}$) contains the integer $s_i+j$. You are to answer $q$ queries, in the $k$-th query you have to answer the number of distinct integers in the matrix from the $l_k$-th to the $r_k$-th columns, inclusive.

## Input

- The first line contains an integer $n$ ($1 \leq n \leq 100000$) — the number of strings.
- The second line contains $n$ integers $s_1,s_2,…,s_n$ ($0 \leq s_i \leq 10^{18}$) — the tuning of the ukulele.
- The third line contains an integer $q$ ($1 \leq q \leq 100000$) — the number of questions.
- The $k$-th among the following $q$ lines contains two integers $l_k$, $r_k$ ($0 \leq l_k \leq r_k \leq 10^{18}$) — a question from the fleas.

## Output

Output one number for each question, separated by spaces — the number of different pitches.

## Examples

### Example 1

**Input:**
```
6
3 1 4 1 5 9
3
7 7
0 2
8 17
```

**Output:**
```
5 10 18
```

### Example 2

**Input:**
```
2
1 500000000000000000
2
1000000000000000000 1000000000000000000
0 1000000000000000000
```

**Output:**
```
2 1500000000000000000
```

## Note

For the first example, the pitches on the 6 strings are as follows.

| Fret | 0    | 1    | 2    | 3    | 4    | 5    | 6    | 7    | 8    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| s₁ : | 3    | 4    | 5    | 6    | 7    | 8    | 9    | 10   | 11   |
| s₂ : | 1    | 2    | 3    | 4    | 5    | 6    | 7    | 8    | 9    |
| s₃ : | 4    | 5    | 6    | 7    | 8    | 9    | 10   | 11   | 12   |
| s₄ : | 1    | 2    | 3    | 4    | 5    | 6    | 7    | 8    | 9    |
| s₅ : | 5    | 6    | 7    | 8    | 9    | 10   | 11   | 12   | 13   |
| s₆ : | 9    | 10   | 11   | 12   | 13   | 14   | 15   | 16   | 17   |

There are 5 different pitches on fret 7 — 8, 10, 11, 12, 16.

There are 10 different pitches on frets 0, 1, 2 — 1, 2, 3, 4, 5, 6, 7, 9, 10, 11.


### ideas
1. 每一行去看， 因为j是递增的，所以， l...r是连续的一段
2. 如果将s按照从小到大排序后，可以看到，第一段是s1 + l, ... s1 + r, 然后第二段是s2+l..s2+r
3. 这两段有可能重叠，重叠了，就有重复的情况了
4. 对于当前的s[i], s[i] + r是它的最大值，可以二分查询到s[j], s[j] + l > s[i] + r => s[j] > s[i] + r - l
5. dp[i][h] = j where s[j] > s[i] + 1 << h
6. 但是当r-l很小的时候，比如1，那么这个就太慢了
7. 
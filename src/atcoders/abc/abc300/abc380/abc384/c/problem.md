# C - Perfect Standings

[Problem link](https://atcoder.jp/contests/abc384/tasks/abc384_c)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 300 points

## Problem Statement

Takahashi decided to hold a programming contest.

The contest consists of five problems: A, B, C, D, E, with scores `a`, `b`,
`c`, `d`, `e`, respectively.

There are 31 participants, and all of them solved at least one problem.

More specifically, for every non-empty subsequence (not necessarily
contiguous) of the string `ABCDE`, there is a participant named after that
subsequence who solved the problems corresponding to the letters in their name
and did not solve the other problems.

For example, participant `A` solved only problem A, and participant `BCE`
solved problems B, C, and E.

Print the names of the participants in order of their obtained scores, from
largest to smallest. The score of a participant is the sum of the scores of
the problems they solved.

If two participants obtained the same score, print the one whose name is
lexicographically smaller first.

## Constraints

- `100 <= a <= b <= c <= d <= e <= 2718`
- All input values are integers

## Input

```
a b c d e
```

## Output

Print 31 lines. The `i`-th line should contain the name of the participant
with the `i`-th highest score (ties broken by lexicographically smaller name
first).

## Samples

### Sample 1

Input:

```
400 500 600 700 800
```

Output:

```
ABCDE
BCDE
ACDE
ABDE
ABCE
ABCD
CDE
BDE
ADE
BCE
ACE
BCD
ABE
ACD
ABD
ABC
DE
CE
BE
CD
AE
BD
AD
BC
AC
AB
E
D
C
B
A
```

### Sample 2

Input:

```
800 800 900 900 1000
```

Output:

```
ABCDE
ACDE
BCDE
ABCE
ABDE
ABCD
CDE
ACE
ADE
BCE
BDE
ABE
ACD
BCD
ABC
ABD
CE
DE
AE
BE
CD
AC
AD
BC
BD
AB
E
C
D
A
B
```

### Sample 3

Input:

```
128 256 512 1024 2048
```

Output:

```
ABCDE
BCDE
ACDE
CDE
ABDE
BDE
ADE
DE
ABCE
BCE
ACE
CE
ABE
BE
AE
E
ABCD
BCD
ACD
CD
ABD
BD
AD
D
ABC
BC
AC
C
AB
B
A
```

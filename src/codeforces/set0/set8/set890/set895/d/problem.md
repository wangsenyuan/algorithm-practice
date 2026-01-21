# Problem Statement

At the Byteland State University, marks are strings of the same length. Mark $x$ is considered better than $y$ if string $y$ is lexicographically smaller than $x$.

Recently at the BSU there was an important test on which Vasya received the mark $a$. It is very hard for the teacher to remember the exact mark of every student, but he knows the mark $b$, such that every student received a mark strictly smaller than $b$.

Vasya isn't satisfied with his mark so he decided to improve it. He can swap characters in the string corresponding to his mark as many times as he likes. Now he wants to know the number of different ways to improve his mark so that his teacher doesn't notice anything suspicious.

More formally: you are given two strings $a$, $b$ of the same length and you need to figure out the number of different strings $c$ such that:

1. $c$ can be obtained from $a$ by swapping some characters, in other words, $c$ is a permutation of $a$.
2. String $a$ is lexicographically smaller than $c$.
3. String $c$ is lexicographically smaller than $b$.

For two strings $x$ and $y$ of the same length it is true that $x$ is lexicographically smaller than $y$ if there exists such $i$, that $x_1 = y_1, x_2 = y_2, \ldots, x_{i-1} = y_{i-1}, x_i < y_i$.

Since the answer can be very large, you need to find the answer modulo $10^9 + 7$.

## Input

The first line contains string $a$, the second line contains string $b$. Strings $a$, $b$ consist of lowercase English letters. Their lengths are equal and don't exceed $10^6$.

It is guaranteed that $a$ is lexicographically smaller than $b$.

## Output

Print one integer — the number of different strings satisfying the conditions of the problem modulo $10^9 + 7$.

## Examples

### Input 1

```text
abc
ddd
```

### Output 1

```text
5
```

### Input 2

```text
abcdef
abcdeg
```

### Output 2

```text
0
```

### Input 3

```text
abacaba
ubuduba
```

### Output 3

```text
64
```

## Note

In the first sample, from string `abc` can be obtained strings `acb`, `bac`, `bca`, `cab`, `cba`, all of them are larger than `abc`, but smaller than `ddd`. So the answer is $5$.

In the second sample, any string obtained from `abcdef` is larger than `abcdeg`. So the answer is $0$.

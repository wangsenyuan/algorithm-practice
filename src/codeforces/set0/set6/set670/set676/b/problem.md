# B. Pyramid of Glasses

[Problem link](https://codeforces.com/problemset/problem/676/B)

**Contest:** [Codeforces Round #354 (Div. 2)](https://codeforces.com/contest/676)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Glasses form a pyramid of height `n`: the top level has 1 glass, the second level 2 glasses, ..., the
bottom level `n` glasses. Each glass on level `i` sits on the two glasses below it on level `i + 1`.

Each second, Vlad pours the volume of exactly one glass into the top glass. When a glass is full and
more champagne flows in, the excess overflows and is split **equally** between the two glasses directly
below it. Overflow from the bottom level spills onto the table. Distribution is instantaneous.

After `t` seconds, count how many glasses are **completely full**.

## Input

A single line with two integers `n` and `t` (`1 <= n <= 10`, `0 <= t <= 10000`) — the pyramid height
and the number of seconds.

## Output

Print one integer — the number of completely full glasses after `t` seconds.

## Example

### Input

```text
3 5
```

### Output

```text
4
```

### Input

```text
4 8
```

### Output

```text
6
```

### Note

In the first sample, after 5 seconds the full glasses are the top glass, both glasses on level 2, and
the middle glass on the bottom level. The left and right bottom glasses are half-empty.

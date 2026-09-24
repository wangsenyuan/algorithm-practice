# A. World Football Cup

[Problem link](https://codeforces.com/problemset/problem/19/A)

**Contest:** [Codeforces Beta Round #19](https://codeforces.com/contest/19)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Everyone knows that the 2010 FIFA World Cup is being held in South Africa now. By the decision of
BFA (Berland's Football Association) the next World Cup will be held in Berland. BFA decided to
change some World Cup regulations:

- the final tournament features `n` teams (`n` is always even)
- the first `n / 2` teams according to the standings come through to the knockout stage
- a team gets 3 points for a win, 1 point for a draw, and 0 points for a defeat
- teams are ordered in the standings by decreasing points; ties are broken by decreasing goal
  difference (scored minus conceded); remaining ties are broken by decreasing goals scored

You are asked to write a program that, given the list of competing teams and the results of all
matches, finds the teams that get through to the knockout stage.

## Input

The first line contains the only integer `n` (`1 ≤ n ≤ 50`) — the number of teams. The following `n`
lines contain the names of these teams. A name is a string of lower-case and upper-case Latin
letters of length at most 30.

The following `n · (n - 1) / 2` lines describe the matches in the format
`name1-name2 num1:num2`, where `name1` and `name2` are team names and `num1`, `num2`
(`0 ≤ num1, num2 ≤ 100`) are the goals scored by the corresponding teams.

It is guaranteed that no two team names coincide even ignoring letter case, no team plays itself,
and each match appears exactly once.

## Output

Print `n / 2` lines — the names of the teams that reach the knockout stage, in lexicographical
order, one name per line, with no extra characters. It is guaranteed that the ranking rules order
the teams without ambiguity.

## Examples

### Input

```text
4
A
B
C
D
A-B 1:1
A-C 2:2
A-D 1:0
B-C 1:0
B-D 0:3
C-D 0:3
```

### Output

```text
A
D
```

### Input

```text
2
a
A
a-A 2:1
```

### Output

```text
a
```

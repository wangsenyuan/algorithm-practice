# C - Fishbones

[Problem link](https://atcoder.jp/contests/abc452/tasks/abc452_c)

**Contest:** [AtCoder Beginner Contest 452](https://atcoder.jp/contests/abc452)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

Artist Takasago has created an object in the shape of a fish skeleton.

The object consists of `N` ribs and one spine. The ribs are numbered `1` through `N`.

He wants to write one string on each of the `N + 1` bones, satisfying all of the
following conditions:

- The length of the string written on the spine is `N`.
- For each rib `i = 1, ..., N`:
  - The length of the string written on rib `i` is `A_i`.
  - The `B_i`-th character of the string written on rib `i` equals the `i`-th
    character of the string written on the spine.
- Each of the strings written on the `N + 1` bones is one of `S_1, ..., S_M`
  (duplicates allowed).

`S_1, ..., S_M` are pairwise distinct lowercase strings.

For each `j = 1, ..., M`, answer whether there exists a valid assignment where the
spine string is `S_j`.

## Constraints

- `1 <= N <= 10`
- `1 <= B_i <= A_i <= 10`
- `1 <= M <= 200000`
- `1 <= |S_j| <= 10`
- `S_1, ..., S_M` are pairwise distinct
- All input values are integers where applicable; strings consist of lowercase
  English letters

## Input

```text
N
A_1 B_1
...
A_N B_N
M
S_1
...
S_M
```

## Output

Print `M` lines. The `j`-th line should be `Yes` if spine string `S_j` is possible,
and `No` otherwise.

## Sample Input 1

```text
5
5 3
5 2
4 1
5 1
3 2
8
retro
chris
itchy
tuna
crab
rock
cod
ash
```

## Sample Output 1

```text
Yes
Yes
No
No
No
No
No
No
```

Writing `chris`, `retro`, `tuna`, `retro`, `cod` on ribs `1..5` makes spine `retro`
valid. Writing `itchy`, `chris`, `rock`, `itchy`, `ash` makes spine `chris` valid.

## Sample Input 2

```text
5
5 1
5 2
5 3
5 4
5 5
8
retro
chris
itchy
tuna
crab
rock
cod
ash
```

## Sample Output 2

```text
Yes
Yes
Yes
No
No
No
No
No
```

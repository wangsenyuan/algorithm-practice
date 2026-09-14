# D. Three Sons

[Problem link](https://codeforces.com/problemset/problem/120/D)

**Contest:** [School Regional Team Contest, Saratov, 2011](https://codeforces.com/contest/120)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: input.txt

output: output.txt

Three sons inherited from their father a rectangular corn field divided into `n × m` squares. For each square we know how many tons of corn grows on it. The father, an old farmer, did not love all three sons equally, which is why he bequeathed to divide his field into three parts containing `A`, `B` and `C` tons of corn.

The field should be divided by two parallel lines. The lines should be parallel to one side of the field and to each other. The lines should go strictly between the squares of the field. Each resulting part of the field should consist of at least one square.

Your task is to find the number of ways to divide the field as is described above, that is, to mark two lines, dividing the field in three parts so that on one of the resulting parts grew `A` tons of corn, `B` on another one and `C` on the remaining one.

## Input

The first line contains space-separated integers `n` and `m` — the sizes of the original field (`1 <= n, m <= 50`, `max(n, m) >= 3`). Then the field's description follows: `n` lines, each containing `m` space-separated integers `c_ij` (`0 <= c_ij <= 100`) — the number of tons of corn each square contains. The last line contains space-separated integers `A`, `B`, `C` (`0 <= A, B, C <= 10^6`).

## Output

Print the answer to the problem: the number of ways to divide the father's field so that one of the resulting parts contained `A` tons of corn, another one contained `B` tons, and the remaining one contained `C` tons. If no such way exists, print `0`.

## Examples

### Input 1

```text
3 3
1 1 1
1 1 1
1 1 1
3 3 3
```

### Output 1

```text
2
```

### Input 2

```text
2 5
1 1 1 1 1
2 2 2 2 2
3 6 6
```

### Output 2

```text
3
```

### Input 3

```text
3 3
1 2 3
3 1 2
2 3 1
5 6 7
```

### Output 3

```text
0
```

### Note

The lines dividing the field can be horizontal or vertical, but they should be parallel to each other.

## Status

I/O and official samples are in place. `solve` is left as a TODO.

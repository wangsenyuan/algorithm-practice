# C. Fifa and Fafa

[Problem link](https://codeforces.com/problemset/problem/935/C)

**Contest:** [Codeforces Round #465 (Div. 2)](https://codeforces.com/contest/935)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Fifa wants to place a Wi-Fi access point inside a circular flat. The access point covers a circle of
radius `r` (chosen by Fifa) around its position. The access point circle must lie entirely inside the
flat, and must **not** cover Fafa's laptop or any point outside the flat.

The flat is a circle centered at `(x1, y1)` with radius `R`. Fafa's laptop is at `(x2, y2)` (not
necessarily inside the flat). Choose the access point position `(xap, yap)` and radius `r` that
minimize the uncovered area inside the flat.

## Input

A single line with five integers `R`, `x1`, `y1`, `x2`, `y2`
(`1 <= R <= 10^5`, `|x1|, |y1|, |x2|, |y2| <= 10^5`).

## Output

Print three space-separated numbers `xap`, `yap`, `r` — the chosen position and radius.

An answer is accepted if the radius is within `10^-6` (absolute or relative) of optimal, and the printed
radius can be adjusted by at most `10^-6` so that Fafa's laptop and every point outside the flat stays
outside the access point circle.

## Example

### Input

```text
5 3 3 1 1
```

### Output

```text
3.7677669529663684 3.7677669529663684 3.914213562373095
```

### Input

```text
10 5 5 5 15
```

### Output

```text
5.0 5.0 10.0
```


### ideas
1. 如果P(fafa的laptop)在flat的外面, 那么就是(x1, y1, R)
2. 如果在内部, 那么找到P最远端的在圆上的且点(这个距离就是2*r), 中点就是(xa, ya)
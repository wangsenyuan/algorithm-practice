# C - Variety (ABC461)

**Contest:** [ABC461](https://atcoder.jp/contests/abc461) — AtCoder Beginner Contest 461  
**Task:** [https://atcoder.jp/contests/abc461/tasks/abc461_c](https://atcoder.jp/contests/abc461/tasks/abc461_c)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 300 points

## Problem Statement

There are `N` gems. Gem `i` has color `C_i` and value `V_i`.

Choose exactly `K` gems such that at least `M` distinct colors appear among the chosen
gems. Find the maximum possible total value. (Such a choice always exists.)

## Constraints

- `1 <= M <= K <= N <= 2 × 10^5`
- `1 <= C_i <= N`
- `1 <= V_i <= 10^9`
- At least `M` distinct colors appear among all `N` gems
- All input values are integers

## Input

```text
N K M
C_1 V_1
C_2 V_2
⋮
C_N V_N
```

## Output

Print the maximum possible total value.

## Sample Input 1

```text
5 3 2
1 30
1 40
1 50
2 10
3 20
```

## Sample Output 1

```text
110
```

Choosing gems 2, 3, and 5 gives colors `{1, 3}` (two colors) and total value
`40 + 50 + 20 = 110`.

## Sample Input 2

```text
5 3 3
1 30
1 40
1 50
2 10
3 20
```

## Sample Output 2

```text
80
```

Choosing gems 3, 4, and 5 gives colors `{1, 2, 3}` and total value `80`.

## Sample Input 3

```text
5 5 1
4 1000000000
5 1000000000
4 1000000000
5 1000000000
4 1000000000
```

## Sample Output 3

```text
5000000000
```

Beware of overflow.

### ideas
1. 假设正好选择了m种宝石, 那么这m种颜色中,最贵的那个宝石肯定会被选中
2. 那不妨先把这些宝石选出来,再添加其他的宝石
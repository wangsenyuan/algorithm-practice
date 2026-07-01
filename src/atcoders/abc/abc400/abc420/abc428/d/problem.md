# D - 183184

[Problem link](https://atcoder.jp/contests/abc428/tasks/abc428_d)

**Contest:** [AtCoder Beginner Contest 428](https://atcoder.jp/contests/abc428)

time limit: 2.5 sec

memory limit: 1024 MiB

score: 400 points

For positive integers `x` and `y`, define `f(x, y)` as follows:

- Let `z` be the string obtained by interpreting `x` and `y` in decimal notation as
  strings and concatenating them in this order. Let `f(x, y)` be the value when `z` is
  interpreted as an integer in decimal notation.

For example, `f(3, 14) = 314` and `f(100, 3) = 1003`.

You are given positive integers `C` and `D`. Find the number of integers `x` that satisfy
the following conditions:

- `1 <= x <= D`
- `f(C, C + x)` is a perfect square

You are given `T` test cases; find the answer for each of them.

## Constraints

- `1 <= T <= 3 * 10^5`
- `1 <= C <= 2 * 10^8`
- `1 <= D <= 5 * 10^9`
- All input values are integers

## Input

```text
T
case_1
case_2
...
case_T
```

Each test case is given in the following format:

```text
C D
```

## Output

Print `T` lines. The `i`-th line should contain the answer to the `i`-th test case.

## Sample Input 1

```text
4
4 80
183 5000
18 10
824 5000000000
```

## Sample Output 1

```text
3
2
0
1421
```

### Note

For the first test case, there are three values of `x` that satisfy the conditions:
`x = 5, 37, 80`.

- When `x = 5`, `f(C, C + 5) = f(4, 9) = 49 = 7^2`
- When `x = 37`, `f(C, C + 37) = f(4, 41) = 441 = 21^2`
- When `x = 80`, `f(C, C + 80) = f(4, 84) = 484 = 22^2`

For the second test case, there are two values of `x` that satisfy the conditions:
`x = 1, 3133`.

- When `x = 1`, `f(C, C + 1) = f(183, 184) = 183184 = 428^2`
- When `x = 3133`, `f(C, C + 3133) = f(183, 3316) = 1833316 = 1354^2`

For the third test case, there are zero values of `x` that satisfy the conditions.

For the fourth test case, there are 1421 values of `x` that satisfy the conditions.

### ideas

1. f(c, c + x) = c#(c+x) 是一个平方数, 那么长度可以到9+8 = 17
2. 那么这个找出所有的完全平方数就不可行了 
3. 如果 x > c 会怎么样?

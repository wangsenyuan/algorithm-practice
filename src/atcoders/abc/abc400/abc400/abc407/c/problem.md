# C - Security 2

[Problem link](https://atcoder.jp/contests/abc407/tasks/abc407_c)

**Contest:** [AtCoder Beginner Contest 407](https://atcoder.jp/contests/abc407)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

At the entrance of AtCoder Inc., there is a special pass-code input device with a screen and two buttons.

Let `t` be the string shown on the screen. Initially, `t` is empty. Pressing a button changes `t` as follows:

- **Button A** appends `0` to the end of `t`.
- **Button B** replaces every digit in `t` with the next digit (`0..8 -> +1`, `9 -> 0`).

For example, if `t = 1984` and button A is pressed, `t` becomes `19840`; if button B is then pressed,
`t` becomes `20951`.

Given a string `S`, starting from the empty string, find the minimum number of button presses needed so
that `t` equals `S`.

## Constraints

- `S` consists of digits `0` through `9`.
- `1 <= |S| <= 5 * 10^5`

## Input

```text
S
```

## Output

Print the answer.

## Sample Input 1

```text
21
```

## Sample Output 1

```text
4
```

One optimal sequence:

1. A → `0`
2. B → `1`
3. A → `10`
4. B → `21`

## Sample Input 2

```text
407
```

## Sample Output 2

```text
17
```

## Sample Input 3

```text
2025524202552420255242025524
```

## Sample Output 3

```text
150
```

### ideas
1. 肯定是有答案的. 可以从尾巴开始处理, 
2. 如果它不是0, 那么肯定进行了x次B操作, 然后通过一次A操作. 这时候可以把它给去掉
3. 对于前方来说, (要记录x次B操作)的影响. 
4. 似乎这样子就是最少操作次数
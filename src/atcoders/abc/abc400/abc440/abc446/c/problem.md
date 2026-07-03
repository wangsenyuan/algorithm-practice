# C - Omelette Restaurant

[Problem link](https://atcoder.jp/contests/abc446/tasks/abc446_c)

**Contest:** [AtCoder Beginner Contest 446](https://atcoder.jp/contests/abc446)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

AtCoder Restaurant was open for `N` days. On day `i` (`1 <= i <= N`):

- **Morning:** purchase `A_i` eggs
- **Noon:** use `B_i` eggs from stock, using eggs in **purchase order** (FIFO)
- **Evening:** discard all eggs that have been in stock for `D` or more days

Before day 1 morning, stock is empty. Eggs never run out at noon on any day.

Find how many eggs remain after the evening action on day `N`.

You are given `T` test cases; solve each.

## Constraints

- `1 <= T <= 2 * 10^5`
- `1 <= D <= N <= 2 * 10^5`
- `1 <= A_i, B_i <= 10`
- Eggs never run out at noon on any day
- For each input, the sum of `N` over all test cases is at most `2 * 10^5`
- All input values are integers

## Input

```text
T
case_1
case_2
...
case_T
```

Each test case:

```text
N D
A_1 A_2 ... A_N
B_1 B_2 ... B_N
```

## Output

Print `T` lines. The `i`-th line is the answer for the `i`-th test case.

## Sample Input 1

```text
3
3 1
7 2 3
1 3 2
3 2
7 2 3
1 3 2
2 1
2 1
1 2
```

## Sample Output 1

```text
3
5
0
```

### Note

In the first test case (`D = 1`), after day 3 evening, 3 eggs remain (all purchased on day 3).

In the second test case (`D = 2`), eggs purchased on day 1 are discarded on day 3 evening; answer is `5`.

### ideas
1. 

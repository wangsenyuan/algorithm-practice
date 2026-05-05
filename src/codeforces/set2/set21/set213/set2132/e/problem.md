# E. Arithmetics Competition

In the arithmetic competition, participants need to achieve the **highest possible sum** from the cards they have. On team **"fst_ezik"**, Vadim has `n` cards with values `a_i`, and Kostya has `m` cards with values `b_i`. In each of the `q` rounds they want to win, but the rules differ slightly from the usual setup.

In each round they are given three integers `x_i`, `y_i`, and `z_i`. The team must choose **exactly** `z_i` cards in total from both piles combined, with Vadim contributing **at most** `x_i` cards from his hand and Kostya **at most** `y_i` from his. For every round, find the **maximum achievable sum** of the chosen cards.

## Input

The input contains several test cases. The first line has a single integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

For each test case:

- The first line has three integers `n`, `m`, `q` (`1 ≤ n, m ≤ 2 · 10^5`, `1 ≤ q ≤ 10^5`) — Vadim’s card count, Kostya’s card count, and the number of rounds.
- The second line has `n` integers `a_i` — Vadim’s card values (`1 ≤ a_i ≤ 10^9`).
- The third line has `m` integers `b_i` — Kostya’s card values (`1 ≤ b_i ≤ 10^9`).
- The next `q` lines each have three integers `x_i`, `y_i`, `z_i` (`0 ≤ x_i ≤ n`, `0 ≤ y_i ≤ m`, `0 ≤ z_i ≤ x_i + y_i`) — caps on how many cards Vadim and Kostya may use, and how many cards must be picked in total.

**Guarantee:** the sum of `n` over all test cases ≤ `2 · 10^5`, the sum of `m` over all test cases ≤ `2 · 10^5`, and the sum of `q` over all test cases ≤ `10^5`.

## Output

For each test case, print `q` numbers — the maximum total sum for each round, in order.

## Example

**Input**

```text
4
3 4 5
10 20 30
1 2 3 4
0 0 0
3 4 7
3 4 4
1 4 4
2 2 4
5 5 2
500000000 300000000 100000000 900000000 700000000
800000000 400000000 1000000000 600000000 200000000
1 4 3
5 2 6
4 4 1
100 100 20 20
100 100 20 20
4 4 5
3 3 6
2 363 711
286 121 102
1 1 1
3 1 1
1 2 0
1 3 2
0 1 0
3 3 3
```

**Output**

```text
0
70
64
39
57
2700000000
4200000000
420
711
711
0
997
0
1360
```

### ideas
1. sum(a[:x]) + sum(b[:y]) 
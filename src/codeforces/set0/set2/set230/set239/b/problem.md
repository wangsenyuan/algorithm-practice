# Problem

There is a programming language in which every program is a non-empty sequence of `<`, `>`, and digits. The interpreter uses:

- **CP** — current character pointer
- **DP** — direction pointer, either left or right

Initially CP points to the leftmost character and DP points to the right.

Repeat until CP leaves the sequence:

1. **Digit:** print that digit, then move CP one step in the direction of DP. After printing, decrease that digit in the sequence by `1`. If the digit was `0`, it cannot be decreased, so it is removed and the sequence length decreases by `1`.
2. **`<` or `>`:** set DP to left or right respectively, then move CP one step according to DP. If the new character under CP is `<` or `>`, erase the **previous** character from the sequence.

If CP ever goes outside the sequence, execution stops.

It is guaranteed that every program terminates after finitely many steps.

You are given a sequence `s_1, s_2, ..., s_n` of `<`, `>`, and digits. Answer `q` queries: each query gives `l` and `r` and asks how many times each digit `0..9` is printed when the substring `s_l, s_{l+1}, ..., s_r` is run as an **independent** program.

## Input

- First line: integers `n` and `q` (`1 <= n, q <= 100`) — length of `s` and number of queries.
- Second line: string `s` — characters `<`, `>`, and digits `0..9` with no spaces.
- Next `q` lines: two integers `l_i` and `r_i` (`1 <= l_i <= r_i <= n`).

## Output

For each query, print **10** space-separated integers `x_0, x_1, ..., x_9`, where `x_i` is how many times digit `i` was printed. Print answers in query order.

## Examples

### Example 1

Input

```text
7 4
1>3>22<
1 3
4 7
7 7
1 7
```

Output

```text
0 1 0 1 0 0 0 0 0 0
2 2 2 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
2 3 2 1 0 0 0 0 0 0
```

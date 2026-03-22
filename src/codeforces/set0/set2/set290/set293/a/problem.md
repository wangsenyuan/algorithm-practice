# Problem

Roman leaves a word for each of them. Each word consists of `2·n` binary characters `"0"` or `"1"`. After that the players start moving in turns. Yaroslav moves first. During a move, a player must choose an integer from `1` to `2·n`, which has not been chosen by anybody up to that moment. Then the player takes a piece of paper and writes out the corresponding character from his string.

Let's represent Yaroslav's word as `s = s1 s2 ... s2n`. Similarly, let's represent Andrey's word as `t = t1 t2 ... t2n`. Then, if Yaroslav chooses number `k` during his move, he writes out character `sk` on the piece of paper. Similarly, if Andrey chooses number `r` during his move, he writes out character `tr` on the piece of paper.

The game finishes when no player can make a move. After the game is over, Yaroslav makes some integer from the characters written on his piece of paper (Yaroslav can arrange these characters as he wants). Andrey does the same. The resulting numbers can contain leading zeroes. The person with the largest number wins. If the numbers are equal, the game ends with a draw.

You are given two strings `s` and `t`. Determine the outcome of the game provided that Yaroslav and Andrey play optimally well.

## Input

- First line: integer `n` (`1 <= n <= 10^6`).
- Second line: string `s` — Yaroslav's word.
- Third line: string `t` — Andrey's word.

It is guaranteed that both words consist of `2·n` characters `"0"` and `"1"`.

## Output

Print `First` if both players play optimally well and Yaroslav wins. If Andrey wins, print `Second` and if the game ends with a draw, print `Draw`. Print the words without the quotes.

## Examples

### Example 1

Input

```text
2
0111
0001
```

Output

```text
First
```

### Example 2

Input

```text
3
110110
001001
```

Output

```text
First
```

### Example 3

Input

```text
3
111000
000111
```

Output

```text
Draw
```

### Example 4

Input

```text
4
01010110
00101101
```

Output

```text
First
```

### Example 5

Input

```text
4
01100000
10010011
```

Output

```text
Second
```

### ideas
1. 就是看1的个数，所以如果(s[i], t[i]) = (0, 0) 那么这个可以最后选
2. 如果(s[i], t[i]) = (1, 1) 那么这个要最早选
3. 如果s[i] != t[i], 遇到了就要选
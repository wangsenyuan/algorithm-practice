# Problem E: CD Games

## Problem Description

Several years ago, Tolya had **n** computer games and at some point decided to burn them to a CD. After that, he wrote down the names of the games one after another in a circle on the CD in clockwise order. The names were distinct, and the length of each name was equal to **k**. The names didn't overlap.

Thus, there is a cyclic string of length **n·k** written on the CD.

Several years have passed, and now Tolya can't remember which games he burned to his CD. He knows that there were **g** popular games in those days. All of the games he burned were among these **g** games, and no game was burned more than once.

You have to restore any valid list of games Tolya could have burned to the CD several years ago.

## Input

The first line of the input contains two positive integers **n** and **k**:
- **n** (1 ≤ n ≤ 10^5) — the amount of games Tolya burned to the CD
- **k** (1 ≤ k ≤ 10^5) — the length of each of the names

The second line of the input contains one string consisting of lowercase English letters — the string Tolya wrote on the CD, split in an arbitrary place. The length of the string is **n·k**. It is guaranteed that the length is not greater than 10^6.

The third line of the input contains one positive integer **g** (n ≤ g ≤ 10^5) — the amount of popular games that could be written on the CD. It is guaranteed that the total length of names of all popular games is not greater than 2·10^6.

Each of the next **g** lines contains a single string — the name of some popular game. Each name consists of lowercase English letters and has length **k**. It is guaranteed that the names are distinct.

## Output

If there is no answer, print `"NO"` (without quotes).

Otherwise, print two lines:
- In the first line, print `"YES"` (without quotes)
- In the second line, print **n** integers — the games whose names were written on the CD. You should print games in the order they could have been written on the CD (i.e., in clockwise order). You can print games starting from any position. Remember that no game was burned to the CD more than once. If there are several possible answers, print any of them.

## Examples

### Example 1

**Input:**
```
3 1
abc
4
b
a
c
d
```

**Output:**
```
YES
2 1 3
```

**Explanation:** The games could be arranged as "abc" which corresponds to games 2, 1, 3 in order.

### Example 2

**Input:**
```
4 2
aabbccdd
4
dd
ab
bc
cd
```

**Output:**
```
NO
```

**Explanation:** It's impossible to arrange the given game names to form the string "aabbccdd".

## ideas
1. 虽然每个名字是unique的，但是可能存在头尾相连组成一个名字的情况
2. 假设 s[i:i+k]在 g中找到了
3. dp[i] = dp[i-k] & s[i-k:i]在g中被找到了
4. 那么就是找到 dp[n], ... dp[n+k]的一个为true的值
5. 所以这里的关键是快速判断s[i:i+k]是否存在，可以用hash
6. 
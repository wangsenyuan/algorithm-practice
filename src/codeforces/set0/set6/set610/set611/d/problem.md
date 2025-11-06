## Problem Statement

Limak is a little polar bear. In the snow he found a scroll with the ancient prophecy. Limak doesn't know any ancient languages and thus is unable to understand the prophecy. But he knows digits!

One fragment of the prophecy is a sequence of $n$ digits. The first digit isn't zero. Limak thinks that it's a list of some special years. It's hard to see any commas or spaces, so maybe ancient people didn't use them. Now Limak wonders what years are listed there.

Limak assumes three things:

- Years are listed in the strictly increasing order;
- Every year is a positive integer number;
- There are no leading zeros.

Limak is going to consider all possible ways to split a sequence into numbers (years), satisfying the conditions above. He will do it without any help. However, he asked you to tell him the number of ways to do so. Since this number may be very large, you are only asked to calculate it modulo $10^9 + 7$.

## Input

The first line of the input contains a single integer $n$ ($1 \le n \le 5000$) — the number of digits.

The second line contains a string of digits and has length equal to $n$. It's guaranteed that the first digit is not '0'.

## Output

Print the number of ways to correctly split the given sequence modulo $10^9 + 7$.

## Examples

### Example 1

**Input:**
```
6
123434
```

**Output:**
```
8
```

### Example 2

**Input:**
```
8
20152016
```

**Output:**
```
4
```

## Note

In the first sample there are 8 ways to split the sequence:

- "123434" = "123434" (maybe the given sequence is just one big number)
- "123434" = "1" + "23434"
- "123434" = "12" + "3434"
- "123434" = "123" + "434"
- "123434" = "1" + "23" + "434"
- "123434" = "1" + "2" + "3434"
- "123434" = "1" + "2" + "3" + "434"
- "123434" = "1" + "2" + "3" + "4" + "34"

Note that we don't count a split "123434" = "12" + "34" + "34" because numbers have to be strictly increasing.

In the second sample there are 4 ways:

- "20152016" = "20152016"
- "20152016" = "20" + "152016"
- "20152016" = "201" + "52016"
- "20152016" = "2015" + "2016"


### ideas
1. dp[i][len] 表示到i为止，长度为len的 suf的方案数
2. 那么dp[j][x] = dp[j-i][y <= x]
3. 对于y < x 的部分，肯定满足 < 的关系，但是对于 y = x的，就比较没法了
4. 必须保证后半部分比前半部分要大，才行
5. 123124 
6. 对于 s[i...] 这样一个后缀，找到一个 s[l...i], s[l...i] < s[i...r]
7. where i - l = r - i
8. 先计算出 f(l, i) = 1， 表示 s[l..i] < s[i...r]
9. f(l, i) 表示 s(l...), s(i) 第一个不同的字符的距离
10. 如果 s[l] != s[i], f(l, i) = 0, 否则的话, f(l, i) = 1 + f(l + 1, i + 1)
11. 那么要判断s[l...i] s[i...r] let k = f(l, i)， 如果 k >= i - l => 两个相等，否则就比较那个不同的字符即可
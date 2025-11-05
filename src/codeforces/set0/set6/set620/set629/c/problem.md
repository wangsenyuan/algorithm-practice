# Valid Bracket Sequences

As Famil Door's birthday is coming, some of his friends (like Gabi) decided to buy a present for him. His friends are going to buy a string consisted of round brackets since Famil Door loves string of brackets of length n more than any other strings!

The sequence of round brackets is called valid if and only if:

1. the total number of opening brackets is equal to the total number of closing brackets;
2. for any prefix of the sequence, the number of opening brackets is greater or equal than the number of closing brackets.

Gabi bought a string s of length m $(m \leq n)$ and want to complete it to obtain a valid sequence of brackets of length n. He is going to pick some strings p and q consisting of round brackets and merge them in a string p + s + q, that is add the string p at the beginning of the string s and string q at the end of the string s.

Now he wonders, how many pairs of strings p and q exists, such that the string p + s + q is a valid sequence of round brackets. As this number may be pretty large, he wants to calculate it modulo $10^9 + 7$.

## Input

First line contains n and m $(1 \leq m \leq n \leq 100000, n - m \leq 2000)$ — the desired length of the string and the length of the string bought by Gabi, respectively.

The second line contains string s of length m consisting of characters '(' and ')' only.

## Output

Print the number of pairs of string p and q such that p + s + q is a valid sequence of round brackets modulo $10^9 + 7$.

## Examples

### Example 1

**Input:**
```
4 1
(
```

**Output:**
```
4
```

### Example 2

**Input:**
```
4 4
(())
```

**Output:**
```
1
```

### Example 3

**Input:**
```
4 3
(((
```

**Output:**
```
0
```

## Note

In the first sample there are four different valid pairs:

- p = "(", q = "))"
- p = "()", q = ")"
- p = "", q = "())"
- p = "", q = ")()"

In the second sample the only way to obtain a desired string is choose empty p and q.

In the third sample there is no way to get a valid sequence of brackets.


### ideas
1. 假设在前面分配了u个，后面分配了v个，那么 u + v = n - m <= 2000
2. 考虑前面得到dp, dp[bal] 表示u个字符后剩余bal个左扩号后的计数， bal >= s中多出来的右括号
3. fp[bal] 表示从后往前右括号-左括号剩余bal时的计数，那么两个乘一下就可以了
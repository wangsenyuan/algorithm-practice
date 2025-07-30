# Problem D: Santa Claus and Palindromes

## Problem Description

Santa Claus likes palindromes very much. There was his birthday recently. **k** of his friends came to him to congratulate him, and each of them presented to him a string **s<sub>i</sub>** having the same length **n**. We denote the beauty of the i-th string by **a<sub>i</sub>**. It can happen that **a<sub>i</sub>** is negative — that means that Santa doesn't find this string beautiful at all.

Santa Claus is crazy about palindromes. He is thinking about the following question: **what is the maximum possible total beauty of a palindrome which can be obtained by concatenating some (possibly all) of the strings he has?** Each present can be used at most once. Note that all strings have the same length **n**.

Recall that a **palindrome** is a string that doesn't change after one reverses it.

Since the empty string is a palindrome too, the answer can't be negative. Even if all **a<sub>i</sub>**'s are negative, Santa can obtain the empty string.

## Input

The first line contains two positive integers **k** and **n** divided by space and denoting the number of Santa friends and the length of every string they've presented, respectively.

**Constraints:**
- 1 ≤ k, n ≤ 100,000
- n·k ≤ 100,000

**k** lines follow. The i-th of them contains the string **s<sub>i</sub>** and its beauty **a<sub>i</sub>** (-10,000 ≤ a<sub>i</sub> ≤ 10,000). The string consists of **n** lowercase English letters, and its beauty is integer. Some of strings may coincide. Also, equal strings can have different beauties.

## Output

In the only line print the required maximum possible beauty.

## Examples

### Example 1

**Input:**
```
7 3
abb 2
aaa -3
bba -1
zyz -4
abb 5
aaa 7
xyx 4
```

**Output:**
```
12
```

### Example 2

**Input:**
```
3 1
a 1
a 2
a 3
```

**Output:**
```
6
```

### Example 3

**Input:**
```
2 5
abcde 10000
abcde 10000
```

**Output:**
```
0
```

## Note

In the first example Santa can obtain `abbaaaxyxaaabba` by concatenating strings 5, 2, 7, 6 and 3 (in this order).


### ideas
1. 有两类字符串，一类是 x . y 是一个回文的，比如 abc . cba 的
2. 还有一类是 x . x 本身就是回文的，比如 aba . aba
3. 第二类的，可以放在最中间。第一类的，只能放在两边
4. 对于第一类，按照分数降序排列，只要分数相加是正数，就添加进去
5. 对于第二类，只需要取正数的部分吗？
6. 似乎有点麻烦，假如有多个第二类的，尽量使用正数的部分
7. 然后剩下最多一个正数 + 多个负数的情况
8. 然后这样子组合，看如果中间放置一个正数，两边放置 正数 + 负数的组合
9. 
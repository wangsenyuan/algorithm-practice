# Problem D: Santa's Letter Algorithm

## Problem Description

Many countries have such a New Year or Christmas tradition as writing a letter to Santa including a wish list for presents. Vasya is an ordinary programmer boy. Like all ordinary boys, he is going to write the letter to Santa on the New Year Eve (we Russians actually expect Santa for the New Year, not for Christmas).

Vasya has come up with an algorithm he will follow while writing a letter. First he chooses two strings, s1 and s2, consisting of uppercase English letters. Then the boy makes string sk, using a recurrent equation sn = sn-2 + sn-1, operation '+' means a concatenation (that is, the sequential record) of strings in the given order. Then Vasya writes down string sk on a piece of paper, puts it in the envelope and sends in to Santa.

Vasya is absolutely sure that Santa will bring him the best present if the resulting string sk has exactly x occurrences of substring AC (the short-cut reminds him of accepted problems). Besides, Vasya decided that string s1 should have length n, and string s2 should have length m. Vasya hasn't decided anything else.

At the moment Vasya's got urgent New Year business, so he asks you to choose two strings for him, s1 and s2 in the required manner. Help Vasya.

## Input

- **First line**: Four integers k, x, n, m (3 ≤ k ≤ 50; 0 ≤ x ≤ 10^9; 1 ≤ n, m ≤ 100)

## Output

- **First line**: String s1, consisting of n uppercase English letters
- **Second line**: String s2, consisting of m uppercase English letters
- If there are multiple valid strings, print any of them
- If the required pair of strings doesn't exist, print `"Happy new year!"` without the quotes

## ideas
1. 假设s1 有x个AC, s2有y个AC
2. 假设 x * 2 <= n and y * 2 <= m, 且f(x, y, k) == w
3. f(x, 0/1, y, 0/1, i) 表示如果s1中有x个AC，且第一个字符是否是C(0不是，1是)
4. s2中有y个AC，且最后一个字符是否是A(0不是,1是)， 在第i个字符中有多少个i
5. f(x, 0, y, 0, i + 1) = f(x, 0, y, 0, i) + f(x, 0, y, 0, i - 1)
6. 不会组合出新的字符
7. f(x, 1, y, 1, i + 1) = f(x, 1, y, 1, i) + f(x, 1, y, 1, i - 1) + 1
8. 
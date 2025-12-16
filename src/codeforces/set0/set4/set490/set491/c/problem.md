### Problem
One day Maria Ivanovna found Sasha's piece of paper with a message dedicated to Olya. The message is ciphered. Students usually cipher their messages by replacing each letter of the original message with another letter: identical letters map to the same letter, and different letters map to different letters.

Maria suspects the message contains answers to the final exam (its length equals the number of exam questions). There are **K** possible answers for each question. Maria knows the correct answers, but Sasha's answers may be wrong. She wants to decipher the message to maximize the number of Sasha's answers that match the correct ones.

### Input
- First line: integers **N** (1 ≤ N ≤ 2,000,000) and **K** (1 ≤ K ≤ 52) — length of both strings and the number of possible answers. Answers are Latin letters in order `abcde...xyzABCDE...XYZ`.  
  For example:  
  - K = 6 → answers are `abcdef`  
  - K = 30 → answers are `abcde...xyzABCD`
- Second line: ciphered message string (length N) of Latin letters.
- Third line: correct answers string (length N) of Latin letters.

### Output
- First line: maximum possible number of Sasha's correct answers.
- Second line: cipher rule as a string of length **K**, where each position (starting from `a` in the students' alphabet) specifies which answer it maps to.
- If multiple rules yield the maximum, output any of them.

### Examples
```
Input
10 2
aaabbbaaab
bbbbabbbbb

Output
7
ba
```

```
Input
10 2
aaaaaaabbb
bbbbaaabbb

Output
6
ab
```

```
Input
9 4
dacbdacbd
acbdacbda

Output
9
cdba
```


### ideas
1. 有n个题目，每个题目都是选择题，有K个备选项
2. a[1], a[2], ... a[k] 被替换成 b[1], b[2], ... b[k]
3. 且一一对应; 在保证f(s) 和 t匹配最大的情况下，找到b
4. f(s) 是s替换前的输入
5. 构建一个 k * k 的图， 1 -> w的cap = 它们在相同位置匹配的数量
6. 然后求最大流 
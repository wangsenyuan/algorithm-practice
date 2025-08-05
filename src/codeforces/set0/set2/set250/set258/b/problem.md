# Little Elephant Political Party

## Problem Description

There have recently been elections in the zoo. Overall there were 7 main political parties: one of them is the Little Elephant Political Party, 6 other parties have less catchy names.

Political parties find their number in the ballot highly important. Overall there are m possible numbers: 1, 2, ..., m. Each of these 7 parties is going to be assigned in some way to exactly one number, at that, two distinct parties cannot receive the same number.

The Little Elephant Political Party members believe in the lucky digits 4 and 7. They want to evaluate their chances in the elections. For that, they need to find out, how many correct assignments are there, such that the number of lucky digits in the Little Elephant Political Party ballot number is strictly larger than the total number of lucky digits in the ballot numbers of 6 other parties.

Help the Little Elephant Political Party, calculate this number. As the answer can be rather large, print the remainder from dividing it by 1000000007 (10^9 + 7).

## Input

A single line contains a single positive integer m (7 ≤ m ≤ 10^9) — the number of possible numbers in the ballot.

## Output

In a single line print a single integer — the answer to the problem modulo 1000000007 (10^9 + 7).

### ideas
1. 这个题目是说，如果分配一些数字给Elephant, 比如，4, 7, 14, 17, 47。。。
2. 这些数字中的4、7的数量 > 其他没有被选中的数量？
3. 这里貌似只需要考虑包含数字4、7的数字，它们被分配时，才会影响到差额
4. 比如数字 1247, 把它分配给E，那么它就多了2,
5. 只有个数有关系，假设1个lucky的数字有cnt[1], 两个的有cnt[2]... cnt[8] (最多cnt[8])
6. 然后对它们进行分配
7. dp[diff]表示分配到目前为止，切 x= diff时的结果（可以是负值）
8. 现在分配4个，假设它有100个，然后其中分配10个给element， 那么diff += （10 - 100） * 4
9. 这个diff似乎是可以很大的吗？
10. 假设全部分配给Element， cnt[1] 4,7,14,17, 41, 71, .... C(8, 1) * 2 * pow(8, 8)
11.  如果只考虑4,7,47, 似乎就没有那么多了
12.  4， 7，47，74，447，。。。。这个数量是很少的，
13.  假设有W个，那么w * C(m - w, 6)


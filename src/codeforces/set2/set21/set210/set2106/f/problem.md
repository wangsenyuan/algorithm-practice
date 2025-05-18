Dr. TC has a new patient called Goblin. He wants to test Goblin's intelligence, but he has gotten bored of his standard test. So, he decided to make it a bit harder.

First, he creates a binary string∗
 𝑠
 having 𝑛
 characters. Then, he creates 𝑛
 binary strings 𝑎1,𝑎2,…,𝑎𝑛
. It is known that 𝑎𝑖
 is created by first copying 𝑠
, then flipping the 𝑖
-th character (𝟷
 becomes 𝟶
 and vice versa). After creating all 𝑛
 strings, he arranges them into an 𝑛×𝑛
 grid 𝑔
 where 𝑔𝑖,𝑗=𝑎𝑖𝑗
.

A set 𝑆
 of size 𝑘
 containing distinct integer pairs {(𝑥1,𝑦1),(𝑥2,𝑦2),…,(𝑥𝑘,𝑦𝑘)}
 is considered good if:

1≤𝑥𝑖,𝑦𝑖≤𝑛
 for all 1≤𝑖≤𝑘
.
𝑔𝑥𝑖,𝑦𝑖=𝟶
 for all 1≤𝑖≤𝑘
.
For any two integers 𝑖
 and 𝑗
 (1≤𝑖,𝑗≤𝑘
), coordinate (𝑥𝑖,𝑦𝑖)
 is reachable from coordinate (𝑥𝑗,𝑦𝑗)
 by traveling through a sequence of adjacent cells (which share a side) that all have a value of 𝟶
.
Goblin's task is to find the maximum possible size of a good set 𝑆
. Because Dr. TC is generous, this time he gave him two seconds to find the answer instead of one. Goblin is not known for his honesty, so he has asked you to help him cheat.

∗
A binary string is a string that only consists of characters 𝟷
 and 𝟶
.

Input
The first line of the input consists of a single integer 𝑡
 (1≤𝑡≤103)
 — the number of test cases.

The first line of each test contains a single integer 𝑛
 (1≤𝑛≤2⋅105)
 — the length of the binary string 𝑠
.

The second line of each test contains a single binary string 𝑠
 of length 𝑛
.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 2⋅105
.

Output
For each test case, output a single number, the maximum possible size of a good set of cells from the grid.

### ideas
1. 例子 0110
`
1110
0010
0100
0111
`
2.  如果有两个连续的1，那么这两个肯定会把这个区域隔开
3.  按列考虑，考虑第一列，如果它是a[0] = 1，那么除了g[0][0] = 0, 其他都是1
4.  如果a[0] = 0, 除了g[0][0] = 1, 其他都是0 （这一列的贡献 = (n - 1) * (n - 2) / 2)
5.  每一列的0的范围最多分两组，对于第一列，要么是[0, 0], 要么是 [1, n -1]
6.  假设a[0] = 1, 那么除了g[1][1] = 0, 其他行都是1，
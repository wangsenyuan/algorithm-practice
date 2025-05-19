There is an apple tree with 𝑛
 nodes, initially with one apple at each node. You have a paper with you, initially with nothing written on it.

You are traversing on the apple tree, by doing the following action as long as there is at least one apple left:

Choose an apple path (𝑢,𝑣)
. A path (𝑢,𝑣)
 is called an apple path if and only if for every node on the path (𝑢,𝑣)
, there's an apple on it.
Let 𝑑
 be the number of apples on the path, write down three numbers (𝑑,𝑢,𝑣)
, in this order, on the paper.
Then remove all the apples on the path (𝑢,𝑣)
.
Here, the path (𝑢,𝑣)
 refers to the sequence of vertices on the unique shortest walk from 𝑢
 to 𝑣
.

Let the number sequence on the paper be 𝑎
. Your task is to find the lexicographically largest possible sequence 𝑎
.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤104
). The description of the test cases follows.

The first line of each test case contains a number 𝑛
 (1≤𝑛≤1.5⋅105
).

The following 𝑛−1
 lines of each test case contain two numbers 𝑢,𝑣
 (1≤𝑢,𝑣≤𝑛
). It's guaranteed that the input forms a tree.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 1.5⋅105
.

Output
For each test case, output the lexicographically largest sequence possible 𝑎1,𝑎2,…,𝑎|𝑎|
. It can be shown that |𝑎|≤3⋅𝑛
.


### ideas
1. 第一步肯定是贪心的，因为path越长，第一个数字越大
2. 但是如果有两个最长呢？甚至有k个最长呢？
3. 那要怎么选择呢
4. 似乎还是要贪心。 因为节点的编号肯定不一样。在k相同的情况下，就选择u最大的那个（起点最大的那个）
5. 所以，先找出直径。然后分割图。再找直径。
6. 复杂性是多少呢？第一次找到长为d的直径。然后剩下的树的直径肯定不会超过d（甚至不会达到d）
7. 感觉衰减会很快的
8. 果然超时
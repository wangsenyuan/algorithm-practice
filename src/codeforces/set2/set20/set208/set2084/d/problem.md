# Problem Statement

You are given three integers $n$, $m$, and $k$, where $m \cdot k < n$.

For a sequence $b$ consisting of non-negative integers, define $f(b)$ as follows:

You may perform the following operation on $b$:

Let $l$ denote the current length of $b$. Choose a positive integer $1 \leq i \leq l - k + 1$, remove the subarray from index $i$ to $i + k - 1$ and concatenate the remaining parts. In other words, replace $b$ with:

$$
[b_1, b_2, \ldots, b_{i-1}, b_{i+k}, b_{i+k+1}, \ldots, b_l]
$$

$f(b)$ is defined as the minimum possible value of $\text{mex}(b)$ after performing the above operation at most $m$ times (possibly zero).

You need to construct a sequence $a$ of length $n$ consisting of non-negative integers, such that:

- For all $1 \leq i \leq n$, $0 \leq a_i \leq 10^9$.
- Over all such sequences $a$, $f(a)$ is maximized.

---

**Note:**

The minimum excluded (MEX) of a collection of integers $c_1, c_2, \ldots, c_k$ is defined as the smallest non-negative integer $x$ which does not occur in the collection $c$.

---

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each test case contains three integers $n$, $m$, and $k$ ($2 \leq n \leq 2 \cdot 10^5$, $1 \leq m < n$, $1 \leq k < n$, $1 \leq m \cdot k < n$).

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output $n$ integers $a_1, a_2, \ldots, a_n$ ($0 \leq a_i \leq 10^9$).

If there are multiple answers, print any of them.


### ideas
1. 假设有两个人，一个是A，一个是B，A设定一个长为n的序列，B进行操作
2. 每次操作，删除连续的k个字符，且最多操作m次。
3. B的目的是，最后保留的序列的mex越小越好。A的目的是让这个mex越大越好
4. 如果不存在0的话， mex = 0. 所以肯定要存在0，且两个0的距离要 > m * k(否则，进行m次操作后，就可以把0给删除掉)
5. 同理两个1的距离要操作 m * k
6. 然后是2，直到n，其他的可以随便
7. 不仅仅和m*k有关系，还和k有关系
# Problem Statement

The **median** of an array $b_1, b_2, \ldots, b_m$, written as $\operatorname{med}(b_1, b_2, \ldots, b_m)$, is the $\lceil m/2 \rceil$-th\* smallest element of array $b$.

You are given an array of integers $a_1, a_2, \ldots, a_n$ and an integer $k$. You need to determine whether there exists a pair of indices $1 \leq l < r < n$ such that:

$$
\operatorname{med}(\operatorname{med}(a_1, \ldots, a_l),\ \operatorname{med}(a_{l+1}, \ldots, a_r),\ \operatorname{med}(a_{r+1}, \ldots, a_n)) \leq k.
$$

In other words, determine whether it is possible to split the array into **three contiguous subarrays**\dagger such that the median of the three subarray medians is less than or equal to $k$.

---

\* $\lceil x \rceil$ is the ceiling function which returns the least integer greater than or equal to $x$.

\dagger An array $x$ is a subarray of an array $y$ if $x$ can be obtained from $y$ by the deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end.

---

## Input
- Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$).
- The description of the test cases follows.
- The first line of each test case contains two integers $n$ and $k$ ($3 \leq n \leq 2 \cdot 10^5$, $1 \leq k \leq 10^9$) — the length of the array $a$ and the constant $k$.
- The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^9$) — the elements of the array $a$.
- It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output
For each testcase, output `YES` if such a split exists, and `NO` otherwise.

You can output the answer in any case (upper or lower). For example, the strings `yEs`, `yes`, `Yes`, and `YES` will be recognized as positive responses.

## ideas
1. 固定l的时候，就可以知道第一段的median x
2. 如果x <= k, 那么需要在后端再找出一个y <= k的median就可以了
3. 如果x > k, 那么就需要在后端再找出两个y <= k
4. 如果存在一个r, r + 1 后缀的median y <= k， 那么第一种情况，只要找到最大的r就可以了
5. 如果不存在这样的r，那么就需要找到中间的部分是否ok
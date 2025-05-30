# Problem Statement

You have an array $a[1], a[2], \ldots, a[n]$, containing **distinct integers from 1 to n**. Your task is to sort this array in increasing order using the following operation (you may need to apply it multiple times):

- **Choose two indexes** $i$ and $j$ ($1 \leq i < j \leq n$; $(j - i + 1)$ is a prime number).
- **Swap** the elements on positions $i$ and $j$; in other words, you are allowed to apply the following sequence of assignments:

  ```
  tmp = a[i]
  a[i] = a[j]
  a[j] = tmp
  ```

You do **not** need to minimize the number of used operations. However, you need to make sure that there are **at most $5n$ operations**.

---

## Input
- The first line contains integer $n$ ($1 \leq n \leq 10^5$).
- The next line contains $n$ **distinct integers** $a[1], a[2], \ldots, a[n]$ ($1 \leq a[i] \leq n$).

## Output
- In the first line, print integer $k$ ($0 \leq k \leq 5n$) — the number of used operations.
- Next, print the operations. Each operation must be printed as `i j` ($1 \leq i < j \leq n$; $(j - i + 1)$ is a prime).

> If there are multiple answers, you can print any of them.


### ideas
1. 2 is prime, and 3 is prime
2. 也就是说，可以使用相邻的交换（但是次数不一定能保证）
3. 那些需要移动的，会组成一个个的环，只交换环上的位置。这样子不会出现多余的交互
4. 如果要交换i, j且 j - i + 1 不是质数，
5. 那必须找到一个合适的位置，来交换它们
6. 假设这个位置是k,如果它在i的前面, i - k + 1 是个质数， j - k + 1 也是个质数
7. 是不是在i的中间，肯定能找到这样一个k, k - i + 1 是个质数, j - k + 1 是个质数
8. 如果不存在，肯定能找到一个序列，将它们链接起来？
9. 然后找到最大的质数比如k, 将j移动到 j1 = j - k + 1
10. 不确定行不行
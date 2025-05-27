# Problem D - Codeforces 279D

You've got a positive integer sequence $a_1, a_2, \ldots, a_n$. All numbers in the sequence are distinct. Let's fix the set of variables $b_1, b_2, \ldots, b_m$. Initially, each variable $b_i$ ($1 \leq i \leq m$) contains the value of zero. Consider the following sequence, consisting of $n$ operations.

The first operation is assigning the value of $a_1$ to some variable $b_x$ ($1 \leq x \leq m$). Each of the following $n-1$ operations is assigning to some variable $b_y$ the value that is equal to the sum of values that are stored in the variables $b_i$ and $b_j$ ($1 \leq i, j, y \leq m$). At that, the value that is assigned on the $t$-th operation must equal $a_t$. For each operation, numbers $y, i, j$ are chosen anew.

Your task is to find the minimum number of variables $m$, such that those variables can help you perform the described sequence of operations.

---

## Input

The first line contains integer $n$ ($1 \leq n \leq 23$).

The second line contains $n$ space-separated integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_k \leq 10^9$).

It is guaranteed that all numbers in the sequence are distinct.

---

## Output

In a single line, print a single number — the minimum number of variables $m$, such that those variables can help you perform the described sequence of operations.

If you cannot perform the sequence of operations at any $m$, print $-1$.


## ideas
1. 什么时候必须有一个新的变量？
2. 假设a[t] = b[i] + b[j]
3. 并且let b[y] = a[t] y是一个已存在变量
4. 如果b[y]在后续的计算中，不被使用到，那么这个y的选择就是安全的
5. 否则就必须使用一个新的变量
6. 
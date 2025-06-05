# Problem: The Little Elephant and the Permutations

The Little Elephant has two permutations $a$ and $b$ of length $n$, consisting of numbers from $1$ to $n$, inclusive. Let's denote the $i$-th ($1 \leq i \leq n$) element of the permutation $a$ as $a_i$, the $j$-th ($1 \leq j \leq n$) element of the permutation $b$ as $b_j$.

The **distance** between permutations $a$ and $b$ is the minimum absolute value of the difference between the positions of the occurrences of some number in $a$ and in $b$. More formally, it's such minimum $|i - j|$, that $a_i = b_j$.

A **cyclic shift** number $i$ ($1 \leq i \leq n$) of permutation $b$ consisting of $n$ elements is a permutation $b_i b_{i+1} \ldots b_n b_1 b_2 \ldots b_{i-1}$. Overall, a permutation has $n$ cyclic shifts.

The Little Elephant wonders, for all cyclic shifts of permutation $b$, what is the distance between the cyclic shift and permutation $a$?

---

## Input

The first line contains a single integer $n$ ($1 \leq n \leq 10^5$) — the size of the permutations.

The second line contains permutation $a$ as $n$ distinct numbers from $1$ to $n$, inclusive. The numbers are separated with single spaces.

The third line contains permutation $b$ in the same format.

---

## Output

In $n$ lines print $n$ integers — the answers for cyclic shifts. Print the answers to the shifts in the order of the shifts' numeration in permutation $b$, that is, first for the 1-st cyclic shift, then for the 2-nd, and so on.

### ideas
1. [1, 2, 3, 4] [a, b, c, d]
2. [1, 2, 3, 4] [d, a, b, c]
3. [1, 2, 3, 4] [c, d, a, b]
4. [1, 2, 3, 4, 1, 2, 3, 4], [a, b, c, d]
5. 在double a的序列上，移动b更容易操作
6. 往后移动一次后， 1从负的堆里面转移到了正的堆里面
7. 原来那些是1的移动到了负数的堆里面
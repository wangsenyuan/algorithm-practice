Sereja has two sequences $a$ and $b$ and number $p$. Sequence $a$ consists of $n$ integers $a_1, a_2, \ldots, a_n$. Similarly, sequence $b$ consists of $m$ integers $b_1, b_2, \ldots, b_m$. 

As usual, Sereja studies the sequences he has. Today he wants to find the number of positions $q$ (where $q + (m - 1) \cdot p \leq n$ and $q \geq 1$), such that sequence $b$ can be obtained from sequence $a_q, a_{q + p}, a_{q + 2p}, \ldots, a_{q + (m - 1)p}$ by rearranging elements.

Sereja needs to rush to the gym, so he asked to find all the described positions of $q$.

### Input

The first line contains three integers $n$, $m$ and $p$ ($1 \leq n, m \leq 2 \cdot 10^5$, $1 \leq p \leq 2 \cdot 10^5$). 

The next line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^9$). 

The next line contains $m$ integers $b_1, b_2, \ldots, b_m$ ($1 \leq b_i \leq 10^9$).

### Output

In the first line print the number of valid $q$s. 

In the second line, print the valid values in the increasing order.

### Examples

**Example 1:**

Input:
```
5 3 1
1 2 3 2 1
1 2 3
```

Output:
```
2
1 3
```

**Example 2:**

Input:
```
6 3 2
1 3 2 2 3 1
1 2 3
```

Output:
```
2
1 2
```


### ideas
1. p是间隔，固定大小，每隔p个元素，选取，选取m个数，这m个数，正好（重排)是b
2. 位置按照求余p进行分组（同一组内，单独处理）
3. 
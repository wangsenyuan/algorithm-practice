**Simon has an array** $a_1, a_2, \ldots, a_n$ **consisting of** $n$ **positive integers. Today Simon asked you to find a pair of integers** $l, r$ **($1 \leq l \leq r \leq n$), such that the following conditions hold:**

1. There is an integer $j$ ($l \leq j \leq r$), such that all integers $a_l, a_{l+1}, \ldots, a_r$ are divisible by $a_j$.
2. The value $r - l$ takes the maximum value among all pairs for which condition 1 is true.

**Help Simon find the required pair(s) of numbers $(l, r)$. If there are multiple required pairs, find all of them.**

---

### Input

- The first line contains integer $n$ ($1 \leq n \leq 3 \cdot 10^5$).
- The second line contains $n$ space-separated integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^6$).

---

### Output

- Print two integers in the first line: the number of required pairs and the maximum value of $r - l$.
- On the following line, print all $l$ values from optimal pairs in increasing order.

---

### ideas
1. a[j]是 a[l...r]的gcd
2. 找到这样的j，找它左边的l, a[l...j]是a[j]的倍数
3. 找到它右边, a[j...r]是a[j]的倍数， 那么这样的a[j],肯定是[l...r]的gcd
4. a[i] <= $10^6$
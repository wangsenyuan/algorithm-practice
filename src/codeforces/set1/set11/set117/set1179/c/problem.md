Serge came to the school dining room and discovered that there is a big queue here. There are **m** pupils in the queue. He's not sure now if he wants to wait until the queue will clear, so he wants to know which dish he will receive if he does. As Serge is very tired, he asks you to compute it instead of him.

Initially there are **n** dishes with costs $a_1, a_2, \ldots, a_n$. As you already know, there is a queue of **m** pupils who have $b_1, \ldots, b_m$ togrogs respectively (pupils are enumerated by queue order, i.e., the first pupil in the queue has $b_1$ togrogs and the last one has $b_m$ togrogs).

Pupils think that the most expensive dish is the most delicious one, so every pupil just buys the most expensive dish for which he has money (every dish has a single copy, so when a pupil has bought it nobody can buy it later), and if a pupil doesn't have money for any dish, he just leaves the queue (so brutal capitalism...)

But money isn't a problem at all for Serge, so Serge is buying the most expensive dish if there is at least one remaining.

Moreover, Serge's school has a very unstable economic situation and the costs of some dishes or number of togrogs of some pupils can change. More formally, you must process **q** queries:

- `1 i x`: Change $a_i$ to $x$. It means that the price of the $i$-th dish becomes $x$ togrogs.
- `2 i x`: Change $b_i$ to $x$. It means that the $i$-th pupil in the queue has $x$ togrogs now.

Nobody leaves the queue during those queries because a saleswoman is late.

After every query, you must tell Serge the price of the dish which he will buy if he has waited until the queue is clear, or $-1$ if there are no dishes at this point, according to rules described above.

---

### Input

The first line contains integers $n$ and $m$ ($1 \leq n, m \leq 300\ 000$) — number of dishes and pupils respectively.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^6$) — elements of array $a$.

The third line contains $m$ integers $b_1, b_2, \ldots, b_m$ ($1 \leq b_i \leq 10^6$) — elements of array $b$.

The fourth line contains integer $q$ ($1 \leq q \leq 300\ 000$) — number of queries.

Each of the following $q$ lines contains as follows:

- If a query changes price of some dish, it contains `1 i x` ($1 \leq i \leq n$, $1 \leq x \leq 10^6$), which means $a_i$ becomes $x$.
- If a query changes number of togrogs of some pupil, it contains `2 i x` ($1 \leq i \leq m$, $1 \leq x \leq 10^6$), which means $b_i$ becomes $x$.

### Output

For each of $q$ queries, print the answer as the statement describes, the answer of the $i$-th query in the $i$-th line (the price of the dish which Serge will buy or $-1$ if nothing remains).

### ideas
1. 如果没有change，只需要模拟就可以了。
2. 但是change后，就很麻烦了
3. 先考虑价格变化的影响。如果a[i] 变小了。那么比a[i]大的那些dish，没有影响（还是按照原来的顺序处理）
4. dish被谁吃掉，其实没有关系。重要的是，有没有被吃掉
5. a[i]被吃掉的条件是, 假设a[i]的位置，从大大小，是x，那么如果存在至少x个小孩的钱，大于等于a[i]， 那么就肯定被吃掉了
6. 也就是找最大的a[i], sum(b[?] >= a[i]) < position of a[i]
7. 进步一下。但怎么维护这个信息呢？
8. 感觉还是得用segment tree？
9. 把a[i]变大，那么就是将这个范围内的数的位置+1
10. a[i]的位置-w
11. 所以感觉上是可以用range update的
12. 如果有一个人有b[i]， 那么代表 <= b[i]的价格（范围）-1
13. 如果有一个产品的价格为a[i], 那么代表 <= a[i]的范围+1
14. 那么就是找到最大的cnt[x] > 0 的那个点
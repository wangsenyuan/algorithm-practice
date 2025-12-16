# Problem Description

As you might remember from our previous rounds, Vova really likes computer games. Now he is playing a strategy game known as Rage of Empires.

In the game Vova can hire $n$ different warriors; $i$-th warrior has the type $a_i$. Vova wants to create a balanced army hiring some subset of warriors. An army is called balanced if for each type of warrior present in the game there are not more than $k$ warriors of this type in the army. Of course, Vova wants his army to be as large as possible.

To make things more complicated, Vova has to consider $q$ different plans of creating his army. $i$-th plan allows him to hire only warriors whose numbers are not less than $l_i$ and not greater than $r_i$.

Help Vova to determine the largest size of a balanced army for each plan.

Be aware that the plans are given in a modified way. See input section for details.

## Input

The first line contains two integers $n$ and $k$ ($1 \leq n, k \leq 10^5$).

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^5$).

The third line contains one integer $q$ ($1 \leq q \leq 10^5$).

Then $q$ lines follow. $i$-th line contains two numbers $x_i$ and $y_i$ which represent $i$-th plan ($1 \leq x_i, y_i \leq n$).

You have to keep track of the answer to the last plan (let's call it $last$). In the beginning $last = 0$. Then to restore values of $l_i$ and $r_i$ for the $i$-th plan, you have to do the following:

$l_i = ((x_i + last) \bmod n) + 1$;

$r_i = ((y_i + last) \bmod n) + 1$;

If $l_i > r_i$, swap $l_i$ and $r_i$.

## Output

Print $q$ numbers. $i$-th number must be equal to the maximum size of a balanced army when considering $i$-th plan.

## Examples

### Example 1

**Input:**

```text
6 2
1 1 1 2 2 2
5
1 6
4 3
1 1
2 6
2 6
```

**Output:**

```text
2
4
1
3
2
```

## Note

In the first example the real plans are:

```text
1 2
1 6
6 6
2 4
4 6
```


### ideas
1. 先不考虑区间，在整个数组上，对于每种type，按照它的freq排序
2. 那么结果就是 min(k, freq[i]) * (m - i)
3. 当freq[i] >= k 的时候，尽量的选,
4. 当 freq[i] < k 的时候，就不一定了，就需要检查每种可能的取值
5. k + k + ... + freq[i] 超过k的，只能选择k，不超过k的全部参加
6. 所以，关键在于找到区间内，超过k的数量，以及它们的数量和总和
7. 用MO + RMQ， q * 330 * 20 = 差不多1e9了
8. 而且也没法用MO
9. 假设将type分成两类，一类是freq[i] <= k的， 一类是freq[i] > k 的
10. 第一类数字应该是比较多。但是很稀疏
11. 第二类数字比较少，但是很稠密
12. 第一类，直接算在区间和里面 pref[r] - pref[l-1]
13. 第二类，在查询区间内，不一定
14. 假设在r的位置， l = f(r) 表示在区间 l...r中间 a[r]出现的次数 = k
15. 这样子，貌似能计数（使用持久化树，查找 >= l 的值）
16. 但是无法知道sum？
One day Natalia was walking in the woods when she met a little mushroom gnome. The gnome told her the following story:

Everybody knows that the mushroom gnomes' power lies in the magic mushrooms that grow in the native woods of the gnomes. There are $n$ trees and $m$ magic mushrooms in the woods: the $i$-th tree grows at a point on a straight line with coordinate $a_i$ and has height $h_i$; the $j$-th mushroom grows at the point with coordinate $b_j$ and has magical power $z_j$.

But one day wild mushroommunchers, the sworn enemies of mushroom gnomes, unleashed a terrible storm on their home forest. As a result, some of the trees began to fall and crush the magic mushrooms. The supreme oracle of mushroom gnomes calculated in advance the probability for each tree that it will fall to the left, to the right, or will stand on. If the tree with coordinate $x$ and height $h$ falls to the left, then all the mushrooms that belong to the right-open interval $[x - h, x)$ are destroyed. If a tree falls to the right, then the mushrooms that belong to the left-open interval $(x, x + h]$ are destroyed. Only those mushrooms that are not hit by a single tree survive.

Knowing that all the trees fall independently of each other (i.e., all the events are mutually independent, and the trees do not interfere with other trees falling in an arbitrary direction), the supreme oracle was also able to quickly calculate the expectation of the total power of the mushrooms which survived after the storm. His calculations ultimately saved the mushroom gnomes from imminent death.

Natalia, as a good Olympiad programmer, got interested in this story, and she decided to come up with a way to quickly calculate the expectation of the sum of the surviving mushrooms' power.

## Input

The first line contains two integers $n$ and $m$ ($1 \le n \le 10^5$, $1 \le m \le 10^4$) — the number of trees and mushrooms, respectively.

Each of the next $n$ lines contains four integers $a_i$, $h_i$, $l_i$, $r_i$ ($|a_i| \le 10^9$, $1 \le h_i \le 10^9$, $0 \le l_i, r_i$, $l_i + r_i \le 100$) which represent the coordinate of the $i$-th tree, its height, and the percentage probabilities that the tree falls to the left and to the right, respectively (the remaining percentage is the probability that the tree will stand on).

Each of the next $m$ lines contains two integers $b_j$, $z_j$ ($|b_j| \le 10^9$, $1 \le z_j \le 10^3$) which represent the coordinate and the magical power of the $j$-th mushroom, respectively.

An arbitrary number of trees and mushrooms can grow at one point.

## Output

Print a real number — the expectation of the total magical power of the surviving mushrooms. The result is accepted with relative or absolute accuracy $10^{-4}$.

## Examples

### Input

```
1 1
2 2 50 50
1 1
```

### Output

```
0.5000000000
```

### Input

```
2 1
2 2 50 50
4 2 50 50
3 1
```

### Output

```
0.2500000000
```

## Note

It is believed that the mushroom with coordinate $x$ belongs to the right-open interval $[l, r)$ if and only if $l \le x < r$. Similarly, the mushroom with coordinate $x$ belongs to the left-open interval $(l, r]$ if and only if $l < x \le r$.

In the first test, the mushroom survives with probability 50%, depending on where the single tree falls.

In the second test, the mushroom survives only if neither of the two trees falls on it. This occurs with probability $50\% \times 50\% = 25\%$.

Pretest №12 is the large test with $10^5$ trees and one mushroom.

### Ideas

1. 因为有多棵树长在同一个位置，所以只考虑一棵树的影响，似乎是不够的
2. 假设一个蘑菇长的位置 $x$，被几棵树覆盖到。它被压到的概率 $= 1.0 -$ 这些树倒向另一边的概率乘积吗？
3. 然后它 survive 的概率 =

### Solution

First of all — the answer is the sum for all mushrooms of the probabilities of not being destroyed, multiplied by that mushroom's power. That is a simple property of random variables' means.

So we come to the equivalent statement: we still have mushrooms, but now instead of trees we have a family of segments with probabilities assigned to them. Every segment "exists" with this probability, otherwise it doesn't, and all these events are independent. We want to count the sum of probabilities (with weights) for each mushroom not to lie in any "existing" segment. (Note that we can reformulate the statement this way because any segments containing any fixed point are truly independent: they can't belong to the same tree. Thus the probability to survive for any point in this statement is equal to the probability for this point in the original statement.)

Now, how do we count this? There are several ways:

**1) Scanning line.** If we go from left to right, we can meet three kinds of events: "the segment $i$ started", "the segment $i$ finished", "the mushroom $j$ found". We can easily maintain the probability of the current point being covered by an "existing" segment if we multiply it by the segment's probability when we find its beginning and divide by it when we find its end. If we find a mushroom along the way, we can add the known probability to the answer (multiplied by its power). To perform the above trick we just sort the array of events by $x$-coordinate and iterate over it.

This solution is good in theory, but in practice it has a flaw: if the number of segments is large, after multiplying lots of real numbers less than 1 we can exceed the negative exponent of the real type used, and thus get a 0 in a variable instead of the desired value. And after any number of divisions it still would be 0, so we couldn't get any sane answer anymore.

This trouble can be resolved in several ways (without changing the solution much):

- **(a)** We can have no more than 101 distinct values of probabilities for segments. So, if we store an array for quantities of segments containing the current point and having a corresponding probability, we just add and subtract 1's from the array's elements. When we find a mushroom we find the product of degrees with exponents stored in the array, spending ~100 operations.

  **Explanation:** Every segment's probability is of the form $k/100$ where $k \in [0, 100]$ (because $l_i, r_i$ are integer percentages). So there are at most 101 distinct probability values. Instead of maintaining a running product (which underflows), we keep a count array $\text{cnt}[0 \ldots 100]$, where $\text{cnt}[k]$ = number of currently active segments whose "exists" probability is $k/100$.

  - When a segment with probability $k/100$ **starts**: $\text{cnt}[k] \mathrel{+}= 1$
  - When a segment with probability $k/100$ **ends**: $\text{cnt}[k] \mathrel{-}= 1$
  - When we encounter a mushroom, compute the survival probability from scratch:
    $$P_{\text{survive}} = \prod_{k=0}^{100} \left(\frac{100 - k}{100}\right)^{\text{cnt}[k]}$$

  Each factor $(100 - k)/100$ is the probability that a single segment with probability $k/100$ does **not** destroy the mushroom, raised to the power of how many such segments currently cover this point. This takes ~100 multiplications per mushroom.

  **Example:** Suppose at some mushroom's position, three segments are active: two with probability $50/100$ and one with probability $30/100$. Then $\text{cnt}[50] = 2$, $\text{cnt}[30] = 1$, and all other counts are 0. The survival probability is:

  $$P = \left(\frac{50}{100}\right)^{2} \times \left(\frac{70}{100}\right)^{1} = 0.5^2 \times 0.7 = 0.175$$

  The key advantage: exponents $\text{cnt}[k]$ are exact integers, so no precision is lost in tracking them, and the product is recomputed fresh each time (avoiding cumulative underflow).

- **(b)** We can store a set of segments containing the current point. Every operation with the set works in $O(\log N)$ time, and iterating over the whole set works in $O(N)$ time. So, upon meeting a mushroom we iterate over the set and multiply the probabilities for all segments in it. The next thing that helps us is that we can drop the answer for the current mushroom if it's too small. If we don't store the segments with probability 1, the most number of segments whose probabilities' product is more than $10^{-8}$ is about 2000 (since $0.99^{2000} < 10^{-8}$). So we can count everything in time.

- **(c)** If we use logs of probabilities instead of themselves, we have to add and subtract them instead of multiplying and dividing. This way we won't encounter any precision troubles.

**2) Segment tree.** Let's sort the mushrooms by their coordinates. Let's also assume we have some set of segments and already counted the desired probabilities. Now we want to add a new segment to the set. What will change? The probabilities of mushrooms lying in this segment (and thus forming a segment in the array) will be multiplied by the segment's probability. Now it's clear we can use a multiplication segment tree (or a simple addition segment tree if we use logs again) to perform the queries for all segments and then sum up the elements in the end.

**About the strange score and pretest:** we discovered the trouble with precision quite late, and realized that it makes the problem way harder (because it's hard to predict during writing and submission phases). What's worse, it won't show itself on the small tests. So we decided to "show up" the test and let the contestants solve this additional problem, for additional score. (However, not all solutions from the above list do actually deal with this problem. Unfortunately, we didn't come up with them beforehand.)
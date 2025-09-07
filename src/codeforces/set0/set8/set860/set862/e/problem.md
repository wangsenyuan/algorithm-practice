# Problem E

## Problem Statement

Dr. Evil is interested in math and functions, so he gave Mahmoud and Ehab array `a` of length `n` and array `b` of length `m`. He introduced a function `f(j)` which is defined for integers `j`, which satisfy $0 \leq j \leq m - n$.

Suppose, $c_i = a_i - b_{i+j}$. Then $f(j) = |c_1 - c_2 + c_3 - c_4 + \ldots + c_n|$. More formally:

$$f(j) = \left|\sum_{i=1}^{n} (-1)^{i+1} \cdot c_i\right|$$

where $c_i = a_i - b_{i+j}$.

Dr. Evil wants Mahmoud and Ehab to calculate the minimum value of this function over all valid `j`. They found it a bit easy, so Dr. Evil made their task harder. He will give them `q` update queries. During each update they should add an integer $x_i$ to all elements in `a` in range $[l_i, r_i]$ i.e. they should add $x_i$ to $a_{l_i}, a_{l_i+1}, \ldots, a_{r_i}$ and then they should calculate the minimum value of $f(j)$ for all valid `j`.

Please help Mahmoud and Ehab.

## Input

The first line contains three integers `n`, `m` and `q` ($1 \leq n \leq m \leq 10^5$, $1 \leq q \leq 10^5$) — number of elements in `a`, number of elements in `b` and number of queries, respectively.

The second line contains `n` integers $a_1, a_2, \ldots, a_n$ ($-10^9 \leq a_i \leq 10^9$) — elements of `a`.

The third line contains `m` integers $b_1, b_2, \ldots, b_m$ ($-10^9 \leq b_i \leq 10^9$) — elements of `b`.

Then `q` lines follow describing the queries. Each of them contains three integers $l_i$, $r_i$, $x_i$ ($1 \leq l_i \leq r_i \leq n$, $-10^9 \leq x_i \leq 10^9$) — range to be updated and added value.

## Output

The first line should contain the minimum value of the function `f` before any update.

Then output `q` lines, the `i`-th of them should contain the minimum value of the function `f` after performing the `i`-th update.

## Example

### Input
```
5 6 3
1 2 3 4 5
1 2 3 4 5 6
1 1 10
1 1 -9
1 5 -1
```

### Output
```
0
9
0
0
```

## Note

For the first example before any updates it's optimal to choose $j = 0$:
$$f(0) = |(1 - 1) - (2 - 2) + (3 - 3) - (4 - 4) + (5 - 5)| = |0| = 0$$

After the first update `a` becomes $\{11, 2, 3, 4, 5\}$ and it's optimal to choose $j = 1$:
$$f(1) = |(11 - 2) - (2 - 3) + (3 - 4) - (4 - 5) + (5 - 6)| = |9| = 9$$

After the second update `a` becomes $\{2, 2, 3, 4, 5\}$ and it's optimal to choose $j = 1$:
$$f(1) = |(2 - 2) - (2 - 3) + (3 - 4) - (4 - 5) + (5 - 6)| = |0| = 0$$

After the third update `a` becomes $\{1, 1, 2, 3, 4\}$ and it's optimal to choose $j = 0$:
$$f(0) = |(1 - 1) - (1 - 2) + (2 - 3) - (3 - 4) + (4 - 5)| = |0| = 0$$


## Solutions

Let's write f(j) in another way:





Now we have 2 sums:
- The first one is constant (doesn't depend on j)
- For the second sum we can calculate all its possible values using sliding window technique
- We want a data-structure that takes the value of the first sum and chooses the best second sum from all the choices

**Observation:** We don't have to try all the possible values of f(j) to minimize the expression, If the first sum is c, We can try only the least value greater than  - c and the greatest value less than  - c ( - c not c because we are minimizing c + second not c - second) because the absolute value means the distance between the two values on the number line, Any other value will be further than at least one of the chosen values, To do this we can keep all the values of f(j) sorted and try the elements numbered lower_bound(-c) and lower_bound(-c)-1 and choose the better (In short we're trying the values close to  - c only).

Now we have a data-structure that can get us the minimum value of the expression once given the value of the first sum in O(log(n)). Now we want to keep track of the value of the first sum.

Let the initial value be c. In each update:
- If the length of the updated interval is even, the sum won't change because x will be added as many times as it's subtracted
- Otherwise x will be added to c or subtracted from c depending on the parity of l (the left-bound of the interval)

**Time complexity:** O(n + (m + q)log(m)) .
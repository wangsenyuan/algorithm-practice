Arpa has found a list containing n numbers. He calls a list bad if and only if it is not empty and gcd (see notes section for more information) of numbers in the list is 1.

Arpa can perform two types of operations:

- Choose a number and delete it with cost x.
- Choose a number and increase it by 1 with cost y.

Arpa can apply these operations to as many numbers as he wishes, and he is allowed to apply the second operation arbitrarily many times on the same number.

Help Arpa to find the minimum possible cost to make the list good.

## Input

First line contains three integers n, x and y (1 ≤ n ≤ 5×10⁵, 1 ≤ x, y ≤ 10⁹) — the number of elements in the list and the integers x and y.

Second line contains n integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁶) — the elements of the list.

## Output

Print a single integer: the minimum possible cost to make the list good.

## Examples

**Input:**
```
4 23 17
1 17 17 16
```

**Output:**
```
40
```

**Input:**
```
10 6 2
100 49 71 73 66 96 8 60 41 63
```

**Output:**
```
10
```

## Note

In example, number 1 must be deleted (with cost 23) and number 16 must increased by 1 (with cost 17).


### ideas
1. dp[x]表示gcd(arr) = x 的数量（这些是不用变的）其他的数，要么递增到x，要么删除
2. 考虑这样一个序列， 很多个3，比较少的4
3. 假设最优解是把所有的4变成6（增加2）
4. 如果还有个2，就更难搞了
5. 对于一个x来说， dp[x] = 把所有的数，通过操作2变成x的倍数的cost
6. 对于i来说，比他小的数, (i - j) *y
7. 比i大的数, j = (j - j / i * i) * y = j * y - (j / i * i) * y
8. 前半部分固定，后半部分 = 在区间 j / i * i 到 下一段中的数量 （这个可以用diff数组）
9. 
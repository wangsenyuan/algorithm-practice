# Problem Description

Given a sequence of integers a₁, ..., aₙ and q queries x₁, ..., xₑ on it. For each query xᵢ you have to count the number of pairs (l, r) such that 1 ≤ l ≤ r ≤ n and gcd(aₗ, aₗ₊₁, ..., aᵣ) = xᵢ.

**Note:** gcd is a greatest common divisor of v₁, v₂, ..., vₙ, that is equal to a largest positive integer that divides all vᵢ.

## Input

The first line of the input contains integer n, (1 ≤ n ≤ 10⁵), denoting the length of the sequence. The next line contains n space separated integers a₁, ..., aₙ, (1 ≤ aᵢ ≤ 10⁹).

The third line of the input contains integer q, (1 ≤ q ≤ 3 × 10⁵), denoting the number of queries. Then follows q lines, each contain an integer xᵢ, (1 ≤ xᵢ ≤ 10⁹).

## Output

For each query print the result in a separate line.

## Examples

### Example 1

**Input:**
```
3
2 6 3
5
1
2
3
4
6
```

**Output:**
```
1
2
2
0
1
```

### Example 2

**Input:**
```
7
10 20 3 15 1000 60 16
10
1
2
3
4
5
6
10
20
60
1000
```

**Output:**
```
14
0
2
2
2
0
2
2
1
1
```

### ideas
1. f(x) 表示能够整除x的子序列的个数
2. g(x) = f(x) - g(2 * x) - g(3 * x), ...
3. 如何快速计算f(x)?
4. 这里x也太大了
5. 假设gcd(a[l...r]) = x, 增加r+1, 那么x要么保持不变，要么变成x的一个因子，这个递减应该是比较快的
6. 假设在每个位置r处，维护一个pair的列表，表示当gcd(l...r) = x时的，最左边的l是哪里
7. 这里的存储空间，应该在n * 30 左右; 然后按照x排序
8. 这样子能找到部分，但是少了一部分，就是考虑 x x x 2 * x,,,,, w 这样的序列
9. 所以，还需要知道l...r 的个数
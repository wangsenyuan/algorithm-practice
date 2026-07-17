# Problem Statement

You are given a permutation P of (1,2,…,N) and an integer K.

Find the expected value, modulo 998244353, of the inversion number of P after performing the following operation:

First, choose an integer i uniformly at random between 1 and N−K+1, inclusive.
Then, shuffle P_i, P_{i+1}, …, P_{i+K−1} uniformly at random.

## What is the inversion number?

The inversion number of a sequence (A_1, A_2, …, A_N) is the number of integer pairs (i,j) satisfying 1≤i<j≤N and A_i > A_j.

## What does "expected value modulo 998244353" mean?

It can be proved that the sought expected value is always rational. Under the constraints of this problem, when this value is represented as an irreducible fraction P/Q, it can also be proved that Q ≢ 0 (mod 998244353). Thus, there is a unique integer R satisfying R×Q≡P (mod 998244353), 0≤R<998244353. Report this integer R.

## Constraints

All input values are integers.

- 1≤K≤N≤2×10^5
- P is a permutation of (1,2,…,N).

## Input

The input is given from Standard Input in the following format:

```text
N K
P_1 P_2 … P_N
```

## Output

Print the answer in one line.

## Sample Input 1

```text
4 2
1 4 2 3
```

## Sample Output 1

```text
166374061
```

The operation changes the permutation P into the following:

- (1,4,2,3) ... probability 1/2
- (4,1,2,3) ... probability 1/6
- (1,2,4,3) ... probability 1/6
- (1,4,3,2) ... probability 1/6

The expected value of the inversion number is 2×1/2 + 3×1/6 + 1×1/6 + 3×1/6 = 13/6.

13/6 modulo 998244353 is 166374061, so print this number.

## Sample Input 2

```text
1 1
1
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
10 6
7 4 10 5 6 1 8 2 3 9
```

## Sample Output 3

```text
499122200
```


### ideas
1. 随机选择连续的长度为k的子串，将其翻转，计算期望的逆序对数量
2. 假设i = 0, 1, ... n - k
3. 然后计算翻转后的逆序对，然后除以 (n - k + 1)
4. 是个数据结构问题，在移动区间的过程中，更新逆序对的数量
5. 将r移入这个区间的时候， 首先减去它的贡献，然后加上它的贡献
6. ...不是翻转，是shuffle～～～
7. 
# Problem B

There is a connected simple undirected graph with N vertices and M edges, where the vertices are numbered 1,…,N and the edges are numbered 1,…,M. Edge i connects vertices Aᵢ and Bᵢ. Also, there are a total of N+1 cards, one each with the integers from 0 to N written on them.

Snuke performed the following operations in order:

1. Place one card on each vertex, and eat the remaining one card.
2. For each i (1≤i≤M), write on edge i the bitwise XOR of the integers written on the two cards placed on vertices Aᵢ and Bᵢ. Let this integer be Xᵢ.
3. Discard all cards placed on the vertices.

Given the information about the graph (N, M, A₁,…,Aₘ, B₁,…,Bₘ, X₁,…,Xₘ), determine the integer written on the card that Snuke ate. If it cannot be uniquely determined, output -1. It is guaranteed that the given X can be obtained by the above operations.

You are given T test cases; solve each of them.

## What is bitwise XOR?

The bitwise XOR of two integers a and b is the integer whose binary representation has 1 in positions where a and b differ, and 0 where they are the same.

## Constraints

- 1≤T≤10⁴
- 1≤N≤2×10⁵
- 0≤M≤2×10⁵
- 1≤Aᵢ,Bᵢ≤N
- The given graph is a connected simple undirected graph.
- X₁,…,Xₘ can be obtained by the operations in the problem statement.
- The sum of N over all test cases is at most 2×10⁵.
- The sum of M over all test cases is at most 2×10⁵.
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```
T
case₁
case₂
⋮
case_T
```

Each test case is given in the following format:

```
N M
A₁ B₁ X₁
A₂ B₂ X₂
⋮
Aₘ Bₘ Xₘ
```

## Output

Output T lines. The i-th line should contain the integer written on the card Snuke ate for the i-th test case if it can be uniquely determined, and -1 otherwise.

## Sample Input 1

```
3
2 1
1 2 3
1 0
4 4
2 1 4
1 3 5
1 4 7
2 3 1
```

## Sample Output 1

```
0
-1
2
```

## Sample Explanation 1

For the first test case, the graph has 2 vertices and 1 edge. The possible pairs of integers written on the cards placed on vertices 1 and 2 are (1,2) and (2,1); in either case, the integer written on the card Snuke ate is 0.

For the second test case, there are two possible integers written on the card Snuke ate: 0 and 1.


### ideas
1. x[i] = a[u] ^ a[v]
2. 考虑一个环 u - v - w - u
3. x[u, v] = a[u] ^ a[v]
4. x[v, w] = a[v] ^ a[w]
5. x[w, u] = a[w] ^ a[u]
6. xor 后 = 0
7. 考虑每一位的0/1的情况？通过最低位的x[i][0], 可以把节点分成2类（red, blue的数量)
8. red + blue = n
9. 且我们知道0...n+1的偶数部分 = (n + 2) / 2, 奇数部分 = (n + 1) / 2
10. 如果 n是偶数，比如 n = 4, (0, 1, 2, 3, 4)
11. 且 red = 2, blue = 2, 那么可以确定缺失的数肯定是偶数(0, 2, 4) 
12. 如果 red = 1, blue = 3, 那么缺失的数肯定是奇数(1, 3)
13. (3, 1)也是奇数
14. 如果 n是奇数，比如 n = 5, (0, 1, 2, 3, 4, 5)
15. 如果 red = 3, blue = 2, （没有答案）
16. 但是如果 red = 4, blue = 1 (不可能)
17. 居然正确～～～～
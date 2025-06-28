# Interactive Problem: Finding Array Lengths

This is an interactive problem.

## Problem Description

You found the numbers $k$ and $n$ in the attic, but lost two arrays $A$ and $B$.

You remember that:

- $|A| + |B| = n$ - the total length of the arrays is $n$
- $|A| \geq k$ and $|B| \geq k$ - the length of each array is at least $k$
- The arrays consist only of numbers from $1$ to $k$
- If you take any $k$ consecutive elements from array $A$, they will all be different. Also, if you take any $k$ consecutive elements from array $B$, they will all be different.

Fortunately, a kind spirit that settled in the attic found these arrays and concatenated them into an array $C$ of length $n$. That is, the elements of array $A$ were first written into array $C$, followed by the elements of array $B$.

You can ask the kind spirit up to **250 questions**. Each question contains an index $i$ ($1 \leq i \leq n$). In response, you will receive the $i$-th element of the concatenated array $C$.

You need to find the lengths of arrays $A$ and $B$, or report that it is impossible to determine them uniquely.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 300$). The description of the test cases follows.

The only line of each test case contains two integers $n$ and $k$ ($1 \leq k \leq 50$, $2k \leq n \leq 10^6$).

**Note:** The sum of $n$ across test cases is not limited.

## Interaction

The interaction for each test case begins with reading the integer $n$.

Then you can make up to **250 queries**.

### Making Queries

To make a query, output a string in the format:
```
? x
```
where $1 \leq x \leq n$. After each query, read an integer — the answer to your query.

**Warning:** If you make too many queries, you will receive a verdict of "Wrong answer".

### Reporting Answer

To report your answer, output a string in the format:
```
! a b
```
where $a$ and $b$ are the lengths of arrays $A$ and $B$ that you found, respectively. The answer is **not counted** when counting the number of queries.

If it is impossible to determine the lengths of the arrays uniquely, output:
```
! -1
```

**Important:** If you answer $-1$ while there is a sequence of at most 250 queries that uniquely determines the lengths of arrays, you will get a "Wrong answer" verdict.

## Guarantees

- It is guaranteed that there are arrays $A$ and $B$ that do not contradict the statement, for which the interactor output is correct.
- The interactor is **not adaptive**, which means that the answer is known before the participant makes queries and does not depend on the queries made by the participant.

## Important Notes

- If your program makes more than 250 queries, your program should **immediately terminate** to receive the verdict "Wrong answer".
- Otherwise, you can get an arbitrary verdict because your solution will continue to read from a closed stream.

## ideas
1. if n == 2 * k => (k, k) no need to query at all
2. if n <= 250 => can brute force
3. n > 250 and at least 5 segments
4. ask A in k (first k), ask B in k (last k)
5. 如果不存在B， 那么 arr[i] = A[i % k], 假设arr[i] = (w = B[?]) != A[i % k]
6. 那么可以二分查找j = i % k的位置，找到最小的r (A[r % k] != arr[r])的位置
7. 那么分割线肯定在(r - k, r)之间
8. 把这一段(k-1)次都问出来，然后检查是否能够在l+1, l+2.... r-1处进行分割，如果只有一个分割点 => ans 否则（有0个或者多个) -1
9. 
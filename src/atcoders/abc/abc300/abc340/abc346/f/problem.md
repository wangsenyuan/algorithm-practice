Let me summarize the problem described in the markdown file.

The problem involves two operations on strings:
1. `f(X,k)`: Repeats the entire string X k times. Example: `f("abc",2) = "abcabc"`
2. `g(X,k)`: Repeats each character of X k times in order. Example: `g("abc",3) = "aaabbbccc"`

Given:
- A positive integer N
- Two strings S and T

The task is to find the largest non-negative integer k such that `g(T,k)` is a subsequence of `f(S,N)`.

Note: By definition, `g(T,0)` (empty string) is always a subsequence of any string.


### ideas
1. 先判断是否能二分？
2. 假设g(T, k) 时 f(S, n)的子序列
3. 那么g(T, k-1) 肯定是 f(S, n)的子序列
4. 符合二分的条件
5. 考察怎么进行判断呢？
6. g(T, k) = T[0] ** k + g(T[1:], k)
7. 先要找到T[0]（重复k次后）在f(S, n)的位置
8. 那么就需要知道T[0]在S中的位置和数量
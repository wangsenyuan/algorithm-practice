# B. Another Problem about Beautiful Pairs

Source: <https://codeforces.com/problemset/problem/2196/B>

For each test case, count index pairs `i < j` such that `a[i] * a[j] = j - i`.


## ideas
1. a[?] > 0 => j > i
2. a[i] * a[j] = j - i
3. i = j - a[i] * a[j]
4. j - x * a[j] > 0 => x < j / a[j]
5. 知道x后, 在x上, 这个地方, 切好要找到 j - x * a[j]的i
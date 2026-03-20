# Problem

The famous writer Velepin is very productive. He signed a contract with a publication and now has to write `k_i` books in the `i`-th year.

He has `n` regular readers. In each year, every reader reads exactly one of the `k` books published that year.

Readers like discussing books. Reader `j` is satisfied in a year if at least `a_j` people read the same book as reader `j` (including reader `j`).

A reading service can choose which book each person reads. Velepin does not want books to be wasted, so each of the `k` books must be read by at least one person.

For each year (query), determine the maximum possible number of satisfied readers.

## Input

- First line: integer `n` (`2 <= n <= 3 * 10^5`) — number of readers.
- Second line: `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= n`) — required group size for each reader.
- Third line: integer `q` (`1 <= q <= 3 * 10^5`) — number of years (queries).
- Next `q` lines: each contains one integer `k_j` (`2 <= k_j <= n`) — number of books in year `j`.

## Output

Print `q` lines. For each query `k_j`, print one integer: the maximum number of satisfied readers if exactly `k_j` books are released and every book is read by at least one person.

## Examples

### Example 1

Input

```text
5
1 2 2 2 2
3
2
3
4
```

Output

```text
5
5
3
```

### Example 2

Input

```text
6
1 2 3 4 5 6
2
2
3
```

Output

```text
5
4
```

### Example 3

Input

```text
6
4 4 1 4 4 4
3
2
3
4
```

Output

```text
6
5
1
```

## Note

In Example 1:

- For `k = 2`, one optimal assignment is `1,2,2,2,2` (book labels). Satisfied readers: `5`.
- For `k = 3`, one optimal assignment is `1,2,2,3,3`. Satisfied readers: `5`.
- For `k = 4`, one optimal assignment is `1,2,3,4,2`. Satisfied readers: `3`.

In Example 2:

- For `k = 2`, one optimal assignment is `1,1,1,1,1,2`; everyone except the 6-th reader is satisfied.
- For `k = 3`, one optimal assignment is `1,1,1,1,2,3`; everyone except the 5-th and 6-th readers is satisfied.

## Summary

Partition the `n` readers into exactly `k` non-empty groups (one book per group). Reader `j` is happy iff their group has size at least `a_j`. For each query `k`, output the largest possible number of happy readers.


### ideas 
1. 一共k本书，每本书都需要被read，分配这些书，让最多的人满意
2. 按照a升序排列，显然越小的人，越容易被满足，假设序列是a = 1, 2, 3, 4, 5, 6, 7
3. 有k = 3, 那么第一个人读第一本书，（2， 3, 4）都第二本书，（5， 6， 7）读第3本书，
4. 那么有（1， 2， 3）个人本满足了
5. 最多是3个人被满足。假设有k本书时，有w个人可以被满足，那么k+1本书的时候，也至少有w个人可以被满足吗？好像不是，应该是反过来，k-1本书的时候，至少有w个人被满足。因为可以把多出来的那本书的读者重新分配，结果不会更差
6. 在k本书的情况下，先满足第一个人的要求, 那么必须有a[1]个人读book 1
7. 剩余k1 =  k - a[1]个人。然后应该满足a[2]（因为他是第二少要求的人）
8. 剩余k2 =  k1 - a[2]个人
9. .. k2 = k[w-1] - a[w]
10. k2 >= 0 => k - sum(a[1], a[w]) >= 0
11. w <= k 也需要成立
12. sum(a[:w]) <= k and w <= k
13. 但是这里不大对， 因为按照前面的考虑k越小，满足的人越多，这里不一致了。漏了条件；
14. 这里w个人是第一层要求，还有一层是，就是和第一个人组队的，其实可以是最少的那些人, 可能也包括a[2]
15. 还有一种策略是，先满足前i个人，a[i] = i, 给他们分配第一本书
16. 然后找到下一个a[i+j] = j，给他们分配第二本书
17. 但减少一本书的时候，前面的分组的边界向后移动，就有更多的人满意。所以这个策略应该是正确的
18. 按照k增加的方式进行处理
19. ans[1] = 所有人都满意
20. ans[2] = 这时候有两本书需要分配，第一个分组要全部满意，第二个分组里面有部分人满意
21. 还不如就让第一个分组都满意，第二个分组都不满意，那么就是 a[i] > n - i + 1（后面的人也不可能满意）
22. 然后考虑3本书，继续往前滑动？好像就不大对了。因为这个时候，假设有free个人，那么这些人，肯定不满意的情况下，完全可以去前面充数
23. 
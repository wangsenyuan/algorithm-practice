# K. King Kog's Reception

[Problem link](https://codeforces.com/problemset/problem/1089/K)

**Contest:** [Codeforces Round #526 (by Moscow Team Olympiad)](https://codeforces.com/contest/1089)

time limit: 2 seconds

memory limit: 512 megabytes

## Problem Statement

King Kog built a reception queue for his knights. Each knight chooses in advance
the time when he will come and how long the visit will take. Knights are served
in increasing order of recorded entrance time, but each knight must wait until
all earlier visits finish.

Princess Keabeanie wants to see her father without interrupting the knights, so
she joins the same queue. Knights frequently join or cancel visits. Help her
answer how long she will wait if she enters at given moments, given the current
records at the reception.

## Input

The first line contains a single integer `q` (`1 <= q <= 3 * 10^5`) — the number
of events.

Each of the next `q` lines is one of three event types:

- `+ t d` (`1 <= t, d <= 10^6`) — a new knight joins; he arrives at time `t` and
  the visit lasts `d`.
- `- i` (`1 <= i <= q`) — cancel the visit from the join event numbered `i`
  (1-based among all events). That join has not been cancelled yet.
- `? t` (`1 <= t <= 10^6`) — query: if Keabeanie comes at time `t`, how long
  will she wait until she sees her father?

It is guaranteed that after each event there are no two knights with the same
entrance time in the queue.

Keabeanie may arrive at the same time as some knight; she is polite and waits
for that knight to pass first.

## Output

For each query, print one line — the waiting time.

## Sample Input 1

```text
19
? 3
+ 2 2
? 3
? 4
+ 5 2
? 5
? 6
+ 1 2
? 2
? 3
? 4
? 5
? 6
? 7
? 9
- 8
? 2
? 3
? 6
```

## Sample Output 1

```text
0
1
0
2
1
3
2
1
2
1
0
0
2
1
1
```

## ideas
1. 如果knight到来的时候t时刻, 没有人排队, 那么他就占据了d的区间(t+d的时刻,又free了)
2. 如果这个时候,有人排队, 那么就需要知道, 最后结束的时间, 假设是tw
3. 那么他结束的时间就是tw+d; 
4. 问题出在取消. 因为取消的时候, 那些在他后面等待的人的时间, 不是减少d
5. 有可能会减少的更少, 所以情况就变得很不确定了~
6. 把i对应的作为一个区间, tl...tr (tr - tl + 1 >= d)
7. tl始终是t的加入时刻, tr = 他的结束时间
8. 当前方删除掉一个区间w时, 后面的都需要减少, 但是不能小于tl
9. 这样的区间更新, 似乎是可行的(lazy)
10. 当取消的时候, 后面的更新 -d
11. 计算t时刻的等待时间, 找到最大的tl <= d, 如果 tr < t, 那么不需要等待
12. 否则, 就是tl开始的一段连续的区间,直到找到一个空的时间点
13. 这一步似乎也有点麻烦

## Solution Summary

Deleting a knight is the difficult part of an interval-update approach: removing
duration `d` does not necessarily move every later finish time by `d`, because
an idle gap may absorb some or all of the removed time.  Instead of storing
actual occupied intervals, store the **effect of a time-ordered block** of
knights on the current queue finish time.

### A block as a function

For one knight with recorded arrival time `t` and duration `d`, if the queue
was previously free at time `x`, its new finish time is:

```text
max(x, t) + d = max(x + d, t + d).
```

Represent every block by a pair `(sum, fixed)` meaning:

```text
F(x) = max(x + sum, fixed).
```

For one knight this pair is `(d, t+d)`.  The empty block is the identity
`(0, -infinity)`.

### Merging two consecutive blocks

Suppose block `A` is earlier than block `B`:

```text
FA(x) = max(x + sumA, fixedA)
FB(x) = max(x + sumB, fixedB)
```

The merged function is `FB(FA(x))`:

```text
max(x + sumA + sumB, fixedA + sumB, fixedB).
```

It has the same form, so the segment-tree merge is:

```go
sum   = sumA + sumB
fixed = max(fixedA + sumB, fixedB)
```

This merge is associative, and its order matters: the left child represents
earlier arrival times and must be merged before the right child.

### Segment tree operations

The constraints allow a direct segment tree indexed by arrival time
`0..10^6`.

- `+ t d`: set leaf `t` to `(d, t+d)`.
- `- i`: look up the original join event `i` and reset its leaf to the
  identity `(0, -infinity)`.
- `? t`: merge the prefix `[0, t]`.  It includes a knight arriving exactly at
  `t`, as required by the statement.  If the aggregate is `(sum, fixed)`, then
  `fixed` is the final finish time, so the answer is:

```text
max(0, fixed - t).
```

For example, knights `(1,2)`, `(2,2)`, `(5,2)` finish at times `3`, `5`, and
`7`.  At `t=2`, the prefix finish is `5`, so the princess waits `3`.

Each update and query takes `O(log 10^6)` time; memory usage is `O(10^6)`.

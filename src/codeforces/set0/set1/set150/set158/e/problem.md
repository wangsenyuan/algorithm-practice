# Problem

- **Given:** `n` incoming calls (start times `tᵢ`, durations `dᵢ`) and a limit `k` of calls that can be ignored.
- **Call handling rules:**
  - If idle when a call arrives, he must answer it immediately.
  - If busy when a call arrives, he may:
    - ignore it (if he still has ignores left), or
    - queue it; after the current call ends, he always takes the **earliest** queued call.
- **Goal:** choose which calls to ignore to **maximize** the length of a single continuous idle interval within `[1, 86400]`.
  
---

## Input

- The first line contains two integers `n`, `k`  
  (`0 ≤ k ≤ n ≤ 4000`) — the number of calls and the maximum number of calls he may ignore.

- Each of the next `n` lines contains two integers `tᵢ`, `dᵢ`  
  (`1 ≤ tᵢ, dᵢ ≤ 86400`) — the scheduled start time and duration of the `i`-th call.

  - All `tᵢ` are distinct.
  - Calls are given in **strictly increasing** order of `tᵢ`.

- Scheduled call intervals `[tᵢ, tᵢ + dᵢ - 1]` may intersect arbitrarily.

---

## Output

Print a single integer from `0` to `86400` — the **maximum possible** number of seconds Mr. Jackson can sleep during the current day.

---

## Examples

### Example 1

Input:

```text
3 2
30000 15000
40000 15000
50000 15000
```

Output:

```text
49999
```

### Example 2

Input:

```text
5 1
1 20000
10000 10000
20000 20000
25000 10000
80000 60000
```

Output:

```text
39999
```

---

## Note

- In the first sample, the best strategy is to **ignore the first two calls**.

- In the second sample, it is best to **ignore the third call**. Then Mr. Jackson talks:

  - first call: from 1st to 20000th second,
  - second call: from 20001st to 30000th second,
  - fourth call: from 30001st to 40000th second (the third call is ignored),
  - fifth call: from 80000th to 139999th second.

  Thus, the longest continuous free time is from 40001st to 79999th second.

---


### ideas
1. 这个题目这么难的～
2. 如果一个人已经在talk了，那么必须是这个talk结束开始休息，那么这个人，开始后的k个call被忽略掉
3. 那么下一个时间 = t[i+k]
4. dp[i][j] 表示对于到目前为止， ignore了j个电话后，最早可以休息的时间
5. dp[i][j] = min(max(dp[i-1][j], t[j]) + d[j], dp[i-1][j-1])
6. ans  = max(t[i] - dp[i-1][?] - 1)
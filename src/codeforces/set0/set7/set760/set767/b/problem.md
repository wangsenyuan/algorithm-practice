# Passport Office Queue Problem

## Problem Description

Finally! Vasya has come of age and that means he can finally get a passport! To do it, he needs to visit the passport office, but it's not that simple. There's only one receptionist at the passport office and people can queue up long before it actually opens. Vasya wants to visit the passport office tomorrow.

He knows that:
- The receptionist starts working after `ts` minutes have passed after midnight
- The receptionist closes after `tf` minutes have passed after midnight (so that `(tf - 1)` is the last minute when the receptionist is still working)
- The receptionist spends exactly `t` minutes on each person in the queue
- If the receptionist would stop working within `t` minutes, he stops serving visitors (other than the one he already serves)

Vasya also knows that exactly `n` visitors would come tomorrow. For each visitor, Vasya knows the point of time when he would come to the passport office. Each visitor queues up and doesn't leave until he was served. If the receptionist is free when a visitor comes (in particular, if the previous visitor was just served and the queue is empty), the receptionist begins to serve the newcomer immediately.

### Important Rules

- For each visitor, the point of time when he would come to the passport office is positive
- Vasya can come to the office at time zero (that is, at midnight) if he needs so, but he can come to the office only at integer points of time
- If Vasya arrives at the passport office at the same time with several other visitors, he yields to them and stands in the queue after the last of them
- Vasya wants to come at such point of time that he will be served by the receptionist, and he would spend the minimum possible time in the queue

## Input

- **Line 1**: Three integers: `ts`, `tf`, and `t`
  - `ts`: The point of time when the receptionist begins to work
  - `tf`: The point of time when the receptionist stops working  
  - `t`: The time the receptionist spends on each visitor
- **Line 2**: One integer `n` — the amount of visitors (0 ≤ n ≤ 100,000)
- **Line 3**: `n` positive integers in non-decreasing order — the points of time when the visitors arrive to the passport office

### Constraints

- All times are set in minutes and do not exceed 10^12
- It is guaranteed that `ts < tf`
- It is also guaranteed that Vasya can arrive at the passport office at such a point of time that he would be served by the receptionist

## Output

Print a single non-negative integer — the point of time when Vasya should arrive at the passport office. If Vasya arrives at the passport office at the same time with several other visitors, he yields to them and queues up last. If there are many answers, you can print any of them.

## Examples

### Example 1

**Input:**
```
10 15 2
2
10 13
```

**Output:**
```
12
```

**Explanation:**
The first visitor comes exactly at the point of time when the receptionist begins to work, and he is served for two minutes. At 12 minutes after midnight, the receptionist stops serving the first visitor, and if Vasya arrives at this moment, he will be served immediately, because the next visitor would only come at 13 minutes after midnight.

### Example 2

**Input:**
```
8 17 3
4
3 4 5 8
```

**Output:**
```
2
```

**Explanation:**
In this example, Vasya has to come before anyone else to be served.

### ideas
1. 假设Vasya在某个时刻x到达去排队
2. 如果他前面有m个人，那么他的等待时间 = m * t + 当前在受理人的剩余时间
3. 这里一共有m个人，那么它们进入，离开的时间共有 m * 2个
4. 可以用优先队列处理
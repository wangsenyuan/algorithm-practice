# Exam Period Problem

## Problem Description

Vasiliy has an exam period which will continue for n days. He has to pass exams on m subjects. Subjects are numbered from 1 to m.

About every day we know exam for which one of m subjects can be passed on that day. Perhaps, some day you can't pass any exam. It is not allowed to pass more than one exam on any day.

On each day Vasiliy can either:
- Pass the exam of that day (it takes the whole day)
- Prepare all day for some exam
- Have a rest

About each subject Vasiliy knows a number ai — the number of days he should prepare to pass the exam number i. Vasiliy can switch subjects while preparing for exams, it is not necessary to prepare continuously during ai days for the exam number i. He can mix the order of preparation for exams in any way.

Your task is to determine the minimum number of days in which Vasiliy can pass all exams, or determine that it is impossible. Each exam should be passed exactly one time.

## Input

The first line contains two integers n and m (1 ≤ n, m ≤ 10^5) — the number of days in the exam period and the number of subjects.

The second line contains n integers d1, d2, ..., dn (0 ≤ di ≤ m), where di is the number of subject, the exam of which can be passed on the day number i. If di equals 0, it is not allowed to pass any exams on the day number i.

The third line contains m positive integers a1, a2, ..., am (1 ≤ ai ≤ 10^5), where ai is the number of days that are needed to prepare before passing the exam on the subject i.

## Output

Print one integer — the minimum number of days in which Vasiliy can pass all exams. If it is impossible, print -1.

## Examples

### Example 1

**Input:**
```
7 2
0 1 0 2 1 0 2
2 1
```

**Output:**
```
5
```

**Explanation:**
Vasiliy can behave as follows:
- Days 1-2: prepare for exam number 1 and pass it on day 5
- Day 3: prepare for exam number 2 and pass it on day 4

### Example 2

**Input:**
```
10 3
0 0 1 2 3 0 2 0 1 2
1 1 4
```

**Output:**
```
9
```

**Explanation:**
- Days 1-4: prepare for exam number 3 and pass it on day 5
- Day 6: prepare for exam number 2 and pass it on day 7
- Day 8: prepare for exam number 1 and pass it on day 9

### Example 3

**Input:**
```
5 1
1 1 1 1 1
5
```

**Output:**
```
-1
```

**Explanation:**
Vasiliy can't pass the only exam because he hasn't enough time to prepare for it.

## ideas
1. d_i 表示在第i天可以通过的考试（每个d_i)可以考虑在最后一次再参加（这样子有充足的时间准备）
2. 但是这样子不一定能满足最短时间（所以还需要二分）
3. 
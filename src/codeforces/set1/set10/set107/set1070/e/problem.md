Polycarp has a lot of work to do. Recently he has learned a new time management rule: “if a task takes five minutes or less, do it immediately”. Polycarp likes the new rule, however he is not sure that five minutes is the optimal value. He supposes that this value $d$ should be chosen based on the existing task list.

Polycarp has a list of $n$ tasks to complete. The $i$-th task has difficulty $p_i$, i.e. it requires exactly $p_i$ minutes to be done. Polycarp reads the tasks one by one from the first to the $n$-th. If a task difficulty is $d$ or less, Polycarp starts the work on the task immediately. If a task difficulty is strictly greater than $d$, he will not do the task at all. It is not allowed to rearrange tasks in the list. Polycarp doesn't spend any time for reading a task or skipping it.

Polycarp has $t$ minutes in total to complete the maximum number of tasks. But he does not want to work all the time. He decides to make a break after each group of $m$ consecutive tasks he was working on. The break should take the same amount of time as it was spent in total on completion of these $m$ tasks.

For example, if $n = 7$, $p = [3, 1, 4, 1, 5, 9, 2]$, $d = 3$ and $m = 2$ Polycarp works by the following schedule:

- Polycarp reads the first task, its difficulty is not greater than $d$ ($p_1 = 3 \le d = 3$) and works for 3 minutes (i.e. the minutes 1, 2, 3).
- Polycarp reads the second task, its difficulty is not greater than $d$ ($p_2 = 1 \le d = 3$) and works for 1 minute (i.e. the minute 4).
- Polycarp notices that he has finished $m = 2$ tasks and takes a break for $3 + 1 = 4$ minutes (i.e. on the minutes 5, 6, 7, 8).
- Polycarp reads the third task, its difficulty is greater than $d$ ($p_3 = 4 > d = 3$) and skips it without spending any time.
- Polycarp reads the fourth task, its difficulty is not greater than $d$ ($p_4 = 1 \le d = 3$) and works for 1 minute (i.e. the minute 9).
- Polycarp reads the tasks 5 and 6, skips both of them ($p_5 > d$ and $p_6 > d$).
- Polycarp reads the 7-th task, its difficulty is not greater than $d$ ($p_7 = 2 \le d = 3$) and works for 2 minutes (i.e. the minutes 10, 11).
- Polycarp notices that he has finished $m = 2$ tasks and takes a break for $1 + 2 = 3$ minutes (i.e. on the minutes 12, 13, 14).

Polycarp stops exactly after $t$ minutes. If Polycarp started a task but has not finished it by that time, the task is not considered as completed. It is allowed to complete less than $m$ tasks in the last group. Also Polycarp considers acceptable to have shorter break than needed after the last group of tasks or even not to have this break at all — his working day is over and he will have enough time to rest anyway.

Please help Polycarp to find such value $d$, which would allow him to complete maximum possible number of tasks in $t$ minutes.

## Input

The first line of the input contains single integer $c$ ($1 \le c \le 5 \cdot 10^4$) — number of test cases. Then description of $c$ test cases follows. Solve test cases separately, test cases are completely independent and do not affect each other.

Each test case is described by two lines. The first of these lines contains three space-separated integers $n$, $m$ and $t$ ($1 \le n \le 2 \cdot 10^5$, $1 \le m \le 2 \cdot 10^5$, $1 \le t \le 4 \cdot 10^{10}$) — the number of tasks in Polycarp's list, the number of tasks he can do without a break and the total amount of time Polycarp can work on tasks. The second line of the test case contains $n$ space separated integers $p_1, p_2, \dots, p_n$ ($1 \le p_i \le 2 \cdot 10^5$) — difficulties of the tasks.

The sum of values $n$ for all test cases in the input does not exceed $2 \cdot 10^5$.

## Output

Print $c$ lines, each line should contain answer for the corresponding test case — the maximum possible number of tasks Polycarp can complete and the integer value $d$ ($1 \le d \le t$) Polycarp should use in time management rule, separated by space. If there are several possible values $d$ for a test case, output any of them.

## Examples

### Input

```
4
5 2 16
5 6 1 4 7
5 3 30
5 6 1 4 7
6 4 15
12 5 15 7 20 17
1 1 50
100
```

### Output

```
3 5
4 7
2 10
0 25
```

### Input

```
3
11 1 29
6 4 3 7 5 3 4 7 3 5 3
7 1 5
1 1 1 1 1 1 1
5 2 18
2 3 3 7 5
```

### Output

```
4 3
3 1
4 5
```

## Note

In the first test case of the first example $n = 5$, $m = 2$ and $t = 16$. The sequence of difficulties is $[5, 6, 1, 4, 7]$. If Polycarp chooses $d = 5$ then he will complete 3 tasks. Polycarp will work by the following schedule:

- Polycarp reads the first task, its difficulty is not greater than $d$ ($p_1 = 5 \le d = 5$) and works for 5 minutes (i.e. the minutes 1, 2, \dots, 5).
- Polycarp reads the second task, its difficulty is greater than $d$ ($p_2 = 6 > d = 5$) and skips it without spending any time.
- Polycarp reads the third task, its difficulty is not greater than $d$ ($p_3 = 1 \le d = 5$) and works for 1 minute (i.e. the minute 6).
- Polycarp notices that he has finished $m = 2$ tasks and takes a break for $5 + 1 = 6$ minutes (i.e. on the minutes 7, 8, \dots, 12).
- Polycarp reads the fourth task, its difficulty is not greater than $d$ ($p_4 = 4 \le d = 5$) and works for 4 minutes (i.e. the minutes 13, 14, 15, 16).
- Polycarp stops work because of $t = 16$.

In total in the first test case Polycarp will complete 3 tasks for $d = 5$. He can't choose other value for $d$ to increase the number of completed tasks.

Polycarp has a lot of work to do. Recently he has learned a new time management rule: "if a task takes five minutes or less, do it immediately". Polycarp likes the new rule, however he is not sure that five minutes is the optimal value. He supposes that this value 𝑑
 should be chosen based on existing task list.

Polycarp has a list of 𝑛
 tasks to complete. The 𝑖
-th task has difficulty 𝑝𝑖
, i.e. it requires exactly 𝑝𝑖
 minutes to be done. Polycarp reads the tasks one by one from the first to the 𝑛
-th. If a task difficulty is 𝑑
 or less, Polycarp starts the work on the task immediately. If a task difficulty is strictly greater than 𝑑
, he will not do the task at all. It is not allowed to rearrange tasks in the list. Polycarp doesn't spend any time for reading a task or skipping it.

Polycarp has 𝑡
 minutes in total to complete maximum number of tasks. But he does not want to work all the time. He decides to make a break after each group of 𝑚
 consecutive tasks he was working on. The break should take the same amount of time as it was spent in total on completion of these 𝑚
 tasks.

For example, if 𝑛=7
, 𝑝=[3,1,4,1,5,9,2]
, 𝑑=3
 and 𝑚=2
 Polycarp works by the following schedule:

Polycarp reads the first task, its difficulty is not greater than 𝑑
 (𝑝1=3≤𝑑=3
) and works for 3
 minutes (i.e. the minutes 1
, 2
, 3
);
Polycarp reads the second task, its difficulty is not greater than 𝑑
 (𝑝2=1≤𝑑=3
) and works for 1
 minute (i.e. the minute 4
);
Polycarp notices that he has finished 𝑚=2
 tasks and takes a break for 3+1=4
 minutes (i.e. on the minutes 5,6,7,8
);
Polycarp reads the third task, its difficulty is greater than 𝑑
 (𝑝3=4>𝑑=3
) and skips it without spending any time;
Polycarp reads the fourth task, its difficulty is not greater than 𝑑
 (𝑝4=1≤𝑑=3
) and works for 1
 minute (i.e. the minute 9
);
Polycarp reads the tasks 5
 and 6
, skips both of them (𝑝5>𝑑
 and 𝑝6>𝑑
);
Polycarp reads the 7
-th task, its difficulty is not greater than 𝑑
 (𝑝7=2≤𝑑=3
) and works for 2
 minutes (i.e. the minutes 10
, 11
);
Polycarp notices that he has finished 𝑚=2
 tasks and takes a break for 1+2=3
 minutes (i.e. on the minutes 12,13,14
).
Polycarp stops exactly after 𝑡
 minutes. If Polycarp started a task but has not finished it by that time, the task is not considered as completed. It is allowed to complete less than 𝑚
 tasks in the last group. Also Polycarp considers acceptable to have shorter break than needed after the last group of tasks or even not to have this break at all — his working day is over and he will have enough time to rest anyway.

Please help Polycarp to find such value 𝑑
, which would allow him to complete maximum possible number of tasks in 𝑡
 minutes.

Input
The first line of the input contains single integer 𝑐
 (1≤𝑐≤5⋅104
) — number of test cases. Then description of 𝑐
 test cases follows. Solve test cases separately, test cases are completely independent and do not affect each other.

Each test case is described by two lines. The first of these lines contains three space-separated integers 𝑛
, 𝑚
 and 𝑡
 (1≤𝑛≤2⋅105,1≤𝑚≤2⋅105,1≤𝑡≤4⋅1010
) — the number of tasks in Polycarp's list, the number of tasks he can do without a break and the total amount of time Polycarp can work on tasks. The second line of the test case contains 𝑛
 space separated integers 𝑝1,𝑝2,…,𝑝𝑛
 (1≤𝑝𝑖≤2⋅105
) — difficulties of the tasks.

The sum of values 𝑛
 for all test cases in the input does not exceed 2⋅105
.

Output
Print 𝑐
 lines, each line should contain answer for the corresponding test case — the maximum possible number of tasks Polycarp can complete and the integer value 𝑑
 (1≤𝑑≤𝑡
) Polycarp should use in time management rule, separated by space. If there are several possible values 𝑑
 for a test case, output any of them.

Examples
InputCopy
4
5 2 16
5 6 1 4 7
5 3 30
5 6 1 4 7
6 4 15
12 5 15 7 20 17
1 1 50
100
OutputCopy
3 5
4 7
2 10
0 25
InputCopy
3
11 1 29
6 4 3 7 5 3 4 7 3 5 3
7 1 5
1 1 1 1 1 1 1
5 2 18
2 3 3 7 5
OutputCopy
4 3
3 1
4 5
Note
In the first test case of the first example 𝑛=5
, 𝑚=2
 and 𝑡=16
. The sequence of difficulties is [5,6,1,4,7]
. If Polycarp chooses 𝑑=5
 then he will complete 3
 tasks. Polycarp will work by the following schedule:

Polycarp reads the first task, its difficulty is not greater than 𝑑
 (𝑝1=5≤𝑑=5
) and works for 5
 minutes (i.e. the minutes 1,2,…,5
);
Polycarp reads the second task, its difficulty is greater than 𝑑
 (𝑝2=6>𝑑=5
) and skips it without spending any time;
Polycarp reads the third task, its difficulty is not greater than 𝑑
 (𝑝3=1≤𝑑=5
) and works for 1
 minute (i.e. the minute 6
);
Polycarp notices that he has finished 𝑚=2
 tasks and takes a break for 5+1=6
 minutes (i.e. on the minutes 7,8,…,12
);
Polycarp reads the fourth task, its difficulty is not greater than 𝑑
 (𝑝4=4≤𝑑=5
) and works for 4
 minutes (i.e. the minutes 13,14,15,16
);
Polycarp stops work because of 𝑡=16
.
In total in the first test case Polycarp will complete 3
 tasks for 𝑑=5
. He can't choose other value for 𝑑
 to increase the number of completed tasks.


 ### ideas
 1. 显然不是d越大越好，因为这样子需要休息的时间似乎也会很长
 2. 也不是越小越好，这样子能够采纳的task会很少
 3. d的值，就是去重后的c的值；
 4. 这里考虑当d变化的时候，如何更新结果？
 5. 假设d = d0的时候，计算出来了task的数量（还有分布情况）
 6. 然后当d增加到d1的时候，原来一部分不能完成的任务就被加入了进来（但是需要休息的时间就更多了）
 7. 假设一共有w个任务被完成，前w1 = w / m * w个任务的总时间是s1, 那么需要的总时间 = 2 * s1
 8. 如果后面还有任务 2 * s1 < t, 然后加上额外的完成的任务
 9. 2 * s1 + suf[w1+1:w] <= t
 10. 如果w == w1, 那么 2 * s1 - 最后一段时间(休息时间) <= t
 11. 2分会处理一些
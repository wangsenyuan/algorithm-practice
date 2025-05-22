# Network Printer Task Scheduling

## Problem Description
Let's consider a network printer that functions like that. It starts working at time 0. In each second it can print one page of a text. At some moments of time the printer receives printing tasks. We know that a printer received n tasks. Let's number the tasks by consecutive integers from 1 to n. Then the task number i is characterised by three integers:
- `ti` is the time when the task came
- `si` is the task's volume (in pages)
- `pi` is the task's priority

The priorities of all tasks are distinct.

## Printer Operation
When the printer receives a task, the task goes to the queue and remains there until all pages from this task are printed. The printer chooses a page to print each time when it either:
1. Stops printing some page, or
2. Is free and receives a new task

Among all tasks that are in the queue at this moment, the printer chooses the task with the highest priority and next second prints an unprinted page from this task. You can assume that a task goes to the queue immediately, that's why if a task has just arrived by time t, the printer can already choose it for printing.

## Challenge
You are given full information about all tasks except for one: you don't know this task's priority. However, we know the time when the last page from this task was finished printing. Given this information, find the unknown priority value and determine the moments of time when the printer finished printing each task.

## Input
- First line: integer n (1 ≤ n ≤ 50000)
- Next n lines: each contains three integers `ti`, `si` and `pi`, separated by single spaces
  - 0 ≤ ti ≤ 10^9
  - 1 ≤ si, pi ≤ 10^9
  - Exactly one task (number x) has -1 written instead of the priority
  - All priorities are different
  - Tasks are written in arbitrary order
- Last line: integer T — the time when the printer finished printing the last page of task x (1 ≤ T ≤ 10^15)

Note: Numbers ti are not necessarily distinct.

## Output
- First line: integer px — the priority of task number x (1 ≤ px ≤ 10^9, all priorities must be distinct)
- Next line: n integers, where the i-th integer represents the moment of time when the last page of task i finished printing

It is guaranteed that at least one solution exists. If there are multiple solutions, print any of them.

## ideas
- 假设知道了所有任务的优先级，如何知道这些任务的结束时间呢？
- 似乎是可以模拟的。随意s[i]很大，但是队列变化的次数是n
- 当一个任务被加入的时候，可以看下一个任务进入的时刻
- 假设这个中间的时间差是diff，那么就可以在diff时间差内，进行模拟
- 找出最高优先级的任务，如果需要执行它的时间少于diff，继续处理下一个
- 直到出现diff时间不够，将它的剩余时间记录下来。处理下个时刻的任务进入事件
- 然后，我们来看看能否计算出x的优先级
- 有一个情况下，x的优先级越大，那么完成它的时间会越早。这个应该是显然的。
- 那么就可以二分它的优先级。找到最小的优先级，在时刻t内完成
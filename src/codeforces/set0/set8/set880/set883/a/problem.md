# Problem A

There is an automatic door at the entrance of a factory. The door works in the following way:

- When one or several people come to the door and it is closed, the door immediately opens automatically and all people immediately come inside,
- When one or several people come to the door and it is open, all people immediately come inside,
- Opened door immediately closes in d seconds after its opening,
- If the door is closing and one or several people are coming to the door at the same moment, then all of them will have enough time to enter and only after that the door will close.

For example, if d = 3 and four people are coming at four different moments of time t₁ = 4, t₂ = 7, t₃ = 9 and t₄ = 13 then the door will open three times: at moments 4, 9 and 13. It will close at moments 7 and 12.

It is known that n employees will enter at moments a, 2·a, 3·a, …, n·a (the value a is a positive integer). Also m clients will enter at moments t₁, t₂, …, tₘ.

Write a program to find the number of times the automatic door will open. Assume that the door is initially closed.

## Input

The first line contains four integers n, m, a and d (1 ≤ n, a ≤ 10⁹, 1 ≤ m ≤ 10⁵, 1 ≤ d ≤ 10¹⁸) — the number of employees, the number of clients, the moment of time when the first employee will come, and the period of time in which the door closes.

The second line contains integer sequence t₁, t₂, …, tₘ (1 ≤ tᵢ ≤ 10¹⁸) — moments of time when clients will come. The values tᵢ are given in non-decreasing order.

## Output

Print the number of times the door will open.

## Examples

### Example 1

**Input:**
```
1 1 3 4
7
```

**Output:**
```
1
```

### Example 2

**Input:**
```
4 3 4 2
7 9 11
```

**Output:**
```
4
```

## Note

In the first example the only employee will come at moment 3. At this moment the door will open and will stay open until the moment 7. At the same moment of time the client will come, so at first he will enter and only after that the door will close. Thus the door will open one time.


### ideas
1. 计算题，不好算哪
2. door不会影响人员进入
3. 这里主要的问题是n很大。那么不考虑m的情况下，假设在某个时刻w开门（这个时刻，只可能是人员到场的时间）
4. 然后经过d时刻关闭，
5. 感觉还是要模拟，但是n次进入是有规律的，所以要利用起来
6. 假设目前门的状态是开启的，且它开始的时刻是w, 当前是t[?] 那么下次关闭的时刻是 w + d, 那么可以找到下一个t[i] > w + d; 但是，需要计算在w + d前，能够完成的员工进入的数量k
7. 这个数量也可以算出来
8. 这样子直到处理到最后一个t[m]; 如果此时n个员工已经完成了, done
9. 否则还是处理到 w + d，时刻接下来就是剩余员工如何处理
10. 如果 d <= a (也就是每次开启后，下次员工进入前)
11. d > a, 好像就是简单的除法了
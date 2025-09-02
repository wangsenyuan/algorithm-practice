# Problem Description

The employees of the F company have lots of ways to entertain themselves. Today they invited a famous magician who shows a trick with plastic cups and a marble.

The point is to trick the spectator's attention. Initially, the spectator stands in front of a line of n plastic cups. Then the magician places a small marble under one cup and shuffles the cups. Then the spectator should guess which cup hides the marble.

But the head coder of the F company isn't easy to trick. When he saw the performance, he noticed several important facts:

- Each cup contains a mark — a number from 1 to n; all marks on the cups are distinct
- The magician shuffles the cups in m operations, each operation looks like that: take a cup marked xi, sitting at position yi in the row of cups (the positions are numbered from left to right, starting from 1) and shift it to the very beginning of the cup row (on the first position)

When the head coder came home after work he wanted to re-do the trick. Unfortunately, he didn't remember the starting or the final position of the cups. He only remembered which operations the magician performed. Help the coder: given the operations in the order they were made find at least one initial permutation of the cups that can go through the described operations in the given order. Otherwise, state that such permutation doesn't exist.

## Input

The first line contains integers n and m (1 ≤ n, m ≤ 10^6). Each of the next m lines contains a couple of integers. The i-th line contains integers xi, yi (1 ≤ xi, yi ≤ n) — the description of the i-th operation of the magician. Note that the operations are given in the order in which the magician made them and the coder wants to make them in the same order.

## Output

If the described permutation doesn't exist (the programmer remembered wrong operations), print -1. Otherwise, print n distinct integers, each from 1 to n: the i-th number should represent the mark on the cup that initially is in the row in position i.

If there are multiple correct answers, you should print the lexicographically minimum one.

## Examples

### Example 1
**Input:**
```
2 1
2 1
```

**Output:**
```
2 1
```

### Example 2
**Input:**
```
3 2
1 2
1 1
```

**Output:**
```
2 1 3
```

### Example 3
**Input:**
```
3 3
1 3
2 3
1 3
```

**Output:**
```
-1
```


### idea1. 
1. x1, y1 （那么表明 x1在位置y1处）
2. 将x1放置到头以后，x1就处于位置1了，在[1..y1)中间的所有的位置+1
3. x2, y2, 如果y2 > y1, 那么x2, 就在位置y2处，然后和x1类似
4. 如果 y2 <= y1, 那么 x2原来的位置，应该是 y2 - 1
5. 这个时候还是清晰的，但是x3, y3怎么处理？
6. 如果 y3 <= min(y1, y2), 那么x3所处的位置就是y3 - 2
7. 如果 y3 > max(y1, y2), 就不需要变化，如果在中间，那么就是 y3 - 1
8. 也就是看yi右边有多少个数字被调整过了？
9. 这是第一个观察。应该是在部分情况下成立的
10. 如果 xi = x1呢？也就是x1又选中了，那么可以用来验证，它的位置，是否真的是yi
11. 要维护一个前端的队列（这个队列是那些已经确定好位置的，交换到前面的cups）
12. 如果一个(xi, yi)它在前端队列中，且是正确的，那么更新前端队列就可以了
13. 如果在队列中，但是不正确，比如位置不对，或者数字不对，那么就是没有答案
14. 插入一个新的元素时，前端队列的位置增加1
15. 这个前端队列的维护，貌似有点麻烦。
16. x1, x2, x3...这些数字它们不是连续的（虽然它们在前端的位置是连续的）
17. 如果记录pos(xi) ，可以很方便的知道，但是更新就很麻烦了（需要一个个的去更新）
18. 如果能很方便的更新，要查询就很麻烦了
19. 貌似要用tree来表示，tree的下标（并不表示位置）每个下标l左边，已经设置的值，表示真实的下表
20. tree的val表示x[i], 如果将x[i]移动到第一位，不去整体移动位置
21. 所以要有足够的空间往前面放；
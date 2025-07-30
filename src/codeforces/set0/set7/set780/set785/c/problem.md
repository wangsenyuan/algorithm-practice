# Problem C: Anton and Fairy Tale

## Problem Description

Anton likes to listen to fairy tales, especially when Danik, Anton's best friend, tells them. Right now Danik tells Anton a fairy tale:

> "Once upon a time, there lived an emperor. He was very rich and had much grain. One day he ordered to build a huge barn to put there all his grain. Best builders were building that barn for three days and three nights. But they overlooked and there remained a little hole in the barn, from which every day sparrows came through. Here flew a sparrow, took a grain and flew away..."

## Problem Statement

More formally, the following takes place in the fairy tale. At the beginning of the first day the barn with the capacity of **n** grains was full. Then, every day (starting with the first day) the following happens:

1. **m** grains are brought to the barn. If **m** grains don't fit in the barn, the barn becomes full and the grains that don't fit are brought back (in this problem we can assume that the grains that don't fit in the barn are not taken into account).

2. Sparrows come and eat grain. On the **i-th** day **i** sparrows come, that is:
   - On the first day: 1 sparrow comes
   - On the second day: 2 sparrows come
   - And so on...
   
   Every sparrow eats one grain. If the barn is empty, a sparrow eats nothing.

Anton is tired of listening how Danik describes every sparrow that eats grain from the barn. Anton doesn't know when the fairy tale ends, so he asked you to determine, by the end of which day the barn will become empty for the first time. Help Anton and write a program that will determine the number of that day!

## Input

The only line of the input contains two integers **n** and **m** (1 ≤ n, m ≤ 10^18) — the capacity of the barn and the number of grains that are brought every day.

## Output

Output one integer — the number of the day when the barn will become empty for the first time. Days are numbered starting with one.

## Examples

### Example 1
**Input:**
```
5 2
```

**Output:**
```
4
```

### Example 2
**Input:**
```
8 1
```

**Output:**
```
5
```

## Note

In the first sample the capacity of the barn is five grains and two grains are brought every day. The following happens:

- **Day 1:**
  - At the beginning: Grain is brought to the barn. It's full, so nothing happens.
  - At the end: One sparrow comes and eats one grain, so 5 - 1 = 4 grains remain.

- **Day 2:**
  - At the beginning: Two grains are brought. The barn becomes full and one grain doesn't fit.
  - At the end: Two sparrows come. 5 - 2 = 3 grains remain.

- **Day 3:**
  - At the beginning: Two grains are brought. The barn becomes full again.
  - At the end: Three sparrows come and eat grain. 5 - 3 = 2 grains remain.

- **Day 4:**
  - At the beginning: Grain is brought again. 2 + 2 = 4 grains remain.
  - At the end: Four sparrows come and eat grain. 4 - 4 = 0 grains remain. The barn is empty.

So the answer is **4**, because by the end of the fourth day the barn becomes empty.

### ideas
1. 在第n天后，肯定能吃完
2. 在开始的时候，因为鸟比较少，1，2,...i只鸟，且明天早上都有补充，
3. 那么在每天早餐都可以补充完整，也就是说 n = m + x
4. x = 前一天剩余的量，也就是说有 n - i = x => i = m 也就是说在有m只鸟的时候，正好第二天呢，可以补充完整
5. (m < n必须成立)
6. 然后开始 m+1只鸟， m + 2... 一直到n只吗？
7. 如果 m >= n => 答案只能是n
8. m < n, 在前m天，都可以填满，且在第m天，后剩余, n - m 克的粮食
9. 然后第m+1天早晨变成n, 有 m + 1只鸟，晚上剩余 n - m - 1 克
10. m+2天，剩余 n - 1, m+2只鸟， n - 1 - (m + 2) = n - m - 3
11. m+3天， 剩余 n - 3, m + 3只鸟 ，n - 3 - (m + 3) = n - m - 6
12. 在n - m - i * (i + 1) / 2 <= 0
13. 
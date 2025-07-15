# Problem C: Socks

## Problem Description

Arseniy is already grown-up and independent. His mother decided to leave him alone for m days and left on a vacation. She have prepared a lot of food, left some money and washed all Arseniy's clothes.

Ten minutes before her leave she realized that it would be also useful to prepare instruction of which particular clothes to wear on each of the days she will be absent. Arseniy's family is a bit weird so all the clothes is enumerated. For example, each of Arseniy's n socks is assigned a unique integer from 1 to n. Thus, the only thing his mother had to do was to write down two integers li and ri for each of the days — the indices of socks to wear on the day i (obviously, li stands for the left foot and ri for the right). Each sock is painted in one of k colors.

When mother already left Arseniy noticed that according to instruction he would wear the socks of different colors on some days. Of course, that is a terrible mistake cause by a rush. Arseniy is a smart boy, and, by some magical coincidence, he posses k jars with the paint — one for each of k colors.

Arseniy wants to repaint some of the socks in such a way, that for each of m days he can follow the mother's instructions and wear the socks of the same color. As he is going to be very busy these days he will have no time to change the colors of any socks so he has to finalize the colors now.

The new computer game Bota-3 was just realised and Arseniy can't wait to play it. What is the minimum number of socks that need their color to be changed in order to make it possible to follow mother's instructions and wear the socks of the same color during each of m days.

## Input

The first line of input contains three integers n, m and k (2 ≤ n ≤ 200 000, 0 ≤ m ≤ 200 000, 1 ≤ k ≤ 200 000) — the number of socks, the number of days and the number of available colors respectively.

The second line contain n integers c1, c2, ..., cn (1 ≤ ci ≤ k) — current colors of Arseniy's socks.

Each of the following m lines contains two integers li and ri (1 ≤ li, ri ≤ n, li ≠ ri) — indices of socks which Arseniy should wear during the i-th day.

## Output

Print one integer — the minimum number of socks that should have their colors changed in order to be able to obey the instructions and not make people laugh from watching the socks of different colors.

## Examples

### Example 1

**Input:**
```
3 2 3
1 2 3
1 2
2 3
```

**Output:**
```
2
```

**Explanation:** Arseniy can repaint the first and the third socks to the second color.

### Example 2

**Input:**
```
3 2 2
1 1 2
1 2
2 1
```

**Output:**
```
0
```

**Explanation:** In the second sample, there is no need to change any colors.

## ideas
1. 每天的选择，就是一条边，连接了u = li, v = ri
2. 这些边组成了一个component，在一个component中，颜色要一致，那么就是保留最多的颜色，改变其他的颜色
3. 
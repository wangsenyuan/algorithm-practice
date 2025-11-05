A flowerbed has many flowers and two fountains.

You can adjust the water pressure and set any values $r_1$ ($r_1 \geq 0$) and $r_2$ ($r_2 \geq 0$), giving the distances at which the water is spread from the first and second fountain respectively. You have to set such $r_1$ and $r_2$ that all the flowers are watered, that is, for each flower, the distance between the flower and the first fountain doesn't exceed $r_1$, or the distance to the second fountain doesn't exceed $r_2$. It's OK if some flowers are watered by both fountains.

You need to decrease the amount of water you need, that is set such $r_1$ and $r_2$ that all the flowers are watered and the $r_1^2 + r_2^2$ is minimum possible. Find this minimum value.

### Input

The first line of the input contains integers $n$, $x_1$, $y_1$, $x_2$, $y_2$ ($1 \leq n \leq 2000$, $-10^7 \leq x_1, y_1, x_2, y_2 \leq 10^7$) — the number of flowers, the coordinates of the first and the second fountain.

Next follow $n$ lines. The $i$-th of these lines contains integers $x_i$ and $y_i$ ($-10^7 \leq x_i, y_i \leq 10^7$) — the coordinates of the $i$-th flower.

It is guaranteed that all $n + 2$ points in the input are distinct.

### Output

Print the minimum possible value $r_1^2 + r_2^2$. Note, that in this problem optimal answer is always integer.

### Examples

**Input**
```
2 -1 0 5 3
0 2
5 2
```

**Output**
```
6
```

**Input**
```
4 0 0 5 0
9 4
8 3
-1 0
1 4
```

**Output**
```
33
```

### Note

The first sample is ($r_1^2 = 5$, $r_2^2 = 1$).

The second sample is ($r_1^2 = 1$, $r_2^2 = 32$).


### ideas
1. 找两个圈，r1, r2把所有的花给包括进去，使的 $r1 * r1 + r2 * r2$ 最小
2. 假设宽度 > 高度， 那么在最优答案中， 两端的点，应该是分属两个圈？
3. 两个圆心是给定的，那么就可以二分了？
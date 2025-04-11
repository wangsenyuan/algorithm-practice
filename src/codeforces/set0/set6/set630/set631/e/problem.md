Blake is the boss of Kris, however, this doesn't spoil their friendship. They often gather at the bar to talk about intriguing problems about maximising some values. This time the problem is really special.

You are given an array a of length n. The characteristic of this array is the value  — the sum of the products of the values ai by i. One may perform the following operation exactly once: pick some element of the array and move to any position. In particular, it's allowed to move the element to the beginning or to the end of the array. Also, it's allowed to put it back to the initial position. The goal is to get the array with the maximum possible value of characteristic.

### ideas
1. y[n] = n * x - p[n]
2. 如果p[n]非常小，那么前面所有的数字，如果要交换，那么只应该交换到位置n
3. 所以，如果存在一组线 y[i], y[i+1], ... y[n]
4. 那么 p[i] < p[i+1] ... < p[n]
5. -p[i] > -p[i+1] > ... > -p[n]
6. 这样子，就可以在加入线的时候，计算出最近的两条线的交叉点
7. 然后根据a[i](也就是x) 找出它应该放在交换到哪段线（j)
8. 
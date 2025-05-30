Jack is working on his jumping skills recently. Currently, he's located at point zero of the number line. He would like to get to the point x.

In order to train, he has decided that he'll first jump by only one unit, and each subsequent jump will be exactly one longer than the previous one. He can go either left or right with each jump.

He wonders how many jumps he needs to reach x.

### ideas
1. always process positive value
2. 1 + 2 + 3 + ... + i - (i + 1) + ... + n
3. sum(1... n) - 2 * (i + 1)
4. x <= 1 + 2 + ... + n 
5. let y = 1 + 2 + ... + n >= x
6. z = y - x， z < n 是否成立？
7. 假设 z > n, 比如 z = n - 1, 那么就可以将n+1
8. 但是问题时 z要变成负数，不能直接使用 -z, 
9. 这样子就变成 -2z了
10. z如果是偶数 z / 2, 那么只要将 z/ 2 变成负数就可以了
11. 如果z是奇数呢？
12. 貌似任何一个变化都是要改变偶数的，所以不大行，必须加1
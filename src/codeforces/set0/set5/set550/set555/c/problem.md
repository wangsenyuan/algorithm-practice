Andrewid the Android is a galaxy-known detective. Now he does not investigate any case and is eating chocolate out of boredom.

A bar of chocolate can be presented as an n × n table, where each cell represents one piece of chocolate. The columns of the table are numbered from 1 to n from left to right and the rows are numbered from top to bottom. Let's call the anti-diagonal to be a diagonal that goes the lower left corner to the upper right corner of the table. First Andrewid eats all the pieces lying below the anti-diagonal. Then he performs the following q actions with the remaining triangular part: first, he chooses a piece on the anti-diagonal and either direction 'up' or 'left', and then he begins to eat all the pieces starting from the selected cell, moving in the selected direction until he reaches the already eaten piece or chocolate bar edge.

After each action, he wants to know how many pieces he ate as a result of this action.

## Input

The first line contains integers n (1 ≤ n ≤ 10<sup>9</sup>) and q (1 ≤ q ≤ 2·10<sup>5</sup>) — the size of the chocolate bar and the number of actions.

Next q lines contain the descriptions of the actions: the i-th of them contains numbers xi and yi (1 ≤ xi, yi ≤ n, xi + yi = n + 1) — the numbers of the column and row of the chosen cell and the character that represents the direction (L — left, U — up).

## Output

Print q lines, the i-th of them should contain the number of eaten pieces as a result of the i-th action.

## Examples

### Input
```
6 5
3 4 U
6 1 L
2 5 L
1 6 U
4 3 U
```

### Output
```
4
3
2
1
2
```

### Input
```
10 6
2 9 U
10 1 U
1 10 U
8 3 L
10 1 L
6 5 U
```

### Output
```
9
1
10
6
0
2
```

### ideas
1. n is too large
2. 第一次假设往上的，那么它肯定吃掉了一整列
3. 然后第二次往左，有可能碰到之前吃掉的那一列
4. 往左的时候，需要知道它左边最近被吃掉的那一列（但是问题出在有些列是断开的）
5. 但是断开的那部分，只可能断在一个点；
6. 虽然n很大，但是q比较小，需要处理的比较少
7. 就是range update

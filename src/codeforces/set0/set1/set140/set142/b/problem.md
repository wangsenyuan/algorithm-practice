# Sir Lancelot's Drill Exercise Problem

## Problem Description

Once upon a time in the Kingdom of Far Far Away lived Sir Lancelot, the chief Royal General. He was very proud of his men and he liked to invite the King to come and watch drill exercises which demonstrated the fighting techniques and tactics of the squad he was in charge of. But time went by and one day Sir Lancelot had a major argument with the Fairy Godmother (there were rumors that the argument occurred after the general spoke badly of the Godmother's flying techniques. That seemed to hurt the Fairy Godmother very deeply).

As the result of the argument, the Godmother put a rather strange curse upon the general. It sounded all complicated and quite harmless: **"If the squared distance between some two soldiers equals to 5, then those soldiers will conflict with each other!"**

## The Curse Explained

The drill exercises are held on a rectangular $n \times m$ field, split into $nm$ square $1 \times 1$ segments for each soldier. 

The square of the distance between the soldiers that stand on squares $(x_1, y_1)$ and $(x_2, y_2)$ equals exactly:
$$(x_1 - x_2)^2 + (y_1 - y_2)^2$$

Now not all $nm$ squad soldiers can participate in the drill exercises as it was before the Fairy Godmother's curse. Unless, of course, the general wants the soldiers to fight with each other or even worse...

### Example of the Curse
If he puts a soldier in the square $(2, 2)$, then he cannot put soldiers in the squares $(1, 4)$, $(3, 4)$, $(4, 1)$ and $(4, 3)$ — each of them will conflict with the soldier in the square $(2, 2)$.

## Task

Your task is to help the general. You are given the size of the drill exercise field. You are asked to calculate the **maximum number of soldiers** that can be simultaneously positioned on this field, so that no two soldiers fall under the Fairy Godmother's curse.

## Input

The single line contains space-separated integers $n$ and $m$ ($1 \leq n, m \leq 1000$) that represent the size of the drill exercise field.

## Output

Print the desired maximum number of warriors.

## Examples

### Example 1
**Input:**
```
2 4
```

**Output:**
```
4
```

**Explanation:** In the first sample test Sir Lancelot can place his 4 soldiers on the $2 \times 4$ court as follows (the soldiers' locations are marked with gray circles on the scheme).

### Example 2
**Input:**
```
3 4
```

**Output:**
```
6
```

**Explanation:** In the second sample test he can place 6 soldiers on the $3 \times 4$ site in the following manner.


### ideas
1. 任何相邻距离为 dx = 1, dy = 2, 或者 dx = 2, dy = 1的，就是有冲突的
2. 也就是棋盘上马的落脚位置
3. 有一个模式是2 * 2 一组，然后每隔 2 * 2， 也就是在 4 * 4 中，有1/4个区域，可以放置一组
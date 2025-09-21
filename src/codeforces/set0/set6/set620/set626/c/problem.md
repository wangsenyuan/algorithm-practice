# Problem C: Block Tower Construction

## Problem Statement

Students in a class are making towers of blocks. Each student makes a (non-zero) tower by stacking pieces lengthwise on top of each other. `n` of the students use pieces made of two blocks and `m` of the students use pieces made of three blocks.

The students don't want to use too many blocks, but they also want to be unique, so no two students' towers may contain the same number of blocks. Find the minimum height necessary for the tallest of the students' towers.

## Input

The first line of the input contains two space-separated integers `n` and `m` (0 ≤ n, m ≤ 1,000,000, n + m > 0) — the number of students using two-block pieces and the number of students using three-block pieces, respectively.

## Output

Print a single integer, denoting the minimum possible height of the tallest tower.

## Examples

### Example 1
**Input:**
```
1 3
```

**Output:**
```
9
```

### Example 2
**Input:**
```
3 2
```

**Output:**
```
8
```

### Example 3
**Input:**
```
5 0
```

**Output:**
```
10
```

## Explanation

- **Example 1:** The student using two-block pieces can make a tower of height 4, and the students using three-block pieces can make towers of height 3, 6, and 9 blocks. The tallest tower has a height of 9 blocks.

- **Example 2:** The students can make towers of heights 2, 4, and 8 with two-block pieces and towers of heights 3 and 6 with three-block pieces, for a maximum height of 8 blocks.

### ideas
1. n个2的倍数, m个3的倍数，组成一个序列，不能重复 
2. 找出那个序列中的最大值
3. 先找出n, 2, 4, 6....
4. 然后找出m个数, 3, 6, 9
5. 然后找出其中6的倍数，k
6. 然后再找出 k个数，k % 2 = 0, or k % 3 = 0
7. 
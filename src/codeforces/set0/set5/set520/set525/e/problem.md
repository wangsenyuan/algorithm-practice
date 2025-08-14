# Problem E: Anya and Cubes

## Problem Description

Anya loves to fold and stick. Today she decided to do just that.

Anya has **n** cubes lying in a line and numbered from 1 to **n** from left to right, with natural numbers written on them. She also has **k** stickers with exclamation marks. We know that the number of stickers does not exceed the number of cubes.

Anya can stick an exclamation mark on a cube and get the factorial of the number written on the cube. For example, if a cube reads 5, then after sticking the sticker it reads 5!, which equals 120.

You need to help Anya count how many ways there are to choose some of the cubes and stick on some of the chosen cubes at most **k** exclamation marks so that the sum of the numbers written on the chosen cubes after the sticking becomes equal to **S**. Anya can stick at most one exclamation mark on each cube.

**Note:** Two ways are considered the same if they have the same set of chosen cubes and the same set of cubes with exclamation marks.

## Input

The first line contains three space-separated integers **n**, **k** and **S**:
- **n** (1 ≤ n ≤ 25) — the number of cubes
- **k** (0 ≤ k ≤ n) — the number of stickers that Anya has  
- **S** (1 ≤ S ≤ 10^16) — the sum that she needs to get

The second line contains **n** positive integers **a_i** (1 ≤ a_i ≤ 10^9) — the numbers written on the cubes. The cubes are described in order from left to right, starting from the first one.

**Note:** Multiple cubes can contain the same numbers.

## Output

Output the number of ways to choose some number of cubes and stick exclamation marks on some of them so that the sum of the numbers becomes equal to the given number **S**.

## Examples

### Example 1
**Input:**
```
2 2 30
4 3
```

**Output:**
```
1
```

**Explanation:** The only way is to choose both cubes and stick an exclamation mark on each of them.

### Example 2
**Input:**
```
2 2 7
4 3
```

**Output:**
```
1
```

**Explanation:** The only way is to choose both cubes but don't stick an exclamation mark on any of them.

### Example 3
**Input:**
```
3 1 1
1 1 1
```

**Output:**
```
6
```

**Explanation:** It is possible to choose any of the cubes in three ways, and also we may choose to stick or not to stick the exclamation mark on it. So, the total number of ways is six.


## ideas
1. 寻找的是b的数量，不是b`的数量
2. 假设将a分成两类，一类是保持原值，一类进行阶乘
3. 那么从原值部分中，假设选择了一个子集，得到和s1, 那么就要在阶乘部分中得到一个子集s2 = S - s1, 且这个子集的大小不能超过k 
4. 
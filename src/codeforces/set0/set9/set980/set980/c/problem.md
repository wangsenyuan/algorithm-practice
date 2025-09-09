# Problem C: Posterization Image Filter

Professor Ibrahim has prepared the final homework for his algorithm's class. He asked his students to implement the Posterization Image Filter.

## Problem Description

Their algorithm will be tested on an array of integers, where the $i$-th integer represents the color of the $i$-th pixel in the image. The image is in black and white, therefore the color of each pixel will be an integer between 0 and 255 (inclusive).

To implement the filter, students are required to divide the black and white color range [0, 255] into groups of consecutive colors, and select one color in each group to be the group's key. In order to preserve image details, the size of a group must not be greater than $k$, and each color should belong to exactly one group.

Finally, the students will replace the color of each pixel in the array with that color's assigned group key.

To better understand the effect, here is an image of a basking turtle where the Posterization Filter was applied with increasing $k$ to the right.

To make the process of checking the final answer easier, Professor Ibrahim wants students to divide the groups and assign the keys in a way that produces the lexicographically smallest possible array.

## Input

The first line of input contains two integers $n$ and $k$ ($1 \leq n \leq 10^5$, $1 \leq k \leq 256$), the number of pixels in the image, and the maximum size of a group, respectively.

The second line contains $n$ integers $p_1, p_2, \ldots, p_n$ ($0 \leq p_i \leq 255$), where $p_i$ is the color of the $i$-th pixel.

## Output

Print $n$ space-separated integers; the lexicographically smallest possible array that represents the image after applying the Posterization filter.

## Examples

### Example 1
**Input:**
```
4 3
2 14 3 4
```

**Output:**
```
0 12 3 3
```

### Example 2
**Input:**
```
5 2
0 2 1 255 254
```

**Output:**
```
0 1 1 254 254
```

## Note

One possible way to group colors and assign keys for the first sample:

- Color 2 belongs to the group [0,2], with group key 0.
- Color 14 belongs to the group [12,14], with group key 12.
- Colors 3 and 4 belong to group [3,5], with group key 3.

Other groups won't affect the result so they are not listed here.


## ideas
1. 先理解一下题目，就是对0...255进行分组，每个分组的size 不超过k, 然后a[i]替换为它的分组号，最后要求得到最小的序列a
2. 从例子2可以看出，没法直接用每个分组k个，然后使用 a[i] / k * k 来表示
3. [0, 2, 1] 按照贪心策略就变成 [0, 2, 0]了
4. a[0]肯定要尽量的小，那么a[0]的分组是确定的
5. x[0] = max(0, a[0] - k)
6. 这样子，就能确定x[0]了
7. a[i] - k, 如果a[i]-k能分配给i, 那么就分配给它
8. 
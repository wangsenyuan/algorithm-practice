# Array Swap Problem

## Problem Description

Little Petya likes arrays of integers a lot. Recently his mother has presented him one such array consisting of n elements. Petya is now wondering whether he can swap any two distinct integers in the array so that the array got unsorted. Please note that Petya can not swap equal integers even if they are in distinct positions in the array. Also note that Petya must swap some two integers even if the original array meets all requirements.

## Sorted Array Definition

Array a (the array elements are indexed from 1) consisting of n elements is called sorted if it meets at least one of the following two conditions:

- a₁ ≤ a₂ ≤ ... ≤ aₙ
- a₁ ≥ a₂ ≥ ... ≥ aₙ

Help Petya find the two required positions to swap or else say that they do not exist.

## Input

- The first line contains a single integer n (1 ≤ n ≤ 10⁵)
- The second line contains n non-negative space-separated integers a₁, a₂, ..., aₙ — the elements of the array that Petya's mother presented him
- All integers in the input do not exceed 10⁹

## Output

If there is a pair of positions that make the array unsorted if swapped, then print the numbers of these positions separated by a space. If there are several pairs of positions, print any of them. If such pair does not exist, print -1. The positions in the array are numbered with integers from 1 to n.

## Examples

### Example 1

**Input:**
```
1
1
```

**Output:**
```
-1
```

### Example 2

**Input:**
```
2
1 2
```

**Output:**
```
-1
```

### Example 3

**Input:**
```
4
1 2 3 4
```

**Output:**
```
1 2
```

### Example 4

**Input:**
```
3
1 1 1
```

**Output:**
```
-1
```

## Note

- In the first two samples the required pairs obviously don't exist
- In the third sample you can swap the first two elements. After that the array will look like this: 2 1 3 4. This array is unsorted


### ideas
1. 假设交换了(i, j), 那么如果 a[i] < a[j], 交换后， a[i] > a[j]
2. 那么如果存在一对数, a[x]  < a[y] (也就是，那么找出两个(i, j) 也满足 a[i] < a[j])进行交换，就可以了
3. 如果不存在x, y（非升序排列），那么进行下一步，或者只有一对
4. 如果是  (x, y, i) 呢？ a[x] < a[y] < a[i] 呢？
5. 似乎也没有问题呐，交换 y, i 即可
# Problem 333D - Table Cropping

## Problem Statement

Gerald found a table consisting of **n rows** and **m columns**. As a prominent expert on rectangular tables, he immediately counted the table's properties, that is, the **minimum of the numbers in the corners** of the table (minimum of four numbers). 

However, he did not like the final value — it seemed to be too small. And to make this value larger, he decided to crop the table a little: delete some columns on the left and some on the right, as well as some rows from the top and some from the bottom. 

Find what the **maximum property** of the table can be after such cropping. Note that the table should have **at least two rows and at least two columns** left in the end. The number of cropped rows or columns from each of the four sides can be zero.

---

## Input

The first line contains two space-separated integers **n** and **m** (2 ≤ n, m ≤ 1000). 

The following n lines describe the table. The i-th of these lines lists the space-separated integers a<sub>i,1</sub>, a<sub>i,2</sub>, ..., a<sub>i,m</sub> (0 ≤ a<sub>i,j</sub> ≤ 10<sup>9</sup>) — the m numbers standing in the i-th row of the table.

---

## Output

Print the answer to the problem.

---

## Examples

### Example 1
**Input:**
```
2 2
1 2
3 4
```
**Output:**
```
1
```

### Example 2
**Input:**
```
3 3
1 0 0
0 1 1
1 0 0
```
**Output:**
```
0
```

---

## Explanation

**Example 1:** Gerald cannot crop the table — table contains only two rows and only two columns.

**Example 2:** If we'll crop the table, the table will contain zero in some corner cell. Also initially it contains two zeros in the corner cells, so the answer is 0.


## ideas
1. 意料之外的难～
2. 假设处理存在一个矩形(r1, c1), (r2, c2)
3. 那么在(r1, c1), (r1, c2), (r2, c1), (r2, c2)都要存在 >= x的数字
4. 处理(r2, c2)的时候，如果存在某个(r1), 它满足某一位(c1)的数量 > 1, 那么就是ok的
5. 为了快速的处理，是需要压缩的
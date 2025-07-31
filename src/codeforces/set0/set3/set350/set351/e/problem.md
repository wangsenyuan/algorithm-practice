# Problem E: Jeff and Inversions

Jeff's friends know full well that the boy likes to get sequences and arrays for his birthday. Thus, Jeff got sequence p₁, p₂, ..., pₙ for his birthday.

Jeff hates inversions in sequences. An inversion in sequence a₁, a₂, ..., aₙ is a pair of indexes i, j (1 ≤ i < j ≤ n), such that an inequality aᵢ > aⱼ holds.

Jeff can multiply some numbers of the sequence p by -1. At that, he wants the number of inversions in the sequence to be minimum. Help Jeff and find the minimum number of inversions he manages to get.

## Input

The first line contains integer n (1 ≤ n ≤ 2000). The next line contains n integers — sequence p₁, p₂, ..., pₙ (|pᵢ| ≤ 10⁵). The numbers are separated by spaces.

## Output

In a single line print the answer to the problem — the minimum number of inversions Jeff can get.

## Examples

### Example 1
**Input:**
```
2
2 1
```

**Output:**
```
0
```

### Example 2
**Input:**
```
9
-2 0 -1 0 -1 2 1 0 -1
```

**Output:**
```
6
``` 


### ideas
1. 无处下手
2. 对a[i] *= -1, 那么也可以当作对其他所有的数 *= -1吗？
3. 按照绝对值从小到大处理
4. 对于当前的数x，根据保持它不变，或者变化后，产生的inversation, 进行对应的变化？
5. 确定顺序后，确实是可以操作的。但是怎么证明能够得到最优解呢？
6. 比如 x < y,（只考虑绝对值），先处理x后， 再处理y，如果x选择了*-1， 那么对于y的选择在最优解中的操作，也是确定的

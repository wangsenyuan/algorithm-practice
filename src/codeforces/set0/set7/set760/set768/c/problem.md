# Problem C: Jon Snow and White Walkers

## Problem Description

Jon Snow now has to fight with White Walkers. He has n rangers, each of which has his own strength. Also Jon Snow has his favourite number x. Each ranger can fight with a white walker only if the strength of the white walker equals his strength. He however thinks that his rangers are weak and need to improve.

Jon now thinks that if he takes the bitwise XOR of strengths of some of rangers with his favourite number x, he might get soldiers of high strength. So, he decided to do the following operation k times:

1. Arrange all the rangers in a straight line in the order of increasing strengths.
2. Take the bitwise XOR (written as ⊕) of the strength of each alternate ranger with x and update its strength.

## Example

Suppose, Jon has 5 rangers with strengths [9, 7, 11, 15, 5] and he performs the operation 1 time with x = 2. He first arranges them in the order of their strengths, [5, 7, 9, 11, 15]. Then he does the following:

- The strength of first ranger is updated to 5 ⊕ 2 = 7
- The strength of second ranger remains the same, i.e. 7
- The strength of third ranger is updated to 9 ⊕ 2 = 11
- The strength of fourth ranger remains the same, i.e. 11
- The strength of fifth ranger is updated to 15 ⊕ 2 = 13

The new strengths of the 5 rangers are [7, 7, 11, 11, 13]

Now, Jon wants to know the maximum and minimum strength of the rangers after performing the above operations k times.

## Input

- First line consists of three integers n, k, x (1 ≤ n ≤ 10^5, 0 ≤ k ≤ 10^5, 0 ≤ x ≤ 10^3) — number of rangers Jon has, the number of times Jon will carry out the operation and Jon's favourite number respectively.
- Second line consists of n integers representing the strengths of the rangers a1, a2, ..., an (0 ≤ ai ≤ 10^3).

## Output

Output two integers, the maximum and the minimum strength of the rangers after performing the operation k times.

## Examples

### Example 1

**Input:**
```
5 1 2
9 7 11 15 5
```

**Output:**
```
13 7
```

### Example 2

**Input:**
```
2 100000 569
605 986
```

**Output:**
```
986 605
```

## ideas
1. 一个数被xor偶数次后，就变成了它自己
2. xor奇数次后，就变成了 a[i] ^ x
3. 所以，问题是每个数被操作了多少次？
4. 虽然n有1e5, 但是 a[i] <= 1e3, 所以，其实只有1000个数
5. 能够brute force吗？
6. 1e8? 似乎是可以的
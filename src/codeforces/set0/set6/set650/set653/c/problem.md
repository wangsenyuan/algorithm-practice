### Nice sequences — single swap to fix

The life goes up and down, just like nice sequences. Sequence t1, t2, ..., tn is called nice if the following two conditions are satisfied:

- For each odd i < n: `t_i < t_{i+1}`
- For each even i < n: `t_i > t_{i+1}`

For example, sequences (2, 8), (1, 5, 1) and (2, 5, 1, 100, 99, 120) are nice, while (1, 1), (1, 2, 3) and (2, 5, 3, 2) are not.

Bear Limak has a sequence of positive integers t1, t2, ..., tn. This sequence is not nice now and Limak wants to fix it by a single swap. He is going to choose two indices i < j and swap elements ti and tj in order to get a nice sequence. Count the number of ways to do so. Two ways are considered different if indices of elements chosen for a swap are different.

### Input
The first line of the input contains one integer n (2 ≤ n ≤ 150000) — the length of the sequence.

The second line contains n integers t1, t2, ..., tn (1 ≤ ti ≤ 150000) — the initial sequence. It's guaranteed that the given sequence is not nice.

### Output
Print the number of ways to swap two elements exactly once in order to get a nice sequence.

### Examples

#### Example 1
Input
```
5
2 8 4 7 7
```
Output
```
2
```

#### Example 2
Input
```
4
200 150 100 50
```
Output
```
1
```

#### Example 3
Input
```
10
3 2 1 4 1 4 1 4 1 4
```
Output
```
8
```

#### Example 4
Input
```
9
1 2 3 4 5 6 7 8 9
```
Output
```
0
```

### Note
In the first sample, there are two ways to get a nice sequence with one swap:

- Swap t2 = 8 with t4 = 7.
- Swap t1 = 2 with t5 = 7.

In the second sample, there is only one way — Limak should swap t1 = 200 with t4 = 50.


### ideas
1. 考虑有一个位置 i(odd), t[i] > t[i+1] 那么必须交换i或者i+1，或者both，使的 t[i] < t[i+1]
2. 同样的也可以找到位置 i(even), t[i] < t[i + 1]
3. 然后如果这样的位置太多，比如有超过3个位置，那么 => 0
4. 否则，就可以brute force
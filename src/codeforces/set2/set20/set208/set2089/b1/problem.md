This is the easy version of the problem. The difference between the versions is that in this version, 𝑘=0
. You can hack only if you solved all versions of this problem.

Ecrade has two sequences 𝑎0,𝑎1,…,𝑎𝑛−1
 and 𝑏0,𝑏1,…,𝑏𝑛−1
 consisting of integers. It is guaranteed that the sum of all elements in 𝑎
 does not exceed the sum of all elements in 𝑏
.

Initially, Ecrade can make exactly 𝑘
 changes to the sequence 𝑎
. It is guaranteed that 𝑘
 does not exceed the sum of 𝑎
. In each change:

Choose an integer 𝑖
 (0≤𝑖<𝑛
) such that 𝑎𝑖>0
, and perform 𝑎𝑖:=𝑎𝑖−1
.
Then Ecrade will perform the following three operations sequentially on 𝑎
 and 𝑏
, which constitutes one round of operations:

For each 0≤𝑖<𝑛
: 𝑡:=min(𝑎𝑖,𝑏𝑖),𝑎𝑖:=𝑎𝑖−𝑡,𝑏𝑖:=𝑏𝑖−𝑡
;
For each 0≤𝑖<𝑛
: 𝑐𝑖:=𝑎(𝑖−1)mod𝑛
;
For each 0≤𝑖<𝑛
: 𝑎𝑖:=𝑐𝑖
;
Ecrade wants to know the minimum number of rounds required for all elements in 𝑎
 to become equal to 0
 after exactly 𝑘
 changes to 𝑎
.

However, this seems a bit complicated, so please help him!

### ideas
1. 考虑操作1以后，要么 a[i]变成0，要么b[i] 变成0， a[i] = a[i] - b[i]
2. 如果a[i]变成了0，那么下一轮，它后面的会变成0。
3. 如果a[i]不是0，那么下一轮， a[i+1] = a[i]
4. b[i]变成0以后，相当于，它对后续的操作，不再起作用了
5. 假设有一个很大很大的b[i], 那么所有的a进过它以后，都会变成0
6. 假想是在移动（环形）数组a，对于a[i], 假设目前它在位置j, 
7. 如果a[i] >= b[j], a[i]损耗b[j], 且b[j]= 0
8. 如果a[i] < b[j], a[i]变成0， 且b[j] 损耗a[i]
9. 对于a[i]来说，如果b[i] > a[i], 那么 a[i]就变成0了
10. 否则的话，假设b[i+1....j] > a[i+1...j] 那么它们就有剩余的b[i+1...j] - a[i+1...j]来损耗a[i]
11. 那对于i来说，是不是找到最近的j，b[i...j] >= a[i...j]？
12. 
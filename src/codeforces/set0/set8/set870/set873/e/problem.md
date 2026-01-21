# Problem E

Alexey recently held a programming contest for students from Berland.  
\(n\) students participated in the contest, the \(i\)-th of them solved \(a_i\) problems.  
Now he wants to award some contestants. Alexey can award the students with diplomas of three different degrees.  
Each student either will receive one diploma of some degree, or won't receive any diplomas at all.

Let \(\text{cnt}_x\) be the number of students that are awarded with diplomas of degree \(x\) \((1 \le x \le 3)\).  
The following conditions must hold:

- For each \(x\) \((1 \le x \le 3)\), \(\text{cnt}_x > 0\);
- For any two degrees \(x\) and \(y\), \(\text{cnt}_x \le 2 \cdot \text{cnt}_y\).

Of course, there are a lot of ways to distribute the diplomas.  
Let \(b_i\) be the degree of diploma the \(i\)-th student will receive (or \(-1\) if the \(i\)-th student won't receive any diploma).  
Also, for any \(x\) such that \(1 \le x \le 3\) let:

- \(c_x\) be the **maximum** number of problems solved by a student that receives a diploma of degree \(x\);
- \(d_x\) be the **minimum** number of problems solved by a student that receives a diploma of degree \(x\).

Alexey wants to distribute the diplomas in such a way that:

1. If student \(i\) solved more problems than student \(j\), then he has to be awarded not worse than student \(j\)  
   (it's impossible that student \(j\) receives a diploma and \(i\) doesn't receive any, and also it's impossible that both of them receive a diploma, but \(b_j < b_i\));
2. \(d_1 - c_2\) is maximum possible;
3. Among all ways that maximize the previous expression, \(d_2 - c_3\) is maximum possible;
4. Among all ways that correspond to the two previous conditions, \(d_3 - c_{-1}\) is maximum possible,  
   where \(c_{-1}\) is the maximum number of problems solved by a student that doesn't receive any diploma  
   (or \(0\) if each student is awarded with some diploma).

Help Alexey to find a way to award the contestants!

## Input

The first line contains one integer number \(n\) \((3 \le n \le 3000)\).

The second line contains \(n\) integer numbers \(a_1, a_2, \dots, a_n\) \((1 \le a_i \le 5000)\).

## Output

Output \(n\) numbers. The \(i\)-th number must be equal to the degree of diploma the \(i\)-th contestant will receive (or \(-1\) if he doesn't receive any diploma).

If there are multiple optimal solutions, print any of them. It is guaranteed that the answer always exists.

## Examples

**Example 1**

Input:
```
4
1 2 3 4
```

Output:
```
3 3 2 1 
```

**Example 2**

Input:
```
6
1 4 3 1 1 2
```

Output:
```
-1 1 2 -1 -1 3 
```


### ideas
1. 好复杂的题目
2. 分数越高，等级越小
3. a按照分数升序排列, 如果i-1有等级x, 那么a[i]可以是[1...x]
4. 这时候d[x-1] - c[x] = 假设在位置i处发生了变化，那么就a[i] - a[i-1]
5. 如果 a[i-1] = 3, 那么是否可以给a[i] = 1; (直接跳过2？)
6. 不可以，因为cnt[2] > 0 的条件无法满足了。如果a[i] = 1, 那么后面只能也赋值为1
7. 假设没有等级是4，
8. 那么就是可以从4开始，或者从3开始（但是不能从2、1开始）， 且都以1结束
9. 对于1，2，3假设w = min(cnt[1...3]), v = max(cnt[1...3])
10. 那么有 v <= 2 * w
11. 考虑例子 [1 4 3 1 1 2]
12. 排序后, 1 1 1 2 3 4， 结果为[-1, -1, -1, 3, 2, 1]
13. 现在的答案cnt[1, 2, 3] = 1
14. d[1] - c[2] = 4 - 3 = 1
15. d[2] - c[1] = 3 - 2 = 1
16. d[3] - c[-1] = 4 - 1 = 3
17. 这里有3个需要优化的目标，从后往前处理(这样子，最后的那个肯定要赋值为1)
18. dp[i][x] = data 表示从i开始，且i的等级为x时的最优解
19. 那么把i...j 分为一组（这组都赋值为x+1）那么更新结果 dp[i][x+1] = dp[j+1][x]
20. 似乎也能搞～
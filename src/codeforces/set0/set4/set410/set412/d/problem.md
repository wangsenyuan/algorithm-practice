# Problem Description

The employees of the R1 company often spend time together: they watch football, they go camping, they solve contests. So, it's no big deal that sometimes someone pays for someone else.

Today is the day of giving out money rewards. The R1 company CEO will invite employees into his office one by one, rewarding each one for the hard work this month. The CEO knows who owes money to whom. And he also understands that if he invites person x to his office for a reward, and then immediately invite person y, who has lent some money to person x, then they can meet. Of course, in such a situation, the joy of person x from his brand new money reward will be much less. Therefore, the R1 CEO decided to invite the staff in such an order that the described situation will not happen for any pair of employees invited one after another.

However, there are a lot of employees in the company, and the CEO doesn't have a lot of time. Therefore, the task has been assigned to you. Given the debt relationships between all the employees, determine in which order they should be invited to the office of the R1 company CEO, or determine that the described order does not exist.

## Input

The first line contains space-separated integers n and m — the number of employees in R1 and the number of debt relations. Each of the following m lines contains two space-separated integers ai, bi (1 ≤ ai, bi ≤ n; ai ≠ bi), these integers indicate that the person number ai owes money to a person a number bi. Assume that all the employees are numbered from 1 to n.

It is guaranteed that each pair of people p, q is mentioned in the input data at most once. In particular, the input data will not contain pairs p, q and q, p simultaneously.

## Output

Print -1 if the described order does not exist. Otherwise, print the permutation of n distinct integers. The first number should denote the number of the person who goes to the CEO office first, the second number denote the person who goes second and so on.

If there are multiple correct orders, you are allowed to print any of them.

## Examples

### Example 1

**Input:**
```
2 1
1 2
```

**Output:**
```
2 1
```

### Example 2

**Input:**
```
3 3
1 2
2 3
3 1
```

**Output:**
```
2 1 3
```

## ideas
1. (a, b) 表示b借钱给了a，（a欠钱）
2. 那么a的后面不能跟着b
3. 那么先邀请b
4. 但是如果出现了loop（因为(a, b), (b, a)不会同时出现）
5. 那么这个环的长度至少是3，
6. 那是不是不用管topo的顺序，就从最少deg的地方入手就好了？
7. 似乎也不对，
8. (1 -> 2), (2 -> 3), (3 -> 4), (4 -> 1), (4 -> 2)
9. 先选择1, 接下来有可能选到2，那么就糟糕了（因为1欠2的钱）
10. 但是选择3就是ok的
11. 假设在一个scc中，能够给这些节点涂色，red, blue 
12. 那么先把red的邀请完，再要求blue，且关注一下分界点的情况，就可以了
13. 假设选择了1，然后选择所有欠钱给1的人
14. (a, b, c), 但是有可能a也欠钱b。
15. 但是如果能找到u，（u欠钱a）先把u加进来；那么a的后面就不会是b
16. 而且似乎u肯定是存在的（不可能是1）因为a是在一个圈里面的
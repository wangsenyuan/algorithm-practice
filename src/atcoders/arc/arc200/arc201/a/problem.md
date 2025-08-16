# CatCoder Double Contest (C2C)

## Problem Description

In the year 2222, CatCoder will hold the CatCoder Double Contest (abbreviated as C2C).

There are $N$ writers who have problem proposals. Each writer's problem proposals are classified into 3 types by difficulty: **Hard**, **Medium**, **Easy**, and the $i$-th writer has $A_i$, $B_i$, $C_i$ problem proposals of Hard, Medium, Easy, respectively.

Each C2C simultaneously holds 2 divisions, **Div.1** and **Div.2**. The problem proposals required for each division are as follows:

- **Div.1**: One Hard and one Medium problem proposal from the same writer
- **Div.2**: One Medium and one Easy problem proposal from the same writer

**Note**: The writers for Div.1 and Div.2 do not necessarily have to be the same. Also, each problem proposal can be used in at most one division of one C2C.

**Find the maximum number of times C2C can be held.**

$T$ test cases are given, so find the answer for each.

## Constraints

- $1 \leq T \leq 10^5$
- $1 \leq N \leq 2 \times 10^5$
- $1 \leq A_i, B_i, C_i \leq 10^9$
- The sum of $N$ over all test cases is at most $2 \times 10^5$
- All input values are integers

## Input

The input is given from Standard Input in the following format:

```
T
case_1
case_2
⋮
case_T
```

Each test case is given in the following format:

```
N
A_1 B_1 C_1
A_2 B_2 C_2
⋮
A_N B_N C_N
```

## Output

Output $T$ lines.

The $i$-th line should contain the maximum number of times C2C can be held for the $i$-th test case.

## Sample Input 1

```
5
2
3 1 4
1 5 3
1
1 1 1
3
5 7 5
1 11 99
3 1 2
5
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000
6
835549144 866512240 105679868
473233032 625162103 823002638
125467290 37501686 380787083
8043910 721085797 254272563
97327826 744196952 18713225
978152989 90127986 33086297
```

## Sample Output 1

```
2
0
7
2500000000
998830769
```

## Explanation

For the first test case, C2C can be held **2 times** by using problem proposals as follows:

| Contest  | Div.1                            | Div.2                            |
| -------- | -------------------------------- | -------------------------------- |
| 1st time | Hard, Medium from the 1st writer | Medium, Easy from the 2nd writer |
| 2nd time | Hard, Medium from the 2nd writer | Medium, Easy from the 2nd writer |


### ideas
1. 每次coc，需要找到两组题目，div1 H + M, div2, M + E
2. 每组题目必须从一个人获得（但两组没有要求）
3. 这里关键好像是M（H的题目，必须出现在div1，E的题目必须出现在div2）
4. 只有M是两边都需要的
5. 如果一个人能够同时出div1 + div2， 似乎让他一个人出，是不是没有问题？
6. 然后剩下的就是要么只能出div1的，要么只能出div2的，还有就是都不能出的
7. 
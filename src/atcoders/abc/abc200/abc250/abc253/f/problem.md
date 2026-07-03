Score : 
500 points

Problem Statement
We have an 
N×M matrix, whose elements are initially all 
0.

Process 
Q given queries. Each query is in one of the following formats.

1 l r x : Add 
x to every element in the 
l-th, 
(l+1)-th, 
…, and 
r-th column.
2 i x: Replace every element in the 
i-th row with 
x.
3 i j: Print the 
(i,j)-th element.
Constraints
1≤N,M,Q≤2×10 
5
 
Every query is in one of the formats listed in the Problem Statement.
For each query in the format 1 l r x, 
1≤l≤r≤M and 
1≤x≤10 
9
 .
For each query in the format 2 i x, 
1≤i≤N and 
1≤x≤10 
9
 .
For each query in the format 3 i j, 
1≤i≤N and 
1≤j≤M.
At least one query in the format 3 i j is given.
All values in input are integers.


### ideas
1. 根据操作次序，如果知道i行的最后一次操作2的时刻，计算这个操作后，操作1的影响。
2. 所以是一个持久化树
3. 但是问题是，持久化树修改的是单个节点，这里修改的是连续的节点，
4. 要怎么处理呢？好像是一样的。试试看
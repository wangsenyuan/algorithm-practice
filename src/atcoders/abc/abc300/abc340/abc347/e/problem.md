# Problem Description

There is an integer sequence A = (A₁, A₂, ..., A_N) of length N, where all elements are initially set to 0. Also, there is a set S, which is initially empty.

Perform the following Q queries in order. Find the value of each element in the sequence A after processing all Q queries.

## Query Format
The i-th query is in the following format:
- An integer xᵢ is given
- If xᵢ is contained in S, remove xᵢ from S
- Otherwise, insert xᵢ to S
- Then, for each j = 1, 2, ..., N:
  - Add |S| to Aⱼ if j ∈ S

## Notes
- |S| denotes the number of elements in the set S
- For example, if S = {3, 4, 7}, then |S| = 3

## Sample Input and Output

### Sample Input 1
```
3 4
1 3 3 2
```

### Sample Output 1
```
6 2 2
```

### Explanation
1. First query (x = 1):
   - Insert 1 to S, making S = {1}
   - Add |S| = 1 to A₁
   - Sequence becomes A = (1, 0, 0)

2. Second query (x = 3):
   - Insert 3 to S, making S = {1, 3}
   - Add |S| = 2 to A₁ and A₃
   - Sequence becomes A = (3, 0, 2)

3. Third query (x = 3):
   - Remove 3 from S, making S = {1}
   - Add |S| = 1 to A₁
   - Sequence becomes A = (4, 0, 2)

4. Fourth query (x = 2):
   - Insert 2 to S, making S = {1, 2}
   - Add |S| = 2 to A₁ and A₂
   - Sequence becomes A = (6, 2, 2)

Final sequence: A = (6, 2, 2)
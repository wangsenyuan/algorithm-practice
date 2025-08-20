# Problem C: Gringotts Vault Security

## Problem Description

Harry, Ron and Hermione have figured out that Helga Hufflepuff's cup is a horcrux. Through her encounter with Bellatrix Lestrange, Hermione came to know that the cup is present in Bellatrix's family vault in Gringott's Wizarding Bank.

The Wizarding bank is in the form of a tree with total $n$ vaults where each vault has some type, denoted by a number between 1 to $m$. A tree is an undirected connected graph with no cycles.

The vaults with the highest security are of type $k$, and all vaults of type $k$ have the highest security.

There can be at most $x$ vaults of highest security.

Also, if a vault is of the highest security, its adjacent vaults are guaranteed to not be of the highest security and their type is guaranteed to be less than $k$.

Harry wants to consider every possibility so that he can easily find the best path to reach Bellatrix's vault. So, you have to tell him, given the tree structure of Gringotts, the number of possible ways of giving each vault a type such that the above conditions hold.

## Input

- The first line of input contains two space separated integers, $n$ and $m$ — the number of vaults and the number of different vault types possible. $(1 \leq n \leq 10^5, 1 \leq m \leq 10^9)$.

- Each of the next $n - 1$ lines contain two space separated integers $u_i$ and $v_i$ $(1 \leq u_i, v_i \leq n)$ representing the $i$-th edge, which shows there is a path between the two vaults $u_i$ and $v_i$. It is guaranteed that the given graph is a tree.

- The last line of input contains two integers $k$ and $x$ $(1 \leq k \leq m, 1 \leq x \leq 10)$, the type of the highest security vault and the maximum possible number of vaults of highest security.

## Output

Output a single integer, the number of ways of giving each vault a type following the conditions modulo $10^9 + 7$.

## Examples

### Example 1
**Input:**
```
4 2
1 2
2 3
1 4
1 2
```

**Output:**
```
1
```

### Example 2
**Input:**
```
3 3
1 2
1 3
2 1
```

**Output:**
```
13
```

### Example 3
**Input:**
```
3 1
1 2
1 3
1 1
```

**Output:**
```
0
```

## Note

In test case 1, we cannot have any vault of the highest security as its type is 1 implying that its adjacent vaults would have to have a vault type less than 1, which is not allowed. Thus, there is only one possible combination, in which all the vaults have type 2.


### ideas
1. 就是说，在树中选择(最多)x个点，赋值为k，且它们的邻居节点 < k 的奇数
2. dp[u][i][0/1] 表示在子树u中，给i个节点赋值为k，u节点是否也是k (0表示不是，1表示是)
3. 这个状态不好搞
4. 把状态改为父节点是否 < k, = k, > k
5. dp[u][i][0] = dp[v][j][0] * (k - 1) + (m - k) * dp[v][j][2] + dp[v][j][1]
6.    在p取值 < k 的时候, 那么u可以取值 < k, 或者 > k, 或者 = k
7. dp[u][i][1] = (k - 1) * dp[v][j][0] (在p为k的时候， u只能取 < k)
8. dp[u][i][2] = (k - 1) * dp[v][j][0] + (m - k) * dp[v][j][2] (当p > k)的时候， u不能取值为k
9. 这个状态转移似乎是正确的
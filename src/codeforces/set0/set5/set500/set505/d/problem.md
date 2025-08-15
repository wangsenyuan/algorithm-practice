# Problem D: Teleportation Pipes in Shuseki Kingdom

## Problem Description

Shuseki Kingdom is the world's leading nation for innovation and technology. There are n cities in the kingdom, numbered from 1 to n.

Thanks to Mr. Kitayuta's research, it has finally become possible to construct teleportation pipes between two cities. A teleportation pipe will connect two cities unidirectionally, that is, a teleportation pipe from city x to city y cannot be used to travel from city y to city x. The transportation within each city is extremely developed, therefore if a pipe from city x to city y and a pipe from city y to city z are both constructed, people will be able to travel from city x to city z instantly.

Mr. Kitayuta is also involved in national politics. He considers that the transportation between the m pairs of city (aᵢ, bᵢ) (1 ≤ i ≤ m) is important. He is planning to construct teleportation pipes so that for each important pair (aᵢ, bᵢ), it will be possible to travel from city aᵢ to city bᵢ by using one or more teleportation pipes (but not necessarily from city bᵢ to city aᵢ). Find the minimum number of teleportation pipes that need to be constructed. So far, no teleportation pipe has been constructed, and there is no other effective transportation between cities.

## Input

The first line contains two space-separated integers n and m (2 ≤ n ≤ 10⁵, 1 ≤ m ≤ 10⁵), denoting the number of the cities in Shuseki Kingdom and the number of the important pairs, respectively.

The following m lines describe the important pairs. The i-th of them (1 ≤ i ≤ m) contains two space-separated integers aᵢ and bᵢ (1 ≤ aᵢ, bᵢ ≤ n, aᵢ ≠ bᵢ), denoting that it must be possible to travel from city aᵢ to city bᵢ by using one or more teleportation pipes (but not necessarily from city bᵢ to city aᵢ). It is guaranteed that all pairs (aᵢ, bᵢ) are distinct.

## Output

Print the minimum required number of teleportation pipes to fulfill Mr. Kitayuta's purpose.

## Examples

### Example 1
**Input:**
```
4 5
1 2
1 3
1 4
2 3
2 4
```

**Output:**
```
3
```

### Example 2
**Input:**
```
4 6
1 2
1 4
2 3
2 4
3 2
3 4
```

**Output:**
```
4
```


## ideas
1. 尽量的使用树，如果要产生回边，就把叶子节点和头连起来？
2. 假设存在（a, b), (a, c), (a, d)
3. 使用了(a, b)后，也使用(a, c), (a, d) 却不是好主意
4. 因为后面有可能存在(b, c), (c, d)，
5. 这里先不管，先把他们都连起来（会形成回路）
6. 然后保留那些必须的边
7. 这个时候，应该尽量的使用dfs
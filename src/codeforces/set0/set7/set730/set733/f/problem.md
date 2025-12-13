In one kingdom there are n cities and m two-way roads. Each road connects a pair of cities, and for each road we know the level of drivers dissatisfaction — the value wi.

For each road we know the value ci — how many lamziks we should spend to reduce the level of dissatisfaction with this road by one. Thus, to reduce the dissatisfaction with the i-th road by k, we should spend k·ci lamziks. And it is allowed for the dissatisfaction to become zero or even negative.

In accordance with the king's order, we need to choose n - 1 roads and make them the main roads. An important condition must hold: it should be possible to travel from any city to any other by the main roads.

The road ministry has a budget of S lamziks for the reform. The ministry is going to spend this budget for repair of some roads (to reduce the dissatisfaction with them), and then to choose the n - 1 main roads.

Help to spend the budget in such a way and then to choose the main roads so that the total dissatisfaction with the main roads will be as small as possible. The dissatisfaction with some roads can become negative. It is not necessary to spend whole budget S.

It is guaranteed that it is possible to travel from any city to any other using existing roads. Each road in the kingdom is a two-way road.

## Input

The first line contains two integers n and m (2 ≤ n ≤ 2·10⁵, n - 1 ≤ m ≤ 2·10⁵) — the number of cities and the number of roads in the kingdom, respectively.

The second line contains m integers w₁, w₂, ..., wₘ (1 ≤ wᵢ ≤ 10⁹), where wᵢ is the drivers dissatisfaction with the i-th road.

The third line contains m integers c₁, c₂, ..., cₘ (1 ≤ cᵢ ≤ 10⁹), where cᵢ is the cost (in lamziks) of reducing the dissatisfaction with the i-th road by one.

The next m lines contain the description of the roads. The i-th of this lines contain a pair of integers aᵢ and bᵢ (1 ≤ aᵢ, bᵢ ≤ n, aᵢ ≠ bᵢ) which mean that the i-th road connects cities aᵢ and bᵢ. All roads are two-way oriented so it is possible to move by the i-th road from aᵢ to bᵢ, and vice versa. It is allowed that a pair of cities is connected by more than one road.

The last line contains one integer S (0 ≤ S ≤ 10⁹) — the number of lamziks which we can spend for reforms.

## Output

In the first line print K — the minimum possible total dissatisfaction with main roads.

In each of the next n - 1 lines print two integers x, vₓ, which mean that the road x is among main roads and the road x, after the reform, has the level of dissatisfaction vₓ.

Consider that roads are numbered from 1 to m in the order as they are given in the input data. The edges can be printed in arbitrary order. If there are several answers, print any of them.

## Examples

### Example 1

**Input:**
```
6 9
1 3 1 1 3 1 2 2 2
4 1 4 2 2 5 3 1 6
1 2
1 3
2 3
2 4
2 5
3 5
3 6
4 5
5 6
7
```

**Output:**
```
0
1 1
3 1
6 1
7 2
8 -5
```

### Example 2

**Input:**
```
3 3
9 5 1
7 7 2
2 1
3 1
3 2
2
```

**Output:**
```
5
3 0
2 5
```


### ideas
1. 使用不超过S，将n-1条路的不满意度，使的它们组成一棵树，然后总的不满意度最低
2. wi - ci * di
3. 假设，我们知道最后的n-1条边，然后这个时候的分配，其实是应该把所有边上使用的预算，都集中在一条边上
4. 且这条边的c最大
5. 那就简单了，就是按照c升序排，在当前边必须被选中的情况下，找到最小值就可以了
6. 但是如果naive的处理，将是n * n
7. 按照w升序（正常的mst处理）
8. 然后添加额外的边，看是否能得到更好的结果
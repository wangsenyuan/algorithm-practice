# Problem D: Bubble Cup Journey

## Problem Description

The competitors of Bubble Cup X gathered after the competition and discussed what is the best way to get to know the host country and its cities.

After exploring the map of Serbia for a while, the competitors came up with the following facts:
- The country has **V** cities which are indexed with numbers from 1 to V
- There are **E** bi-directional roads that connect the cities
- Each road has a weight (the time needed to cross that road)
- There are **N** teams at the Bubble Cup

The competitors came up with the following plan: each of the N teams will start their journey in one of the V cities, and some of the teams may share the starting position.

They want to find the shortest time **T**, such that:
- Every team can move in these T minutes
- The number of different cities they end up in is at least **K** (because they will only get to know the cities they end up in)
- A team doesn't have to be on the move all the time - if they like it in a particular city, they can stay there and wait for the time to pass

Please help the competitors determine the shortest time T so it's possible for them to end up in at least K different cities, or print -1 if that is impossible no matter how they move.

**Note:** There can exist multiple roads between some cities.

## Input

- **Line 1:** Four integers: V, E, N, and K
  - 1 ≤ V ≤ 600 (number of cities)
  - 1 ≤ E ≤ 20000 (number of roads)
  - 1 ≤ N ≤ min(V, 200) (number of teams)
  - 1 ≤ K ≤ N (minimum number of different cities to end up in)

- **Line 2:** N integers representing the starting cities for each team

- **Lines 3 to E+2:** Information about the roads in format: `Ai Bi Ti`
  - 1 ≤ Ai, Bi ≤ V (cities connected by the road)
  - 1 ≤ Ti ≤ 10000 (time in minutes to cross the road)

## Output

Output a single integer representing the minimal time the teams can move for, such that they end up in at least K different cities, or output -1 if there is no solution.

**Note:** If the solution exists, the result will be no greater than 1731311.

## Example

### Input
```
6 7 5 4
5 5 2 2 5
1 3 3
1 5 2
1 6 5
2 5 4
2 6 7
3 4 11
3 5 3
```

### Output
```
3
```

### Explanation

- Three teams start from city 5, and two teams start from city 2
- If they agree to move for 3 minutes, one possible situation would be:
  - Two teams in city 2
  - One team in city 5
  - One team in city 3
  - One team in city 1
- This results in four different cities where the teams end their journey, which satisfies K = 4.


## ideas
1. 考虑图是连通的，始终能够得到答案，也就是在时间T内，存在一个节点的集合，大小至少为K
2. 总能安排，这N个人，从他们的开始位置，移动到这k个节点上面去
3. 先考虑一个*笨办法*，对于每个人，可以知道在时间T内，它能访问到的节点的集合
4. 那么就是n个集合的并集，就是能够访问的节点的位置
5. 就是多源bfs？
6. 那这样子，是不是也不用二分查询了？
7. 但是问题，似乎出现在当多个人处在同一个节点上的时候，就有问题了（也就是怎么并起来）
8. 这里N比较小（200）应该是个突破口
9. 那么二分，然后在T时间内，找出N个人能够访问到的集合，然后merge在一起，
10. 然后看这个size >= k? 这样子，大概可以在 N * (E + V) * log(T) 内完成？ 
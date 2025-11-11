# The Two Routes

## Problem Statement

In Absurdistan, there are n towns (numbered 1 through n) and m bidirectional railways. There is also an absurdly simple road network — for each pair of different towns x and y, there is a bidirectional road between towns x and y if and only if there is no railway between them. Travelling to a different town using one railway or one road always takes exactly one hour.

A train and a bus leave town 1 at the same time. They both have the same destination, town n, and don't make any stops on the way (but they can wait in town n). The train can move only along railways and the bus can move only along roads.

You've been asked to plan out routes for the vehicles; each route can use any road/railway multiple times. One of the most important aspects to consider is safety — in order to avoid accidents at railway crossings, the train and the bus must not arrive at the same town (except town n) simultaneously.

Under these constraints, what is the minimum number of hours needed for both vehicles to reach town n (the maximum of arrival times of the bus and the train)? Note, that bus and train are not required to arrive to the town n at the same moment of time, but are allowed to do so.

## Input

The first line of the input contains two integers n and m (2 ≤ n ≤ 400, 0 ≤ m ≤ n(n - 1) / 2) — the number of towns and the number of railways respectively.

Each of the next m lines contains two integers u and v, denoting a railway between towns u and v (1 ≤ u, v ≤ n, u ≠ v).

You may assume that there is at most one railway connecting any two towns.

## Output

Output one integer — the smallest possible time of the later vehicle's arrival in town n. If it's impossible for at least one of the vehicles to reach town n, output -1.

## Examples

### Example 1
**Input**
```
4 2
1 3
3 4
```
**Output**
```
2
```

### Example 2
**Input**
```
4 6
1 2
1 3
1 4
2 3
2 4
3 4
```
**Output**
```
-1
```

### Example 3
**Input**
```
5 5
4 2
3 5
4 5
5 1
1 2
```
**Output**
```
3
```

## Note

In the first sample, the train can take the route 1→3→4 and the bus can take the route 1→2→4 or 1→4. Note that they can arrive at town 4 at the same time.

In the second sample, Absurdistan is ruled by railwaymen. There are no roads, so there's no way for the bus to reach town n.


### ideas
1. 有铁路的地方，没有公路；
2. 也就是所有的节点间，都被连通的，要么被公路，要么被铁路
3. 所以，先看（1，n)有没有铁路，有的话，就考虑公路部分，否则考虑铁路部分
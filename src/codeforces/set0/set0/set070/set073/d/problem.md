Vasya plays FreeDiv. In this game he manages a huge state, which has n cities and m two-way roads between them. Unfortunately, not from every city you can reach any other one moving along these roads. Therefore Vasya decided to divide the state into provinces so that in every province, one could reach from every city all the cities of the province, but there are no roads between provinces.

Unlike other turn-based strategies, in FreeDiv a player has the opportunity to build tunnels between cities. The tunnels are two-way roads along which one can move armies undetected by the enemy. However, no more than one tunnel can be connected to each city. As for Vasya, he wants to build a network of tunnels so that any pair of cities in his state were reachable by some path consisting of roads and a tunnels. But at that no more than k tunnels are connected to each province (otherwise, the province will be difficult to keep in case other provinces are captured by enemy armies).

Vasya discovered that maybe he will not be able to build such a network for the current condition of the state. Maybe he'll have first to build several roads between cities in different provinces to merge the provinces. Your task is to determine the minimum number of roads Vasya needs to build so that it was possible to build the required network of tunnels in the resulting state.

Input
The first line contains three integers n, m and k (1 ≤ n, k ≤ 106, 0 ≤ m ≤ 106). Each of the next m lines contains two integers. They are the numbers of cities connected by a corresponding road. No road connects city to itself and there is at most one road between each pair of cities.

Output
Print a single number, the minimum number of additional roads.


### ideas
1. 先根据已有的road，构建出province
2. 假设有w个prvince，那么需要 w - 1条tunnels把它们连接起来
3. 这个k是什么意思呢？
4. he wants to build a network of tunnels so that any pair of cities in his state were reachable by some path consisting of roads and a tunnels.
5. 这句话比较重要，就是任意两个provice中间应该有一条tunnel。
6. 假设有x个provice, 那么按照这个要求
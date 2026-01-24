# Problem E

Egor is a famous Russian singer, rapper, actor and blogger, and finally he decided to give a concert in the sunny Republic of Dagestan.

There are 𝑛 cities in the republic, some of them are connected by 𝑚 directed roads without any additional conditions. In other words, road system of Dagestan represents an arbitrary directed graph. Egor will arrive to the city 1, travel to the city 𝑛 by roads along some path, give a concert and fly away.

As any famous artist, Egor has lots of haters and too annoying fans, so he can travel only by safe roads. There are two types of the roads in Dagestan, black and white: black roads are safe at night only, and white roads — in the morning. Before the trip Egor's manager's going to make a schedule: for each city he'll specify it's color, black or white, and then if during the trip they visit some city, the only time they can leave it is determined by the city's color: night, if it's black, and morning, if it's white. After creating the schedule Egor chooses an available path from 1 to 𝑛, and for security reasons it has to be the shortest possible.

Egor's manager likes Dagestan very much and wants to stay here as long as possible, so he asks you to make such schedule that there would be no path from 1 to 𝑛 or the shortest path's length would be greatest possible.

A path is one city or a sequence of roads such that for every road (excluding the first one) the city this road goes from is equal to the city previous road goes into. Egor can move only along paths consisting of safe roads only.

The path length is equal to the number of roads in it. The shortest path in a graph is a path with smallest length.

## Input

The first line contains two integers 𝑛, 𝑚 (1≤𝑛≤500000, 0≤𝑚≤500000) — the number of cities and the number of roads.

The 𝑖-th of next 𝑚 lines contains three integers — 𝑢𝑖, 𝑣𝑖 and 𝑡𝑖 (1≤𝑢𝑖,𝑣𝑖≤𝑛, 𝑡𝑖∈{0,1}) — numbers of cities connected by road and its type, respectively (0 — night road, 1 — morning road).

## Output

In the first line output the length of the desired path (or −1, if it's possible to choose such schedule that there's no path from 1 to 𝑛).

In the second line output the desired schedule — a string of 𝑛 digits, where 𝑖-th digit is 0, if the 𝑖-th city is a night one, and 1 if it's a morning one.

If there are multiple answers, print any.

## Examples

### Example 1

**Input:**
```
3 4
1 2 0
1 3 1
2 3 0
2 3 1
```

**Output:**
```
2
011
```

### Example 2

**Input:**
```
4 8
1 1 0
1 3 0
1 3 1
3 2 0
2 1 0
3 4 1
2 4 0
2 4 1
```

**Output:**
```
3
1101
```

### Example 3

**Input:**
```
5 10
1 2 0
1 3 1
1 4 0
2 3 0
2 3 1
2 5 0
3 4 0
3 4 1
4 2 1
4 5 0
```

**Output:**
```
-1
11111
```

## Note

For the first sample, if we paint city 1 white, the shortest path is 1→3. Otherwise, it's 1→2→3 regardless of other cities' colors.

For the second sample, we should paint city 3 black, and there are both black and white roads going from 2 to 4. Note that there can be a road connecting a city with itself.


### ideas
1. 通过选定一些节点的颜色，使的从1.。。n的最短距离最长
2. 把所有节点分成两个u0, u1; u0, 更新所有可以在晚上可以通行的下一个节点
3. u1更新在白天通行的路径
4. 但是问题出现在，到达一个城市后，才需要选择它的颜色
5. 
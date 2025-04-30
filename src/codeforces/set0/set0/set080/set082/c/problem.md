The Berland Kingdom is a set of n cities connected with each other with n - 1 railways. Each road connects exactly two different cities. The capital is located in city 1. For each city there is a way to get from there to the capital by rail.

In the i-th city there is a soldier division number i, each division is characterized by a number of ai. It represents the priority, the smaller the number, the higher the priority of this division. All values of ai are different.

One day the Berland King Berl Great declared a general mobilization, and for that, each division should arrive in the capital. Every day from every city except the capital a train departs. So there are exactly n - 1 departing trains each day. Each train moves toward the capital and finishes movement on the opposite endpoint of the railway on the next day. It has some finite capacity of cj, expressed in the maximum number of divisions, which this train can transport in one go. Each train moves in the direction of reducing the distance to the capital. So each train passes exactly one railway moving from a city to the neighboring (where it stops) toward the capital.

In the first place among the divisions that are in the city, division with the smallest number of ai get on the train, then with the next smallest and so on, until either the train is full or all the divisions are be loaded. So it is possible for a division to stay in a city for a several days.

The duration of train's progress from one city to another is always equal to 1 day. All divisions start moving at the same time and end up in the capital, from where they don't go anywhere else any more. Each division moves along a simple path from its city to the capital, regardless of how much time this journey will take.

Your goal is to find for each division, in how many days it will arrive to the capital of Berland. The countdown begins from day 0.

### ideas
1. 和capital直接相连的，第一天就可以到达
2. 感觉要从叶子节点开始，在某个节点u，目前聚集了一组人，（这个节点上，也已经通过了很多人）
3. 似乎有点混乱
4. 模拟每个train。在每条边上设置一个等待的queue（queue整体的size = n）
5. 每轮至少会有一部分train，会被舍弃掉。考虑从1到n是一条直线的情况下，最多需要模拟n-1次
6. 所以不大行。
7. 假设只考虑子树u，这个子树中到达节点u的顺序已经知道了（考虑了层级，以及a[i])的关系
8. n是5000. 所以模拟是可以的
9. 
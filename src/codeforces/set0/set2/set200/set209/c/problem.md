Vasya went for a walk in the park. The park has `n` glades, numbered from `1` to `n`. There are `m` trails between the glades. The trails are numbered from `1` to `m`, where the `i`-th trail connects glades `xi` and `yi`. The numbers of the connected glades may be the same (`xi = yi`), which means that a trail connects a glade to itself. Also, two glades may have several non-intersecting trails between them.

Vasya is on glade `1`. He wants to walk on **all** trails of the park **exactly once**, so that he can eventually return to glade `1`. Unfortunately, Vasya does not know whether this walk is possible or not. Help Vasya determine whether the walk is possible or not. If such a walk is impossible, find the **minimum** number of trails the authorities need to add to the park in order to make the described walk possible.

Vasya can shift from one trail to another only on glades. He can move on the trails in both directions. If Vasya started going on the trail that connects glades `a` and `b` from glade `a`, then he must finish this trail on glade `b`.

## Input

The first line contains two integers `n` and `m` (`1 <= n <= 10^6`; `0 <= m <= 10^6`) — the number of glades in the park and the number of trails in the park, respectively.

The next `m` lines specify the trails. The `i`-th line specifies the `i`-th trail as two space-separated numbers `xi`, `yi` (`1 <= xi, yi <= n`) — the numbers of the glades connected by this trail.

## Output

Print a single integer — the answer to the problem. If Vasya's walk is possible without adding extra trails, print `0`; otherwise print the minimum number of trails the authorities need to add to the park in order to make Vasya's walk possible.

## Examples

### Example 1

**Input**

```
3 3
1 2
2 3
3 1
```

**Output**

```
0
```

### Example 2

**Input**

```
2 5
1 1
1 2
1 2
2 2
1 2
```

**Output**

```
1
```

## Note

In the first test case the described walk is possible without building extra trails. For example, first go on the first trail, then on the second one, and finally on the third one.

In the second test case the described walk is impossible without adding extra trails. To make the walk possible, it is enough to add one trail, for example, between glades number one and two.


### ideas
1. 需要一个欧拉回路，条件是，图是联通的，每个节点的deg是偶数，或者只有起点和另外一个点的deg是奇数
2. 不用考虑self loop（这个始终可以在需要的时候，被访问到）
3. 考虑fix两种情况（都是偶数，或者起点和某个点的deg是奇数的情况）
4. 然后查看需要添加多少条边（还需要考虑图不联通的情况）
5. 主要是图不联通的情况，
6. 假设有x个部分，那么需要使用x条边去联通他们（正常是x-1）但是因为要回到起点，所以，需要x条边（x>1)
7. 这样子，如果这块里面，有c个奇数个deg的，那么可以减去1，（如果c = 0, 那么 c = 1)因为必须添加一条边
8. 剩余奇数deg的节点数量 = sum of (c), 两两配对
9. 因为加一条边 = 增加deg 2（整体的效果）所以如果一开始 deg是奇数的，似乎就无解？
10. 但是答案没说无解的情况，所以一定有解？是的，还是因为一条边增加2，所以肯定是偶数
11. 第二种情况无法返回1
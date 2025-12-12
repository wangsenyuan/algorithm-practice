# Problem C

Vladik often travels by trains. He remembered some of his trips especially well and I would like to tell you about one of these trips:

Vladik is at initial train station, and now n people (including Vladik) want to get on the train. They are already lined up in some order, and for each of them the city code ai is known (the code of the city in which they are going to).

Train chief selects some number of disjoint segments of the original sequence of people (covering entire sequence by segments is not necessary). People who are in the same segment will be in the same train carriage. The segments are selected in such way that if at least one person travels to the city x, then all people who are going to city x should be in the same railway carriage. This means that they can't belong to different segments. Note, that all people who travel to the city x, either go to it and in the same railway carriage, or do not go anywhere at all.

Comfort of a train trip with people on segment from position l to position r is equal to XOR of all distinct codes of cities for people on the segment from position l to position r. XOR operation also known as exclusive OR.

Total comfort of a train trip is equal to sum of comfort for each segment.

Help Vladik to know maximal possible total comfort.

## Input

First line contains single integer n (1 ≤ n ≤ 5000) — number of people.

Second line contains n space-separated integers a1, a2, ..., an (0 ≤ ai ≤ 5000), where ai denotes code of the city to which i-th person is going.

## Output

The output should contain a single integer — maximal possible total comfort.

## Examples

### Example 1

**Input:**
```
6
4 4 2 5 2 3
```

**Output:**
```
14
```

### Example 2

**Input:**
```
9
5 1 3 1 5 2 4 2 5
```

**Output:**
```
9
```

## Note

In the first test case best partition into segments is: [4, 4] [2, 5, 2] [3], answer is calculated as follows: 4 + (2 xor 5) + 3 = 4 + 7 + 3 = 14

In the second test case best partition into segments is: 5 1 [3] 1 5 [2, 4, 2] 5, answer calculated as follows: 3 + (2 xor 4) = 3 + 6 = 9.


### ideas
1. 假设要去x城市的人，要么他们都在一节车上，要么都不出发
2. 假设处理当前的人i, a[i], 如果a[i]第一次出现，那么有两种选择，放弃它（同时后续的a[?] = a[i]都需要放弃）
3. 或者把它加入到已有的车厢中, dp[j] + fp[j][i-1] ^ a[i]
4. 或者在x = a[i]最后一次出现的时候，进行处理
5. dp[j-1] + fp[j][i-1] ^ a[i] (在最后一次出现时，再考虑a[i]是否加入进去)
6. fp[l][r]表示，这个区间里面不同code最大的xor值
7. 所以j可以用所有在i之前，第一次某个数出现的位置，且它们最后一次出现的位置，在i的前面
8. 这个区间内的值，要用到线性基来处理吗？
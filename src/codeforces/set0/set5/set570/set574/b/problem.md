## Three Musketeers

Do you know a story about the three musketeers? Anyway, you will learn about its origins now.

Richelimakieu is a cardinal in the city of Bearis. He is tired of dealing with crime by himself. He needs three brave warriors to help him to fight against bad guys.

There are $n$ warriors. Richelimakieu wants to choose three of them to become musketeers but it's not that easy. The most important condition is that musketeers must know each other to cooperate efficiently. And they shouldn't be too well known because they could be betrayed by old friends. For each musketeer his recognition is the number of warriors he knows, excluding other two musketeers.

Help Richelimakieu! Find if it is possible to choose three musketeers knowing each other, and what is minimum possible sum of their recognitions.

### Input

The first line contains two space-separated integers, $n$ and $m$ ($3 \le n \le 4000$, $0 \le m \le 4000$) — respectively number of warriors and number of pairs of warriors knowing each other.

$i$-th of the following $m$ lines contains two space-separated integers $a_i$ and $b_i$ ($1 \le a_i, b_i \le n$, $a_i \ne b_i$). Warriors $a_i$ and $b_i$ know each other. Each pair of warriors will be listed at most once.

### Output

If Richelimakieu can choose three musketeers, print the minimum possible sum of their recognitions. Otherwise, print `-1` (without the quotes).

### Examples

#### Example 1

**Input:**
```
5 6
1 2
1 3
2 3
2 4
3 4
4 5
```

**Output:**
```
2
```

#### Example 2

**Input:**
```
7 4
2 1
3 6
5 1
1 7
```

**Output:**
```
-1
```

### Note

In the first sample Richelimakieu should choose a triple 1, 2, 3. The first musketeer doesn't know anyone except other two musketeers so his recognition is 0. The second musketeer has recognition 1 because he knows warrior number 4. The third musketeer also has recognition 1 because he knows warrior 4. Sum of recognitions is $0 + 1 + 1 = 2$.

The other possible triple is 2, 3, 4 but it has greater sum of recognitions, equal to $1 + 1 + 1 = 3$.

In the second sample there is no triple of warriors knowing each other.

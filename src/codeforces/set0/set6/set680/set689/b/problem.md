Recently, Mike was very busy with studying for exams and contests. Now he is going to chill a bit by doing some sight seeing in the city.

City consists of $n$ intersections numbered from $1$ to $n$. Mike starts walking from his house located at the intersection number $1$ and goes along some sequence of intersections. Walking from intersection number $i$ to intersection $j$ requires $|i - j|$ units of energy. The total energy spent by Mike to visit a sequence of intersections $p_1 = 1, p_2, \ldots, p_k$ is equal to $\sum_{i=1}^{k-1}|p_i - p_{i+1}|$ units of energy.

Of course, walking would be boring if there were no shortcuts. A shortcut is a special path that allows Mike walking from one intersection to another requiring only $1$ unit of energy. There are exactly $n$ shortcuts in Mike's city, the $i$-th of them allows walking from intersection $i$ to intersection $a_i$ ($i \le a_i \le a_i + 1$) (but not in the opposite direction), thus there is exactly one shortcut starting at each intersection. Formally, if Mike chooses a sequence $p_1 = 1, p_2, \ldots, p_k$ then for each $1 \le i < k$ satisfying $p_{i + 1} = a_{p_i}$ and $a_{p_i} \ne p_i$ Mike will spend only $1$ unit of energy instead of $|p_i - p_{i + 1}|$ walking from the intersection $p_i$ to intersection $p_{i + 1}$. For example, if Mike chooses a sequence $p_1 = 1, p_2 = a_{p_1}, p_3 = a_{p_2}, \ldots, p_k = a_{p_{k - 1}}$, he spends exactly $k - 1$ units of total energy walking around them.

Before going on his adventure, Mike asks you to find the minimum amount of energy required to reach each of the intersections from his home. Formally, for each $1 \le i \le n$ Mike is interested in finding minimum possible total energy of some sequence $p_1 = 1, p_2, \ldots, p_k = i$.

## Input

The first line contains an integer $n$ ($1 \le n \le 200000$) — the number of Mike's city intersection.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($i \le a_i \le n$), describing shortcuts of Mike's city, allowing to walk from intersection $i$ to intersection $a_i$ using only $1$ unit of energy. Please note that the shortcuts don't allow walking in opposite directions (from $a_i$ to $i$).

## Output

In the only line print $n$ integers $m_1, m_2, \ldots, m_n$, where $m_i$ denotes the least amount of total energy required to walk from intersection $1$ to intersection $i$.

## Examples

### Input
```
3
2 2 3
```

### Output
```
0 1 2
```

### Input
```
5
1 2 3 4 5
```

### Output
```
0 1 2 3 4
```

### Input
```
7
4 4 4 4 7 7 7
```

### Output
```
0 1 2 1 2 3 3
```

## Note

In the first sample case desired sequences are:

1: $1$; $m_1 = 0$;

2: $1, 2$; $m_2 = 1$;

3: $1, 3$; $m_3 = |3 - 1| = 2$.

In the second sample case the sequence for any intersection $1 < i$ is always $1, i$ and $m_i = |1 - i|$.

In the third sample case — consider the following intersection sequences:

1: $1$; $m_1 = 0$;

2: $1, 2$; $m_2 = |2 - 1| = 1$;

3: $1, 4, 3$; $m_3 = 1 + |4 - 3| = 2$;

4: $1, 4$; $m_4 = 1$;

5: $1, 4, 5$; $m_5 = 1 + |4 - 5| = 2$;

6: $1, 4, 6$; $m_6 = 1 + |4 - 6| = 3$;

7: $1, 4, 5, 7$; $m_7 = 1 + |4 - 5| + 1 = 3$.

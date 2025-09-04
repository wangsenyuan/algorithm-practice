# Problem Description

A boy named Gena really wants to get to the "Russian Code Cup" finals, or at least get a t-shirt. But the offered problems are too complex, so he made an arrangement with his n friends that they will solve the problems for him.

The participants are offered m problems on the contest. For each friend, Gena knows what problems he can solve. But Gena's friends won't agree to help Gena for nothing: the i-th friend asks Gena xi rubles for his help in solving all the problems he can. Also, the friend agreed to write a code for Gena only if Gena's computer is connected to at least ki monitors, each monitor costs b rubles.

Gena is careful with money, so he wants to spend as little money as possible to solve all the problems. Help Gena, tell him how to spend the smallest possible amount of money. Initially, there's no monitors connected to Gena's computer.

## Input

The first line contains three integers n, m and b (1 ≤ n ≤ 100; 1 ≤ m ≤ 20; 1 ≤ b ≤ 10^9) — the number of Gena's friends, the number of problems and the cost of a single monitor.

The following 2n lines describe the friends. Lines number 2i and (2i + 1) contain the information about the i-th friend. The 2i-th line contains three integers xi, ki and mi (1 ≤ xi ≤ 10^9; 1 ≤ ki ≤ 10^9; 1 ≤ mi ≤ m) — the desired amount of money, monitors and the number of problems the friend can solve. The (2i + 1)-th line contains mi distinct positive integers — the numbers of problems that the i-th friend can solve. The problems are numbered from 1 to m.

## Output

Print the minimum amount of money Gena needs to spend to solve all the problems. Or print -1, if this cannot be achieved.

## Examples

### Example 1

**Input:**
```
2 2 1
100 1 1
2
100 2 1
1
```

**Output:**
```
202
```

### Example 2

**Input:**
```
3 2 5
100 1 1
1
100 1 1
2
200 1 2
1 2
```

**Output:**
```
205
```

### Example 3

**Input:**
```
1 2 1
1 1 1
1
```

**Output:**
```
-1
```


### ideas
1. 如果第i个朋友参与，那么他解决的问题 + 已经解决的问题，会得到一个新的状态
2. 但是这里，有个问题，就是假设，解决了3个问题，花费了w，那么解决其中2个问题，花费不会超过w，这是因为，我可i多解决1个问题的情况下，也是ok的
3. 还有个条件，假设mask表示的题目，使用了x个monitor，目前新的朋友需要y个monitor, 那么就需要max(x, y)个monitor
4. 所以，要按照k倒序排列，如果已经有朋友解决了问题，当前就不需要额外花费monitore的钱，
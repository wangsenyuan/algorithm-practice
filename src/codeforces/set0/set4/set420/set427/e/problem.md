Imagine that your city is an infinite 2D plane with Cartesian coordinate system. The only crime-affected road of your city is the x-axis. Currently, there are n criminals along the road. No police station has been built on this road yet, so the mayor wants to build one.

As you are going to be in charge of this new police station, the mayor has asked you to choose a suitable position (some integer point) for building it. You should choose the best position for the police station, so that you could minimize the total time of your criminal catching mission. Your mission of catching the criminals will operate only from this station.

The new station will have only one patrol car. You will go to the criminals by this car, carry them on the car, bring them back to the police station and put them in prison. The patrol car can carry at most m criminals at a time. Note that, the criminals don't know about your mission. So, they will stay where they are instead of running away.

Your task is to find the position for the police station, so that total distance you need to cover to catch all the criminals will be minimum possible. Note that, you also can built the police station on the positions where one or more criminals already exist. In such a case all these criminals are arrested instantly.

## Input

The first line of the input will have two integers n (1 ≤ n ≤ 10^6) and m (1 ≤ m ≤ 10^6) separated by spaces. The next line will contain n integers separated by spaces. The ith integer is the position of the ith criminal on the x-axis. Absolute value of positions will not exceed 10^9. If a criminal has position x, he/she is located in the point (x, 0) of the plane.

The positions of the criminals will be given in non-decreasing order. Note, that there can be more than one criminal standing at some point of the plane.

**Note:** since the size of the input/output could be very large, don't use slow input/output techniques in your language. For example, do not use input/output streams (cin, cout) in C++.

## Output

Print a single integer, that means the minimum possible distance you need to cover to catch all the criminals.

## Examples

**Input:**
```
3 6
1 2 3
```

**Output:**
```
4
```

**Input:**
```
5 5
-7 -6 -3 -1 1
```

**Output:**
```
16
```

**Input:**
```
1 369
0
```

**Output:**
```
0
```

**Input:**
```
11 2
-375 -108 1336 1453 1598 1892 2804 3732 4291 4588 4822
```

**Output:**
```
18716
```


### ideas
1. 显然警局应该建在某个具体的犯罪地点。 
2. 考虑警局在位置x，它应该从远处开始处理，还是应该从近处开始处理呢？因为它肯定要到最远处一次（至少一次）如果从近处开始处理
3. 那么某一次去最远处时，能够装的罪犯就变少了，从而不得不多跑一趟。反过来却不会（即使近处多跑一趟，也不是很糟糕）
4. 那么先处理最远处的，然后处理近处
5. 此外，这个是关于位置x的一次函数（所以符合单调性）且应该是中间更快，两头更慢。可以用3路分支查询
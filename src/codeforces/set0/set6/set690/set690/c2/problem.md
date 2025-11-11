# Zombie's Brain Latency

## Problem Description

Further research on zombie thought processes yielded interesting results. As we know from the previous problem, the nervous system of a zombie consists of $n$ brains and $m$ brain connectors joining some pairs of brains together. It was observed that the intellectual abilities of a zombie depend mainly on the topology of its nervous system. 

More precisely, we define the **distance** between two brains $u$ and $v$ ($1 \leq u, v \leq n$) as the minimum number of brain connectors used when transmitting a thought between these two brains. The **brain latency** of a zombie is defined to be the maximum distance between any two of its brains. Researchers conjecture that the brain latency is the crucial parameter which determines how smart a given zombie is. Help them test this conjecture by writing a program to compute brain latencies of nervous systems.

In this problem you may assume that any nervous system given in the input is valid, i.e., it satisfies conditions (1) and (2) from the easy version.

## Input

The first line of the input contains two space-separated integers $n$ and $m$ ($1 \leq n, m \leq 100000$) denoting the number of brains (which are conveniently numbered from 1 to $n$) and the number of brain connectors in the nervous system, respectively. 

In the next $m$ lines, descriptions of brain connectors follow. Every connector is given as a pair of brains $a$ $b$ it connects ($1 \leq a, b \leq n$ and $a \neq b$).

## Output

Print one number – the brain latency.

## Examples

### Example 1

**Input:**
```
4 3
1 2
1 3
1 4
```

**Output:**
```
2
```

### Example 2

**Input:**
```
5 4
1 2
2 3
3 4
3 5
```

**Output:**
```
3
```

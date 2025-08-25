# Problem I: Noise Pollution in Berland's Capital

## Problem Description

The Berland's capital has the form of a rectangle with sizes $n \times m$ quarters. All quarters are divided into three types:

- **Regular quarters** (labeled with the character `.`) — such quarters do not produce noise but are not obstacles to the propagation of noise
- **Sources of noise** (labeled with an uppercase Latin letter from 'A' to 'Z') — such quarters are noise sources and are not obstacles to the propagation of noise  
- **Heavily built-up quarters** (labeled with the character `*`) — such quarters are soundproofed, the noise does not penetrate into them and they themselves are obstacles to the propagation of noise

## Noise Production Rules

A quarter labeled with letter 'A' produces $q$ units of noise. A quarter labeled with letter 'B' produces $2 \cdot q$ units of noise. And so on, up to a quarter labeled with letter 'Z', which produces $26 \cdot q$ units of noise. There can be any number of quarters labeled with each letter in the city.

## Noise Propagation Rules

When propagating from the source of noise, the noise level is halved when moving from one quarter to a quarter that shares a side with it (when an odd number is to be halved, it's rounded down). The noise spreads along the chain. 

For example, if some quarter is located at a distance 2 from the noise source, then the value of noise which will reach the quarter is divided by 4. So the noise level that comes from the source to the quarter is determined solely by the length of the shortest path between them. Heavily built-up quarters are obstacles, the noise does not penetrate into them.

The values in the cells of the table on the right show the total noise level in the respective quarters for $q = 100$, the first term in each sum is the noise from the quarter 'A', the second — the noise from the quarter 'B'.

The noise level in a quarter is defined as the sum of the noise from all sources. To assess the quality of life of the population of the capital of Berland, it is required to find the number of quarters whose noise level exceeds the allowed level $p$.

## Input

The first line contains four integers $n$, $m$, $q$ and $p$ ($1 \leq n, m \leq 250$, $1 \leq q, p \leq 10^6$) — the sizes of Berland's capital, the number of noise units that a quarter 'A' produces, and the allowable noise level.

Each of the following $n$ lines contains $m$ characters — the description of the capital quarters, in the format that was described in the statement above. It is possible that in the Berland's capital there are no quarters of any type.

## Output

Print the number of quarters in which the noise level exceeds the allowed level $p$.

## Examples

### Example 1

**Input:**
```
3 3 100 140
...
A*.
.B.
```

**Output:**
```
3
```

### Example 2

**Input:**
```
3 3 2 8
B*.
BB*
BBB
```

**Output:**
```
4
```

### Example 3

**Input:**
```
3 4 5 4
..*B
..**
D...
```

**Output:**
```
7
```

## Note

The illustration to the first example is in the main part of the statement.

## ideas
1. 超过30距离的source，不需要处理
2. 每个噪音源，假设只访问30 * 30范围内，或者减少到q的格子
3. n * m * 20 * 20 的距离
4. 250 * 250 * 20 * 20 = 似乎也ok
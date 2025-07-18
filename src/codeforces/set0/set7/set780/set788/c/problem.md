# Sasha and Kolya's Coke Mixing Problem

Sasha and Kolya decided to get drunk with Coke, again. This time they have $k$ types of Coke. The $i$-th type is characterized by its carbon dioxide concentration $a_i$. Today, on the party in honour of Sergiy of Vancouver, they decided to prepare a glass of Coke with carbon dioxide concentration $n$. The drink should also be tasty, so the glass can contain only integer number of liters of each Coke type (some types can be not presented in the glass). Also, they want to minimize the total volume of Coke in the glass.

Carbon dioxide concentration is defined as the volume of carbon dioxide in the Coke divided by the total volume of Coke. When you mix two Cokes, the volume of carbon dioxide sums up, and the total volume of Coke sums up as well.

Help them find the minimal natural number of liters needed to create a glass with carbon dioxide concentration $n$. Assume that the friends have unlimited amount of each Coke type.

## Input

The first line contains two integers $n, k$ ($0 \leq n \leq 1000$, $1 \leq k \leq 10^6$) — carbon dioxide concentration the friends want and the number of Coke types.

The second line contains $k$ integers $a_1, a_2, \ldots, a_k$ ($0 \leq a_i \leq 1000$) — carbon dioxide concentration of each type of Coke. Some Coke types can have same concentration.

## Output

Print the minimal natural number of liters needed to prepare a glass with carbon dioxide concentration $n$, or $-1$ if it is impossible.

## Examples

### Example 1

**Input:**
```
400 4
100 300 450 500
```

**Output:**
```
2
```

**Explanation:** We can achieve concentration $400$ using one liter of Coke of types $100$ and $700$: $\frac{100 + 700}{2} = 400$.

### Example 2

**Input:**
```
50 2
100 25
```

**Output:**
```
3
```

**Explanation:** We can achieve concentration $50$ using two liters of type $100$ and one liter of type $25$: $\frac{2 \times 100 + 1 \times 25}{3} = \frac{225}{3} = 75$. Wait, this doesn't give us $50$. Let me correct this: using two liters of type $25$ and one liter of type $100$: $\frac{2 \times 25 + 1 \times 100}{3} = \frac{150}{3} = 50$.


## ideas
1. 从 n / 1000开始吗？
# Problem E

## Description

The Kingdom of Kremland is a tree (a connected undirected graph without cycles) consisting of **n** vertices. Each vertex **i** has its own value **aᵢ**. All vertices are connected in series by edges. Formally, for every **1 ≤ i < n** there is an edge between the vertices of **i** and **i+1**.

Denote the function **f(l, r)**, which takes two integers **l** and **r** (**l ≤ r**):

1. We leave in the tree only vertices whose values range from **l** to **r**.
2. The value of the function will be the number of connected components in the new graph.

Your task is to calculate the following sum:

**∑ₗ₌₁ⁿ ∑ᵣ₌ₗⁿ f(l, r)**

## Input

The first line contains a single integer **n** (**1 ≤ n ≤ 10⁵**) — the number of vertices in the tree.

The second line contains **n** integers **a₁, a₂, …, aₙ** (**1 ≤ aᵢ ≤ n**) — the values of the vertices.

## Output

Print one number — the answer to the problem.

## Solution

First of all, assign **0** to **a₀**.

### How to find the value of f(l,r) in O(n)?

For each **i** from **0** to **n**, set **bᵢ** to **1** if **l ≤ aᵢ ≤ r**, otherwise set it to **0**. Now we can see that the value of **f(l,r)** is equal to the number of adjacent pairs **(0,1)** in array **b**.

So now we can find the answer in **O(n)** using the technique of contribution to the sum. For every adjacent pair of elements in array **a** (including the zero-indexed element), we must find its contribution to the overall answer. Considering the thoughts above about **f(l,r)**, we must find the number of pairs **(l,r)** such that **aᵢ** is in the range **[l,r]** and **aᵢ₋₁** is not in the range **[l,r]**. There are two cases:

#### Case 1: aᵢ > aᵢ₋₁
Then **l** must be in the range from **aᵢ₋₁ + 1** to **aᵢ** and **r** must be in the range from **aᵢ** to **n**. The contribution is:

```
(aᵢ - aᵢ₋₁) × (n - aᵢ + 1)
```

#### Case 2: aᵢ < aᵢ₋₁
Then **l** must be in the range from **1** to **aᵢ** and **r** must be in the range from **aᵢ** to **aᵢ₋₁ - 1**. The contribution is:

```
aᵢ × (aᵢ₋₁ - aᵢ)
```

Sum up the contributions of all adjacent pairs to find the answer.

**Complexity:** O(n)
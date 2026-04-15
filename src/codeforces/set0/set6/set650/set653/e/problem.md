# E — Bear and Forgotten Tree 2

## Problem description

A **tree** is a connected undirected graph with **n** vertices and **n − 1** edges. Vertices are numbered **1 … n**.

Limak once had a tree on **n** vertices but lost it. He still remembers:

- **m** pairs **(a₁, b₁), …, (aₘ, bₘ)** such that for each **i**, the tree had **no** edge between **aᵢ** and **bᵢ** (these pairs are **forbidden**).
- Vertex **1** had degree exactly **k** (exactly **k** incident edges).

Decide whether **some** tree exists that satisfies **all** of this. Print whether it is **possible** or **impossible**.

## Input

- The first line contains three integers **n**, **m**, **k** (1 ≤ n ≤ 3·10⁵, 0 ≤ m ≤ 3·10⁵, 1 ≤ k < n) — number of vertices, number of forbidden pairs, and required degree of vertex **1**.
- Each of the next **m** lines contains two distinct integers **aᵢ**, **bᵢ** (1 ≤ aᵢ, bᵢ ≤ n, aᵢ ≠ bᵢ) — a forbidden edge. Each unordered pair appears **at most once**.

## Output

Print **`possible`** if at least one tree satisfies the conditions; otherwise print **`impossible`** (lowercase, no quotes).

## Examples

### Example 1

**Input:**

```
5 4 2
1 2
2 3
4 2
4 1
```

**Output:**

```
possible
```

### Example 2

**Input:**

```
6 5 3
1 2
1 3
1 4
1 5
1 6
```

**Output:**

```
impossible
```

## Note

In the first example, **n = 5** and vertex **1** must have degree **k = 2**. One valid tree has edges **1–5**, **5–2**, **1–3**, and **3–4** (among others).

In the second example, Limak forbids every edge **1–v** for **v = 2, …, 6**, so vertex **1** cannot connect to any other vertex — no tree is possible.


### ideas
1. no ideas at all
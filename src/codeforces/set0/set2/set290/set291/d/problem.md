# D. Parallel program model

## Setup

Polycarpus has a computer with `n` **processors** and `n` **memory cells**. Processors and cells are numbered `1, 2, …, n`.

He wants a **parallel program** such that, after it finishes, cell `i` holds the value **`n − i`** (the “distance” from cell `i` to cell `n` in index space).

Let **`a[i]`** be the value in cell `i`.

- **Initially:** `a[i] = 1` for `1 ≤ i < n`, and `a[n] = 0`.
- **Who may write:** only **processor `i`** may write to **cell `i`**.
- **Reading:** every processor may read any cell; many processors may read the same cell at once.

## One step (parallel increment)

The program runs in **several steps**. In **each step** you run one **parallel increment**:

1. **Choose reads.** Each processor `i` independently picks a cell index **`c[i]`** with `1 ≤ c[i] ≤ n` (the cell whose value it will add).
2. **Add in parallel.** All processors apply, **simultaneously**:

   `a[i] ← a[i] + a[c[i]]`

## Task

You are given **`n`** and **`k`**. Construct a valid program that finishes in **exactly `k` steps**, so that after those `k` steps,

`a[i] = n − i` for every `i`.

Output the choices **`c[1], …, c[n]`** for each of the `k` steps.

## Input

The first line contains two integers **`n`** and **`k`**:

- `1 ≤ n ≤ 10^4`
- `1 ≤ k ≤ 20`

It is **guaranteed** that a valid sequence of operations exists for the given `n` and `k`.

## Output

Print **`n · k`** integers in **`k`** lines.

- Line `t` (`1 ≤ t ≤ k`): `n` integers **`c[1], c[2], …, c[n]`** for step `t`, with `1 ≤ c[i] ≤ n`.

After performing the printed operations, **`a[i] = n − i`** must hold for all `i`.

## Examples

### Example 1

**Input**

```
1 1
```

**Output**

```
1
```

### Example 2

**Input**

```
3 2
```

**Output**

```
2 3 3
3 3 3
```

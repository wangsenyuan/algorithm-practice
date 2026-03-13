# Problem

Demonstrative competitions will be held in the run-up to the 20NN Berlatov Olympic Games. Today is the day for the running competition!

Berlatov team has **2n** runners placed on **two** tracks: **n** runners on each track. Runners on each track are numbered from **1** to **n**. Runner **i** runs the entire track in **i** seconds.

The competition works as follows:

- First, the first runners on both tracks start at the same time.
- When the **slower** of them finishes, the second runners on both tracks start, and everyone waits until the slower of that pair finishes.
- This repeats until all **n** pairs have run.

The organizers want the run to be as long as possible, but if it lasts more than **k** seconds, the crowd gets bored. As the coach, you may choose **any order** of runners on each track (you cannot change how many runners are on each track or swap runners between tracks).

**Task:** Choose the order on each track so that the **duration is as long as possible** but **does not exceed k** seconds.

Formally: find two permutations **p** and **q** (each of length n) such that

\[
\mathrm{sum} = \sum_{i=1}^{n} \max(p_i, q_i)
\]

is **maximum** subject to \(\mathrm{sum} \le k\). If no such pair exists, report it.

---

## Input

- One line with two integers **n** and **k**  
  (\(1 \le n \le 10^6\), \(1 \le k \le n^2\)) — number of runners per track and maximum allowed duration (seconds).

---

## Output

- If it is **impossible** to reorder so that the duration does not exceed **k**, print **-1**.

- Otherwise print **three lines**:
  1. One integer **sum** — the maximum possible duration not exceeding **k**.
  2. A permutation **p₁, p₂, …, pₙ** (\(1 \le p_i \le n\), all distinct) — runner order on the first track.
  3. A permutation **q₁, q₂, …, qₙ** (\(1 \le q_i \le n\), all distinct) — runner order on the second track.

  The value \(\mathrm{sum} = \sum_{i=1}^{n} \max(p_i, q_i)\) must be maximum possible and \(\le k\). If multiple answers exist, print any.

---

## Examples

### Example 1

Input:

```text
5 20
```

Output:

```text
20
1 2 3 4 5
5 2 4 3 1
```

### Example 2

Input:

```text
3 9
```

Output:

```text
8
1 2 3
3 2 1
```

### Example 3

Input:

```text
10 54
```

Output:

```text
-1
```

---

## Note

- **Example 1:** Order on first track [5, 3, 2, 1, 4], on second [1, 4, 2, 5, 3]. Duration = max(5,1)+max(3,4)+max(2,2)+max(1,5)+max(4,3) = 5+4+2+5+4 = 20.
- **Example 2:** Order on first track [2, 3, 1], on second [2, 1, 3]. Duration = 8, which is the maximum possible for n = 3.

---

## Short summary

- **Given:** n (runners per track), k (max duration). Runner i takes i seconds.
- **Rule:** In each round, both tracks run one runner; round ends when the slower of the two finishes. Total duration = sum over rounds of max(pᵢ, qᵢ).
- **Goal:** Choose permutations p, q of [1..n] to **maximize** sum = Σ max(pᵢ, qᵢ) subject to sum ≤ k.
- **Output:** If impossible → -1; else maximum sum and two permutations achieving it.

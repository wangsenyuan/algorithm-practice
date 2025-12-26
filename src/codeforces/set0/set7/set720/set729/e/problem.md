# Problem E

There are n workers in a company, each of them has a unique id from 1 to n. Exactly one of them is a chief, his id is s. Each worker except the chief has exactly one immediate superior.

There was a request to each of the workers to tell how many superiors (not only immediate). Worker's superiors are his immediate superior, the immediate superior of his immediate superior, and so on. For example, if there are three workers in the company, from which the first is the chief, the second worker's immediate superior is the first, the third worker's immediate superior is the second, then the third worker has two superiors, one of them is immediate and one not immediate. The chief is a superior to all the workers except himself.

Some of the workers were in a hurry and made a mistake. You are to find the minimum number of workers that could make a mistake.

## Input

The first line contains two positive integers n and s (1 ≤ n ≤ 2·10⁵, 1 ≤ s ≤ n) — the number of workers and the id of the chief.

The second line contains n integers a₁, a₂, ..., aₙ (0 ≤ aᵢ ≤ n - 1), where aᵢ is the number of superiors (not only immediate) the worker with id i reported about.

## Output

Print the minimum number of workers that could make a mistake.

## Examples

### Example 1

**Input:**

```text
3 2
2 0 2
```

**Output:**

```text
1
```

### Example 2

**Input:**

```text
5 3
1 0 0 4 1
```

**Output:**

```text
2
```

## Note

In the first example it is possible that only the first worker made a mistake. Then:

- the immediate superior of the first worker is the second worker,
- the immediate superior of the third worker is the first worker,
- the second worker is the chief.


### ideas
1. chef 的答案 如果不等于0, +1
2. chef是root，现在把所有a[i] = 1 的找出来；如果没有 +1， 但是选择谁呢？pending， 消耗一个人
3. 现在选择a[i] = 2 的人, 如果没有, pending (但是消耗一个人) 
4. 。。until left > pending 
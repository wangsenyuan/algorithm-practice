# P9588 「MXOI Round 2」队列

[Problem link](https://www.luogu.com.cn/problem/P9588)

time limit per test: 1.00s

memory limit per test: 512.00MB

Maintain a queue through `q` operations:

| Operation | Meaning |
| --- | --- |
| `1 x` | Append `1, 2, 3, …, x` to the back |
| `2 y` | Pop the first `y` elements |
| `3 z` | Query the `z`-th element, counting from the front |
| `4` | Query the maximum element currently in the queue |

## Input

The first line contains `c` and `q`. `c` is the test-point id. `c = 0` means this file is a sample.

Each of the next `q` lines is one operation, as above.

## Output

For every operation `3` and `4`, print the answer on its own line.

## Sample

### Input

```text
0 9
1 5
1 3
2 2
1 4
3 6
3 8
2 4
4
3 3
```

### Output

```text
3
2
4
1
```

After the fourth operation the queue is `3, 4, 5, 1, 2, 3, 1, 2, 3, 4`.

After the seventh operation the queue is `2, 3, 1, 2, 3, 4`.

## Constraints

Let `sum x` be the total of every `x` in operation `1`.

- `1 <= q <= 2 * 10^5`
- `1 <= x, y, z <= 10^9`
- `0 <= sum x <= 2 * 10^14`
- operation `2` never pops more elements than the queue holds
- operation `3` never asks past the end of the queue
- operation `4` is only asked when the queue is non-empty

Some test points omit operation `2`, `3`, or `4`.

## Status

I/O and the official sample are in place. `solve` is left as a TODO.

# F. Unjust Binary Life

## Key observation

Cell `(i, j)` is `0` exactly when

`a_i xor b_j = 0`, i.e. `a_i = b_j`.

So every zero cell on a valid path says: the row variable `a_i` and the column variable `b_j` must have the same final bit.

Now fix one target `(x, y)`. Any monotone path from `(1, 1)` to `(x, y)` must:

- visit every row `1..x` at least once, because it starts at row `1` and ends at row `x`;
- visit every column `1..y` at least once, because it starts at column `1` and ends at column `y`.

The cells on the path form a connected chain. If every cell on this chain is zero, every touched `a_i` and `b_j` is connected by equality constraints. Therefore all variables in

`a_1..a_x` and `b_1..b_y`

must become one common bit.

This condition is also sufficient: if all of these prefix bits are equal, then every cell inside the prefix rectangle is zero, so any monotone path works.

Therefore the shape of the path does not matter. For target `(x, y)`, we only need to decide whether to make all bits in the two prefixes become `0` or become `1`.

## Cost for one target

Let:

- `onesA[x]` be the number of `1`s in `a[1..x]`;
- `zeroesA[x]` be the number of `0`s in `a[1..x]`;
- `onesB[y]` and `zeroesB[y]` be defined similarly.

If the common bit is `0`, every original `1` in the two prefixes must be flipped, so the cost is

`onesA[x] + onesB[y]`.

If the common bit is `1`, every original `0` in the two prefixes must be flipped, so the cost is

`zeroesA[x] + zeroesB[y]`.

Thus

`f(x, y) = min(onesA[x] + onesB[y], zeroesA[x] + zeroesB[y])`.

The answer is the sum of this value over all `n^2` prefix pairs.

## Turning the min into a sorted split

For every prefix define

`diff = ones - zeroes`.

For a pair `(x, y)`, choosing final bit `0` is no worse than choosing final bit `1` when

`onesA[x] + onesB[y] <= zeroesA[x] + zeroesB[y]`.

Move terms around:

`(onesA[x] - zeroesA[x]) + (onesB[y] - zeroesB[y]) <= 0`.

That is:

`diffA[x] + diffB[y] <= 0`.

So for a fixed `x`, among all `b` prefixes:

- if `diffB[y] <= -diffA[x]`, contribution is `onesA[x] + onesB[y]`;
- otherwise, contribution is `zeroesA[x] + zeroesB[y]`.

This is why the code sorts all `b` prefixes by `diff`.

## Matching the code

The `item` struct stores one prefix:

```go
type item struct {
	diff  int
	ones  int
	zeros int
}
```

For prefix ending at zero-based index `i`, its length is `i + 1`, so:

```go
zeros := i + 1 - ones
diff := ones - zeros
```

The code writes the same `diff` as:

```go
2*ones - i - 1
```

because `ones - zeros = ones - (i + 1 - ones) = 2*ones - i - 1`.

After sorting all `b` prefixes by `diff`, the code builds prefix sums:

```go
sumOnes[i+1] = sumOnes[i] + int64(cur.ones)
sumZeros[i+1] = sumZeros[i] + int64(cur.zeros)
```

For one `a` prefix `cur`, it finds the first `b` prefix that should use the `zeroes` cost:

```go
cnt := sort.Search(n, func(i int) bool {
	return y[i].diff > -cur.diff
})
```

So indices `[0, cnt)` satisfy:

`y[i].diff <= -cur.diff`.

These pairs use the common final bit `0`, and each contributes:

`cur.ones + y[i].ones`.

Their total contribution is:

```go
int64(cnt*cur.ones) + sumOnes[cnt]
```

The remaining `n - cnt` prefixes use the common final bit `1`, and each contributes:

`cur.zeros + y[i].zeros`.

Their total contribution is:

```go
int64((n-cnt)*cur.zeros) + sumZeros[n] - sumZeros[cnt]
```

Adding these two parts for every `a` prefix gives the required sum.

## Complexity

For each test case:

- building prefix records is `O(n)`;
- sorting `b` prefixes is `O(n log n)`;
- each `a` prefix does one binary search, so `O(n log n)` total.

Overall complexity is `O(n log n)` per test case, with `O(n)` memory.

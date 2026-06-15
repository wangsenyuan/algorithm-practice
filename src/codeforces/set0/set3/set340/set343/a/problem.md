# A. Rational Resistance

[Problem link](https://codeforces.com/problemset/problem/343/A)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: stdin

output: stdout

Mad scientist Mike is building a time machine. To finish it, he needs a resistor with
a certain resistance value.

Mike has many identical unit resistors with resistance `R0 = 1`. Other resistance
values can be built from them. A valid element is:

1. one resistor;
2. one element and one resistor connected **in series**;
3. one element and one resistor connected **in parallel**.

For a series connection, if the existing element has resistance `Re`, the new
resistance is:

```text
R = Re + 1
```

For a parallel connection:

```text
R = Re / (Re + 1)
```

Mike needs an element with resistance exactly `a / b`. Find the minimum number of
unit resistors required to build such an element.

## Input

One line contains two space-separated integers `a` and `b` (`1 <= a, b <= 10^18`).
The fraction `a / b` is irreducible. A solution always exists.

In C++, do not use `%lld` for 64-bit integers; prefer `cin`/`cout` or `%I64d`.

## Output

Print one integer — the minimum number of resistors needed.

## Examples

### Input

```text
1 1
```

### Output

```text
1
```

### Input

```text
3 2
```

### Output

```text
3
```

### Input

```text
199 200
```

### Output

```text
200
```

## Note

In the first sample, one resistor is enough.

In the second sample, connect two resistors in parallel to get resistance `1/2`,
then connect that element in series with a third resistor to get `1/2 + 1 = 3/2`.
Two resistors are not enough.

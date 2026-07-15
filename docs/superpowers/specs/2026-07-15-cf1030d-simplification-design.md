# Codeforces 1030D solution simplification

## Goal

Simplify the accepted construction in
`src/codeforces/set1/set10/set103/set1030/d` while preserving its output
contract and the user's current in-progress changes.

## Algorithm

Use the standard axis-aligned triangle with vertices `(0, 0)`, `(a, 0)`, and
`(0, b)`. Its doubled area is `a*b`, so the required equation is:

```text
a * b = 2 * n * m / k
```

Start with `a = n` and `b = m`. Remove the factors of `k` from `a` and `b`
using two GCD operations. After those reductions, the remaining part of `k`
must divide `2`; otherwise no integer-coordinate triangle exists.

- If the remaining factor is `2`, the reduced `a` and `b` already give the
  required doubled area.
- If it is `1`, double either `a` or `b`, choosing a dimension that still fits
  inside the rectangle.

This replaces divisor enumeration with `O(log(max(n,m,k)))` arithmetic.

## Code changes

- Replace the divisor loop in `solve` with the two-GCD construction.
- Keep the existing `result`, `drive`, output format, and `gcd` helper.
- Preserve unrelated user changes and the package's existing coding style.

## Tests

- Keep the sample-based checks.
- Remove the accidental duplicate sample.
- Add an exhaustive small-range test that checks:
  - feasibility agrees with `(2*n*m) % k == 0`;
  - every returned point lies inside the rectangle;
  - the returned triangle's doubled area is exactly `2*n*m/k`.

## Documentation

Replace the unfinished `ideas` section in `problem.md` with the reduction,
construction cases, correctness proof, and complexity analysis matching the
final code.

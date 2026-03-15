### Problem

You are given two integers `l` and `r`.

**Definitions:**

- **Strictly increasing:** each element is strictly greater than the previous one (if any).
- **Strictly decreasing:** each element is strictly less than the previous one (if any).
- **Good number:** an integer whose digits form a **strictly monotonic** sequence (strictly increasing or strictly decreasing). All single-digit numbers are considered good.
- **Wonderful number:** an integer that is either a good number, or whose **digit sum** is a good number.

Return the number of wonderful numbers in the range `[l, r]` (inclusive).

### Input / Output

- **Input:** two integers `l` and `r`.
- **Output:** one integer — the count of wonderful numbers in `[l, r]`.

### Examples

**Example 1**

- **Input:** `l = 8`, `r = 10`
- **Output:** `3`
- **Explanation:** 8 and 9 are single-digit, so they are good and wonderful. 10 has digits `[1, 0]`, which is strictly decreasing, so 10 is good and wonderful. Answer is 3.

**Example 2**

- **Input:** `l = 12340`, `r = 12341`
- **Output:** `1`
- **Explanation:**
  - 12340: digits `[1, 2, 3, 4, 0]` are not strictly monotonic, so not good. Digit sum = 10; 10 is good (digits `[1, 0]` strictly decreasing). So 12340 is wonderful.
  - 12341: digits `[1, 2, 3, 4, 1]` are not strictly monotonic. Digit sum = 11; 11 is not good (digits `[1, 1]`). So 12341 is not wonderful. Answer is 1.

**Example 3**

- **Input:** `l = 123456788`, `r = 123456788`
- **Output:** `0`
- **Explanation:** 123456788 is not good (digits not strictly monotonic). Digit sum = 44; 44 is not good (digits `[4, 4]`). So 123456788 is not wonderful. Answer is 0.

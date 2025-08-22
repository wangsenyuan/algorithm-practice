Here is the formatted markdown content for your file. Please copy and paste it into `problem.md`:

---

# Unkillable Slug

Vasya plays **The Elder Trolls IV: Oblivon**. Oh, those creators of computer games! What they do not come up with! Absolutely unique monsters have been added to the game. One of these monsters is **Unkillable Slug**.

**Why is it "Unkillable"?**
- It can be killed with cutting weapon only, so lovers of two-handed amber hammers should find a suitable knife themselves.
- It is necessary to make *so many* cutting strokes to Unkillable Slug. Extremely many. Too many!

Vasya has already promoted his character to 80-th level and, in order to gain level 81, he was asked to kill Unkillable Slug. The monster has a very interesting shape. It looks like a rectangular parallelepiped with size $x \times y \times z$, consisting of undestructable cells $1 \times 1 \times 1$. At one stroke Vasya can cut the Slug along an imaginary grid, i.e. cut with a plane parallel to one of the parallelepiped sides. The monster dies when the amount of parts it is divided into reaches some critical value.

All parts of the monster do not fall after each cut; they remain exactly in their places. I.e. Vasya can cut several parts with one cut.

Vasya wants to know what is the **maximum number of pieces** he can cut the Unkillable Slug into, striking him at most $k$ times.

Vasya's character uses an absolutely thin sword with infinite length.

---

## Input

The first line of input contains four integer numbers $x, y, z, k$ ($1 \leq x, y, z \leq 10^6$, $0 \leq k \leq 10^9$).

## Output

Output the only number — the answer for the problem.

*Please, do not use `%lld` specificator to read or write 64-bit integers in C++. It is preferred to use `cout` (also you may use `%I64d`).*

---

## Examples

### Input
```
2 2 2 3
```
### Output
```
8
```

### Input
```
2 2 2 1
```
### Output
```
2
```

---

## Note

In the first sample, Vasya makes 3 pairwise perpendicular cuts. He cuts the monster into two parts with the first cut, then divides each part into two with the second cut, and finally divides each of the 4 parts into two.

---

Let me know if you want this saved to the file automatically once the technical issue is resolved.
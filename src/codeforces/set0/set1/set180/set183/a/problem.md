# A. Headquarters

[Problem link](https://codeforces.com/problemset/problem/183/A)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

Sensation in the two-dimensional kingdom! The police have caught a highly
dangerous outlaw, member of the notorious "Pihters" gang. The law department
states that the outlaw was driving from the gang's headquarters in his car when
he crashed into an ice cream stall. The stall, the car, and the headquarters
each occupies exactly one point on the two-dimensional kingdom.

The outlaw's car was equipped with a GPS transmitter. The transmitter showed that
the car made exactly `n` movements on its way from the headquarters to the stall.
A movement can move the car from point `(x, y)` to one of these four points:

- `(x - 1, y)` — letter `L`
- `(x + 1, y)` — letter `R`
- `(x, y - 1)` — letter `D`
- `(x, y + 1)` — letter `U`

The GPS transmitter is very inaccurate and does not preserve the exact sequence
of the car's movements. Instead, it keeps records of the car's possible
movements. Each record is a string of one of these types: `UL`, `UR`, `DL`,
`DR`, or `ULDR`. Each such string means that the car made a single movement
corresponding to one of the characters of the string. For example, string `UL`
means that the car moved either `U` or `L`.

You have received the journal with the outlaw's possible movements from the
headquarters to the stall. The journal records are given in chronological order.
Given that the ice-cream stall is located at point `(0, 0)`, print the number
of different points that can contain the gang headquarters (that is, the number
of different possible locations of the car's origin).

## Input

The first line contains a single integer `n` (`1 <= n <= 2 * 10^5`) — the number
of the car's movements from the headquarters to the stall.

Each of the following `n` lines describes the car's possible movements. It is
guaranteed that each possible movement is one of the following strings: `UL`,
`UR`, `DL`, `DR`, or `ULDR`.

All movements are given in chronological order.

## Output

Print a single integer — the number of different possible locations of the gang's
headquarters.

## Example

### Input

```text
3
UR
UL
ULDR
2
DR
DL
```

### Output

```text
9
4
```

## Note

The figure in the statement shows the nine possible positions of the gang
headquarters in the first sample.

For example, the following movements can get the car from point `(1, 0)` to point
`(0, 0)`.


### ideas
1. 假设移动了l,r,u,d长度
2. 在只考虑(UL, UR, DL, DR)的情况下,能够得到一个菱形的区域,这个菱形区域的边上,可以移动到它们的原点
3. 现在考虑ULRD的影响, 它可以让这个圈扩大一圈,
4. 如果是个ULRD呢? 似乎是从最外面的那圈,一个ULRD,往里面移动一格
# D. Cinema

[Problem link](https://codeforces.com/problemset/problem/234/D)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: input.txt

output: output.txt

There are `m` actors in Berland. Each actor has a personal identifier — an integer
from `1` to `m` (distinct actors have distinct identifiers). Vasya likes to watch
Berland movies with Berland actors, and he has `k` favorite actors.

He watched movie trailers for the next month and wrote down, for every movie:

- the movie title;
- the number of actors who starred in it;
- the identifiers of these actors.

He copied the titles and actor counts, but for some movies he failed to write
down every actor identifier.

Vasya wonders which movies may be his favorites, and which may not. Once the
exact cast of all movies is known, his favorite movies are determined as
follows: a movie becomes a favorite movie if no other movie from Vasya's list
has more favorite actors.

For each movie, determine:

- whether it **surely will** be his favorite movie;
- whether it **surely will not** be his favorite movie;
- whether it **can either** be favorite or not favorite.

## Input

The first line contains two integers `m` and `k` (`1 <= m <= 100`, `1 <= k <= m`)
— the number of actors in Berland and the number of Vasya's favorite actors.

The second line contains `k` distinct integers `a_i` (`1 <= a_i <= m`) — the
identifiers of Vasya's favorite actors.

The third line contains a single integer `n` (`1 <= n <= 100`) — the number of
movies in Vasya's list.

Then follow `n` blocks of lines, each block describing one movie. The `i`-th
movie description contains three lines:

- the first line contains string `s_i` (lowercase English letters, length from
  1 to 10) — the movie title;
- the second line contains a non-negative integer `d_i` (`1 <= d_i <= m`) — the
  number of actors in this movie;
- the third line contains `d_i` integers `b_{i,j}` (`0 <= b_{i,j} <= m`) — the
  identifiers of the actors. If `b_{i,j} = 0`, Vasya does not remember the
  identifier of the `j`-th actor.

It is guaranteed that the actor list of a movie contains no duplicate
identifiers. All movie titles are distinct. Numbers on a line are separated by
single spaces.

## Output

Print `n` lines. In the `i`-th line print:

- `0` if the `i`-th movie will surely be a favorite;
- `1` if the `i`-th movie will surely not be a favorite;
- `2` if the `i`-th movie can either be a favorite or not a favorite.

## Example

### Input

```text
5 3
1 2 3
6
firstfilm
3
0 0 0
secondfilm
4
0 0 4 5
thirdfilm
1
2
fourthfilm
1
5
fifthfilm
1
4
sixthfilm
2
1 0
```

### Output

```text
2
2
1
1
1
2
```

### Input

```text
5 3
1 3 5
4
jumanji
3
0 0 0
theeagle
5
1 2 3 4 0
matrix
3
2 4 0
sourcecode
2
2 4
```

### Output

```text
2
0
1
1
```

## Note

In the second sample:

- Movie `jumanji` can theoretically have from 1 to 3 of Vasya's favorite actors.
- Movie `theeagle` has all three favorite actors, because the forgotten actor can
  only have identifier `5`.
- Movie `matrix` can have exactly one favorite actor.
- Movie `sourcecode` has no favorite actors.

So `theeagle` will surely be favorite, `matrix` and `sourcecode` will surely not
be favorite, and `jumanji` can either be favorite (if it has all three favorite
actors) or not favorite.

## Solution Summary

For every movie, compute an interval:

```text
[minimum possible favorite actors, maximum possible favorite actors]
```

Let:

- `stars` be the number of known favorite actors in the movie;
- `zero` be the number of unknown actor identifiers;
- `used` be the known identifiers already appearing in that movie.

Actor identifiers inside one movie must be distinct, so unknown positions can
only use identifiers not in `used`.

### Minimum Possible Count

To minimize the number of favorite actors, assign as many unknown positions as
possible to available non-favorite actors.

If `nonFavoriteAvailable` is the number of unused non-favorite actors, then:

```text
minimum = stars + max(0, zero - nonFavoriteAvailable)
```

The remaining unknown positions, if any, are forced to be favorite actors.

This value is stored as `s1` in the implementation.

### Maximum Possible Count

To maximize the number of favorite actors, assign as many unknown positions as
possible to available favorite actors.

If `favoriteAvailable` is the number of unused favorite actors, then:

```text
maximum = stars + min(zero, favoriteAvailable)
```

This value is stored as `s2`.

### Classifying a Movie

Suppose movie `i` has interval `[s1_i, s2_i]`.

A movie is favorite when no other movie has strictly more favorite actors. Ties
are allowed.

Movie `i` is surely favorite if even its minimum is at least every other
movie's maximum:

```text
s1_i >= max(s2_j), for every j != i
```

Then no completion of the unknown identifiers can make another movie strictly
better, so the answer is `0`.

Movie `i` is surely not favorite if some other movie's minimum is strictly
greater than its maximum:

```text
s2_i < max(s1_j), for some j != i
```

Then that other movie always has more favorite actors, so the answer is `1`.

If neither condition holds, the result depends on how the unknown identifiers
are assigned, so the answer is `2`.

The implementation keeps the two movies with the largest `s2` values and the
two with the largest `s1` values. Two are needed because, when classifying a
movie, that movie itself must be excluded from the competitor maximum.

When there is only one movie, it is always favorite, so the answer is `0`.

## Correctness

The computed `s1` is attainable because it first fills unknown positions with
all possible unused non-favorite actors and uses favorite actors only when
necessary. No assignment can use fewer favorite actors because there are not
enough available non-favorites.

Similarly, `s2` is attainable by filling unknown positions with as many unused
favorite actors as possible, and no assignment can contain more favorites.
Thus `[s1, s2]` gives the exact minimum and maximum for each movie.

If `s1_i` is at least every competitor's `s2`, movie `i` can never be strictly
beaten and is surely favorite. If `s2_i` is smaller than some competitor's
`s1`, it is always strictly beaten and is surely not favorite. In every other
case, both outcomes are possible, giving classification `2`.

## Complexity

For each movie, the code scans at most all `m` actor identifiers to compute its
minimum and maximum. The total time complexity is `O(n * m)`, and the memory
usage is `O(n + m)`.

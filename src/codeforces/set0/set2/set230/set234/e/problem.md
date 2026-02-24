# Problem E

In the autumn of this year, two Russian teams came into the group stage of the most prestigious football club competition in the world — the UEFA Champions League. Now, these teams have already started to play in the group stage and are fighting for advancing to the playoffs. In this problem we are interested in the draw stage: the process of sorting teams into groups.

The process of the draw goes as follows (the rules here are simplified compared to real life). Suppose $n$ teams take part in the group stage ($n$ is divisible by four). The teams are divided into groups of four. Denote the number of groups by $m$ ($m = n/4$). Each team has a **rating** — an integer characterizing the team's achievements. The teams are sorted by rating in **decreasing** order (no two teams have the same rating).

Then four "baskets" are formed, each containing $m$ teams: the first $m$ teams (highest rating) go to the first basket, the next $m$ to the second, and so on.

The following procedure is repeated $m - 1$ times:

1. A team is chosen at random from each basket — first from basket 1, then basket 2, then 3, then 4. The four chosen teams form a new group and are removed from their baskets.

The four teams left in the baskets after these $m - 1$ rounds form the last group.

**Random number generator.** In reality the draw uses people to pick teams; here we use a generator with four positive integer parameters $x$, $a$, $b$, $c$. On each call it:

1. Computes $y = (a \cdot x + b) \bmod c$ (remainder of $a \cdot x + b$ when divided by $c$).
2. Sets $x \leftarrow y$.
3. Returns (the new) $x$ as the random number.

When we need to choose a team from a basket, the generator yields a number $k$. The teams still in the basket are numbered $0$ to $s - 1$ in order of decreasing rating, where $s$ is the current basket size. We take the team with number $k \bmod s$.

Given the list of teams and the generator parameters, determine the result of the draw.

---

## Input

- First line: integer $n$ ($4 \le n \le 64$, $n$ divisible by $4$) — the number of teams.
- Second line: four integers $x$, $a$, $b$, $c$ ($1 \le x, a, b, c \le 1000$) — the generator parameters.
- Next $n$ lines: each line is the team name and its rating, separated by a space. Names use uppercase/lowercase English letters, length 1–20. Rating is an integer from $0$ to $1000$. All names are distinct; all ratings are distinct.

## Output

Print the groups in the order they are formed. Label groups with consecutive uppercase letters starting from `A`. Inside each group, print the four team names one per line, in order of **decreasing** rating. See the samples for the exact format.

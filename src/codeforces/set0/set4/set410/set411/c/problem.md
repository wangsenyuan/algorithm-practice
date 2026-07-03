# C. Kicker

[Problem link](https://codeforces.com/problemset/problem/411/C)

**Contest:** [Codeforces Round #239 (Div. 2)](https://codeforces.com/contest/411)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Two teams play table football (kicker), two players per team.

- Team 1: players 1 and 2
- Team 2: players 3 and 4

Player `i` has defence skill `a_i` and attack skill `b_i`.

Before the match, each team assigns one player to defence and one to attack. **Team 1 chooses first**; **Team 2 chooses after seeing Team 1's assignment**.

Team X beats Team Y if both hold:

- `defence_X > attack_Y`
- `attack_X > defence_Y`

Both teams play optimally. Determine which team is **guaranteed** to win, or print `Draw` if neither can force a win.

## Input

Four lines. Line `i` contains two integers `a_i` and `b_i` (`1 <= a_i, b_i <= 100`).

## Output

Print `Team 1`, `Team 2`, or `Draw`.

## Example

### Input

```text
1 10
10 1
9 9
9 9
```

### Output

```text
Team 1
```

### Input

```text
1 1
2 2
3 3
2 2
```

### Output

```text
Team 2
```

### Input

```text
3 3
2 2
1 1
2 2
```

### Output

```text
Draw
```

### Note

In sample 1, Team 1 can assign player 1 to attack and player 2 to defence, guaranteeing victory.

In sample 2, Team 2 can respond optimally after Team 1 reveals its roles.

## Solution

There are only two possible role assignments for each team:

1. First player defends, second player attacks.
2. Second player defends, first player attacks.

For one concrete assignment `x` against another assignment `y`, team `x` beats team `y` exactly when:

```text
x.defence > y.attack
x.attack  > y.defence
```

The important part is the order of decisions. Team 1 chooses first, then Team 2 sees that choice and
can choose its own best response.

So Team 1 can force a win only if it has at least one assignment that beats **both** possible Team 2
assignments. In logic:

```text
exists team1_assignment such that
    for every team2_assignment,
        team1_assignment beats team2_assignment
```

If Team 1 cannot force a win, Team 2 can force a win if it can respond successfully to every Team 1
choice:

```text
for every team1_assignment,
    exists team2_assignment such that
        team2_assignment beats team1_assignment
```

If neither condition holds, neither team is guaranteed to win, so the result is `Draw`.

The implementation builds the two assignments for Team 1 and Team 2, then checks these two conditions
directly.

### Correctness

The two generated assignments for a team cover every legal choice, because each team has exactly two
players and must choose one defender and one attacker.

If the algorithm prints `Team 1`, it found one Team 1 assignment that beats every Team 2 assignment.
Since Team 2 chooses after seeing Team 1's choice, Team 2 can only choose one of those checked
assignments, and all of them lose. Therefore Team 1 is guaranteed to win.

If the algorithm prints `Team 2`, then for every possible first choice by Team 1, there exists at
least one Team 2 assignment that beats it. Team 2 moves second and knows Team 1's assignment, so it can
choose the corresponding winning response. Therefore Team 2 is guaranteed to win.

If the algorithm prints `Draw`, Team 1 has no assignment that beats all Team 2 responses, and Team 2
does not have a winning response against every Team 1 assignment. Thus neither side can guarantee a
win under optimal play.

Therefore the algorithm returns exactly the required result.

### Complexity

Each team has only `2` assignments, so the checks are constant time. The time complexity is `O(1)`,
and the memory usage is `O(1)`.

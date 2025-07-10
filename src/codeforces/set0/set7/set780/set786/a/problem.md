# Problem A: Rick and Morty's Berzerk Game

## Problem Description

Rick and Morty are playing their own version of Berzerk (which has nothing in common with the famous Berzerk game). This game needs a huge space, so they play it with a computer.

In this game there are **n** objects numbered from 1 to n arranged in a circle (in clockwise order). Object number 1 is a black hole and the others are planets. There's a monster in one of the planets. Rick and Morty don't know on which one yet, only that he's not initially in the black hole, but Unity will inform them before the game starts. But for now, they want to be prepared for every possible scenario.

## Game Rules

Each one of them has a set of numbers between 1 and n - 1 (inclusive):
- Rick's set is **s₁** with **k₁** elements
- Morty's set is **s₂** with **k₂** elements

One of them goes first and the players change alternatively. In each player's turn, he should choose an arbitrary number like **x** from his set and the monster will move to his x-th next object from its current position (clockwise). If after his move the monster gets to the black hole, he wins.

## Task

Your task is that for each of monster's initial positions and who plays first, determine if the starter wins, loses, or the game will get stuck in an infinite loop. In case when a player can lose or make the game infinite, it's more profitable to choose the infinite game.

## Input

- The first line of input contains a single integer **n** (2 ≤ n ≤ 7000) — number of objects in game.
- The second line contains integer **k₁** followed by **k₁** distinct integers **s₁,₁, s₁,₂, ..., s₁,k₁** — Rick's set.
- The third line contains integer **k₂** followed by **k₂** distinct integers **s₂,₁, s₂,₂, ..., s₂,k₂** — Morty's set.

**Constraints:**
- 1 ≤ kᵢ ≤ n - 1
- 1 ≤ sᵢ,₁, sᵢ,₂, ..., sᵢ,kᵢ ≤ n - 1 for 1 ≤ i ≤ 2

## Output

- In the first line print **n - 1** words separated by spaces where the i-th word is:
  - **"Win"** (without quotations) if in the scenario that Rick plays first and monster is initially in object number i + 1 he wins
  - **"Lose"** if he loses
  - **"Loop"** if the game will never end

- Similarly, in the second line print **n - 1** words separated by spaces where the i-th word is:
  - **"Win"** (without quotations) if in the scenario that Morty plays first and monster is initially in object number i + 1 he wins
  - **"Lose"** if he loses
  - **"Loop"** if the game will never end

## Examples

### Example 1

**Input:**
```
5
2 3 2
3 1 2 3
```

**Output:**
```
Lose Win Win Loop
Loop Win Win Win
```

### Example 2

**Input:**
```
8
4 6 2 3 4
2 3 6
```

**Output:**
```
Win Win Win Win Win Win Win
Lose Win Lose Lose Win Lose Lose
```



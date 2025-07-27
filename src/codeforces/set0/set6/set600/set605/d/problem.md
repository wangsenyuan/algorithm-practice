# Problem D: Board Card Game

## Problem Description

You are playing a board card game. In this game the player has two characteristics:
- **x**: the white magic skill
- **y**: the black magic skill

There are **n** spell cards lying on the table, each with four characteristics: `ai`, `bi`, `ci`, and `di`.

## Game Mechanics

In one move, a player can pick one of the cards and cast the spell written on it, but only if the first two characteristics meet the requirement:
- `ai ≤ x` 
- `bi ≤ y`

This means the player must have enough magic skill to cast the spell.

After casting the spell, the player's characteristics change to:
- `x = ci`
- `y = di`

## Game Objective

- At the beginning of the game, both characteristics are equal to zero
- The goal is to cast the **n-th spell**
- You must achieve this in as few moves as possible
- You can use spells in any order and any number of times (including not using some spells at all)

## Input

- **Line 1**: A single integer `n` (1 ≤ n ≤ 100,000) — the number of cards on the table
- **Lines 2 to n+1**: Four integers `ai`, `bi`, `ci`, `di` (0 ≤ ai, bi, ci, di ≤ 10^9) — the characteristics of the corresponding card

## Output

- **Line 1**: A single integer `k` — the minimum number of moves needed to cast the n-th spell
- **Line 2**: `k` numbers — the indices of the cards in the order you should cast them

**Note**: If there are multiple possible solutions, print any of them.

If it is impossible to cast the n-th spell, print **-1**.
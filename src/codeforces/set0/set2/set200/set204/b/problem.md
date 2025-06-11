# Funny Card Set

## Problem Statement
The Little Elephant loves to play with color cards. He has n cards, each with exactly two colors (front and back). Initially, all cards are placed on the table with the front side up. In one move, the Little Elephant can turn any card to its other side.

A set of cards is considered "funny" if at least half of the cards have the same color (considering the color of the upper side for each card).

## Task
Find the minimum number of moves needed to make the set of n cards funny. If it's impossible to make the set funny, output -1.

## Input Format
- First line: A single integer n ($1 \leq n \leq 10^5$) — the number of cards
- Next n lines: Description of each card
  - Each line contains two positive integers (not exceeding $10^9$)
  - First number: Color of the front side
  - Second number: Color of the back side
  - Note: Front and back colors may be the same
  - Numbers are separated by single spaces

## Output Format
- A single integer
  - The minimum number of moves needed to make the set funny
  - Output -1 if it's impossible to make the set funny


# Fox Ciel's Card Game

Fox Ciel is playing a card game with her friend Jiro.

## Game Setup

- **Jiro** has $n$ cards, each with two attributes: position (Attack or Defense) and strength
- **Ciel** has $m$ cards, each with two attributes: position and strength
- All of Ciel's cards have Attack position

## Battle Phase Rules

During Ciel's battle phase, she can perform the following operation multiple times:

1. Choose one of her cards $X$ (must not have been chosen before)
2. If Jiro has no alive cards:
   - Jiro takes damage equal to $X$'s strength
3. If Jiro has alive cards:
   - Ciel must choose one of Jiro's alive cards $Y$
   - **If $Y$'s position is Attack:**
     - $X$'s strength $\geq$ $Y$'s strength must hold
     - Card $Y$ dies
     - Jiro takes damage equal to $(X$'s strength$) - (Y$'s strength$)$
   - **If $Y$'s position is Defense:**
     - $X$'s strength $>$ $Y$'s strength must hold
     - Card $Y$ dies
     - Jiro takes no damage

Ciel can end her battle phase at any moment (she doesn't need to use all her cards).

**Goal:** Calculate the maximal sum of damage Jiro can receive.

## Input

- First line: Two integers $n$ and $m$ $(1 \leq n, m \leq 100)$ — number of cards Jiro and Ciel have
- Next $n$ lines: String `position` and integer `strength` $(0 \leq \text{strength} \leq 8000)$ — Jiro's cards
  - `position` is either "ATK" (attack) or "DEF" (defense)
- Next $m$ lines: Integer `strength` $(0 \leq \text{strength} \leq 8000)$ — Ciel's cards

## Output

Output an integer: the maximal damage Jiro can receive.

## Examples

### Example 1
**Input:**
```
2 3
ATK 2000
DEF 1700
2500
2500
2500
```

**Output:**
```
3000
```

**Explanation:** Ciel has 3 cards with same strength. Best strategy:
1. Use one card to attack "ATK 2000" → destroys it, Jiro gets $2500 - 2000 = 500$ damage
2. Use second card to destroy "DEF 1700" → Jiro gets no damage
3. Use third card to attack (no cards left) → Jiro gets $2500$ damage
4. Total: $500 + 2500 = 3000$

### Example 2
**Input:**
```
3 4
ATK 10
ATK 100
ATK 1000
1
11
101
1001
```

**Output:**
```
992
```

**Explanation:** Use "1001" card to attack "ATK 100", then "101" card to attack "ATK 10". Total damage: $(1001 - 100) + (101 - 10) = 992$

### Example 3
**Input:**
```
2 4
DEF 0
ATK 0
0
0
1
1
```

**Output:**
```
1
```

**Note:** Ciel can destroy "ATK 0" card with a card of strength 0, but cannot destroy "DEF 0" card with strength 0 (must be strictly greater).
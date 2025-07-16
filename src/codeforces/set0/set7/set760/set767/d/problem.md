# Problem D: Milk Cartons

## Problem Description

Olya likes milk very much. She drinks k cartons of milk each day if she has at least k and drinks all of them if she doesn't. But there's an issue — expiration dates. Each carton has a date after which you can't drink it (you still can drink it exactly at the date written on the carton). Due to this, if Olya's fridge contains a carton past its expiry date, she throws it away.

Olya hates throwing out cartons, so when she drinks a carton, she chooses the one which expires the fastest. It's easy to understand that this strategy minimizes the amount of cartons thrown out and lets her avoid it if it's even possible.

> **Note:** Milk. Best before: 20.02.2017.

The main issue Olya has is the one of buying new cartons. Currently, there are n cartons of milk in Olya's fridge, for each one an expiration date is known (how soon does it expire, measured in days). In the shop that Olya visited there are m cartons, and the expiration date is known for each of those cartons as well.

**Find the maximum number of cartons Olya can buy so that she wouldn't have to throw away any cartons.** Assume that Olya drank no cartons today.

## Input

- **First line:** Three integers n, m, k (1 ≤ n, m ≤ 10^6, 1 ≤ k ≤ n + m)
  - n = amount of cartons in Olya's fridge
  - m = amount of cartons in the shop
  - k = number of cartons Olya drinks each day

- **Second line:** n integers f₁, f₂, ..., fₙ (0 ≤ fᵢ ≤ 10^7)
  - Expiration dates of the cartons in Olya's fridge
  - Expiration date is expressed by the number of days the drinking of this carton can be delayed
  - Example: 0 means it must be drunk today, 1 means no later than tomorrow, etc.

- **Third line:** m integers s₁, s₂, ..., sₘ (0 ≤ sᵢ ≤ 10^7)
  - Expiration dates of the cartons in the shop in a similar format

## Output

- **If there's no way for Olya to drink the cartons she already has in her fridge:** Print -1

- **Otherwise:**
  - **First line:** Print the maximum number x of cartons which Olya can buy so that she wouldn't have to throw a carton away
  - **Second line:** Print exactly x integers — the numbers of the cartons that should be bought
    - Cartons are numbered in the order they are written in the input, starting with 1
    - Numbers should not repeat, but can be in arbitrary order
    - If there are multiple correct answers, print any of them

### ideas
1. 每瓶牛奶，可以在它被过期前的任何日期被喝掉。但是最好是delay到最后一天再决定
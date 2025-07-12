# Problem C: Felicity and Evolution Camp

## Problem Description

It's that time of the year, Felicity is around the corner and you can see people celebrating all around the Himalayan region. The Himalayan region has **n** gyms. The **i-th** gym has **gi** Pokemon in it. There are **m** distinct Pokemon types in the Himalayan region numbered from 1 to m.

There is a special evolution camp set up in the fest which claims to evolve any Pokemon. The type of a Pokemon could change after evolving, subject to the constraint that:

- If two Pokemon have the **same type** before evolving, they will have the **same type** after evolving.
- If two Pokemon have **different types** before evolving, they will have **different types** after evolving.
- It is also possible that a Pokemon has the same type before and after evolving.

## Formal Definition

An **evolution plan** is a permutation **f** of {1, 2, ..., m}, such that f(x) = y means that a Pokemon of type x evolves into a Pokemon of type y.

The gym leaders are intrigued by the special evolution camp and all of them plan to evolve their Pokemons. The protocol of the mountain states that in each gym, for every type of Pokemon, the number of Pokemon of that type before evolving any Pokemon should be equal to the number of Pokemon of that type after evolving all the Pokemons according to the evolution plan.

Two evolution plans **f1** and **f2** are **distinct** if they have at least one Pokemon type evolving into a different Pokemon type in the two plans, i.e., there exists an i such that f1(i) ≠ f2(i).

## Task

Your task is to find how many distinct evolution plans are possible such that if all Pokemon in all the gyms are evolved, the number of Pokemon of each type in each of the gyms remains the same. As the answer can be large, output it modulo 10^9 + 7.

## Input

- The first line contains two integers **n** and **m** (1 ≤ n ≤ 10^5, 1 ≤ m ≤ 10^6) — the number of gyms and the number of Pokemon types.
- The next **n** lines contain the description of Pokemons in the gyms. The **i-th** of these lines begins with the integer **gi** (1 ≤ gi ≤ 10^5) — the number of Pokemon in the i-th gym. After that **gi** integers follow, denoting types of the Pokemons in the i-th gym. Each of these integers is between 1 and m.
- The total number of Pokemons (the sum of all gi) does not exceed 5·10^5.

## Output

Output the number of valid evolution plans modulo 10^9 + 7.

## Examples

### Example 1
**Input:**
```
2 3
2 1 2
2 2 3
```
**Output:**
```
1
```
**Explanation:** The only possible evolution plan is the identity permutation.

### Example 2
**Input:**
```
1 3
3 1 2 3
```
**Output:**
```
6
```
**Explanation:** Any permutation of (1, 2, 3) is valid.

### Example 3
**Input:**
```
2 4
2 1 2
3 2 3 4
```
**Output:**
```
2
```
**Explanation:** There are two possible plans.

### Example 4
**Input:**
```
2 2
3 2 2 1
2 1 2
```
**Output:**
```
1
```
**Explanation:** The only possible evolution plan is the identity permutation.

### Example 5
**Input:**
```
3 7
2 1 2
2 3 4
3 5 6 7
```
**Output:**
```
24
```
**Explanation:** There are 24 valid evolution plans.


## ideas
1. 不同的类型，可以组成一个cycle，在一个cycle内，假设它的长度是n, 贡献 = P[n]
2. 同一个cycle内，满足条件 = 其中的元素c[1], c[2], ... c[n] 在所有的gym中，freq要相同
3. 比如x, y 在gym1中，出现了0次，在gym2中，出现1次，在gym3中出现了2次，那么它们可以放在一个cycle中
4. 但是如果 x和y，只要在一个gym中，它们的freq不一样，那么它们就不能放在一个cycle中
5. 感觉无从下手啊
6. 假设对于任何一个x，对它在所有gym中的freq，作为一个向量化，那么就是所有相同向量值的，分在一起
7. ok。用这个做个hash就可以了
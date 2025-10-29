# Problem C: Fruit Salad

## Problem Statement

Dima, Inna and Seryozha have gathered in a room. That's right, someone's got to go. To cheer Seryozha up and inspire him to have a walk, Inna decided to cook something.

Dima and Seryozha have n fruits in the fridge. Each fruit has two parameters: the taste and the number of calories. Inna decided to make a fruit salad, so she wants to take some fruits from the fridge for it. Inna follows a certain principle as she chooses the fruits: the total taste to the total calories ratio of the chosen fruits must equal k. In other words, $\frac{\sum a_j}{\sum b_j} = k$, where $a_j$ is the taste of the j-th chosen fruit and $b_j$ is its calories.

Inna hasn't chosen the fruits yet, she is thinking: what is the maximum taste of the chosen fruits if she strictly follows her principle? Help Inna solve this culinary problem — now the happiness of a young couple is in your hands!

Inna loves Dima very much so she wants to make the salad from at least one fruit.

## Input

The first line of the input contains two integers n, k (1 ≤ n ≤ 100, 1 ≤ k ≤ 10). The second line of the input contains n integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 100) — the fruits' tastes. The third line of the input contains n integers b₁, b₂, ..., bₙ (1 ≤ bᵢ ≤ 100) — the fruits' calories. Fruit number i has taste aᵢ and calories bᵢ.

## Output

If there is no way Inna can choose the fruits for the salad, print in the single line number -1. Otherwise, print a single integer — the maximum possible sum of the taste values of the chosen fruits.

## Examples

### Example 1

**Input:**
```
3 2
10 8 1
2 7 1
```

**Output:**
```
18
```

### Example 2

**Input:**
```
5 3
4 4 4 4 4
2 2 2 2 2
```

**Output:**
```
-1
```

## Note

In the first test sample we can get the total taste of the fruits equal to 18 if we choose fruit number 1 and fruit number 2, then the total calories will equal 9. The condition $\frac{18}{9} = 2$ fulfills, that's exactly what Inna wants.

In the second test sample we cannot choose the fruits so as to follow Inna's principle.


### ideas
1. 但是没法使用200个数，把它给表示出来～
2. 假设，目标是选择x个item的时候，使的 sum(a) = sum(b) * k时最大的sum(a), 是否可行呢？
3. 这个x似乎也没啥用～
4. 100个item，k的可能性有多少个？理论上是 pow(2, n)， 但感觉似乎不是
5. 是不是这个范围是比较少的。如果能够计算出所有的k，那么就可以给它们编码，从而去计算在指定k的情况下的最优解
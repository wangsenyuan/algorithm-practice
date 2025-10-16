## Problem Description

There was a big bank robbery in Tablecity. In order to catch the thief, the President called none other than Albert – Tablecity's Chief of Police. Albert does not know where the thief is located, but he does know how he moves.

Tablecity can be represented as a 1000 × 2 grid, where every cell represents one district. Each district has its own unique name "(X, Y)", where X and Y are the coordinates of the district in the grid.

### Thief's Movement

Every hour the thief will leave the district (X, Y) he is currently hiding in, and move to one of the districts:
- (X - 1, Y)
- (X + 1, Y)
- (X - 1, Y - 1)
- (X - 1, Y + 1)
- (X + 1, Y - 1)
- (X + 1, Y + 1)

as long as it exists in Tablecity.

For example, if the thief is located in district (7, 1), he can move to any of the valid neighboring districts listed above.

### Task

Albert has enough people so that every hour he can pick any two districts in Tablecity and fully investigate them, making sure that if the thief is located in one of them, he will get caught. Albert promised the President that the thief will be caught in no more than 2015 hours and needs your help in order to achieve that.

## Input

There is no input for this problem.

## Output

The first line of output contains integer N – duration of police search in hours.

Each of the following N lines contains exactly 4 integers `Xi1, Yi1, Xi2, Yi2` separated by spaces, that represent 2 districts `(Xi1, Yi1)`, `(Xi2, Yi2)` which got investigated during i-th hour.

Output is given in chronological order (i-th line contains districts investigated during i-th hour) and should guarantee that the thief is caught in no more than 2015 hours, regardless of thief's initial position and movement.

## Constraints

- N ≤ 2015
- 1 ≤ X ≤ 1000
- 1 ≤ Y ≤ 2


### ideas
1. 从左往右，每列都查询两次？
2. 因为thief没法在原来的位置停留，所以假设第一次检查的时候，他正好在下一列，然后检查完以后，马上移动到了这一列，那么第二次检查就抓住他了
3. 不行，存在一种情况是， thief正好在第i+2列，这样子，检查完以后他到了i+1列，检查完以后，再到i列，就完蛋了～
4. 
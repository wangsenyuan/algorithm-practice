# Photo Problem

Many years have passed, and n friends met at a party again. Technologies have leaped forward since
the last meeting, cameras with timer appeared and now it is not obligatory for one of the friends
to stand with a camera, and, thus, being absent on the photo.

Simply speaking, the process of photographing can be described as follows. Each friend occupies a
rectangle of pixels on the photo: the i-th of them in a standing state occupies a wi pixels wide
and a hi pixels high rectangle. But also, each person can lie down for the photo, and then he will
occupy a hi pixels wide and a wi pixels high rectangle.

The total photo will have size W × H, where W is the total width of all the people rectangles, and
H is the maximum of the heights. The friends want to determine what minimum area the group photo
can they obtain if no more than n / 2 of them can lie on the ground (it would be strange if more
than n / 2 gentlemen lie on the ground together, isn't it?..)

Help them to achieve this goal.

## Input

The first line contains integer n (1 ≤ n ≤ 1000) — the number of friends.

The next n lines have two integers wi, hi (1 ≤ wi, hi ≤ 1000) each, representing the size of the
rectangle, corresponding to the i-th friend.

## Output

Print a single integer equal to the minimum possible area of the photo containing all friends if
no more than n / 2 of them can lie on the ground.

## Examples

### Example 1

```text
3
10 1
20 2
30 3
```

Output:

```text
180
```

### Example 2

```text
3
3 1
2 2
4 3
```

Output:

```text
21
```

### Example 3

```text
1
5 10
```

Output:

```text
50
```

## Solution

In an online mirror version the problem was slightly harder. Let's call people with $w \leq h$ *high*, and remaining people *wide*. Let's fix photo height $H$. Let's consider several following cases:

1. If a high person fits into height $H$, we leave him as is.
2. If a high person doesn't fit into height $H$, then we have to ask him to lie down, increasing the counter of such people by 1.
3. If a wide person fits into height $H$, but doesn't fit lying on the ground, then we leave him staying.
4. If a wide person fits into height $H$ in both ways, then we first ask him to stay and write into a separate array value of answer decrease if we ask him to lie on the ground: $w - h$.
5. If somebody doesn't fit in both ways, then such value of $H$ is impossible.

Now we have several people that have to lie on the ground (from case 2) and if there are too many of them (more than $n / 2$) then such value of $H$ is impossible.

After we put down people from case 2 there can still be some vacant ground positions, we distribute them to the people from case 4 with highest values of $w - h$. Then we calculate the total area of the photo and relax the answer.

Having bought his own apartment, Boris decided to paper the walls in every room. Boris's flat has n rooms, each of which has the form of a rectangular parallelepiped. For every room we known its length, width and height of the walls in meters (different rooms can have different dimensions, including height).

Boris chose m types of wallpaper to paper the walls of the rooms with (but it is not necessary to use all the types). Each type of wallpaper is sold in rolls of a fixed length and width (the length, naturally, shows how long the unfolded roll will be). In addition, for each type we know the price of one roll of this type.

The wallpaper of each type contains strips running along the length of the roll. When gluing the strips must be located strictly vertically (so the roll cannot be rotated, even if the length is less than the width). Besides, a roll can be cut in an arbitrary manner, but the joints of glued pieces should also be vertical. In addition, each room should be papered by only one type of wallpaper. And pieces of the same roll cannot be used to paper different rooms. That is, for each room the rolls are purchased separately. Also, some rolls can be used not completely.

After buying an apartment Boris is short of cash, so he wants to spend the minimum money on wallpaper. Help him.

Input
The first line contains a positive integer n (1 ≤ n ≤ 500) — the number of rooms in Boris's apartment.

Each of the next n lines contains three space-separated positive integers — the length, width and height of the walls in a given room in meters, respectively.

The next line contains a positive integer m (1 ≤ m ≤ 500) — the number of available wallpaper types.

Each of the following m lines contains three space-separated positive integers — the length and width in meters of a given wallpaper and the price of one roll, respectively.

All numbers in the input data do not exceed 500. It is guaranteed that each room can be papered using these types of wallpaper.

Output
Print a single number — the minimum total cost of the rolls.

### ideas
1. 每个房间可以换成成长度 2 * (length + width), height 
2. 贴纸的时候，是从上到下（垂直）的贴，且不能有水平方向的连接
3. 
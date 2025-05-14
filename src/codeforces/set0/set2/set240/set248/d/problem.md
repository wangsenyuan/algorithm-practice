Christmas celebrations are coming to Whoville. Cindy Lou Who and her parents Lou Lou Who and Betty Lou Who decided to give sweets to all people in their street. They decided to give the residents of each house on the street, one kilogram of sweets. So they need as many kilos of sweets as there are homes on their street.

The street, where the Lou Who family lives can be represented as n consecutive sections of equal length. You can go from any section to a neighbouring one in one unit of time. Each of the sections is one of three types: an empty piece of land, a house or a shop. Cindy Lou and her family can buy sweets in a shop, but no more than one kilogram of sweets in one shop (the vendors care about the residents of Whoville not to overeat on sweets).

After the Lou Who family leave their home, they will be on the first section of the road. To get to this section of the road, they also require one unit of time. We can assume that Cindy and her mom and dad can carry an unlimited number of kilograms of sweets. Every time they are on a house section, they can give a kilogram of sweets to the inhabitants of the house, or they can simply move to another section. If the family have already given sweets to the residents of a house, they can't do it again. Similarly, if they are on the shop section, they can either buy a kilo of sweets in it or skip this shop. If they've bought a kilo of sweets in a shop, the seller of the shop remembered them and the won't sell them a single candy if they come again. The time to buy and give sweets can be neglected. The Lou Whos do not want the people of any house to remain without food.

The Lou Whos want to spend no more than t time units of time to give out sweets, as they really want to have enough time to prepare for the Christmas celebration. In order to have time to give all the sweets, they may have to initially bring additional k kilos of sweets.

Cindy Lou wants to know the minimum number of k kilos of sweets they need to take with them, to have time to give sweets to the residents of each house in their street.

Your task is to write a program that will determine the minimum possible value of k.

Input
The first line of the input contains two space-separated integers n and t (2 ≤ n ≤ 5·105, 1 ≤ t ≤ 109). The second line of the input contains n characters, the i-th of them equals "H" (if the i-th segment contains a house), "S" (if the i-th segment contains a shop) or "." (if the i-th segment doesn't contain a house or a shop).

It is guaranteed that there is at least one segment with a house.

Output
If there isn't a single value of k that makes it possible to give sweets to everybody in at most t units of time, print in a single line "-1" (without the quotes). Otherwise, print on a single line the minimum possible value of k.

### ideas
1. 二分check没有问题
2. 但是如何check呢？假设到达某个H的时候，这时候，没有礼物，那么只能往前移动
3. 到达shop的时候，买到一个礼物。那么有两个选择，
4. 1是返回头去赠送这个礼物
5. 2是向前移动
6. 对于策略2来说，时间似乎是固定的
7. 就是最前面没有被满足的H的距离 + 最后面的H/S的距离，在这个范围内，在回头的时候能够满足所有的H + 目前移动的距离
8. 后面这个策略，应该是将所有的礼物都买下来，然后将看，这个差值，是否能够到达那个第一个未满足的地方
9. 如果策略2能成立，那么就返回true；
10. 如果策略2不能成立，那么就必须选择策略1
11. 先返回去满足前面，然后回到这里，继续往前
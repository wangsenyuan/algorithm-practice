In ABBYY a wonderful Smart Beaver lives. This time, he began to study history. When he read about the Roman Empire, he became interested in the life of merchants.

The Roman Empire consisted of n cities numbered from 1 to n. It also had m bidirectional roads numbered from 1 to m. Each road connected two different cities. Any two cities were connected by no more than one road.

We say that there is a path between cities c1 and c2 if there exists a finite sequence of cities t1, t2, ..., tp (p ≥ 1) such that:

t1 = c1
tp = c2
for any i (1 ≤ i < p), cities ti and ti + 1 are connected by a road
We know that there existed a path between any two cities in the Roman Empire.

In the Empire k merchants lived numbered from 1 to k. For each merchant we know a pair of numbers si and li, where si is the number of the city where this merchant's warehouse is, and li is the number of the city where his shop is. The shop and the warehouse could be located in different cities, so the merchants had to deliver goods from the warehouse to the shop.

Let's call a road important for the merchant if its destruction threatens to ruin the merchant, that is, without this road there is no path from the merchant's warehouse to his shop. Merchants in the Roman Empire are very greedy, so each merchant pays a tax (1 dinar) only for those roads which are important for him. In other words, each merchant pays di dinars of tax, where di (di ≥ 0) is the number of roads important for the i-th merchant.

The tax collection day came in the Empire. The Smart Beaver from ABBYY is very curious by nature, so he decided to count how many dinars each merchant had paid that day. And now he needs your help.

### ideas
1. 就是要找出所有的bridge
2. 联通块变成一个超级节点。看这两个连接属于哪两个超级节点，看他们之间的lca距离
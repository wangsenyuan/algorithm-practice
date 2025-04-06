Inside the large kingdom, there is an infinite dining hall. It can be represented as a set of cells (𝑥,𝑦
), where 𝑥
 and 𝑦
 are non-negative integers. There are an infinite number of tables in the hall. Each table occupies four cells (3𝑥+1,3𝑦+1
), (3𝑥+1,3𝑦+2
), (3𝑥+2,3𝑦+1
), (3𝑥+2,3𝑦+2
), where 𝑥
 and 𝑦
 are arbitrary non-negative integers. All cells that do not belong to any of the tables are corridors.

There are 𝑛
 guests that come to the dining hall one by one. Each guest appears in the cell (0,0)
 and wants to reach a table cell. In one step, they can move to any neighboring by side corridor cell, and in their last step, they must move to a neighboring by side a free table cell. They occupy the chosen table cell, and no other guest can move there.

Each guest has a characteristic 𝑡𝑖
, which can either be 0
 or 1
. They enter the hall in order, starting to walk from the cell (0,0
). If 𝑡𝑖=1
, the 𝑖
-th guest walks to the nearest vacant table cell. If 𝑡𝑖=0
, they walk to the nearest table cell that belongs to a completely unoccupied table. Note that other guests may choose the same table later.

The distance is defined as the smallest number of steps needed to reach the table cell. If there are multiple table cells at the same distance, the guests choose the cell with the smallest 𝑥
, and if there are still ties, they choose among those the cell with the smallest 𝑦
.

For each guest, find the table cell which they choose.


### ideas
1. 模拟
2. 距离不是等于 x + y, 而是要从空道走, 所以，右上角的要多+2
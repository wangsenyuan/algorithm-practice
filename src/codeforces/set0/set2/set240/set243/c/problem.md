Old MacDonald has a farm and a large potato field, (1010 + 1) × (1010 + 1) square meters in size. The field is divided into square garden beds, each bed takes up one square meter.

Old McDonald knows that the Colorado potato beetle is about to invade his farm and can destroy the entire harvest. To fight the insects, Old McDonald wants to spray some beds with insecticides.

So Old McDonald went to the field, stood at the center of the central field bed and sprayed this bed with insecticides. Now he's going to make a series of movements and spray a few more beds. During each movement Old McDonald moves left, right, up or down the field some integer number of meters. As Old McDonald moves, he sprays all the beds he steps on. In other words, the beds that have any intersection at all with Old McDonald's trajectory, are sprayed with insecticides.

When Old McDonald finished spraying, he wrote out all his movements on a piece of paper. Now he wants to know how many beds won't be infected after the invasion of the Colorado beetles.

It is known that the invasion of the Colorado beetles goes as follows. First some bed on the field border gets infected. Than any bed that hasn't been infected, hasn't been sprayed with insecticides and has a common side with an infected bed, gets infected as well. Help Old McDonald and determine the number of beds that won't be infected by the Colorado potato beetle.

### ideas
1. 就是要找出，在n步后，形成了多少被围起来的区域
2. 没啥想法～
  
### solution
Firstly tou should emulate all process in "idle mode". Let farmer was in points (x0 + 1 / 2, y0 + 1 / 2), (x1 + 1 / 2, y1 + 1 / 2), ... , (xn + 1 / 2, yn + 1 / 2). Coordinates x0, x0 + 1, x1, x1 + 1, ... xn, xn + 1 on axis x are interesting for us. Coordinates y0, y0 + 1, y1, y1 + 1, ... yn, yn + 1 on axis y are interesting too. You should store all of them into two arrays. Also you should add to these coordinates bounds of the field.

Then you should build "compressed" field: every cell of this field means rectangle of the starting field that placed between neighbour interesting coordinates. Size of the compressed field is O(n) × O(n). Let's initally all cells painted into the white color. Now you should emulate all process again and paint all visited cells in the compressed field by the black color. During every move you will paint O(n) cells, so all process will be done in O(n2).

Now you should emulate bugs' actions. You should run DFS (ot BFS) from some border cell of the compressed field and paint all reached cells be red color.

At the end you should iterate over all compressed field and find sum of areas of rectengles corresponding to black and white cells. So, you will receive the answer.

This solution works in O(n2).
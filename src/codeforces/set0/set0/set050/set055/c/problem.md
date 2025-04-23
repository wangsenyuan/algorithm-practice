Volodya and Vlad play the following game. There are k pies at the cells of n  ×  m board. Each turn Volodya moves one pie to the neighbouring (by side) cell. If the pie lies at the border of the board then Volodya can move it outside the board, get the pie and win. After Volodya's move, Vlad bans some edge at the border of the board of length 1 (between two knots of the board) so that Volodya is not able to move the pie outside the board through this edge anymore. The question is: will Volodya win this game? We suppose both players follow the optimal strategy.

### ideas
1. 从Vlad的角度看，是不是每次都要封堵离某个Pie最近的边？
2. 似乎也不是。如果Volodya一步就获胜了，那么就没有对方什么事了
3. 否则的话， Vlad只需要看对方动了哪个pie，在哪个方向，就封堵对应方向的边就可以了
4. 还有一种情况，就是Volodya，可以在有限步内到达角落
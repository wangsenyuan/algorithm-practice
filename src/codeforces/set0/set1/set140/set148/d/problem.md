The dragon and the princess are arguing about what to do on the New Year's Eve. The dragon suggests flying to the mountains to watch fairies dancing in the moonlight, while the princess thinks they should just go to bed early. They are desperate to come to an amicable agreement, so they decide to leave this up to chance.

They take turns drawing a mouse from a bag which initially contains w white and b black mice. The person who is the first to draw a white mouse wins. After each mouse drawn by the dragon the rest of mice in the bag panic, and one of them jumps out of the bag itself (the princess draws her mice carefully and doesn't scare other mice). Princess draws first. What is the probability of the princess winning?

If there are no more mice in the bag and nobody has drawn a white mouse, the dragon wins. Mice which jump out of the bag themselves are not considered to be drawn (do not define the winner). Once a mouse has left the bag, it never returns to it. Every mouse is drawn from the bag with the same probability as every other one, and every mouse jumps out of the bag with the same probability as every other one.

### ideas
1. dp[x][y]表示在剩余x个白色，y个黑色mouse的情况下，当前由公主draw时的胜率
2. 如果 x = 0 => dp[x][y] = 0 (龙胜)
3. 如果 y = 0 => dp[x][y] = 1 (公主胜)
4.    dp[x][y] = (x) / (x + y) (获得白色)
5.      + y / (x + y) * (?) (获得黑色)
6.    ？ = (y) / (x + y - 1) (龙获得黑色) (x > 0)的时候
7.       * (dp[x - 2][y-1] 跑掉的是白老鼠 + dp[x-1][y-2] 跑掉的是黑老鼠)
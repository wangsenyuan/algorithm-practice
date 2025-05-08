Vasya plays the Need For Brake. He plays because he was presented with a new computer wheel for birthday! Now he is sure that he will win the first place in the championship in his favourite racing computer game!

n racers take part in the championship, which consists of a number of races. After each race racers are arranged from place first to n-th (no two racers share the same place) and first m places are awarded. Racer gains bi points for i-th awarded place, which are added to total points, obtained by him for previous races. It is known that current summary score of racer i is ai points. In the final standings of championship all the racers will be sorted in descending order of points. Racers with an equal amount of points are sorted by increasing of the name in lexicographical order.

Unfortunately, the championship has come to an end, and there is only one race left. Vasya decided to find out what the highest and lowest place he can take up as a result of the championship.

### ideas
1. 要获得最低排名，最好是自己不得分，然后让在自己后面的人尽量的得分，跑到自己前面去
2. 已经在自己前面的没有必要浪费
3. 但是如何分配分数，也是挺麻烦的
4. 如果要获得最高排名，自己肯定要获得最高的分数，然后在这个分数上面的部分（已经在自己前面的，尽量的分配）不改变自己的排名
5. 然后离自己越远的分配越大的分数
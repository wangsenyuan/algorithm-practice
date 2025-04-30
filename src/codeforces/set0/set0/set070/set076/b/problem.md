Modern researches has shown that a flock of hungry mice searching for a piece of cheese acts as follows: if there are several pieces of cheese then each mouse chooses the closest one. After that all mice start moving towards the chosen piece of cheese. When a mouse or several mice achieve the destination point and there is still a piece of cheese in it, they eat it and become well-fed. Each mice that reaches this point after that remains hungry. Moving speeds of all mice are equal.

If there are several ways to choose closest pieces then mice will choose it in a way that would minimize the number of hungry mice. To check this theory scientists decided to conduct an experiment. They located N mice and M pieces of cheese on a cartesian plane where all mice are located on the line y = Y0 and all pieces of cheese — on another line y = Y1. To check the results of the experiment the scientists need a program which simulates the behavior of a flock of hungry mice.

Write a program that computes the minimal number of mice which will remain hungry, i.e. without cheese.

### ideas
1. 对于任何一个mice，最近的chess，最多有两个，前面/后面
2. 
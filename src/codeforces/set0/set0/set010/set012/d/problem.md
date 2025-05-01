N ladies attend the ball in the King's palace. Every lady can be described with three values: beauty, intellect and richness. King's Master of Ceremonies knows that ladies are very special creatures. If some lady understands that there is other lady at the ball which is more beautiful, smarter and more rich, she can jump out of the window. He knows values of all ladies and wants to find out how many probable self-murderers will be on the ball. Lets denote beauty of the i-th lady by Bi, her intellect by Ii and her richness by Ri. Then i-th lady is a probable self-murderer if there is some j-th lady that Bi < Bj, Ii < Ij, Ri < Rj. Find the number of probable self-murderers.

### ideas
1. 按照B降序排，然后维护一个结构，如果能够查到某个I对应的最大的R是多少，那么对于当前的lady，如果存在I+1的R比她大，那么+1
2. 所以，range query + point update
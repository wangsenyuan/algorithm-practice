One must train much to do well on wizardry contests. So, there are numerous wizardry schools and magic fees.

One of such magic schools consists of n tours. A winner of each tour gets a huge prize. The school is organised quite far away, so one will have to take all the prizes home in one go. And the bags that you've brought with you have space for no more than k huge prizes.

Besides the fact that you want to take all the prizes home, you also want to perform well. You will consider your performance good if you win at least l tours.

In fact, years of organizing contests proved to the organizers that transporting huge prizes is an issue for the participants. Alas, no one has ever invented a spell that would shrink the prizes... So, here's the solution: for some tours the winner gets a bag instead of a huge prize. Each bag is characterized by number ai — the number of huge prizes that will fit into it.

You already know the subject of all tours, so you can estimate the probability pi of winning the i-th tour. You cannot skip the tour under any circumstances.

Find the probability that you will perform well on the contest and will be able to take all won prizes home (that is, that you will be able to fit all the huge prizes that you won into the bags that you either won or brought from home).

### ideas
1. dp[i][j][x] 表示在前i个项目中，其中j个项目获胜，且得到了x个bag的概率
On a cold winter evening our hero Vasya stood in a railway queue to buy a ticket for Codeforces championship final. As it usually happens, the cashier said he was going to be away for 5 minutes and left for an hour. Then Vasya, not to get bored, started to analyze such a mechanism as a queue. The findings astonished Vasya.

Every man is characterized by two numbers: ai, which is the importance of his current task (the greater the number is, the more important the task is) and number ci, which is a picture of his conscience. Numbers ai form the permutation of numbers from 1 to n.

Let the queue consist of n - 1 people at the moment. Let's look at the way the person who came number n behaves. First, he stands at the end of the queue and the does the following: if importance of the task ai of the man in front of him is less than an, they swap their places (it looks like this: the man number n asks the one before him: "Erm... Excuse me please but it's very important for me... could you please let me move up the queue?"), then he again poses the question to the man in front of him and so on. But in case when ai is greater than an, moving up the queue stops. However, the man number n can perform the operation no more than cn times.

In our task let us suppose that by the moment when the man number n joins the queue, the process of swaps between n - 1 will have stopped. If the swap is possible it necessarily takes place.

Your task is to help Vasya model the described process and find the order in which the people will stand in queue when all the swaps stops.
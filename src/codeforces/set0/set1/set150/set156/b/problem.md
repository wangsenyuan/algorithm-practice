# Problem Statement

As Sherlock Holmes was investigating a crime, he identified $n$ suspects. He knows for sure that exactly one of them committed the crime. To find out which one did it, the detective lines up the suspects and numbers them from $1$ to $n$. After that, he asked each one: "Which one committed the crime?" Suspect number $i$ answered either "The crime was committed by suspect number $a_i$", or "Suspect number $a_i$ didn't commit the crime". Also, the suspect could say so about himself ($a_i = i$).

Sherlock Holmes understood for sure that exactly $m$ answers were the truth and all other answers were a lie. Now help him understand this: which suspect lied and which one told the truth?

## Input

The first line contains two integers $n$ and $m$ ($1 \leq n \leq 10^5$, $0 \leq m \leq n$) — the total number of suspects and the number of suspects who told the truth. Next $n$ lines contain the suspects' answers. The $i$-th line contains either `+a_i` (without the quotes), if the suspect number $i$ says that the crime was committed by suspect number $a_i$, or `-a_i` (without the quotes), if the suspect number $i$ says that the suspect number $a_i$ didn't commit the crime ($a_i$ is an integer, $1 \leq a_i \leq n$).

It is guaranteed that at least one suspect exists, such that if he committed the crime, then exactly $m$ people told the truth.

## Output

Print $n$ lines. Line number $i$ should contain `Truth` if suspect number $i$ has told the truth for sure. Print `Lie` if the suspect number $i$ lied for sure and print `Not defined` if he could lie and could tell the truth, too, depending on who committed the crime.

### ideas
1. 假设i是罪犯（只有一个罪犯）
2. 那么所有说它不是罪犯的肯定在说谎
3. 所有所它是罪犯的就是在说真话，以及说有那些说其他人（j != i)不是罪犯的人，在说真话
4. 如果只有一个i满足这个条件，那么就可以找到答案
5. 如果有两个i满足这个条件，那么就是not defined
6. 还有一个，假设嫌疑犯在一个范围内，所有指向这个范围的，属于not defined，否则就是lie
7. 针对i的真话，包括那些说它是嫌犯的，还包括那些说j(j != i)的是嫌犯的
8. 如果不等于m，那么它肯定不是嫌犯（所有那些说它是嫌犯的，可以直接判定为lie）
9. 如果等于m，那么待定
10. 
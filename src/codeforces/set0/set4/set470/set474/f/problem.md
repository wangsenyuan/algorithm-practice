Mole is hungry again. He found one ant colony, consisting of n ants, ordered in a row. Each ant i (1 ≤ i ≤ n) has a strength si.

In order to make his dinner more interesting, Mole organizes a version of «Hunger Games» for the ants. He chooses two numbers l and r (1 ≤ l ≤ r ≤ n) and each pair of ants with indices between l and r (inclusively) will fight. When two ants i and j fight, ant i gets one battle point only if si divides sj (also, ant j gets one battle point only if sj divides si).

After all fights have been finished, Mole makes the ranking. An ant i, with vi battle points obtained, is going to be freed only if vi = r - l, or in other words only if it took a point in every fight it participated. After that, Mole eats the rest of the ants. Note that there can be many ants freed or even none.

In order to choose the best sequence, Mole gives you t segments [li, ri] and asks for each of them how many ants is he going to eat if those ants fight.
Yet another education system reform has been carried out in Berland recently. The innovations are as follows:

An academic year now consists of n days. Each day pupils study exactly one of m subjects, besides, each subject is studied for no more than one day. After the lessons of the i-th subject pupils get the home task that contains no less than ai and no more than bi exercises. Besides, each subject has a special attribute, the complexity (ci). A school can make its own timetable, considering the following conditions are satisfied:

the timetable should contain the subjects in the order of the complexity's strict increasing;
each day, except for the first one, the task should contain either k times more exercises, or more by k compared to the previous day (more formally: let's call the number of home task exercises in the i-th day as xi, then for each i (1 < i ≤ n): either xi = k + xi - 1 or xi = k·xi - 1 must be true);
the total number of exercises in all home tasks should be maximal possible.
All limitations are separately set for each school.

It turned out that in many cases ai and bi reach 1016 (however, as the Berland Minister of Education is famous for his love to half-measures, the value of bi - ai doesn't exceed 100). That also happened in the Berland School №256. Nevertheless, you as the school's principal still have to work out the timetable for the next academic year...

### ideas
1. 按照ci升序排列
2. 不考虑j的限制，dp[i][j]表示到i为止，在第i个task上，做j个exercises时的最多
3. dp[i][j] = dp[x][j - k] + j or dp[x][j / k] + j
4. j = a[i] + v
5. 变成 dp[i][v], 那么v的范围就是100 
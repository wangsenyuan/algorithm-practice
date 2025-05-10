Dr. Moriarty is about to send a message to Sherlock Holmes. He has a string s.

String p is called a substring of string s if you can read it starting from some position in the string s. For example, string "aba" has six substrings: "a", "b", "a", "ab", "ba", "aba".

Dr. Moriarty plans to take string s and cut out some substring from it, let's call it t. Then he needs to change the substring t zero or more times. As a result, he should obtain a fixed string u (which is the string that should be sent to Sherlock Holmes). One change is defined as making one of the following actions:

Insert one letter to any end of the string.
Delete one letter from any end of the string.
Change one letter into any other one.
Moriarty is very smart and after he chooses some substring t, he always makes the minimal number of changes to obtain u.

Help Moriarty choose the best substring t from all substrings of the string s. The substring t should minimize the number of changes Moriarty should make to obtain the string u from it.


### ideas
1. 对于给定的t，要得到u，= len(u) - 那些可以保留的部分（不变的）= u和t的公共部门
2. 如果是s的子串，删除和插入没有意义，直接使用替换就可以了
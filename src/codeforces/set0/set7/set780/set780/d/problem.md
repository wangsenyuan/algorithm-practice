Innokenty is a president of a new football league in Byteland. The first task he should do is to assign short names to all clubs to be shown on TV next to the score. Of course, the short names should be distinct, and Innokenty wants that all short names consist of three letters.

Each club's full name consist of two words: the team's name and the hometown's name, for example, "DINAMO BYTECITY". Innokenty doesn't want to assign strange short names, so he wants to choose such short names for each club that:

the short name is the same as three first letters of the team's name, for example, for the mentioned club it is "DIN",
or, the first two letters of the short name should be the same as the first two letters of the team's name, while the third letter is the same as the first letter in the hometown's name. For the mentioned club it is "DIB".
Apart from this, there is a rule that if for some club x the second option of short name is chosen, then there should be no club, for which the first option is chosen which is the same as the first option for the club x. For example, if the above mentioned club has short name "DIB", then no club for which the first option is chosen can have short name equal to "DIN". However, it is possible that some club have short name "DIN", where "DI" are the first two letters of the team's name, and "N" is the first letter of hometown's name. Of course, no two teams can have the same short name.

Help Innokenty to choose a short name for each of the teams. If this is impossible, report that. If there are multiple answer, any of them will suit Innokenty. If for some team the two options of short name are equal, then Innokenty will formally think that only one of these options is chosen.

### ideas
1. 对于队伍u，按照规则1取短名称是true，规则2为false
2. 如果队伍u，按照规则2（也就是取false时），所有其他的队伍，如果按照规则1，和u有相同的名称（规则1的名称）时，它们必须也按照规则2
3. 还有一个点时，如何保证最后大家的名称都不一样？
4. 按照规则1取名时，所有按照规则1，规则2相同的，都不能取
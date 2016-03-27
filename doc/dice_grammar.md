Definition of the Dices Grammar
===============================

Basic elements
--------------

digit excluding zero = "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;

digit = "0" | digit excluding zero ;

dice = "d"

explode = "e"

keep = "k"

open = "o"

reroll = "r"

success = "s"

explodingSuccess = "es"

modifier =  discard natural | explode  | explodingSuccess natural | open | reroll  natural| success natural

natural  = digit excluding zero, { digit } ;

integer = "0" | [ "-" ], natural ;

numDices = natural

numSides = natural

constant = Integer

basicDiceExpression = [numDices,] dice, numSides, [modifier]

expressionOperators = "+"| "-" | "*" | "/"


ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .

References
----------

* https://github.com/Bernardo-MG/tabletop-dice-java/blob/master/src/main/antlr4/DiceNotationExtended.g4
* https://github.com/Bernardo-MG/tabletop-dice-java/blob/master/src/main/antlr4/DiceNotation.g4
* https://en.wikipedia.org/wiki/Dice_notation
* https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_Form
* https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_Form

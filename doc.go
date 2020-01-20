/*Package scalc provides a basic sets calculator that parses an expression string.

Grammar of calculator is given:
expression := “[“ operator sets “]”
sets := set | set sets
set := file | expression
operator := “SUM” | “INT” | “DIF”

Each file should contain sorted integers, one integer in a line.
Meaning of operators:
SUM - returns union of all sets
INT - returns intersection of all sets
DIF - returns difference of first set and the rest ones
*/
package scalc

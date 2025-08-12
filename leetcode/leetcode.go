package main

import (
	"fmt"
)

func main() {
	var t, i uint
	t, i = 1, 1

	for i = 1; i < 10; i++ {
		fmt.Printf("%d << %d = %d \n", t, i, t<<i)
	}

	fmt.Println()

	t = 512
	for i = 1; i < 10; i++ {
		fmt.Printf("%d >> %d = %d \n", t, i, t>>i)
	}

}

func myPow(x float64, n int) float64 {
	isNeg := false
	var exp int
	if n < 0 {
		isNeg = true
		exp = n * -1
	} else {
		exp = n
	}
	var res float64
	res = 1
	for i := 1; i <= exp; i++ {
		res = x * res
	}
	if isNeg {
		return (1 / res)
	}
	return res
}

func powHelper(num float64, exp int) float64 {
	if exp < 0 {
		return (1 / num) * powHelper(num, exp+1)
	} else if exp > 0 {
		return num * powHelper(num, exp-1)
	} else {
		return 1
	}
}

/*
We will use two stacks:

Operand stack: to keep values (numbers)  and

Operator stack: to keep operators (+, -, *, . and ^).


In the following, “process” means, (i) pop operand stack once (value1) (ii) pop operator stack once (operator) (iii) pop operand stack again (value2) (iv) compute value1 operator  value2 (v) push the value obtained in operand stack.


Algorithm:


Until the end of the expression is reached, get one character and perform only one of the steps (a) through (f):

	(a) If the character is an operand, push it onto the operand stack.

	(b) If the character is an operator, and the operator stack is empty then push it onto the operator stack.

	(c) If the character is an operator and the operator stack is not empty, and the character's precedence is greater than the precedence of the stack top of operator stack, then push the character onto the operator stack.

	(d) If the character is "(", then push it onto operator stack.

	(e) If the character is ")", then "process" as explained above until the corresponding "(" is encountered in operator stack.  At this stage POP the operator stack and ignore "(."

	(f) If cases (a), (b), (c), (d) and (e) do not apply, then process as explained above.


	 When there are no more input characters, keep processing until the operator stack becomes empty.  The values left in the operand stack is the final result of the expression.

	 (1+2)*3+(4*2) = 17

	 Operand       Operator
         8           +
			         9
*/

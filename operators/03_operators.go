package operators

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	fmt.Println("Addition: ", a+b)
	fmt.Println("Subtraction: ", a-b)
	fmt.Println("Multiplication: ", a*b)
	fmt.Println("Division: ", a/b)
	fmt.Println("Modulus: ", a%b)

	// Increment and Decrement
	a++
	b--
	fmt.Println("Incremented a: ", a)
	fmt.Println("Decremented b: ", b)

	// Comparison Operators
	fmt.Println("a == b: ", a == b)
	fmt.Println("a != b: ", a != b)
	fmt.Println("a < b: ", a < b)
	fmt.Println("a > b: ", a > b)
	fmt.Println("a <= b: ", a <= b)
	fmt.Println("a >= b: ", a >= b)

	// Logical Operators
	fmt.Println("a > 5 && b < 30: ", a > 5 && b < 30)
	fmt.Println("a > 5 || b < 30: ", a > 5 || b < 30)
	fmt.Println("!(a > 5): ", !(a > 5))

	// Bitwise Operators
	var x uint = 5
	var y uint = 3
	fmt.Println("Bitwise AND: ", x&y)
	fmt.Println("Bitwise OR: ", x|y)
	fmt.Println("Bitwise XOR: ", x^y)
	fmt.Println("Bitwise AND NOT: ", x&^y)
	fmt.Println("Left Shift: ", x<<1)
	fmt.Println("Right Shift: ", x>>1)

}

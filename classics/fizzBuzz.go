package classics

import "fmt"

/*
Write a short program that prints each number from 1 to 100 on a new line.

For each multiple of 3, print "Fizz" instead of the number.

For each multiple of 5, print "Buzz" instead of the number.

For numbers which are multiples of both 3 and 5, print "FizzBuzz" instead of the number.

Solution:
For loop from 1 to 100. Each iteration check if number mod 3 and 5 or number mod 3 or number mod 5.
Print number otherwise.
 */
func fizzBuzz() {
	for i := 1; i < 101; i += 1 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}


package warmup

/*
You are choreographing a circus show with various animals. For one act, you are given two kangaroos on a number line ready to jump in the positive direction (i.e, toward positive infinity).

The first kangaroo starts at location  and moves at a rate of  meters per jump.
The second kangaroo starts at location  and moves at a rate of  meters per jump.
You have to figure out a way to get both kangaroos at the same location at the same time as part of the show. If it is possible, return YES, otherwise return NO.

For example, kangaroo  starts at  with a jump distance  and kangaroo  starts at  with a jump distance of . After one jump, they are both at , (, ), so our answer is YES.

Function Description

Complete the function kangaroo in the editor below. It should return YES if they reach the same position at the same time, or NO if they don't.

kangaroo has the following parameter(s):

x1, v1: integers, starting position and jump distance for kangaroo 1
x2, v2: integers, starting position and jump distance for kangaroo 2
Input Format

A single line of four space-separated integers denoting the respective values of , , , and .

Constraints

Output Format

Print YES if they can land on the same location at the same time; otherwise, print NO.

Note: The two kangaroos must land at the same location after making the same number of jumps.

Sample Input 0

0 3 4 2
Sample Output 0

YES
 */

// Solution: As x2 > x1 we only need to check v2 > v1. In that case we return false.
// Then we need to  find if the equation exists x1 + y * v1 = x2 + y * v2 with y an integer.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v2 > v1 {
		return "NO"
	}
	if v1 > v2 && ((x2-x1) % (v1-v2)) == 0 {
		return "YES"
	}
	return "NO"
}

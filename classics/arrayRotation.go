package classics

/*
A left rotation operation on an array shifts each of the array's elements
unit to the left. For example, if  left rotations are performed on array ,
then the array would become .

Given an array  of  integers and a number, , perform  left rotations on the array.
Return the updated array to be printed as a single line of space-separated integers.

Function Description

Complete the function rotLeft in the editor below. It should return the resulting
array of integers.

rotLeft has the following parameter(s):

An array of integers .
An integer , the number of rotations.
Input Format

The first line contains two space-separated integers  and , the size of
and the number of left rotations you must perform.
The second line contains  space-separated integers .

Constraints

Output Format

Print a single line of  space-separated integers denoting the final
state of the array after performing  left rotations.

Sample Input

5 4
1 2 3 4 5
Sample Output

5 1 2 3 4
Explanation

When we perform  left rotations, the array undergoes the following sequence of changes:
*/
func RotLeft(a []int32, d int32) []int32 {
	size1 := int32(len(a))
	if size1 > d {
		first := a[:d]
		rest := a[d:]
		res := append(rest, first...)
		return res
	}
	div := d % size1
	first := a[:div]
	rest := a[div:]
	res := append(rest, first...)
	return res
}

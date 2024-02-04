/*

Leetcode Problem 9: Palindrome Number

https://leetcode.com/problems/palindrome-number/

*/

package leetcode


func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversed := 0

	for num := x; num > 0; num /= 10 {
		reversed = reversed*10 + num%10
	}

	return x == reversed
}

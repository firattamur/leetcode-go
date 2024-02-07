package p0020

/*

Leetcode Problem 20: Valid Parentheses

https://leetcode.com/problems/valid-parentheses/

*/

func IsValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

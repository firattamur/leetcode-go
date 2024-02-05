/*

Leetcode Problem 13: Roman to Integer

https://leetcode.com/problems/roman-to-integer/

*/

package leetcode

func RomanToInt(s string) int {
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	result, lastValue := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[s[i]]

		if value < lastValue {
			result -= value
		} else {
			result += value
		}
		lastValue = value

	}

	return result
}

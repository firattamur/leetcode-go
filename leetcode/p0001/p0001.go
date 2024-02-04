/*

Leetcode Problem 1: Two Sum
https://leetcode.com/problems/two-sum/

*/

package leetcode

func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for idx, num := range nums {
		j, exists := numsMap[target-num]
		if exists {
			return []int{j, idx}
		}

		numsMap[num] = idx
	}

	return nil
}

package p0014

/*

Leetcode Problem 14: Longest Common Prefix

https://leetcode.com/problems/longest-common-prefix/

*/

func LongestCommonPrefix(strs []string) string {
	result := ""

	if len(strs) == 0 {
		return result
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}

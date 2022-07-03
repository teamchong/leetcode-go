//Given an array of strings strs, group the anagrams together. You can return
//the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a
//different word or phrase, typically using all the original letters exactly once.
//
//
// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:
// Input: strs = [""]
//Output: [[""]]
// Example 3:
// Input: strs = ["a"]
//Output: [["a"]]
//
//
// Constraints:
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.
//
// Related Topics Array Hash Table String Sorting 👍 10054 👎 337

package main

//leetcode submit region begin(Prohibit modification and deletion)
import "sort"

func groupAnagrams(strs []string) [][]string {
	sortedToGroup := make(map[string][]string, len(strs))
	for _, str := range strs {
		runes := sortRunes(str)
		sort.Sort(runes)
		sortedToGroup[string(runes)] = append(sortedToGroup[string(runes)], str)
	}
	groups := make([][]string, 0, len(sortedToGroup))
	for _, group := range sortedToGroup {
		groups = append(groups, group)
	}
	return groups
}

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Less(i, j int) bool { return s[i] > s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//leetcode submit region end(Prohibit modification and deletion)

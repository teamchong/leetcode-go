/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
import "sort"

// We sort each string by rune, and then use the sorted string as a key in a map
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	// Sorting each string by rune, and then using the sorted string as a key in a map.
	for _, str := range strs {
		r := ByRune(str)
		sort.Sort(r)
		m[string(r)] = append(m[string(r)], str)
	}
	// Converting the map into a slice of slices.
	groups := make([][]string, 0)
	for _, group := range m {
		groups = append(groups, group)
	}
	return groups
}

type ByRune []rune

// A method of the type ByRune. It returns the length of the ByRune type.
func (s ByRune) Len() int { return len(s) }

// Sorting the string in reverse order.
func (s ByRune) Less(i, j int) bool { return s[i] > s[j] }

// Swapping the values of the two indexes.
func (s ByRune) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// @lc code=end


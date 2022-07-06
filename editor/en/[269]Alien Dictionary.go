//There is a new alien language that uses the English alphabet. However, the
//order among the letters is unknown to you.
//
// You are given a list of strings words from the alien language's dictionary,
//where the strings in words are sorted lexicographically by the rules of this new
//language.
//
// Return a string of the unique letters in the new alien language sorted in
//lexicographically increasing order by the new language's rules. If there is no
//solution, return "". If there are multiple solutions, return any of them.
//
// A string s is lexicographically smaller than a string t if at the first
//letter where they differ, the letter in s comes before the letter in t in the alien
//language. If the first min(s.length, t.length) letters are the same, then s is
//smaller if and only if s.length < t.length.
//
//
// Example 1:
//
//
//Input: words = ["wrt","wrf","er","ett","rftt"]
//Output: "wertf"
//
//
// Example 2:
//
//
//Input: words = ["z","x"]
//Output: "zx"
//
//
// Example 3:
//
//
//Input: words = ["z","x","z"]
//Output: ""
//Explanation: The order is invalid, so return "".
//
//
//
// Constraints:
//
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] consists of only lowercase English letters.
//
// Related Topics Array String Depth-First Search Breadth-First Search Graph
//Topological Sort 👍 3694 👎 767

package main

//leetcode submit region begin(Prohibit modification and deletion)
import (
	"strings"
)

func alienOrder(words []string) string {
	// Step 0: Create data structures and find all unique letters.
	adjList := make(map[rune][]rune)
	counts := make(map[rune]int)

	for _, word := range words {
		for _, c := range word {
			counts[c] = 0
			adjList[c] = make([]rune, 0)
		}
	}

	// Step 1: Find all edges.
	for i, word1 := range words[:len(words)-1] {
		word2 := words[i+1]
		len1, len2 := len(word1), len(word2)

		// Check that word2 is not a prefix of word1.
		if len1 > len2 && strings.HasPrefix(word1, word2) {
			return ""
		}

		// Find the first non match and insert the corresponding relation.
		var min int
		if min = len1; min > len2 {
			min = len2
		}
		for j, ch1 := range word1[:min] {
			if ch2 := rune(word2[j]); ch1 != ch2 {
				adjList[ch1] = append(adjList[ch1], ch2)
				counts[ch2]++
				break
			}
		}
	}

	// Step 2: Breadth-first search.
	queue := make(queue, 0)
	for c, count := range counts {
		if count == 0 {
			queue.enqueue(c)
		}
	}
	var sb strings.Builder
	for len(queue) > 0 {
		c := queue.dequeue()
		sb.WriteRune(c)
		for _, next := range adjList[c] {
			if counts[next]--; counts[next] == 0 {
				queue.enqueue(next)
			}
		}
	}

	if sb.Len() < len(counts) {
		return ""
	}
	return sb.String()
}

type queue []rune

func (s *queue) enqueue(item rune) {
	*s = append(*s, item)
}
func (s *queue) dequeue() rune {
	c := (*s)[0]
	*s = (*s)[1:]
	return c
}

//leetcode submit region end(Prohibit modification and deletion)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
import "container/heap"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		pq := PriorityQueue(str)
		heap.Init(&pq)
		key := string(pg)
		m[key] = append(m[key], str)
	}
	groups := make([][]string, 0)
	for _, group := range m {
		groups = append(groups, group)
	}
	return groups
}

type PriorityQueue []rune

func (pq PriorityQueue) Len() int          		 { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool 	 { return pq[i] > pq[j] }
func (pq PriorityQueue) Swap(i, j int)  	     { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(value interface)   { *pg = append(*pg, value.(rune)) }
func (pq *PriorityQueue) Pop() interface{} {
	last := *pg[len(*pg)-1]
	*pg = (*pg)[:len(*pg)-1]
	return last
}

// @lc code=end


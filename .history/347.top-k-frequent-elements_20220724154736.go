/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	numToFreq := make(map[int]int, len(nums))
	for _, num := range nums {
		numToFreq[num]++
	}

	f := &NumFreqs{}
	heap.Init(f)
	for num, freq := range numToFreq {
		heap.Push(f, NumFreq{num, freq})
	}
	mostFrequent := make([]int, 0, k)
	for i := 0; i < k; i++ {
		if last, ok := heap.Pop(f).(NumFreq); ok {
			mostFrequent = append(mostFrequent, last.num)
		} else {
			break
		}
	}
	return mostFrequent
}

type NumFreq = struct {
	num  int
	freq int
}
type NumFreqs []NumFreq

func (f NumFreqs) Len() int            { return len(f) }
func (f NumFreqs) Less(i, j int) bool  { return f[i].freq > f[j].freq }
func (f NumFreqs) Swap(i, j int)       { f[i], f[j] = f[j], f[i] }
func (f *NumFreqs) Push(x interface{}) { *f = append(*f, x.(NumFreq)) }
func (f *NumFreqs) Pop() interface{} {
	last := (*f)[len(*f)-1]
	*f = (*f)[:len(*f)-1]
	return last
}

// @lc code=end

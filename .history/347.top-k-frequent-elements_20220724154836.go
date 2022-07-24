/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
package main

import "container/heap"

// We use a heap to keep track of the k most frequent numbers
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

// NumFreq is a struct with two fields, num and freq, both of type int.
// @property {int} num - The number that is being counted.
// @property {int} freq - The frequency of the number.
type NumFreq = struct {
	num  int
	freq int
}
type NumFreqs []NumFreq

// A method that is part of the `NumFreqs` type. It is a method that returns the length of the
// `NumFreqs` type.
func (f NumFreqs) Len() int { return len(f) }

// A method that is part of the `NumFreqs` type. It is a method that returns a boolean value
// that is true if the frequency of the number at index i is greater than the frequency of the number
// at
// index j.
func (f NumFreqs) Less(i, j int) bool { return f[i].freq > f[j].freq }

// Swapping the elements at index i and j.
func (f NumFreqs) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

// A method that is part of the `NumFreqs` type. It is a method that appends the `NumFreq` type to the
// `NumFreqs` type.
func (f *NumFreqs) Push(x interface{}) { *f = append(*f, x.(NumFreq)) }

// A method that is part of the `NumFreqs` type. It is a method that returns the last element of the
// `NumFreqs` type.
func (f *NumFreqs) Pop() interface{} {
	last := (*f)[len(*f)-1]
	*f = (*f)[:len(*f)-1]
	return last
}

// @lc code=end

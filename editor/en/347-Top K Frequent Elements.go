//Given an integer array nums and an integer k, return the k most frequent
//elements. You may return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
// Example 2:
// Input: nums = [1], k = 1
//Output: [1]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.
//
//
//
// Follow up: Your algorithm's time complexity must be better than O(n log n),
//where n is the array's size.
// Related Topics Array Hash Table Divide and Conquer Sorting Heap (Priority
//Queue) Bucket Sort Counting Quickselect 👍 9722 👎 384

package main

//leetcode submit region begin(Prohibit modification and deletion)
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

//leetcode submit region end(Prohibit modification and deletion)

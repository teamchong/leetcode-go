/*
 * @lc app=leetcode id=23 lang=rust
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    /// It takes a vector of linked lists and merges them into one linked list.
    /// 
    /// Arguments:
    /// 
    /// * `list`: Vec<Option<Box<ListNode>>>
    /// 
    /// Returns:
    /// 
    /// A `Box<ListNode>`
    pub fn merge_k_lists(mut list: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        let i = list.iter().enumerate().min_by_key(|(_, x)| x.as_ref().map_or(std::i32::MAX, |x| x.val))?.0;
        let mut h = list[i].take()?;
        list[i] = h.next;
        h.next = Self::merge_k_lists(list);
        Some(h)
    }
}
// @lc code=end


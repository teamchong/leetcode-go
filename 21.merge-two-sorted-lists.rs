/*
 * @lc app=leetcode id=21 lang=rust
 *
 * [21] Merge Two Sorted Lists
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
    pub fn merge_two_lists(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let (mut list1, mut list2, mut head) = (&list1, &list2, None);
        let mut current = &mut head;
        loop {
            match (list1, list2) {
                (None, None) => {
                    *current = None;
                    break head;
                }
                (Some(n), None) | (None, Some(n)) => {
                    *current = Some(n.clone());
                    break head;
                }
                (Some(n1), Some(n2)) => {
                    if n1.val < n2.val {
                        *current = Some(n1.clone());
                        list1 = &n1.next;
                    } else {
                        *current =  Some(n2.clone());
                        list2 = &n2.next;
                    }
                    current = &mut current.as_mut().unwrap().next;
                }
            }
        }
    }
}
// @lc code=end


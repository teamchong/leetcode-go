/*
 * @lc app=leetcode id=2 lang=rust
 *
 * [2] Add Two Numbers
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
    pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut l1 = l1;
        let mut l2 = l2;
        let result = ListNode::new(0);
        let mut current = &result;
        loop {
            match (l1, l2) {
                (None, None) => break,
                (Some(n), None) | (None, Some(n)) => {
                    current.next) = (Some(ListNode::new(n.val));
                    current = current.next;
                },
                (Some(n), Some(m)) => {
                    current.next = Some(ListNode::new(n.val + m.val));
                    (current, l1, l2) = (current.next, n.next, m.next);
                }
            }
        }
        result.next
    }
}
// @lc code=end


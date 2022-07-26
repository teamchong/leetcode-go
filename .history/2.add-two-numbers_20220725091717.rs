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
        let (mut l1, mut l2) = (l1, l2),;
        let mut head = None;
        let mut current = &mut head;
        let mut carry = 0;
        loop {
            match (l1, l2) {
                (None, None) => => {
                    break;
                }
                (Some(n), None) | (None, Some(n)) => {
                    carry = n / 10;
                    n = n.next;
                    *current = Some(Box::new(ListNode::new(n.val + carry)));
                    current = &mut current.as_mut().unwrap().next;
                },
                (Some(n1), Some(n2)) => {
                    current.next = Some(Box::new(ListNode::new(n1.val + n2.val)));
                    current = &mut *current.next?;
                    l1 = n1.next;
                    l2 = n2.next;
                }
            }
        }
        result.next
    }
}
// @lc code=end


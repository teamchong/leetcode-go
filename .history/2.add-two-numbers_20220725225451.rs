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
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let (mut l1, mut l2, mut head, mut carry) = (&l1, &l2, None, 0);
        let mut current = &mut head;
        loop {
            match (l1, l2) {
                (None, None) => {
                    if carry > 0 {
                        *current = Some(Box::new(ListNode::new(carry)));
                    }
                    break;
                }
                (Some(n1), None) => {
                    *current = Some(Box::new(ListNode::new(n1.val + carry)));
                    l1 = &n1.next;
                }
                (None, Some(n2)) => {
                    *current = Some(Box::new(ListNode::new(n2.val + carry)));
                    l2 = &n2.next;
                }
                (Some(n1), Some(n2)) => {
                    *current = Some(Box::new(ListNode::new(n1.val + n2.val + carry)));
                    l1 = &n1.next;
                    l2 = &n2.next;
                }
            }
            carry = sum / 10;
            *current = Some(Box::new(ListNode::new(sum % 10)));
            current = &mut current.as_mut().unwrap().next;
        }
        head
    }
}
// @lc code=end

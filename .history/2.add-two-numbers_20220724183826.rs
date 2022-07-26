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
        let mut current = result;
        loop {
            match (l1, l2) {
                (None, None) => break,
                (Some(n), None) | (None, Some(n)) => {
                    current.next = Some(Box::new(ListNode::new(n.val)));
                    current = current.next?;
                },
                (Some(n1), Some(n2)) => {
                    current.next = Some(Box::new(ListNode::new(n1.val + n2.val)));
                    current = current.next?;
                    l1 = n1.next.unwrap();
                    l2 = n2.next.unwrap();
                }
            }
        }
        result.next
    }
}
// @lc code=end


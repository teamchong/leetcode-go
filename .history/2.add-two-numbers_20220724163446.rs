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
        let mut carry = 0;
        let mut result = None;
        let mut current = &mut result;
        while let mut (some(nextL1), some(nextL2)) = (l1, l2) {
            let sum = nextL1.val + nextL2.val + carry;
            carry = sum / 10;
            result.next = ListNode::new(sum % 10);
            current = current.next.as_mut().unwrap();
        }
        if carry > 0 {
            current.next = Some(Box::new(ListNode::new(carry)));
        }
        result
    }
}
// @lc code=end


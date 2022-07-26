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
//
// Constraints:
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.
impl Solution {
    /// We create a mutable reference to the head of the list, and a mutable reference to the carry. We
    /// loop through the list, and for each iteration, we add the carry to the sum, and then we match on
    /// the current node of each list. If both lists are None, we break out of the loop. If one list is
    /// None, we add the carry to the other list, and then we update the carry and the current node of
    /// the list. If both lists are not None, we add the carry to the sum of the two lists, and then we
    /// update the carry and the current node of the list
    /// 
    /// Arguments:
    /// 
    /// * `l1`: Option<Box<ListNode>>, l2: Option<Box<ListNode>>
    /// * `l2`: Option<Box<ListNode>>
    /// 
    /// Returns:
    /// 
    /// A linked list.
    pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let (mut l1, mut l2, mut head, mut carry) = (&l1, &l2, None, 0);
        let mut current = &mut head;
        loop {
            match (l1, l2) {
                (None, None) => {
                    if carry > 0 {
                      *current = Some(Box::new(ListNode::new(carry)));
                    }
                    break head;
                }
                (Some(n1), None) => {
                    *current = Some(Box::new(ListNode::new((n1.val + carry) % 10)));
                    carry = (n1.val + carry) / 10;
                    l1 = &n1.next;
                },
                (None, Some(n2)) => {
                    *current = Some(Box::new(ListNode::new((n2.val + carry) % 10)));
                    carry = (n2.val + carry) / 10;
                    l2 = &n2.next;
                },
                (Some(n1), Some(n2)) => {
                    *current = Some(Box::new(ListNode::new((n1.val + n2.val + carry) % 10)));
                    carry = (n1.val + n2.val + carry) / 10;
                    l1 = &n1.next;
                    l2 = &n2.next;
                }
            }
            current = &mut current.unwrap().next;
        }
    }
}
// @lc code=end


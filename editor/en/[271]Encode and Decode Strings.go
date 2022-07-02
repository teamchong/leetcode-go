//Design an algorithm to encode a list of strings to a string. The encoded
//string is then sent over the network and is decoded back to the original list of
//strings.
//
// Machine 1 (sender) has the function:
//
//
//string encode(vector<string> strs) {
//  // ... your code
//  return encoded_string;
//}
//Machine 2 (receiver) has the function:
//
//
//vector<string> decode(string s) {
//  //... your code
//  return strs;
//}
//
//
// So Machine 1 does:
//
//
//string encoded_string = encode(strs);
//
//
// and Machine 2 does:
//
//
//vector<string> strs2 = decode(encoded_string);
//
//
// strs2 in Machine 2 should be the same as strs in Machine 1.
//
// Implement the encode and decode methods.
//
// You are not allowed to solve the problem using any serialize methods (such
//as eval).
//
//
// Example 1:
//
//
//Input: dummy_input = ["Hello","World"]
//Output: ["Hello","World"]
//Explanation:
//Machine 1:
//Codec encoder = new Codec();
//String msg = encoder.encode(strs);
//Machine 1 ---msg---> Machine 2
//
//Machine 2:
//Codec decoder = new Codec();
//String[] strs = decoder.decode(msg);
//
//
// Example 2:
//
//
//Input: dummy_input = [""]
//Output: [""]
//
//
//
// Constraints:
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] contains any possible characters out of 256 valid ASCII characters.
//
//
//
// Follow up: Could you write a generalized algorithm to work on any possible
//set of characters?
// Related Topics Array String Design 👍 889 👎 242

package main

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(fmt.Sprintf("%02x", len(str)))
		b.WriteString(str)
	}
	return b.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	output := make([]string, 0)
	for l, ok := strToInt(strs); ok; l, ok = strToInt(strs) {
		output = append(output, strs[2:2+l])
		strs = strs[2+l:]
	}
	return output
}

func strToInt(str string) (int, bool) {
	if len(str) < 2 {
		return 0, false
	} else if i, err := strconv.ParseInt(str[:2], 16, 32); err == nil {
		return int(i), true
	}
	return 0, false
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
//leetcode submit region end(Prohibit modification and deletion)

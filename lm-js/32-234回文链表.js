
// 234. 回文链表
// 请判断一个链表是否为回文链表。

// 示例 1:

// 输入: 1->2
// 输出: false
// 示例 2:

// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function (head) {
    //利用数组解决这一问题的
    let w = 0, cur = head, arr = []
    //遍历链表，将value存进数组
    while (cur) {
        w++
        arr.push(cur.val)
        cur = cur.next
    }
    //如果数组长度小于2,则为回文链表
    if (w < 2) {
        return true
    }
    let start = 0, end = arr.length - 1
    //双指针，判断头尾指针是否相等
    while (start >= 0 && end < w) {
        //如果有不相等的直接返回false
        if (arr[start] !== arr[end]) {
            return false
        }
        else {
            //如果数组已经比较结束，至中间位置，则返回true
            if (parseInt(w / 2) == end+1 || parseInt(w / 2) == start+1) {
                return true
            }
            start++;
            end--;
        }
    }
};

  /** 题解
利用数组头尾指针来解决。遍历一遍先将链表的value存储到数组中，然后遍历数组的头尾指针。
当头或尾某一个先跑至中间位置，则遍历结束，最后一次依然头尾指针相等，则返回true，
复杂度分析：
	时间复杂度：O(N)
    两次遍历，一次存储，一次对比

	空间复杂度：O(n)
	变量声明了一个数组
 */


 //其实还可以用翻转链表，再遍历对比链表value的方法。这周有点赶就没写，用了自己擅长的方式，随后再补刷。。。
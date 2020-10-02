// 输入两个链表，找出它们的第一个公共节点。

// 如下面的两个链表：



// 在节点 c1 开始相交。

//  

// 示例 1：



// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
// 输出：Reference of the node with value = 8
// 输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
//  

// 示例 2：



// 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// 输出：Reference of the node with value = 2
// 输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
//  

// 示例 3：



// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null
// 输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
// 解释：这两个链表不相交，因此返回 null。
//  

// 注意：

// 如果两个链表没有交点，返回 null.
// 在返回结果后，两个链表仍须保持原有的结构。
// 可假定整个链表结构中没有循环。
// 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
// 本题与主站 160 题相同：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/


/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {
    //利用map对象，先把出现过的节点设值为true
    const map = new Map();
    let node = headA;
    while (node) {
        map.set(node, true);
        node = node.next;
    }

    node = headB;
    while (node) {//再次遍历第二个链表，出现过的节点则直接返回
        if (map.has(node)) return node;
        node = node.next;
    }
    return null;//没有相交节点，则返回null
};


 /**题解
   利用map对象，记录遍历过的节点，设置为true，再次遍历第二个链表时，出现过的节点直接返回


  复杂度分析：
    时间复杂度是:O(N)
    一层循环将该最短链表遍历完毕。N取决与两条链表的最短长度。
    空间复杂度：O(N）
    定义map对象，进行设值。
 */


 /**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

 /**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {
    //如果某一条链表为空，则返回null
     if(!headA || !headB) {
        return null
    }

    var nodeA = headA
    var nodeB = headB

    while(nodeA !== nodeB) {
        //nodeA和nodeB分别走一遍 headA和headB, 如果headA和headB相交，nodeA和nodeB最终会一起走到最后一个节点，否则headA和headB不相交
        nodeA = (nodeA === null) ? headB : nodeA.next
        nodeB = (nodeB === null) ? headA : nodeB.next
    }
    
    return nodeA
};

 /**题解
   双指针实现法。直接考虑到空间复杂度为O(1)，设置两条链表的头指针，让其分别走一遍两条链表，最后相交，两个头指针肯定会走到最后一个节点，
   否则不相交，返回为null

  复杂度分析：
    时间复杂度是:O(M+N)
    一层循环将该最短链表遍历完毕。N取决与两条链表的最短长度。
    空间复杂度：O(1）这个有些不确定
    虽然是两个头指针，但是它所占据的内存应该也是链表。
 */
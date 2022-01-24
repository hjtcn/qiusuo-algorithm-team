/*
147. 对链表进行插入排序
对链表进行插入排序。


插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

 

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。
 

示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/

/*
    思路:记录最后一个排序的节点lastSorted。
        同时保持指针不断后移，避免出现环。
        原理：如果lastSorted<cur,则当前节点直接排到后面就行。
            反之，从头寻找第一个cur大于的值，并替换
        记录prev节点，同时替换时，及时断掉next链接
    


*/
/*function ListNode(val, next) {
    *     this.val = (val===undefined ? 0 : val)
    *     this.next = (next===undefined ? null : next)
    * }
    */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
 var insertionSortList = function(head) {
    if (head === null) {
        return head;
    }
    const dummyHead = new ListNode(0);
    dummyHead.next = head;
    //虚拟节点
    let lastSorted = head, curr = head.next;
    //最后一个排序的节点
    while (curr !== null) {
        if (lastSorted.val <= curr.val) {
            //最后一个排序节点小于当前节点，那就比较下一个节点
            lastSorted = lastSorted.next;
        } else {
            //从头寻找第一个比当前节点大的值
            let prev = dummyHead;
            while (prev.next.val <= curr.val) {
                prev = prev.next;
            }
            lastSorted.next = curr.next;//去掉当前节点的链接。
            curr.next = prev.next;
            prev.next = curr;
        }
        curr = lastSorted.next;
    }
    return dummyHead.next;
};

/*
    时间复杂度：O(N^2)
    空间复杂度：O(1)
*/

/*
    思考：长时间不写，把自己绕里面了，只要抱着记录最后一个排序过的节点，就会好理解一些
*/
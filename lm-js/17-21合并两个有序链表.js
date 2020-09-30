// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 



// 示例：

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4


/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function (l1, l2) {
    let res = new ListNode(), newRes = res//依然定义两个变量，分别遍历链表，并赋值，另一个等待返回头指针
    while (l1 && l2) {//保证其中一个链表遍历完毕
        if (l1.val < l2.val) {//如果链表l1的当前节点比l2小
            res.next = l1//则目标链表的next节点即为链表l1的当前节点
            l1 = l1.next//链表l1指针也指向下一个节点
            res = res.next//寻找目标链表的下一个节点
        }
        else if (l1.val >=  l2.val) {//反之亦然
            res.next = l2
            l2 = l2.next
            res = res.next
        }

    }
    l1 && (res.next = l1)//如果l1没遍历完，res的next指针指向l1
    l2 && (res.next = l2)//如果是l2没遍历完，res的next指针指向l2
    return newRes.next//返回目标指针的头节点
};

 /**题解
   遍历两条链表，直至其中一条链表遍历结束，如果剩余一个没有遍历结束，再将目标链表的指针指向该链表即可

   那么如何控制链表的遍历呢？
   目标链表每次赋值并指向最小的节点，然后调整res的指针，去查询下一个节点。而原链表l1,l2则直接l1=l1.next;l2=l2.next就可以。
   注意：原本我在给每个节点赋值或者比大小的时候依然用了&&去确保准确。
   但后来发现，这次是只要最短的链表遍历结束就结束遍历，不会出现值为null或为空的报错情况了。

  复杂度分析：
    时间复杂度是:O(N)
    一层循环将该最短链表遍历完毕。N取决与两条链表的最短长度。
    空间复杂度：O(N+M）
    定义了res和newRes变量，存放新的和链表.
 */



 /**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function (l1, l2) {
    if(!l1){//如果l1指针指向null了，直接返回l2
        return l2
    }
    if(!l2){//如果l2指针指向null了，直接返回l1
        return l1
    }
    if(l1.val<l2.val){//比较当前l1,l2的value值的大小，
        l1.next=mergeTwoLists(l1.next,l2)//如果l1小，则调整l1的指针，再次调用mergeTwoLists方法，返回l1当前节点
        return l1
    }
    else{
        l2.next=mergeTwoLists(l1,l2.next)//如果l2小，则调整l2的指针，再次调用mergeTwoLists方法，返回l2当前节点
        return l2
    }
};

 /**题解
  递归解决。某一条当前指针指向的节点value较小，就移动指针，递归调用mergeTwoLists方法，直到该条链表遍历结束。
  而比大小后，仅仅对l1的next的指针进行了赋值，此时返回的l1的真实指针指向还是指向头指针的。

  复杂度分析：
    时间复杂度是:O(N+M)
    递归调用，直到两条链表都遍历完毕，
    空间复杂度：O(1）
    未定义额外变量
 */
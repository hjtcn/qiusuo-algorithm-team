// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807


/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
    let carry = 0, sum = new ListNode(), res = sum
    //同时两条链表查询结束，以及进位等于0之后，则循环结束。。。
    while (l1 || l2 || carry) {
        let currentSum = (l1 && l1.val) + (l2 && l2.val) + carry//记录和
        carry = currentSum > 9 ? 1 : 0//如何和大于9进位1，反之则不进位
        sum.next = new ListNode(currentSum % 10)//这里卡了一下，我发现自己不知道，怎么用js来进行链表的下一步来。要理解sum.next是一个节点
        sum = sum.next//链表调整next指针
        //原有的两条链表进行next指针调整，进行下一位的求和
        l1 && (l1 = l1.next)
        l2 && (l2 = l2.next)
    }
    //为什么没有直接返回sum链表，因为sum链表的当前指针现在已经链表的最后了。
    //res虽然引用了sum对象，但是，它的原指针还是指向0初始位置的，因此返回res.next
    return res.next
};


/**题解
   遍历两条链表，直至两条链表遍历结束，以及和的进位也不存在之后，则循环结束。
   而求和的过程即每条链表的当前值相加再加上进位(carry),求和之后sum.next即为new ListNode(当前和%10),而下次的进位则根据当前和是否大于9判断。

   那么如何控制链表的遍历呢？
   对新的和链表res.next进行赋值，再重新赋值给sum，而原链表l1,l2则直接l1=l1.next;l2=l2.next就可以。
   注意：两条原有链表长度不一的情况，要将它们的val和next进行与(&&)的处理。

  复杂度分析：
    时间复杂度是:O(N)
    一层循环将该链表遍历完毕。N取决与两条链表的最大长度。
    空间复杂度：O(N）
    定义了res和sum变量，存放新的和链表。N取决与两条链表的最大长度。
 */


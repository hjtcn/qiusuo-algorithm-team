// 92. 反转链表 II
// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。

// 示例:

// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL



/*
    思路：思路其实很简单，但是细节好难弄啊，在指定范围内反转，且能正确拼接。
    最后自己还是抠出来了，点个赞。。
    思路还是比之前清晰一点点的，起码敢想了，比如next节点的提前记录，prev节点的提前记录
    不过反转链表代码的可读性确实不太好
    我的基本思路就是：1.利用num++记录当前节点位于第几个
                   2.<m之前，遍历就行，记录结束时的prev节点，等待连接
                     >=m&&<=n之间，交换节点
                  3.跳出循环之后，开始拼接
    最初的时候，我是直接res就定位在head，但是发现只有一个元素的时候就失效了。。。
    所以链表最好养成的习惯，就是构造空节点，结果返回res.next
*/


/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} m
 * @param {number} n
 * @return {ListNode}
 */
var reverseBetween = function(head, m, n) {
    let res=node=new ListNode(),num=1,prev=new ListNode(),reversePrev=null,next=head
        node.next=head
        while(num<m){
            prev=head
            head=head.next
            num++;
        }  
        let waitNext=head //第三段的头，第一段结束的下一个
           next=head.next
       while(m<=num&&n>num){
          last=head
          head=next
          next=head.next
          head.next=last
          num++;
       }
       if(m==1){
           node.next=head//这个位置一直是我卡的点，我发现从第一个就开始反转的时候，第一次赋值是失效的
       }
        prev.next=head //第一段的尾指向第三段开始的前一个
        waitNext.next=next   
        return res.next
};


/*
时间复杂度：O(N)
一次遍历
空间复杂度:O(N)
声明了一些变量记录节点
*/
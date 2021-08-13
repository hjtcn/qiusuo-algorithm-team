/*
142. 环形链表 II
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

进阶：

你是否可以使用 O(1) 空间解决此题？
 

示例 1：



输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：



输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：



输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。
 

提示：

链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引
*/

/*
    思路：一开始理解错题意了，以为要返回的是出现环的索引。后来才理解是要返回节点本身.
        可以打个标记，当前节点如果遍历过，就设置为true,再次到这个节点，发现已经为true,就返回当前节点。
        没有找到环，返回null.

    题解：快慢指针。想到了指针，但是没有想到快慢指针，以及其中的原理。
        1.fast每次走两步
          slow每次走一步
          能碰上,就说明存在环。
        2.寻找环形的入口节点.涂涂画画
        有两个重要的位置(入口节点，相遇点)
        head-------入口节点-----相遇点-----
               x     ｜     y           ｜
                     --------------------            z   
        fast节点=slow节点走的路径*2
        x+y+n*(y+z)=(x+y)*2
        x=(n-1)(y+z)+z
        故y+z,为环周距离，n>=1,x最小直接等于z,再或是等于转了好几圈的指针+(从相遇点出发,入口截止的距离）
        
        看完代码随想录的推导。自己也跟着多推推，加强记忆。

        理解完思路后，处理细节也踩了一些坑
        1.while边界条件及初始化处理。
        fast和slow节点一开始的初始化就应该为
        fast=head.next.next
        slow=head.next
        否则在while(!fast||!fast.next||fast!=slow)都无法进入while循环。
        2.为什么边界条件处理的是fast.next而不是slow？？？
        fast=fast.next.next
        最后发现思路跑偏了，肯定fast跑的快呀，fast还能接着跑，slow不可能停止的。
        因此if(!fast||!fast.next){
            没有环，返回null
        }

        
    

*/

var detectCycle = function(head) {
   let map=new Map()
   while(head){
       if(map.has(head)){
           return head
       }
       map.set(head,true)
       head=head.next
   } 
   return null
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/


var detectCycle = function(head) {
    if(!head||!head.next||!head.next.next) return null
    let slow=head.next,fast=head.next.next
    while(fast&&fast.next&&fast!=slow){
        fast=fast.next.next
        slow=slow.next
    }
    if(!fast||!fast.next){
        //快指针走完了，说明没有环
        return null
    } 
    //如果是因为fast==slow跳出循环的，说明遇到相遇点了，相遇点为slow
    let node1=fast,node2=head
    while(node1!=node2){//肯定会相遇才可以这样处理。
        node1=node1.next
        node2=node2.next
    }
    return node2
 };

 /*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：一开始想的是入度出度。
         看题解加深了快慢指针的了解，解决环型问题不要怕，多练练就是了。
*/
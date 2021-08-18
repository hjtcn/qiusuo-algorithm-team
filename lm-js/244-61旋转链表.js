/**
61. 旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

 

示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：


输入：head = [0,1,2], k = 4
输出：[2,0,1]
 

提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109

 */
/*
    思路：1.k值正好是链表长度的倍数
        直接返回head，不操作
        2.k值非链表长度的倍数
        中间断掉，头尾拼接

        我选择用map记录prevNode节点,和curNode节点
        map.set(count,[prevNode,curNode])
        这样一次遍历，得到链表长度count和尾节点tailNode。并用k%count来区分是否操作
        if(k%count==0){
            return head
        }
        else{
            let diff=count-k%count  //操作次数
            let [prevNode,curNode]=map.get(diff)
            prevNode.next=null//断开前面
            tailNode.next=head//头尾拼接
            return curNode  //从当前节点返回。
        }
*/

var rotateRight = function(head, k) {
    if(k==0||!head) return head
    let curNode=moveNode=head
    let count=0,map=new Map()
    let oldTailNode=null
    while(curNode){
        if(!curNode.next){
            oldTailNode=curNode
        }
        let prevNode=curNode
        curNode=curNode.next
        count++;
        map.set(count,[prevNode,curNode])
    }
    let sum=k%count
    if(sum){
        let diff=count-sum
        let [tailNode,node]=map.get(diff)
        tailNode.next=null
        oldTailNode.next=head
        return node
    }
    else{
        return head
    }
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

var rotateRight = function(head, k) {
    if(k==0||!head) return head
    let curNode=moveNode=head
    let count=0
    let oldTailNode=null
    while(curNode){
        if(!curNode.next){
            oldTailNode=curNode
        }
        curNode=curNode.next
        count++;
    }
    let sum=k%count
    if(sum){
        let diff=count-sum
        while(moveNode){
            let prevNode=moveNode
            moveNode=moveNode.next
            diff--
            if(diff==0){
                prevNode.next=null
                oldTailNode.next=head
                return moveNode
            }
        }
    }
    else{
        return head
    }
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：记录衔接点，进行拼接
*/
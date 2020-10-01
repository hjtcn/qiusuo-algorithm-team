// 445. 两数相加 II
// 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。

 

// 进阶：

// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

 

// 示例：

// 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 8 -> 0 -> 7



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
var addTwoNumbers = function(l1, l2) {
    let len=0,newL1=l1,newL2=l2,sum=[]
    //遍历判断链表谁的比较长，长度差为多少
    while(l1){
        l1=l1.next
        len++
    }
    while(l2){
        l2=l2.next
        len--
    }
    while(newL1||newL2){
        if(len>0){//如果l1比较长，就把长度差外面的这些数字，直接push进入和的数组
            while(len)
            {sum.push(newL1.val)
            newL1&&(newL1=newL1.next)
            len--
            }
        }
        else if(len<0){//反之将l2的长度差之外前面的这些数字，push进入和的数组。
            len=Math.abs(len)
            while(len)
            {
                sum.push(newL2.val)
            newL2&&(newL2=newL2.next)
            len--
            }
        }
        else{//剩下就是长度一致，才开始求和，此时不考虑进位，只求和
            (newL1&&newL2)&&sum.push(newL1.val+newL2.val)
            newL1=newL1.next
            newL2=newL2.next
        }
    }
    //利用到了数组存值，然后翻转一下
  sum.reverse()
  let carry=0
  for(let i=0;i<sum.length;i++){//此时再把和超过9的元素，carry进位1，等待下一次遍历再加上carry
      let curSum=sum[i]+carry  //当前和为元素+carry
      sum[i]=curSum>9?curSum%10:curSum //元素对10取余
      carry=curSum>9?1:0 //下一位进位为1
  }
  //这里则是当最高位相加和也进一了，但是已经循环结束了，则新增一个最高位，且值为1
  if(carry==1){
      sum.push(1)
  }
  //当和数组全部调整完毕，再根据pop()来遍历构建一个新的链表
  let res=new ListNode(),newRes=res
  while(sum.length){
      res.next=new ListNode(sum.pop())
      res=res.next
  }
  return newRes.next
};

/**题解
   首先对于长的链表，在长度差之外这些数字即为求和之后的值(即len之前，和即本身)，然后遍历两个数组，进行求和赋值
   此题的关键是我借用了数组，方便他进行翻转，进位之后，再进行构建链表


  复杂度分析：
    时间复杂度是:O(N)
    一层循环将该链表遍历完毕。N取决与两条链表的最大长度。
    空间复杂度：O(N）
    定义了sum数组，进行数据处理，定义了res和newRes变量，存放新的和链表。N取决与两条链表的最大长度。
 */

<?php
/*
 * @Descripttion: 周一，好久没刷题了
 * @Author: tacks321@qq.com
 * @LastEditTime: 2021-08-10 17:55:26
 */



/*
 * @lc app=leetcode.cn id=707 lang=php
 *
 * [707] 设计链表
 *
 * https://leetcode-cn.com/problems/design-linked-list/description/
 *
 * algorithms
 * Medium (31.13%)
 * Likes:    269
 * Dislikes: 0
 * Total Accepted:    55.1K
 * Total Submissions: 174.5K
 * Testcase Example:  '["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]\n' +
  '[[],[1],[3],[1,2],[1],[1],[1]]'
 *
 * 设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next
 * 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
 * 
 * 在链表类中实现这些功能：
 * 
 * 
 * get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
 * addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
 * addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
 * addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index
 * 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
 * deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
 * 
 * 
 * 
 * 
 * 示例：
 * 
 * MyLinkedList linkedList = new MyLinkedList();
 * linkedList.addAtHead(1);
 * linkedList.addAtTail(3);
 * linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
 * linkedList.get(1);            //返回2
 * linkedList.deleteAtIndex(1);  //现在链表是1-> 3
 * linkedList.get(1);            //返回3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 所有val值都在 [1, 1000] 之内。
 * 操作次数将在  [1, 1000] 之内。
 * 请不要使用内置的 LinkedList 库。
 * 
 * 
 */



// @lc code=start
// 节点定义
// class ListNode {
//     public $val  = 0;
//     public $next = null;
//     function __construct($val = 0, $next = null) {
//         $this->val  = $val;
//         $this->next = $next;
//     }
// }


// /**
//  * 单链表
//  * 
//  * 执行用时：136 ms, 在所有 PHP 提交中击败了32.35%的用户
//  * 内存消耗：16.3 MB, 在所有 PHP 提交中击败了100.00%的用户
//  *
//  * @date  2021-08-09 19:04:55
//  */
// class MyLinkedList {

//     private $size;      // 链表长度
//     private $dummyHead; // 链表哑节点，哨兵节点，伪头节点

//     /**
//      * Initialize your data structure here.
//      */
//     function __construct() {
//         // 链表初始化
//         $this->size = 0;
//         $this->dummyHead = new ListNode();
//     }

//     /**
//      * Get the value of the index-th node in the linked list. If the index is invalid, return -1.
//      * @param Integer $index
//      * @return Integer
//      */
//     // 获取某个位置节点
//     // index是从0开始的，第0个节点就是头结点
//     function get($index) {
//         // 边界判断
//         if($index < 0 || $index >= $this->size) {
//             return -1;
//         }
//         // 头节点
//         $cur = $this->dummyHead->next;
//         // 向后遍历
//         while($index--) {
//             $cur = $cur->next;
//         }
//         return $cur->val;

//     /*
//         // 另一种遍历方法
//         $head = $this->dummyHead->next;
//         $i = 0;
//         // 循环读取
//         while($head) {
//             if($i == $index) {
//                 return $head->val;
//             }
//             // 不断向后移动
//             $head = $head->next;
//             $i++;
//         }
//     */
    

//     }

//     /**
//      * Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
//      * @param Integer $val
//      * @return NULL
//      */
//     // O(1)
//     // 头插入节点
//     // 在链表最前面插入一个节点，插入完成后，新插入的节点为链表的新的头结点
//     function addAtHead($val) {
//         // 创建一个节点
//         $node = new ListNode($val);
//         $node->next = $this->dummyHead->next; // 插入头节点

//         // 维护链表
//         $this->dummyHead->next = $node;
//         $this->size++; // 链表长度++
//     }

//     /**
//      * Append a node of value val to the last element of the linked list.
//      * @param Integer $val
//      * @return NULL
//      */
//     // O(n)
//     // 尾插入节点
//     // 在链表最后面添加一个节点
//     function addAtTail($val) {
//         // 节点
//         $node = new ListNode($val);
//         $cur  = $this->dummyHead;
//         // 遍历链表尾部
//         while($cur->next != null ) {
//             $cur = $cur->next;
//         }
//         // 从尾部插入一个节点
//         $cur->next = $node;
//         $this->size++;
//     }

//     /**
//      * Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
//      * @param Integer $index
//      * @param Integer $val
//      * @return NULL
//      */
//     // 在第index个节点后面插入一个新的节点
//     function addAtIndex($index, $val) {
//         // 头部边界判断
//         if($index <= 0) {
//             // index为0，那么新插入的节点为 链表插入头节点。
//             $this->addAtHead($val);
//             return ;
//         }
//         // 尾部边界判断
//         if($index >= $this->size) {                                         
//             // index 等于链表的长度，则说明是新插入的节点为 链表尾部插入节点
//             $this->addAtTail($val);
//             return ;
//         }
        
//         // 中间插入节点
//         $node = new ListNode($val);

//         $cur = $this->dummyHead; // 哑节点
//         while($index--) {
//             $cur = $cur->next;
//         }
//         // 中间插入节点
//         $node->next = $cur->next;
//         $cur->next  = $node;

//         $this->size++;
//     }

//     /**
//      * Delete the index-th node in the linked list, if the index is valid.
//      * @param Integer $index
//      * @return NULL
//      */
//     // 删除在第index个节点
//     function deleteAtIndex($index) {
//         if($index < 0 || $index >= $this->size) {
//             return ;
//         }

//         $cur = $this->dummyHead;
//         // 不停向后移动
//         while($index--) {
//             $cur = $cur->next;
//         }
        
//         $temp = $cur->next;
//         // 将index节点 改为 后一个节点
//         $cur->next = $cur->next->next;
//         // 释放第index个节点
//         unset($temp);
//         $this->size--;
//     }
    
// }



class ListNode {
    public $val;    // 当前节点元素 值
    public $next; // next指针 指向后一个节点的指针
    public $prev; // prev指针 指向前一个节点的指针
    function __construct($val = 0) {
        $this->val  = $val;
        $this->next = null;
        $this->prev = null;
        
    }
}

/**
 * 双向链表
 * 
 * 执行用时：56 ms, 在所有 PHP 提交中击败了100.00%的用户
 * 内存消耗：20.8 MB, 在所有 PHP 提交中击败了8.82%的用户
 * 
 * 
 * @date  2021-08-09 19:04:55
 */
class MyLinkedList {

    public $size = 0;         // 链表长度
    public $dummyHead = null; // 链表伪头节点
    public $dummyTail = null; // 链表伪尾节点

    /**
     * Initialize your data structure here.
     */
    function __construct() {
        // 链表初始化
        $this->size = 0;
        // 伪头 伪尾 总是存在
        $this->dummyHead = new ListNode();
        $this->dummyTail = new ListNode();
        // 刚开始就是两个伪节点互相连接
        $this->dummyHead->next = $this->dummyTail;
        $this->dummyTail->prev = $this->dummyHead;
    }

    /**
     * Get the value of the index-th node in the linked list. If the index is invalid, return -1.
     * @param Integer $index
     * @return Integer
     */
    // 获取某个位置节点
    function get($index) {
        // 边界判断
        if($index < 0 || $index >= $this->size) {
            return -1;
        }

        // 选择距离最近的进行遍历
        $lenHeadToIndex = $index + 1;
        $lenTailToIndex = $this->size - $index;

        if($lenHeadToIndex < $lenTailToIndex) {
            // 伪头开始
            $cur = $this->dummyHead;
            // 从左向右遍历
            for ($i=0; $i < $lenHeadToIndex; $i++) { 
                $cur = $cur->next;
            }
        } else {
            // 伪尾开始
            $cur = $this->dummyTail;
            // 从右向左遍历
            for ($i=0; $i < $lenTailToIndex; $i++) { 
                $cur = $cur->prev;
            }
        }
        
        return $cur->val;
    }

    /**
     * Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
     * @param Integer $val
     * @return NULL
     */
    // O(1)
    // 头插入节点
    function addAtHead($val) {
        $this->size++; // 链表长度++

        // 头部节点
        // dummy - headNode  （ 伪头 - 头节点 ）
        $dummy = $this->dummyHead;
        $headNode   = $this->dummyHead->next;

        // 新增一个节点
        $addnode = new ListNode($val);
        // 通过改变前驱结点和后继节点的链接关系添加元素
        // dummy - addNode - headNode （ 伪头 - 新头节点 - 头节点 ）
        $addnode->prev = $dummy;
        $addnode->next = $headNode; 

        // 维护链表
        $dummy->next     = $addnode;
        $headNode->prev  = $addnode;

      
    }

    /**
     * Append a node of value val to the last element of the linked list.
     * @param Integer $val
     * @return NULL
     */
    // O(1)
    // 尾插入节点
    function addAtTail($val) {
        $this->size++; // 链表长度++

        // 尾部节点
        // tailNode - dummy （ 尾节点 - 伪尾 ）
        $dummy     = $this->dummyTail;
        $tailNode  = $this->dummyTail->prev;

        // 节点
        $addNode = new ListNode($val);
        // tailNode - addNode - dummy  （ 尾节点 - 新尾节点 - 伪尾 ）
        $addNode->prev = $tailNode;
        $addNode->next = $dummy;

        $dummy->prev    = $addNode;
        $tailNode->next = $addNode;

      
    }

    /**
     * Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
     * @param Integer $index
     * @param Integer $val
     * @return NULL
     */
    // 在第index个节点后面插入一个新的节点
    function addAtIndex($index, $val) {
        // 头部边界判断
        if($index < 0) {
            $index = 0;
        }
        // 尾部边界判断
        if($index > $this->size) {                                         
            return ;
        }
        
        // 选择距离最近的进行遍历
        $lenHeadToIndex = $index;
        $lenTailToIndex = $this->size - $index;

        if($lenHeadToIndex < $lenTailToIndex) {
            // 伪头开始
            $preNode = $this->dummyHead;
            // 从左向右遍历
            for ($i=0; $i < $lenHeadToIndex; $i++) { 
                $preNode = $preNode->next;
            }
            // 应该插入的位置
            $nextNode = $preNode->next;

        } else {
            // 伪尾开始
            $nextNode = $this->dummyTail;
            // 从右向左遍历
            for ($i=0; $i < $lenTailToIndex; $i++) { 
                $nextNode = $nextNode->prev;
            }
            $preNode = $nextNode->prev;
        }

        // 链表长度 +1
        $this->size++;
        
        // 中间插入节点
        $addNode = new ListNode($val);
        $addNode->prev = $preNode;  
        $addNode->next = $nextNode;

        // 维持链表
        $preNode->next  = $addNode;
        $nextNode->prev = $addNode;
      
    }

    /**
     * Delete the index-th node in the linked list, if the index is valid.
     * @param Integer $index
     * @return NULL
     */
    // 删除在第index个节点
    function deleteAtIndex($index) {
        if($index < 0 || $index >= $this->size) {
            return ;
        }

        // 选择距离最近的进行遍历 删除
        $lenHeadToIndex = $index;
        $lenTailToIndex = $this->size - $index;

        // 找到 删除节点的前驱结点、删除节点的后继节点
        if($lenHeadToIndex < $lenTailToIndex) {
            // 伪头开始
            $preNode = $this->dummyHead;
            // 从左向右遍历
            for ($i = 0; $i < $lenHeadToIndex; $i++) { 
                $preNode = $preNode->next;
            }
            // 注意
            $nextNode = $preNode->next->next;

        } else {
            // 伪尾开始
            $nextNode = $this->dummyTail;
            // 从右向左遍历
            for ($i=0; $i < $lenTailToIndex - 1; $i++) { 
                $nextNode = $nextNode->prev;
            }
            // 注意
            $preNode = $nextNode->prev->prev;
        }

        // 链表长度 -1
        $this->size--;

        // 越过第 index个元素  通过改变前驱结点和后继节点的链接关系删除元素
        $preNode->next = $nextNode; 
        $nextNode->prev = $preNode; 
        
    
    }
    

    
    // 打印链表
    function printLinkedList() {
        var_dump($this->size);
        echo "\n===================\n";
        $cur = $this->dummyHead;
        while($cur->next !== null) {
            var_dump($cur->val);
            $cur = $cur->next;
        }
    }
}

 
$obj = new MyLinkedList();
$obj->addAtHead(1);
$obj->addAtTail(3);
$obj->addAtIndex(1,2);
$obj->get(1);

$obj->deleteAtIndex(0);
$obj->printLinkedList();

$obj->get(0);


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * $obj = MyLinkedList();
 * $ret_1 = $obj->get($index);
 * $obj->addAtHead($val);
 * $obj->addAtTail($val);
 * $obj->addAtIndex($index, $val);
 * $obj->deleteAtIndex($index);
 */
// @lc code=end



/*

【题目】
    链表：
        设计一个链表，实现基本操作（头插节点、尾插节点、获取某个位置节点的值、在某个位置后面插入节点、删除某个位置节点）
        
        链表是一种线性数据结构，其中每个节点元素是单独的对象，然后这些对象通过一个引用的字段（指针）链接在一起，串起来，从而形成链表结构
        大多时候利用头节点（第一个节点）来表示整个列表

        跟数组的不同是，数组是一块顺序的内存，而链表则可以是任意的

        主要有，单链表、双链表

    单链表：
        每个节点元素（包含 元素值 val、下一个节点元素的引用字段 next）
    双链表：
        每个节点元素（包含 元素值 val、下一个节点元素的引用字段 next 、上一个节点元素的引用字段 prev）

    伪节点、哨兵节点、哑节点：
        哨兵节点被用作伪头始终存在，这样结构中永远不为空，它将至少包含一个伪头。
        使用这个节点可以让程序更加简化，而不需要每次判断是否是头插，还是中间插入。    
【解析】

单链表：
//////////////////////////////////[单链表]////////////////////////////////////////////////////////////////////////////////////////////////////

=================================时间复杂度 O(1) 
        addAtHead($val)
            头插节点
                1）初始化一个新的节点 node
                    $node = new ListNode($val);
                2）将新的节点链接到原始的头节点
                    $node->next = $this->dummyHead->next; // 插入头节点
                3）将 node 指定为 头节点
                    $this->dummyHead->next = $node;
                4）链表长度 +1
                    $this->size++; // 链表长度++


=================================时间复杂度 O(n) 
        addAtTail($val)
            尾插节点
                1）初始化一个新的节点 node
                   $node = new ListNode($val);
                2）遍历到链表尾部
                    $cur  = $this->dummyHead; // 哑节点
                    while($cur->next != null ) {
                        $cur = $cur->next;
                    }
                3）尾部插入一个节点
                    $cur->next = $node;
                4）链表长度 +1
                    $this->size++;
                    

                    
=================================时间复杂度 O(k) 
        get($index)
            获取某个位置节点的值
                1）边界判断
                    由于输入index位置，需要判断是否越界
                    [0, $this->size)
                2）利用头节点开始，向后遍历index步
                    $cur = $this->dummyHead->next;
                    // 向后遍历
                    while($index--) {
                        $cur = $cur->next;
                    }
                3）返回当前节点的值
                    $cur->val;

        addAtIndex($index, $val)    【核心】就是找到要插入节点的前驱节点
            在某个位置插入节点
                1）边界判断
                    [0, $this->size)
                2）判断如果是 index == 0 ，直接利用头插入
                3）判断如果是 index == $this->size , 直接利用 尾部插入
                4）初始化一个新的节点 node
                    $node = new ListNode($val);
                5）利用哑节点，找到对应前驱节点 cur
                    // 先后顺序不能换 ======================================================================================【重点】
                    $node->next = $cur->next    //  改变node next的指向。先知道cur的next然后保存起来
                    $cur->next  = $node         //  改变前驱节点cur的next = node。然后再去改变cur的next为新的node
                6）链表长度+1
                    $this->size++;

        deleteAtIndex($index)       【核心】就是要找到删除节点的前驱节点
            删除某个位置节点
                1）边界判断
                    [0, $this->size)
                2）利用哑节点，向后不断遍历到index个节点，找到要删除节点的前驱节点cur
                    $cur = $this->dummyHead;
                    while($index--) {
                        $cur = $cur->next;
                    }
                3）将前驱节点的next改为 next->next的节点 =====================================================================【重点】
                    $temp = $cur->next;              // 临时保存一下
                    $cur->next = $cur->next->next;   // 改变删除节点的前驱节点的next
                4）释放 前驱节点的next节点
                    unset($temp);
                5）链表长度 -1
                    $this->size--;











/////////////////////////////////////[双链表]///////////////////////////////////////////////////////////////////////
双链表：      
    它会更快一些，通过更为复杂的链表元素结构，保存了前驱节点 prev

=================================时间复杂度 O(1) 
        addAtHead($val)            
            头插节点 
            【核心】插入节点的前驱节点是头部哑节点；插入节点的后继节点是真实的头节点
                    // dummy - headNode  （ 伪头 - 头节点 ）
                    =>
                    // dummy - addNode - headNode （ 伪头 - 新头节点 - 头节点 ）
              
        addAtTail($val)
            尾插节点
            【核心】插入节点的前驱节点是真实尾部节点；插入节点的后继节点是尾部哑节点
                    // tailNode  - dummy （ 尾节点 - 伪尾 ）
                    =>
                    // tailNode - addNode - dummy （ 尾节点 - 新尾节点 - 伪尾 ）
                    
                    
=================================时间复杂度 O(min(k, N-k)) 
通过比较是从头遍历快，还是从尾遍历快来加快处理
$lenHeadToIndex = $index;
$lenTailToIndex = $this->size - $index;


        get($index)
            获取某个位置节点的值
        addAtIndex($index, $val)    【核心】就是找到要插入节点的前驱节点，和后继节点
            在某个位置插入节点
        deleteAtIndex($index)       【核心】就是要找到删除节点的前驱节点，和后继节点
            删除某个位置节点
              


                    
*/
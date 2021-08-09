//设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针
///引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
//
// 在链表类中实现这些功能：
//
//
// get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
// addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
// addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val 的节点。如果 index 等于链表的长度，则该节点将附加
//到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
// deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
//
//
//
//
// 示例：
//
// MyLinkedList linkedList = new MyLinkedList();
//linkedList.addAtHead(1);
//linkedList.addAtTail(3);
//linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
//linkedList.get(1);            //返回2
//linkedList.deleteAtIndex(1);  //现在链表是1-> 3
//linkedList.get(1);            //返回3
//
//
//
//
// 提示：
//
//
// 所有val值都在 [1, 1000] 之内。
// 操作次数将在 [1, 1000] 之内。
// 请不要使用内置的 LinkedList 库。
//
// Related Topics 设计 链表
// 👍 269 👎 0
package main

import (
	"fmt"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
type MyLinkedList struct {
	Val int
	Next *MyLinkedList
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	//arr := []int{3, 4, 6, 8, 9, 22}
	arr := []int{1}

	link_list := &MyLinkedList{0, nil}
	curr_list := link_list

	for _, val := range arr {
		curr_list.Next = &MyLinkedList{
			val,
			nil,
		}

		curr_list = curr_list.Next
	}

	return *link_list.Next
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	for i := 0; i <= index; i++ {
		if this == nil {
			return -1
		}

		if i == index {
			return this.Val
		}

		this = this.Next
	}

	return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	temp_list := this
	node := &MyLinkedList{val, nil}
	node.Next = temp_list
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	node := &MyLinkedList{val, nil}
	temp_list := this
	temp_list.Next = node
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	temp := this

	node := &MyLinkedList{val, nil}

	if index <= 0 {
		node.Next = temp
	} else {
		for i := 0; i <= index; i++ {
			if (temp == nil) {
				break
			}

			if i == (index - 1) {
				node.Next = temp.Next
				temp.Next = node
				break
			}

			temp = temp.Next
		}
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	temp := this

	if index == 0 {
		this = this.Next
	} else if index >= 0 {
		for i := 0; i <= index; i++ {
			if temp.Next != nil {
				break
			}

			if i == (index - 1) {
				delete_node := temp.Next
				temp.Next = delete_node.Next
				delete_node.Next = nil
			}
		}
	}
}

func (this *MyLinkedList) dy()  {
	str := ""
	for this != nil {
		temp_num := this.Val
		str = str + strconv.Itoa(temp_num) + "->"
		this = this.Next
	}
	fmt.Println(str)
}

func main()  {
	obj := Constructor();
	new_obj := &obj;

	new_obj.dy()
	//nil nil

	new_obj.AddAtHead(1)
	new_obj.dy()
	//1 nil

	new_obj.AddAtTail(3)
	new_obj.dy()
	//1->3 nil

	new_obj.AddAtIndex(1, 2)
	new_obj.dy()
	//1->2->3 nil

	new_obj.Get(1)
	new_obj.dy()
	//1->2->3 2

	new_obj.DeleteAtIndex(1)
	new_obj.dy()
	//1->3 nil

	new_obj.Get(1)
	new_obj.dy()
	//1->3 3
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//leetcode submit region end(Prohibit modification and deletion)

不使用任何库函数，设计一个跳表。

跳表是在 O(log(n)) 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，并且跳表的代码长度相较下更短，其设计思想与链表相似。

例如，一个跳表包含 [30, 40, 50, 60, 70, 90]，然后增加 80、45 到跳表中，以下图的方式操作：

Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons

跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。跳表的每一个操作的平均时间复杂度是 O(log(n))，空间复杂度是 O(n)。

在本题中，你的设计应该要包含这些函数：

bool search(int target) : 返回target是否存在于跳表中。
void add(int num): 插入一个元素到跳表。
bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。
了解更多 : https://en.wikipedia.org/wiki/Skip_list

注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。

样例:

Skiplist skiplist = new Skiplist();

skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // 返回 false
skiplist.add(4);
skiplist.search(1);   // 返回 true
skiplist.erase(0);    // 返回 false，0 不在跳表中
skiplist.erase(1);    // 返回 true
skiplist.search(1);   // 返回 false，1 已被擦除

约束条件:
0 <= num, target <= 20000
最多调用 50000 次 search, add, 以及 erase操作。

1. Clarification:


2. 看题解：

看题解也有点懵逼
const MaxLevel = 32

const Probability = 0.25 // 基于时间与空间综合 best practice 值, 越上层概率越小

func randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level < MaxLevel; level++ {
	}
	return
}

type node struct {
	nextNodeArray []*node
	key           int
}

func newNode(key, level int) *node {
	return &node{key: key, nextNodeArray: make([]*node, level)}
}

type Skiplist struct {
	head  *node
	level int
}

func Constructor() Skiplist {
	return Skiplist{head: newNode(0, MaxLevel), level: 1}
}

func (this *Skiplist) Search(target int) bool {
	// 类似 delete
	current := this.head
	for i := this.level - 1; i >= 0; i-- {
		for current.nextNodeArray[i] != nil {
			if current.nextNodeArray[i].key == target {
				return true
			} else if current.nextNodeArray[i].key > target {
				break
			} else {
				current = current.nextNodeArray[i]
			}
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	current := this.head
	update := make([]*node, MaxLevel) // 新节点插入以后的前驱节点
	for i := this.level - 1; i >= 0; i-- {
		if current.nextNodeArray[i] == nil || current.nextNodeArray[i].key > num {
			update[i] = current
		} else {
			for current.nextNodeArray[i] != nil && current.nextNodeArray[i].key < num {
				current = current.nextNodeArray[i] // 指针往前推进
			}
			update[i] = current
		}
	}

	level := randLevel()
	if level > this.level {
		// 新节点层数大于跳表当前层数时候, 现有层数 + 1 的 head 指向新节点
		for i := this.level; i < level; i++ {
			update[i] = this.head
		}
		this.level = level
	}
	node := newNode(num, level)
	for i := 0; i < level; i++ {
		node.nextNodeArray[i] = update[i].nextNodeArray[i]
		update[i].nextNodeArray[i] = node
	}
}

func (this *Skiplist) Erase(num int) bool {
	current := this.head
	flag := false
	for i := this.level - 1; i >= 0; i-- {
		for current.nextNodeArray[i] != nil {
			if current.nextNodeArray[i].key == num {
				tmp := current.nextNodeArray[i]
				current.nextNodeArray[i] = tmp.nextNodeArray[i]
				tmp.nextNodeArray[i] = nil
				flag = true
				// 这里要 break, 因为leetcode 的判定是，在有两个 num 9时， erase 只能删除一个
				break
			} else if current.nextNodeArray[i].key > num {
				break
			} else {
				current = current.nextNodeArray[i]
			}
		}
	}
	return flag
}

3. 复杂度分析：
时间复杂度：O(log(n))
空间复杂度：O(1)

4. 总结：
4.1: 对复杂的问题还是需要注意和分析

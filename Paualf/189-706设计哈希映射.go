不使用任何内建的哈希表库设计一个哈希映射（HashMap）。
实现 MyHashMap 类：

MyHashMap() 用空映射初始化对象
void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，则更新其对应的值 value 。
int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。
 

示例：

输入：
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
输出：
[null, null, null, 1, -1, null, 1, null, -1]

解释：
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
 

提示：

0 <= key, value <= 10^6
最多调用 104 次 put、get 和 remove 方法


Clarification:
一开始有点懵逼，这是要干啥？

又仔细想了下，既然要设计简单的hash映射，那么我们从它的存储开始分析

Q: 使用数组存储？
不知道数据会有多长，数组长度没有办法确定，使用数组不是一个好的选择，因为题目中给的key的范围是 0 到 10^6 次方，如果使用数组的话太浪费空间了

Q: 使用链表进行存储？
使用链表进行存储的话

怎么删除元素呢？

remove?


Q: 使用链表的话怎么判断它是否存在呢？
  需要引入hash 去帮我们进行判断key是否存在

设计的有点懵懵的。。。


2.看题解：

const base = 769

type entry struct {
    key, value int
}

type MyHashMap struct {
    data []list.List
}

func Constructor() MyHashMap {
    return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
    return key % base
}

func (m *MyHashMap) Put(key, value int) {
    h := m.hash(key)
    for e := m.data[h].Front(); e != nil; e = e.Next() {
        if et := e.Value.(entry); et.key == key {
            e.Value = entry{key, value}
            return
        }
    }
    m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
    h := m.hash(key)
    for e := m.data[h].Front(); e != nil; e = e.Next() {
        if et := e.Value.(entry); et.key == key {
            return et.value
        }
    }
    return -1
}

func (m *MyHashMap) Remove(key int) {
    h := m.hash(key)
    for e := m.data[h].Front(); e != nil; e = e.Next() {
        if e.Value.(entry).key == key {
            m.data[h].Remove(e)
        }
    }
}


链地址法：
还用了标准库的 list

3. 复杂度分析：
时间复杂度：O(n): 需要扫描链地址
空间复杂度：O(n)

4. 总结：
4.1: 有的时候一种数据结构满足不了我们的实现，我们需要引入另外一种数据结构帮助我们解决我们的问题

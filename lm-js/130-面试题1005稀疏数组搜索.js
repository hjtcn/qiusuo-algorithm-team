// 面试题 10.05. 稀疏数组搜索
// 稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

// 示例1:

//  输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
//  输出：-1
//  说明: 不存在返回-1。
// 示例2:

//  输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
//  输出：4
// 提示:

// words的长度在[1, 1000000]之间
/* 
    思路：一开始看到题，脑子里直接就想到indexOf方法了，然后懵了一下，想起来，对，我们这种主要是练习二分，哈哈哈哈
        然后思考了下如何二分：while(l<=r)
                        {
                            let mid=parseInt((l+r)/2)
                            if(){//向右查
                                l=mid+1
                            }
                            if(){//向左查
                                r=mid-1
                            }
                        }

        把架子搭好后，发现事情不是我想象的那样简单。是存在空字符串的
        那么没有匹配上的时候该咋处理呢？我想过同时向两边发散，找到非空值之后再匹配，但是匹配过后方向又无法确定了。
        然后看了看题解，发现仅仅选择一边发散即可，提前记录mid,如果找到最边界，依然无法找到非空值。则修改l或者r即可

*/

/**
 * @param {string[]} words
 * @param {string} s
 * @return {number}
 */
var findString = function(words, s) {
    return words.indexOf(s)
};


/**
 * @param {string[]} words
 * @param {string} s
 * @return {number}
 */
var findString = function(words, s) {
    let l=0,r=words.length-1
    while(l<=r)
    {  let flag=mid=parseInt((l+r)/2)
        //这是选择向右查询，向左查询也可以，及时控制左边的边界即可
        while(!words[mid]&&mid<r){
            mid++
        }
        //到边了
        if(!words[mid]){
            r=flag-1
            continue;
        }
        // 向左查询
        // while(!words[mid]&&mid>l){
        //     mid--
        // }
        // if(!words[mid]){
        //     l=flag+1
        //     continue;
        // }
       if(words[mid]==s){
           return mid
       }
       else if(words[mid]<s){
           l=mid+1
       }
       else{
           r=mid-1
       }
    }
    return -1
}

/*
    时间复杂度：普通遍历为O(N)
              二分查找为O(logN).但此题为拓展题，存在空值，也会存在极端情况，比如正确目标在第二个，或者倒数第二个，中间都为空。也会存在将所有值遍历一遍的情况
    空间复杂度：O(1)
             仅声明一些变量，存储l,r,mid

*/
/*
面试题 10.05. 稀疏数组搜索
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

示例1:

 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
 输出：-1
 说明: 不存在返回-1。
示例2:

 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
 输出：4
提示:

words的长度在[1, 1000000]之间

*/

/*
    思路：这次还没有上次敢思考呢，没敢想mid发散直至边界值。
    如如果为空，mid不停右移。
    需要控制的边界值r
    while(!word[mid]&&mid<r){在范围内
        mid++;
    }
    if(!word[mid]){//mid==r,此时最右边依然为空
        r--;
        continue;
    }
*/

var findString = function(words, s) {
    let l=0,r=words.length-1
    while(l<=r){
        let mid=parseInt((l+r)/2)
        while(!words[mid]&&mid<r){
            mid++;
        }
        if(!words[mid]){
            r--
            continue
        }
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
    时间复杂度：O(logN)
    空间复杂度：O(1)
*/

/*
    思路：掺杂其它事物的有序数组，二分之外可mid发散，同时控制边界。
*/
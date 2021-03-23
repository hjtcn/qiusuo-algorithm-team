// 696. 计数二进制子串
// 给定一个字符串 s，计算具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是连续的。

// 重复出现的子串要计算它们出现的次数。

 

// 示例 1 :

// 输入: "00110011"
// 输出: 6
// 解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

// 请注意，一些重复出现的子串要计算它们出现的次数。

// 另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
// 示例 2 :

// 输入: "10101"
// 输出: 4
// 解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
 

// 提示：

// s.length 在1到50,000之间。
// s 只包含“0”或“1”字符。


/*
    思路：自己做的时候，这是简单题？看完题解的时候，这确实是简单题
        好吧，自己没想出来，遍历的时候，计算了半天，总有漏掉的。
        
    题解思路：统计相同数字出现的次数。
            如："00110011"   次数可记录为[2,2,2,2]
            然后记录前后值的最小值，代表能抵消的数值，然后相加和为目标子串个数

*/

/**
 * @param {string} s
 * @return {number}
 */
 var countBinarySubstrings = function(s) {
    let arr=[],curNode=s[0],k=1,sum=0
  for(let i=1;i<s.length;i++){
      if(curNode!=s[i]){
          arr.push(k)
          k=1
          curNode=s[i]
      }
      else{
        k++
      }
  }
  arr.push(k)
  for(let i=0;i<arr.length-1;i++){
      sum+=Math.min(arr[i],arr[i+1])
  }
  return sum
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)

*/


/*
    后来又看了看官方题解，人家把空间复杂度优化到O(1)了
    记录第一个值，然后开始往后遍历，记录个数count。
    然后统计count和上一个count的最小值，进行相加。
*/

var countBinarySubstrings = function(s) {
    let pre=0,last=0,len=s.length,sum=0
    while(pre<len){
      let preNode=s[pre],count=0
      while(pre<len&&preNode==s[pre]){
          pre++;
          count++;
      }
      sum+=Math.min(count,last) //统计前后出现次数的最小值，进行相加
      last=count //记录上一个连续数字的出现次数
    }
    return sum
};


/*
    时间复杂度：O(N)
    空间复杂度：O(1)

*/
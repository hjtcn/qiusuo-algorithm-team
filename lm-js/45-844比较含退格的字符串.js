// 844. 比较含退格的字符串
// 给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

// 注意：如果对空文本输入退格字符，文本继续为空。

 

// 示例 1：

// 输入：S = "ab#c", T = "ad#c"
// 输出：true
// 解释：S 和 T 都会变成 “ac”。
// 示例 2：

// 输入：S = "ab##", T = "c#d#"
// 输出：true
// 解释：S 和 T 都会变成 “”。
// 示例 3：

// 输入：S = "a##c", T = "#a#c"
// 输出：true
// 解释：S 和 T 都会变成 “c”。
// 示例 4：

// 输入：S = "a#c", T = "b"
// 输出：false
// 解释：S 会变成 “c”，但 T 仍然是 “b”。
 

// 提示：

// 1 <= S.length <= 200
// 1 <= T.length <= 200
// S 和 T 只含有小写字母以及字符 '#'。
 

// 进阶：

// 你可以用 O(N) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？


/*
 * @lc app=leetcode.cn id=844 lang=javascript
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
/**
 * @param {string} S
 * @param {string} T
 * @return {boolean}
 */
var backspaceCompare = function(S, T) {
    let stack1=doBackspace(S),stack2=doBackspace(T)
    return stack1===stack2
  
  };
  function doBackspace(arr){
      let stack=[]
      //遍历处理数组，遇到#，出栈，反之将该元素入栈即可
      for(let i=0;i<arr.length;i++){
          if(arr[i]=='#'){
              stack.length>0?stack.pop():[]
          }
          else{
              stack.push(arr[i])
          }
      }
      //stack数组变为字符串比较是否相等，因为js中数组为对象，就算两个数组所有元素都相等，引用地址不同，这两个数组也是不相等的。
      return stack.join('')
  }


  /**
   *封装一个数组的处理函数，遍历数组，遇到#，则数组pop,非#，则数组push当前元素，并将数组拼接为字符串返回
    再比较处理过后的两个字符串是否相等

    复杂度分析：
    时间复杂度：O(N)
    数组遍历过程

    空间复杂度：O(N)
    变量声明数组，模拟入栈出栈操作。
   */
  // @lc code=end
  
  var backspaceCompare = function(S, T) {
    let sLen=S.length-1,tLen=T.length-1,sFlag=0,tFlag=0
    while(sLen>=0||tLen>=0){
        //遇到#标记+1，指针左移
        if(S[sLen]=='#'){
            sLen--;
            sFlag++;
            continue;
        }
        if(T[tLen]=='#'){
            tLen--;
            tFlag++;
            continue;
        }
        //标记大于0，则指针接着左移，标记减1(去除#影响力的过程)
        if(sFlag){
            sLen--;
            sFlag--;
            continue;
        }
        if(tFlag){
            tLen--;
            tFlag--;
            continue;
        }
        //战胜前面两种情况,双指针指向的两个元素是否相等，不等直接返回false
        if(S[sLen]!=T[tLen]){
            return false
        }
        else{//相等，指针同时左移
            sLen--;
            tLen--
        }
    }
    return true
  };

    /**
    双指针，原理：从后向前遍历数组元素不会受退格的影响。
    同时遍历遍历两个数组，会出现三种情况
    1.遇到'#',则两个数组的标记变量flag都+1，指针左移，continue，跳入下次循环；
    2.遇到两个数组的标记变量flag都不为0，说明#号的影响力还没消失，因此，指针再次左移，flag也可以-1(减去一个#的影响)，continue,跳入下次循环
    3.以上都不是，则比较当前指针指向的元素是否相等，不想等，则返回false,相等，则指针再次左移。
    如果跳出循环，依然没有返回值，则返回true.

    复杂度分析：
    时间复杂度：O(N)
    一层while循环，这里本来觉得时间复杂度取决于数组较长的一方(maxLen).
    后来仔细思考continue，时间复杂度应该大于等于O(maxLen),小于等于O(m+n)

    空间复杂度：O(1)
    定义变量存储指针位置
   */
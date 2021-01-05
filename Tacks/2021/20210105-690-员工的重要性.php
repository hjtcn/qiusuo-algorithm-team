<?php
/*
 * @Descripttion: 周二，今天头有点懵逼
 * @Author: tacks321@qq.com
 * @Date: 2021-01-05 17:24:44
 * @LastEditTime: 2021-01-05 20:39:27
 */


/*
 * @lc app=leetcode.cn id=690 lang=php
 *
 * [690] 员工的重要性
 *
 * https://leetcode-cn.com/problems/employee-importance/description/
 *
 * algorithms
 * Easy (59.22%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    20.6K
 * Total Submissions: 34.7K
 * Testcase Example:  '[[1,2,[2]], [2,3,[]]]\n2'
 *
 * 给定一个保存员工信息的数据结构，它包含了员工唯一的id，重要度 和 直系下属的id。
 * 
 * 比如，员工1是员工2的领导，员工2是员工3的领导。他们相应的重要度为15, 10, 5。那么员工1的数据结构是[1, 15,
 * [2]]，员工2的数据结构是[2, 10, [3]]，员工3的数据结构是[3, 5,
 * []]。注意虽然员工3也是员工1的一个下属，但是由于并不是直系下属，因此没有体现在员工1的数据结构中。
 * 
 * 现在输入一个公司的所有员工信息，以及单个员工id，返回这个员工和他所有下属的重要度之和。
 * 
 * 示例 1:
 * 
 * 
 * 输入: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
 * 输出: 11
 * 解释:
 * 员工1自身的重要度是5，他有两个直系下属2和3，而且2和3的重要度均为3。因此员工1的总重要度是 5 + 3 + 3 = 11。
 * 
 * 
 * 注意:
 * 
 * 
 * 一个员工最多有一个直系领导，但是可以有多个直系下属
 * 员工数量不超过2000。
 * 
 * 
 */

// @lc code=start
/**
 * Definition for Employee.
 * class Employee {
 *     public $id = null;
 *     public $importance = null;
 *     public $subordinates = array();
 *     function __construct($id, $importance, $subordinates) {
 *         $this->id = $id;
 *         $this->importance = $importance;
 *         $this->subordinates = $subordinates;
 *     }
 * }
 */

class Solution {

    private $map = []; // 快速查找到员工信息
    /**
     * @param Employee[] $employees
     * @param Integer $id
     * @return Integer
     */
    // 执行用时：16 ms, 在所有 PHP 提交中击败了88.89%的用户
    // 内存消耗：17.3 MB, 在所有 PHP 提交中击败了59.26%的用户
    function getImportance($employees, $id) {
        foreach($employees as $employee) {
            // 建立map哈希，快速查询员工数据结构
            $this->map[$employee->id] = $employee;
        }

        return $this->dfs($id);
    }

    function dfs($id) {
        // 找到对应成员的 数据结构
        $employee = $this->map[$id];
        $result = $employee->importance; // 当前成员的重要值
        foreach($employee->subordinates as $eId) {
            $result += $this->dfs($eId); // 深搜
        }
        return $result;
    }
}
// @lc code=end


/*

一个员工 => 一个直系领导
一个领导 => 多个直系下属

一对多

题意：找到某个员工的重要性+所有下属重要性

做法：如何找到某个员工，深搜 DFS


*/

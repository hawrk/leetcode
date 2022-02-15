/**
* @Author:hawrkchen
* @Date:2021/3/6 2:43 下午
* @desc:  两数之和
 */

package main

// 1 暴力解法
func twoSum(nums []int, target int) []int {
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if v1+v2 == target {
				return []int{k1, k2}
			}
		}
	}
	return nil
}

// 2. hash map 保存数据
func towSum2(nums []int, target int) []int {
	//TODO
	return []int{}
}

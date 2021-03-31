/**
* @Author:hawrkchen
* @Date:2021/3/6 2:43 下午
* @desc:
 */

package main

func twoSum(nums []int, target int) []int {
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if v1 + v2 == target {
				return []int{k1, k2}
			}
		}
	}
	return nil
}
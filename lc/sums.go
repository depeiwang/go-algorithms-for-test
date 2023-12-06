package lc

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	numWithIdx := map[int]int{}
	for i1 := range nums {
		n1 := nums[i1]
		n2 := target - n1
		if i2, found := numWithIdx[n2]; found {
			result[0] = i1
			result[1] = i2
			break
		}
		numWithIdx[n1] = i1
	}

	return result
}

// 返回不重复的三元组列表，每一个三元组内元素之和为0
func ThreeSum(nums []int) [][]int {
	cache := map[string]struct{}{}
	result := [][]int{}
	for i := range nums {
		for j := range nums {
			for k := range nums {
				if i == j || j == k || i == k {
					continue
				}
				vi, vj, vk := nums[i], nums[j], nums[k]
				cacheKey := keyWitOrder(vi, vj, vk)
				if _, found := cache[cacheKey]; found {
					continue
				}
				cache[cacheKey] = struct{}{}
				if vi+vj+vk == 0 {
					result = append(result, []int{vi, vj, vk})
				}
			}
		}
	}
	return result
}

func keyWitOrder(v1, v2, v3 int) string {
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	if v1 > v3 {
		v1, v3 = v3, v1
	}
	if v2 > v3 {
		v2, v3 = v3, v2
	}
	return fmt.Sprintf("%d,%d,%d", v1, v2, v3)
}

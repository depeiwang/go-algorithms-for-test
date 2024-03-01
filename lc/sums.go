package lc

func twoSum(nums []int, target int, skipIdx map[int]struct{}) (int, int) {
	numWithIdx := map[int]int{}

	for i1 := 0; i1 < len(nums); i1++ {
		n1 := nums[i1]
		n2 := target - n1
		if i2, found := numWithIdx[n2]; found {
			return i1, i2
		}
		numWithIdx[n1] = i1
	}

	return -1, -1
}

// 返回不重复的三元组列表，每一个三元组内元素之和为0
func ThreeSum(nums []int) [][]int {
	cache := map[int]map[int]map[int]struct{}{}
	result := [][]int{}

	for i1 := range nums {
		v1 := nums[i1]
		skipIds := map[int]struct{}{i1: {}}
		i2, i3 := twoSum(nums, -v1, skipIds)
		if i2 < 0 || i1 != i2 && i1 != i3 && i2 != i3 {
			continue
		}
		v2, v3 := nums[i2], nums[i3]
		if exists(v1, v2, v3, cache) {
			continue
		}
		result = append(result, []int{v1, v2, v3})
	}
	return result
}

func exists(v1, v2, v3 int, cache map[int]map[int]map[int]struct{}) bool {
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	if v1 > v3 {
		v1, v3 = v3, v1
	}
	if v2 > v3 {
		v2, v3 = v3, v2
	}
	l2 := cache[v1]
	if l2 == nil {
		l2 = map[int]map[int]struct{}{}
	}
	l1 := l2[v2]
	if l1 == nil {
		l1 = map[int]struct{}{}
	}
	if _, found := l1[v3]; found {
		return true
	}
	l1[v3] = struct{}{}
	l2[v2] = l1
	cache[v1] = l2
	return false
}

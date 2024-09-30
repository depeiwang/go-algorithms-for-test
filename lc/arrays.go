package lc

/*
4. 寻找两个正序数组的中位数
困难
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6
*/
// https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	indexMin, indexMax, halfLen := 0, m, (m+n+1)/2
	for indexMin <= indexMax {
		indexMiddle := (indexMin + indexMax) / 2
		j := halfLen - indexMiddle
		if indexMiddle < m && nums2[j-1] > nums1[indexMiddle] {
			indexMin = indexMiddle + 1
		} else if indexMiddle > 0 && nums1[indexMiddle-1] > nums2[j] {
			indexMax = indexMiddle - 1
		} else {
			var max_of_left, min_of_right int
			if indexMiddle == 0 {
				max_of_left = nums2[j-1]
			} else if j == 0 {
				max_of_left = nums1[indexMiddle-1]
			} else {
				max_of_left = nums1[indexMiddle-1]
				if nums2[j-1] > max_of_left {
					max_of_left = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(max_of_left)
			}
			if indexMiddle == m {
				min_of_right = nums2[j]
			} else if j == n {
				min_of_right = nums1[indexMiddle]
			} else {
				min_of_right = min(nums1[indexMiddle], nums2[j])
			}
			return float64(max_of_left+min_of_right) / 2.0
		}
	}
	return 0.0
}

// https://leetcode.cn/problems/next-permutation/
func nextPermutation(nums []int) {
	lastIdx := len(nums) - 1
	lastVal := nums[lastIdx]

	reverse := func(arr []int) {
		for i, j := 0, len(arr)-1; i < j; {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	// 从后往前找到第一个递增的数
	for i := lastIdx - 1; i >= 0; i-- {
		if nums[i] < lastVal {
			// 从后往前找到第一个比nums[i]大的数
			for j := lastIdx; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					// 交换后，i后面的数需要重新排序
					reverse(nums[i+1:])
					return
				}
			}
		}
		lastVal = nums[i]
	}
	reverse(nums)
}

package search

func searchRange(nums []int, target int) []int {
	index := binarySearch(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}
	start, end := index, index
	// 从指定位置向前向后查找
	for i := index; i < len(nums); i++ {
		end = i
		if nums[i] != nums[index] {
			end--
			break
		}
	}
	for j := index; j >= 0; j-- {
		start = j
		if nums[j] != nums[index] {
			start++
			break
		}
	}
	return []int{start, end}
}

func binarySearch(nums []int, target int) int {

	low, mid, high := 0, 0, 0
	high = len(nums) - 1
	for low <= high {
		mid = (high-low)/2 + low
		if target < nums[mid] { //目标值大于中间值,则从前一半中查找
			high = mid - 1
		} else if target > nums[mid] { //目标值大于中间值,则从后一半中查找
			low = mid + 1
		} else { //相等
			return mid
		}
	}
	return -1
}

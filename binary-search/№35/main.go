package main

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	var mid int
	for {
		if nums[low] > target {
			return low
		}
		
		if nums[high] < target {
			return high+1
		}
	
		if low <= high {
			mid = int((low+high)/2)
			guess := nums[mid]
			if guess == target {
				return mid
			}
			if guess > target {
				high = mid - 1
			}
			if guess < target {
				low = mid + 1
			}
		} else {
			break
		}
	}
	return mid
}

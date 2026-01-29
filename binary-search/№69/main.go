package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	low := 0
	high := x 

	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == x {
			return int(mid)
		}
		if mid*mid < x {
			low = int(mid) + 1
		} else {
			high = int(mid) - 1
		}
	}
	return high
}
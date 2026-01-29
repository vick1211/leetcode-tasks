package main

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    low := 1
    high := n
    var mid int
    for low <= high {
        mid = int((low+high)/2)
        if (isBadVersion(mid) && !isBadVersion(mid-1)) || (low == high) {
            return mid
        }
        if isBadVersion(mid) {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return 0
}
package sliceutil

import "sort"

func HasUint64(slice []uint64, target uint64) bool {
	for i := range slice {
		if slice[i] == target {
			return true
		}
	}
	return false
}

func HasUint64BS(slice []uint64, target uint64) bool {
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	low, high := 0, len(slice)-1
	for low <= high { // binary search
		mid := low + (high-low)>>1
		if slice[mid] == target {
			return true
		} else if slice[mid] > target { // move to low
			high = mid - 1
		} else { // move to high
			low = mid + 1
		}
	}
	return false
}

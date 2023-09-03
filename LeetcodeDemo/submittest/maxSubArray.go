package submittest

//nolint:golint,unused
func maxSubArray(nums []int) int {
	var f int
	var s int
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	if numsLen == 1 {
		return nums[0]
	}
	if numsLen == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}

	}
	f = nums[0]
	s = nums[1]
	if s > f {
		f = nums[1]
		s = nums[0]
	}
	for i := 2; i < len(nums)-3; i++ {
		n := nums[i]
		if n > s {
			s = n
		}
		if n > f {
			s = f
			f = n
		}
	}
	return f + s
}

package submittest

func maxSubArray(nums []int) int {
	var f int
	var s int
	numsLen := len(nums)
	if numsLen < 2{
		return nums[0] + nums[1]
	}
	f = nums[0]
	s = nums[1]
	if s > f {
		f = nums[1]
		s = nums[0]
	}
	for i := 2 ; i < len(nums) -3 ; i++ {
		n := nums[i]
		if n > f {
			f = n
			s = f
		}
	}
	return f + s
}

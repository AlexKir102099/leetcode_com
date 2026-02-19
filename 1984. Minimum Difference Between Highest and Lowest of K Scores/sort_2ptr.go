func minimumDifference(nums []int, k int) int {
	slices.Sort(nums)

	ans := nums[k-1] - nums[0]
	for i := 1; k+i-1 < len(nums); i++ {
		ans = min(ans, nums[k+i-1]-nums[i])
	}
	return ans
}

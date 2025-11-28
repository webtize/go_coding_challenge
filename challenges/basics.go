package challenges

// ProcessNumbers analyzes a slice of integers and returns:
// - sum: the sum of all numbers
// - avg: the average of the numbers (as float64)
// - max: the maximum number in the slice
// - min: the minimum number in the slice
//
// If the slice is empty, return 0 for all values.
//
// Examples:
// Input: []int{1, 2, 3, 4, 5}
// Output: sum=15, avg=3.0, max=5, min=1
//
// Input: []int{10, -2, 0}
// Output: sum=8, avg=2.666..., max=10, min=-2
func ProcessNumbers(nums []int) (sum int, avg float64, max int, min int) {
	// TODO: Implement this function
	if len(nums) == 0 {
		return 0, 0.0, 0, 0
	}
	max, min = nums[0], nums[0]

	for i = 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	avg = float64 (sum) / float64 (len(nums))
	return
		
}

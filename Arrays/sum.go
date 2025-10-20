package arrays

func Sum(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return
}

func SumAll(numsToSum ...[]int) []int {
	length := len(numsToSum)
	sums := make([]int, length)
	for i, nums := range numsToSum {
		sums[i] = Sum(nums)
	}
	return sums
}

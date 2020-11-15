package leetcode

func twoSum(nums []int, target int) []int {
  indexes := make([]int, 2)
  for i1, v1 := range nums {
    offset := i1 + 1
    for i2, v2 := range nums[offset:] {
      if (v1 + v2 == target) {
        indexes = []int{ i1, i2 + offset }
      }
    }
  }
  return indexes
}

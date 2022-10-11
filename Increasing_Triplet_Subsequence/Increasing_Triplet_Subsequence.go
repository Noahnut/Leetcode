package increasingtripletsubsequence

func IncreasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }

    left, mid := 2147483647, 2147483647

    for i := 0; i < len(nums); i++ {
        if nums[i] > mid {
            return true
        } else if nums[i] > left && nums[i] < mid {
            mid = nums[i]
        } else if nums[i] < left {
            left = nums[i]
        }
    }

    return false
}
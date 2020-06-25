package practice
/**
 153.寻找旋转排序数组中的最小值
 */
import "testing"

func findMin(nums []int) int {
	left :=0
	right :=len(nums)-1
	for left<right{
		mid :=left+(right-left)>>1
		if nums[mid]>nums[right]{
			left = mid+1
		}else {
			right= right-1
		}
	}
	return nums[left]
}

func TestFindMin(t *testing.T){
	nums := []int{3,4,5,1,2}
	t.Log(findMin(nums))
}
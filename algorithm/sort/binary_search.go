package sort

/**
 *@Author tudou
 *@Date 2020/12/7
 **/



func BinarySearch(arr []int,target int)int{
	left,right:=0,len(arr)-1
	for left<=right{
		mid:=(right-left)/2+left
		if arr[mid]<target{
			left = mid+1
		}else if arr[mid]>target{
			right = mid-1
		}else{
			return mid
		}
	}
	return -1
}

func LeftBoundBinarySearch(arr []int,target int)int{
	left,right:=0,len(arr)-1
	for left<=right{
		mid:=(right-left)/2+left
		if arr[mid]<target{
			left = mid+1
		}else if arr[mid]>=target{
			right = mid-1
		}
	}
	if left>=len(arr)&&arr[left]!=target{
		return -1
	}
	return left
}

func RightBoundBinarySearch(arr []int,target int)int{
	left,right:=0,len(arr)-1
	for left<=right{
		mid:=(right-left)/2+left
		if arr[mid]<=target{
			left = mid+1
		}else if arr[mid]>target{
			right = mid-1
		}
	}
	if right>=len(arr)&&arr[left]!=target{
		return -1
	}
	return right
}
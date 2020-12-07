package sort

/**
 *@Author tudou
 *@Date 2020/12/7
 **/


func MergeSort(arr []int)[]int{
	if len(arr)<=1{
		return arr
	}
	middle:=len(arr)/2
	left:=MergeSort(arr[:middle])
	right:=MergeSort(arr[middle:])
	return Merge(left,right)
}


func Merge(left,right []int)(result []int){
	l,r:=0,0

	for l<len(left)&&r<len(right){
		if left[l]>right[r]{
			result=append(result,right[r])
			r++
		}else{
			result=append(result,left[l])
			l++
		}
	}

	result=append(result,left[l:]...)
	result=append(result,right[r:]...)
	return
}
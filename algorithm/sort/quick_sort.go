package sort

/**
 *@Author tudou
 *@Date 2020/12/7
 **/

func QuickSort(arr []int,start,end int){
	if start<end{
		i,j:=start,end
		key:=arr[(start+end)/2]
		for i<=j{
			for arr[i]<key{
				i++
			}
			for arr[j]>key{
				j--
			}
			if i<=j{
				arr[i],arr[j]=arr[j],arr[i]
				i++
				j--
			}
		}
		if start<j{
			QuickSort(arr,start,j)
		}
		if end>i{
			QuickSort(arr,i,end)
		}
	}

}

func QuickSortA(arr []int,start,end int){
	if start>=end{
		return
	}
	partition:=Partition(arr,start,end)
	QuickSort(arr,start,partition-1)
	QuickSort(arr,partition+1,end)
}

func Partition(arr []int,start,end int)int{
	pivot:=arr[end]
	i,j:=start,start
	for ;j<end;j++{
		if arr[j]<pivot{
			arr[i],arr[j]=arr[j],arr[i]
			i++
		}
	}
	arr[i],arr[j]=arr[j],arr[i]
	return i
}

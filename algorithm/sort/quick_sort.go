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

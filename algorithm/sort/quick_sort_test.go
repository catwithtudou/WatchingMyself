package sort

import (
	"fmt"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/12/7
 **/

func TestQuickSort(t *testing.T) {
	arr:=[]int{2,3,5,1,6,10,0}
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}
func TestQuickSortA(t *testing.T) {
	arr:=[]int{2,3,5,1,6,10,0}
	QuickSortA(arr,0,len(arr)-1)
	fmt.Println(arr)
}


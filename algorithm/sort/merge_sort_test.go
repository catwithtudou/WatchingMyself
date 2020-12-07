package sort

import (
	"fmt"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/12/7
 **/

func TestMergeSort(t *testing.T) {
	arr:=[]int{10,3,4,1,5,6,8,11}
	fmt.Println(MergeSort(arr))
}
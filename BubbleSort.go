package main

import (
	"fmt"
)

func BubbleSort(sli []int,n int){
	for i:=0;i<n;i++ {
		for j:=1;j<n;j++ {
			if(sli[j-1]>sli[j]){
				swap(sli,j)
			}
		}
	}
}
func swap(sli []int,j int){
	temp:=sli[j]
	sli[j]=sli[j-1]
	sli[j-1]=temp
}
var n int
var num int
var arr_len int
func main() {
	fmt.Printf("Enter Length: ")
	fmt.Scan(&n)
	fmt.Printf("Enter %d Integers:",n)
	arr:=make([]int,n)
	for i:=0;i<n;i++ {
		fmt.Scan(&arr[i])
	}
	BubbleSort(arr,n)
	fmt.Printf("%d",arr)
}	

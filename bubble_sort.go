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
func main() {
	arr:= []int{5,6,4,3,2,10,1,8,7,9}
	n:=len(arr)
	BubbleSort(arr,n)
	fmt.Printf("%d",arr)
}	

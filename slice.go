package main
import (
	"fmt"
	"sort"
	"strconv"
)
var s string
var i int
func main(){
	sli:=make([]int,3)
	i=0
	for i >= 0{
		fmt.Printf("\nEnter Integer : ")
		fmt.Scan(&s)
		if(s=="X"){
			sort.Ints(sli)
			fmt.Printf("Final Slice : %d",sli)
			break
		}else{
			temp,err:=strconv.Atoi(s)
			if(err == nil){
				sli[i]=temp
			}
			i++
			if(i>=3){
				sort.Ints(sli)
				fmt.Printf("\nSlice : %d",sli)
				sli=append(sli,i+1)
			}
		}
	}
}
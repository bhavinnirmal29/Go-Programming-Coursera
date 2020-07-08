package main

import (
	"encoding/csv"
	"fmt"
	"os"
	)

type Person struct{
	fname string
	lname string
}

var file string
func main() {
	//idMap:=make(map[string]string)
	var p1 []Person
	var p2 Person
	fmt.Printf("\nEnter File Name : ")
	fmt.Scan(&file)
	f, err := os.Open(file)
	if err != nil	{
		fmt.Printf("Error:%s",err)
	}
	// Read lines into a variable
	reader := csv.NewReader(f)
	reader.Comma = ' '
	d, err := reader.ReadAll()

	if err != nil {
		fmt.Printf("Error : %s",err)
	}
	
	for _,v := range d{
		if(len(v[0]) < 20 && len(v[1])<20){
			p2.fname=v[0]
			p2.lname=v[1]
			p1 = append(p1,p2)
		}
	}
	for _,v := range p1{
		fmt.Printf("\nFName : %s LName : %s",v.fname,v.lname)
	}
}

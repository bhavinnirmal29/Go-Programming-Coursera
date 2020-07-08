package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)
type Person struct{
	name string
	addr string
}
var name1 string
var addr1 string
var p1 Person
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	idMap:=make(map[string]string)
	fmt.Printf("Enter First Name : ")
	scanner.Scan()
	name1 = scanner.Text()
	fmt.Printf("\nEnter Address : ")
	scanner.Scan()
	addr1 = scanner.Text()
	p1.name=name1
	p1.addr=addr1
	idMap["name"]=p1.name
	idMap["address"]=p1.addr
	barr,err:=json.Marshal(idMap)
	if(err!=nil){
		fmt.Printf("\nError %s",err)
	}
	fmt.Printf("\nJson Object : ")
	os.Stdout.Write(barr)
}

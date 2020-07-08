package main

import (
	"fmt"
	"os"
)

type thename struct {
	fname string
	lname string
}

func main() {
	var name string
	fmt.Printf("Enter name of your file: ")
	fmt.Scan(&name)
	var fname string
	var lname string
	myslice := make([]thename, 0)
	file, err := os.Create(name + ".txt")
	err = err
	for fname != "1" {
		fmt.Printf("Enter firstname: ")
		fmt.Scan(&fname)
		if fname == "1" {
			break
		}
		f1, err := file.WriteString(fname + " ")
		err = err
		f1 = f1
		fmt.Printf("Enter lastname: ")
		err = err
		fmt.Scan(&lname)
		f21, err := file.WriteString(lname)
		f2, err := file.WriteString("\n")
		f21 = f21
		f2 = f2
		n1 := thename{fname: fname, lname: lname}
		myslice = append(myslice, n1)
		fmt.Printf("%s\n", fname)
		fmt.Printf("%s\n", lname)
	}
}

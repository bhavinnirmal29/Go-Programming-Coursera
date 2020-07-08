package main

import (
	"encoding/csv"
	"fmt"
	"os"
	)

type Name struct {
	fname string
	lname string
}

func main() {

	//Gets the filename from user and stores it in fileName
	var fileName string
	fmt.Printf("Enter a filename to read names: ")
	fmt.Scan(&fileName)
	lines, err := ReadCsv(fileName)

	if err != nil	{
		panic(err)
	}
	var person Name
	var persons []Name

	// Loops through the returned type and stores data in person struct
	// The struct then is stored in a slice
	for _, name := range lines	{
		if len(name[0]) < 20 && len(name[1]) < 20 {
			person.fname = name[0]
			person.lname = name[1]
			persons = append(persons, person)
		}
	}
	// Looping on the struct to print the contents
	for _, fullname := range persons	{
		fmt.Println("First Name: " + fullname.fname + ", Last Name: " + fullname.lname)
	}
}

//ReadCsv accepts a filename and returns its contents as a multidimensional type
// With lines in each column is a string type
func ReadCsv(fileName string) ([][]string, error)	{

	f, err := os.Open(fileName)
	if err != nil	{
		return [][]string{}, err
	}
	defer f.Close()

	// Read lines into a variable
	reader := csv.NewReader(f)
	reader.Comma = ' '
	lines, err := reader.ReadAll()

	if err != nil	{
		return [][]string{}, err
	}
	// return the lines as multidimensional type
	return lines, nil
}
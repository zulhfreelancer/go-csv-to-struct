package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCSVFile(input string) (data [][]string, err error) {
	sf, err := os.Open(input)
	if err != nil {
		fmt.Println(err, "[", input, "]")
		log.Fatal(err, input)

	}
	data, err = csv.NewReader(sf).ReadAll()
	return
}

// Generate - call this from outside of this project
func main() {

	data, err := readCSVFile("prdata/pull-requests.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("type csvData struct {")
	for _, item := range data[0] {
		myItem := item
		if item == "#" {
			myItem = "ObjectID"
		}

		fmt.Printf("    %v string\n", myItem)
	}
	fmt.Println("}")
	fmt.Println()
	fmt.Println("for _, line := range data {")
	for num, item := range data[0] {
		myItem := item
		if item == "#" {
			myItem = "ObjectID"
		}

		fmt.Printf("    myObject.%v = line[%v]\n", myItem, num)
	}
	fmt.Println("}")
	fmt.Printf("\n// INSERT PROCESSING HERE ---\n")
}

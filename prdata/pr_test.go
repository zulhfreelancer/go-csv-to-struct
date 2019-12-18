package prdata

import (
	"fmt"
	"go-csv-to-struct/prdata"
	"testing"
)

func TestPR(t *testing.T) {
	data, err := prdata.ReadPRCSVFile("pull-requests.csv")
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range data {
		fmt.Println(obj.ObjectID)
	}
}

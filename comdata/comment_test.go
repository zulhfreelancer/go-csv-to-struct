package comdata

import (
	"csv-to-struct/comdata"
	"fmt"
	"testing"
)

func TestPR(t *testing.T) {
	data, err := comdata.ReadPRCSVFile("comments.csv")
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range data {
		fmt.Println(obj.User)
	}
}

package comdata

import (
	"csv-to-struct/comdata"
	"fmt"
	"testing"
)

func TestComment(t *testing.T) {
	data, err := comdata.ReadCommentCSVFile("comments.csv")
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range data {
		fmt.Println(obj.User)
	}
}

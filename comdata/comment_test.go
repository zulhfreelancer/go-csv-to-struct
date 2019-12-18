package comdata

import (
	"fmt"
	"github.com/zulhfreelancer/go-csv-to-struct/comdata"
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

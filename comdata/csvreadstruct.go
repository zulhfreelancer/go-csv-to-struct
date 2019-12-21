package comdata

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// CommentData - BitBucket comment data
type CommentData struct {
	Repository  string
	Type        string
	PRNumber    string
	User        string
	CommentType string
	BodyRaw     string
	BodyHTML    string
	CreatedAt   string
	IsDeleted   string
	ToLine      string
	FromLine    string
	FilePath    string
	Diff        string
	CommentID   string
}

// ReadCommentCSVFile - return array of objects
func ReadCommentCSVFile(input string) (objs []CommentData, err error) {
	var data [][]string
	sf, err := os.Open(input)
	if err != nil {
		fmt.Println(err, "[", input, "]")
		log.Fatal(err, input)

	}
	data, err = csv.NewReader(sf).ReadAll()
	for i, line := range data {
		if i == 0 {
			continue
		}

		var myObject CommentData
		myObject.Repository = line[0]
		myObject.Type = line[1]
		myObject.PRNumber = line[2]
		myObject.User = line[3]
		myObject.CommentType = line[4]
		myObject.BodyRaw = line[5]
		myObject.BodyHTML = line[6]
		myObject.CreatedAt = line[7]
		myObject.IsDeleted = line[8]
		myObject.ToLine = line[9]
		myObject.FromLine = line[10]
		myObject.FilePath = line[11]
		myObject.Diff = line[12]
		myObject.CommentID = line[13]
		objs = append(objs, myObject)
	}
	return
}

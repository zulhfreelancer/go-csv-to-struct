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
	CommentID   string
	BodyRaw     string
	BodyHTML    string
	CreatedAt   string
	IsDeleted   string
	ToLine      string
	FromLine    string
	FilePath    string
	Diff        string
	ParentID    string
	CommitHash  string
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
		myObject.CommentID = line[5]
		myObject.BodyRaw = line[6]
		myObject.BodyHTML = line[7]
		myObject.CreatedAt = line[8]
		myObject.IsDeleted = line[9]
		myObject.ToLine = line[10]
		myObject.FromLine = line[11]
		myObject.FilePath = line[12]
		myObject.Diff = line[13]
		myObject.ParentID = line[14]
		myObject.CommitHash = line[15]
		objs = append(objs, myObject)
	}
	return
}

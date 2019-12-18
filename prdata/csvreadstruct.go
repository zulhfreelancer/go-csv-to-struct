package prdata

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// PRData - BitBucket comment data
type PRData struct {
	Repository        string
	Type              string
	ObjectID          string
	User              string
	Title             string
	State             string
	CreatedAt         string
	UpdatedAt         string
	URL               string
	BodyRaw           string
	BodyHTML          string
	SourceCommit      string
	DestinationCommit string
	SourceBranch      string
	DestinationBranch string
	DeclineReason     string
	MergeCommit       string
}

// ReadPRCSVFile - return array of objects
func ReadPRCSVFile(input string) (objs []PRData, err error) {
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

		var myObject PRData
		myObject.Repository = line[0]
		myObject.Type = line[1]
		myObject.ObjectID = line[2]
		myObject.User = line[3]
		myObject.Title = line[4]
		myObject.State = line[5]
		myObject.CreatedAt = line[6]
		myObject.UpdatedAt = line[7]
		myObject.URL = line[8]
		myObject.BodyRaw = line[9]
		myObject.BodyHTML = line[10]
		myObject.SourceCommit = line[11]
		myObject.DestinationCommit = line[12]
		myObject.SourceBranch = line[13]
		myObject.DestinationBranch = line[14]
		myObject.DeclineReason = line[15]
		myObject.MergeCommit = line[16]
		objs = append(objs, myObject)
	}
	return
}

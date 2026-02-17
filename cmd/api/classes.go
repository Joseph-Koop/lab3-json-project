package main

import (
	//   "encoding/json"
	"fmt"
	"net/http"
	// import the data package which contains the definition for Comment
	// "github.com/Joseph-Koop/json-project/internal/data"
)

func (a *applicationDependencies) postClassHandler(w http.ResponseWriter,
	r *http.Request) {
	// create a struct to hold a comment
	// we use struct tags[``] to make the names display in lowercase
	var incomingData struct {
		Studio  string `json:"studio"`
		Trainer  string  `json:"trainer"`
		Capacity_limit  int64 `json:"capacity_limit"`
		Membership_tier  string `json:"membership_tier"`
		Name  string `json:"name"`
		Status  string `json:"status"`
	}
	// perform the decoding
	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	// for now display the result
	fmt.Fprintf(w, "+%v\n", incomingData)
}

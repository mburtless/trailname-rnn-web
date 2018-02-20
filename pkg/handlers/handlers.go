package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"bytes"
)

type TrailName struct {
	Result	[]string		`json:"result,omitempty"`
}

func GetTestTrailName(w http.ResponseWriter, r *http.Request) {
	//log.Println("Responding to GET request")
	trailNames := TrailName{Result: []string{"GlacialFreeze", "RiptideRush"}}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(trailNames); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func GetTrailName(w http.ResponseWriter, r *http.Request) {
	//log.Println("Responding to GET request")
	err := r.ParseForm()

	if err != nil {
		log.Fatalf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	startText := r.Form.Get("starttext")

	trailNames := apiNameReq(startText)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(trailNames); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// apiNameReq sends the passed startText via JSON request to the backend API
// and returns an array of trailnames returned by the API
func apiNameReq(startText string) TrailName {
	log.Printf("Recieved request for trailname with starttext of %s", startText)
	// Contstruct our json request, convert to bytes and send via post to api
	jsonString := `{"starttext":"` + startText  + `"}`
	jsonBytes := []byte(jsonString)
	log.Printf("Request is %+v", jsonString)
	var n TrailName
	res, err := http.Post("http://127.0.0.1:6788/api", "application/json; charset=utf-8", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Decode the api response
	if err := json.NewDecoder(res.Body).Decode(&n); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("First returned  name is %s", n.Result[1])
	//trailNames := TrailName{Result: []string{"GlacialFreeze", "RiptideRush"}}
	return n
}

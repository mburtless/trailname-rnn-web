package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"bytes"
	"github.com/mburtless/trailname-rnn-web/pkg/namerank"
    "github.com/AntoineAugusti/wordsegmentation/corpus"
	"github.com/mburtless/trailname-rnn-web/pkg/configs"
	"time"
)

type TrailName struct {
	Result	[]string		`json:"result,omitempty"`
}

type PollRequest struct {
	Poll	string			`json:"poll"`
}

type pollResponse struct {
	Result	bool			`json:"result"`
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/pages/index.html")
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	/*contentType := r.Header.Get("Content-type")
	if contentType == "application/x-www-form-urlencoded" {
		GetTrailName(w, r)
	}*/
	switch contentType := r.Header.Get("Content-type"); contentType {
	case "application/x-www-form-urlencoded; charset=UTF-8":
		GetTrailName(w, r)
	case "application/json; charset=UTF-8":
		PollAPI(w, r)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If content is JSON with attribute poll=true
	// call function to poll the API, else if content is form
	// submit to GetTrailName(), else return error
}

func PollAPI(w http.ResponseWriter, r *http.Request) {
	var pr PollRequest
	var jsonResponse pollResponse
	if r.Body == nil {
		http.Error(w, "Please submit a JSON request", 400)
		return
	}
	// Inspect JSON, make sure poll is true
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// If poll attribute was set to true
	if pr.Poll == "true" {
		result := sendPoll(configs.ConfigVars["apiHost"].ParsedVal)
		if result {
			jsonResponse = pollResponse{Result: true}
		} else {
			jsonResponse = pollResponse{Result: false}
		}
		// Generate JSON response and reply to ajax request
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}

func sendPoll(dest string) bool {
	log.Printf("Send poll to %s", dest)
	if len(dest) < 1 {
		log.Printf("Polling destination must be provided")
		return false
	}
	// Craft a http request to dest
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
    rs, err := client.Get("http://" + dest)
    // Process response
    if err != nil {
		log.Printf("Error durring poll %v", err)
		return false
	}
    defer rs.Body.Close()
	log.Printf("StatusCode of poll response was %d", rs.StatusCode)
	// If header of response is 200, return true
	if rs.StatusCode == 200 {
		return true
	}
	// Anything else, including timeout, return false
	return false
}

func GetTrailName(w http.ResponseWriter, r *http.Request) {
	//log.Println("Responding to GET request")
	err := r.ParseForm()

	if err != nil {
		log.Fatalf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	startText := r.Form.Get("starttext")

	trailNames := apiNameReq(startText)
	trailNames.filterNames()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(trailNames); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// filterNames method itterates through TrailNames struct and removes
// names that are nonsensical
func (tn *TrailName) filterNames() {
	finalTn := TrailName{Result: []string{}}
	englishCorpus := corpus.NewEnglishCorpus()
	for _, name := range tn.Result {
		//log.Printf("Language of %s is %v", name, detector.GetClosestLanguage(name))
		//tn.Result[i] = name + "_test"
		//score := checkSubstrings(name)
		//seg, score := wordsegmentation.Segment(englishCorpus, name)
		//log.Printf("Seg of %s is %v and score is ", name, seg, score)
		score := namerank.SegmentAndRank(name, &englishCorpus)
		if score >= float64(-10) {
			finalTn.Result = append(finalTn.Result, name)
		}
	}
	if len(finalTn.Result) > 0 {
		log.Printf("Names to keep are %v", finalTn)
	} else {
		log.Printf("namerank didn't rank any names high enough to use")
	}
	*tn = finalTn
}


func checkSubstrings (name string) bool {
	length := len(name)
	//return base case
	if length == 0 {
		return true
	}
	/*for i := 0; i < length; i++ {
		if dictContains(name[0:i]) && checkSubstring(name[i:length-i]) {
			return true
		}
	}*/

	return false
	/*for i := 0; i < length; i++ {
		for j := i+3; j < length; j++ {
			substring := name[i:j + 1]
			log.Printf("%s", substring)
		}
	}
	return score*/
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
	res, err := http.Post("http://" + configs.ConfigVars["apiHost"].ParsedVal + "/api", "application/json; charset=utf-8", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Decode the api response
	if err := json.NewDecoder(res.Body).Decode(&n); err != nil {
		log.Fatalf("Error: %v", err)
	}
	//log.Printf("First returned  name is %s", n.Result[0])
	//trailNames := TrailName{Result: []string{"GlacialFreeze", "RiptideRush"}}
	return n
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abhilashr1/linearequations/compute"
	"github.com/abhilashr1/linearequations/stringfilter"

	"github.com/gorilla/mux"
)

type inp struct {
	First  string
	Second string
}

// responseOk displays the Http Response with the output back to the user
func responseOk(w http.ResponseWriter, res inp, vars []string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		vars[0]: res.First,
		vars[1]: res.Second,
	}
	json.NewEncoder(w).Encode(body)
}

// responseError displays the formatted error message in case of invalid input or error.
func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}

// solvelinear is the main integration function which :
// 1. Processes the string and captures the co-efficients and variables.
// 2. Computes the solution
// 3. Displays the relevant result or error, as needed.
func solvelinear(w http.ResponseWriter, r *http.Request) {
	input := inp{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		// If there are problems with json formatting, let the user know
		responseError(w, http.StatusBadRequest, "Invalid request payload. Please enter the data in the form of 	{\"First\": \"ax + by = c\" , \"Second\": \"ax + by = c\" }")
		return
	}

	filtered, err := stringfilter.Filter(input.First, input.Second)
	if err != nil {
		// If there are problems with json formatting, let the user know
		responseError(w, http.StatusBadRequest, `Invalid request payload. Please enter the data in the form of {"First": "ax + by = c" , "Second": "ax + by = c" }`)
		return
	}

	vars := stringfilter.GetVariables(input.First)

	d, d1, d2 := compute.Coeffmatrix(filtered)

	var result inp
	first, second := compute.Solution(d, d1, d2)
	result.First = strconv.FormatFloat(float64(first), 'f', 6, 32)
	result.Second = strconv.FormatFloat(float64(second), 'f', 6, 32)

	if result.First == "-9999.000000" || result.Second == "-9999.000000" {
		responseError(w, http.StatusBadRequest, `Solution cannot be found for this equation`)
		return
	}

	responseOk(w, result, vars)
}

// responseOk displays the Http Response with the output back to the user
func getHelp(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `To begin, send a POST call to /api/linear with body in the following format: 
	{
		"First" : "ax+by=c",
		"Second": "px+qy=r"
	}
	
	For example, 
	{		
		"First" : "x+y=5",
		"Second": "x-y=3"
	}
		`)
}

// Routing function which listens to POST calls at 80
func main() {
	log.Println("Starting Server")
	r := mux.NewRouter()
	r.HandleFunc("/api/linear", solvelinear).Methods("POST")
	r.HandleFunc("/api/help", getHelp).Methods("GET")

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}

}

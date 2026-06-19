package asciiweb

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

)

func EndPoints(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "pong")
}

func GetName(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "Hello, Guest!", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Hello, "+name+"!")
	}

}

func Count(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Send a POST request with text to count words")
	}

	if r.Method == http.MethodPost {
		wordlen, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Empty file", http.StatusBadRequest)
		}
		cnvt := len(string(wordlen))

		fmt.Fprintln(w, cnvt)
	} 
}


func Calculate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var result int

	op := r.URL.Query().Get("op")
	a, err1 := strconv.Atoi(r.URL.Query().Get("a"))

	if err1 != nil {
		http.Error(w, "Invalid Operation", 400)
	}

	b, err2 := strconv.Atoi(r.URL.Query().Get("b"))
	if err2 != nil {
		http.Error(w, "Invalid Operation", 400)
	}

	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "divive":
		result = a / b
	}	
	fmt.Fprintf(w, "Result: %d", result)
}



func Agent(w http.ResponseWriter, r *http.Request){

	
	userInfo := r.Header.Get("User-Agent")

	// if userInfo == ""{
	// 	http.Error(w, "Not found", 404)
	// 	return
	// }

	fmt.Fprintf(w, "You are visiting us using: %s", userInfo)

	

	
}
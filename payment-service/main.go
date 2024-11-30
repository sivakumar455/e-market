package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct to hold configuration
type Config struct {
	Port string `yaml:"port"`
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "this home page")
	producer()
	fmt.Fprintf(w, "Done publishign")
}

func consumeMsg(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Consuming Message \n")
	go func() {
		consumer()
	}()
	fmt.Fprintf(w, "\nDone Consuming")
}

// Handler for the query parameters
func queryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	age := queryParams.Get("age")

	// Respond with the parsed parameters
	fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
}

// Handler for form data
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		r.ParseForm()
		name := r.FormValue("name")
		age := r.FormValue("age")

		// Respond with the parsed form data
		fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
	} else {
		// Respond with a simple form for testing
		fmt.Fprintf(w, `<form method="POST" action="/form">
            Name: <input type="text" name="name">
            Age: <input type="text" name="age">
            <input type="submit">
        </form>`)
	}
}

func main() {
	fmt.Println("In Main")
	// register patterns
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/consume", consumeMsg)

	// Read the configuration file
	file, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Unmarshal the YAML content into the Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Println("Error unmarshalling config file:", err)
		return
	}

	// Start the server on port 8080
	fmt.Println("Starting server on :", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}

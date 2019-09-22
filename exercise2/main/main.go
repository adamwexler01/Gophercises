package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"flag"

	"github.com/gophercises/exercise2"
)

func main() {
	mux := defaultMux()

	yamlFileName := flag.String("yaml", "example.yml", "provide an appropriate any *.yml or *.yaml and its name as input")
	jsonFileName := flag.String("json", "example.json", "provide an appropriate any *.json for its name as input")
	flag.Parse()

	json, error := ioutil.ReadFile(*jsonFileName)
	if(error != nil){
		panic(error)
	}
	yaml, error := ioutil.ReadFile(*yamlFileName)
	if error != nil {
		panic(error)
	}

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	//Build the JSONHandler using the mapHandler as the fallback
	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"flag"
	"log"

	"github.com/gophercises/exercise2"
)

func main() {
	mux := defaultMux()

	yamlFileName := flag.String("yaml", "example.yml", "provide an appropriate yml file and its name as input")
	flag.Parse()
	yaml, error := ioutil.ReadFile(*yamlFileName)
	if error != nil {
		log.Fatal("Can't read the incoming yml file")
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

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

package main

import (
	"flag"
	"fmt"
	"net/http"
)

const DEFAULT_FILENAME = "paths.yml"

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/maria":          "https://mariainesserra.com",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	fn := flag.String("yaml", DEFAULT_FILENAME, "the paths file!")

	yamlHandler, err := YAMLHandler(fn, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

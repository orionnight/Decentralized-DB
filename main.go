package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/ece1770/router"
	"github.com/gorilla/mux"
)

func urlCheck(endpoint func(http.ResponseWriter, *http.Request, map[string]interface{}), url string, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != url {
			http.Error(w, "404 not found.", http.StatusNotFound)
		} else if r.Method != method {
			fmt.Fprintf(w, "Method err: Only %s is supported for url: %s", method, url)
		} else if method == "POST" {
			if r.Body != nil {
				defer r.Body.Close()
			}

			body, readErr := ioutil.ReadAll(r.Body)
			if readErr != nil {
				fmt.Fprintf(w, "Cannot read from body err: %v\n", readErr)
				return
			}
			var result map[string]interface{}
			err := json.Unmarshal([]byte(body), &result)
			if err != nil {
				fmt.Fprintf(w, "Wrong JSON format() err: %v\n", err)
				return
			}
			endpoint(w, r, result)
		} else {
			endpoint(w, r, nil)
		}
	})
}

func main() {
	r := mux.NewRouter()
	r.Handle("/db/create", urlCheck(router.DBCreate, "/db/create", "POST")).Methods("POST")
	r.Handle("/create_ks", urlCheck(router.CreateKeystore, "/create_ks", "POST")).Methods("POST")

	fmt.Printf("Starting server at port 8080\n")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())
}

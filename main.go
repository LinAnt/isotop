package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type state struct {
	currentTemp float32
	stableTemp  float32
	isOpen      bool
}

func main() {
	var port int

	flag.IntVar(&port, "p", 8080, "Port to listen to")
	flag.Parse()

	// Initiate router, be harsh to
	// people who can't type URL's
	router := &httprouter.Router{
		RedirectTrailingSlash:  false,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: true,
		// PanicHandler:           errors.PanicHandler,
		// NotFound: new(errors.NotFoundHandler),
	}

	router.Handle("GET", "/dummy", dummyHandler)

	log.Printf("listening on 0.0.0.0:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router))
}

func dummyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Print()
	dummy := state{
		currentTemp: rand.Float32()*10 + 10,
		stableTemp:  rand.Float32()*10 + 10,
		isOpen:      (rand.Intn(100) < 50),
	}
	b, err := json.Marshal(dummy)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var (
	overallState = state{}
)

type state struct {
	CurrentTemp float32
	StableTemp  float32
	IsOpen      bool
	StartTime   time.Time
}

func main() {
	var port int

	flag.IntVar(&port, "p", 8080, "Port to listen to")
	flag.Parse()
	probe := thermoprobe.NewPT100()
	overallState.CurrentTemp = thermoprobe.Read()
	overallState.StableTemp = nil
	overallState.IsOpen = false
	overallState.CurrentTemp = 0

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
	router.Handle("GET", "/current", currentHandler)

	log.Printf("listening on 0.0.0.0:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router))
}

func currentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	stateAsJSON, err := json.Marshal(overallState)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(stateAsJSON))
}

func dummyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dummy := &state{
		CurrentTemp: rand.Float32()*10 + 70,
		StableTemp:  rand.Float32()*10 + 70,
		IsOpen:      (rand.Intn(100) < 50),
		StartTime:   time.Now(),
	}
	b, err := json.Marshal(dummy)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", b)
}

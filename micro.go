package ServeAPI

import "net/http"

// go get github.com/gorilla/mux

//	"net/http"
//import (
//	"get github.com/gorilla/mux"
//)

http.ListenAndServe(":8181",r)

func ServeAPI(endpoint string) error {
	handler := &eventServicehandler{}
	r := mux.NewRouter()

	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}

type eventServicehandler struct {}

func (eh *eventServicehandler) findEventHandler (w http.ResponseWriter,r *http.Request) {}
func (eh *eventServicehandler) allEventHandler (w http.ResponseWriter,r *http.Request) {}
func (eh *eventServicehandler) newEventHandler (w http.ResponseWriter,r *http.Request) {}
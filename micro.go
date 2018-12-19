package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Event struct {
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  *Location
}
type Location struct {
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

type Hall struct {
	Name     string `json:"Name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}

var events []Event

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/events", GetAllEvent).Methods("GET")
	router.HandleFunc("/event/{id}", GetEventWithID).Methods("GET")
	router.HandleFunc("/people/{name}", GetEventWithName).Methods("GET")
	router.HandleFunc("", CreateEvent).Methods("POST")

	events = append(events, Event{Name: "Test", Duration: 12, EndDate: 5492456, Location: &Location{Name: "Test", Address: "sdfsd", CloseTime: 895425, Country: "Iran", OpenTime: 54264646}})
	events = append(events, Event{Name: "Test", Duration: 12, EndDate: 5492456, Location: &Location{Name: "Test", Address: "sdfsd", CloseTime: 895425, Country: "Iran", OpenTime: 54264646}})
	events = append(events, Event{Name: "Test", Duration: 12, EndDate: 5492456, Location: &Location{Name: "Test", Address: "sdfsd", CloseTime: 895425, Country: "Iran", OpenTime: 54264646}})
	events = append(events, Event{Name: "Test", Duration: 12, EndDate: 5492456, Location: &Location{Name: "Test", Address: "sdfsd", CloseTime: 895425, Country: "Iran", OpenTime: 54264646}})
	http.ListenAndServe(":8000", router)
}

func GetAllEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}
func GetEventWithID(w http.ResponseWriter, r *http.Request)   {}
func GetEventWithName(w http.ResponseWriter, r *http.Request) {}
func CreateEvent(w http.ResponseWriter, r *http.Request)      {}

//http.ListenAndServe(":8181",r)
//func Main(endpoint string) error {
//	r := mux.NewRouter()
//	eventsrouter := r.PathPrefix("/events").Subrouter()
//	eventsrouter.Methods("GET").Path("/{searchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
//	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
//	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
//	return http.ListenAndServe(endpoint, r)
//
//}
//type	eventserviceHandler struct {}
//
//func (eh *eventserviceHandler) findEventHandler(w http.ResponseWriter,r http.Request){}
//func (eh *eventserviceHandler) allfEventHandler(w http.ResponseWriter,r http.Request){}
//func (eh *eventserviceHandler) newEventHandler(w http.ResponseWriter,r http.Request){}

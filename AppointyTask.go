package main

import (
	"fmt"
	"html/template"
	"net/http"
	"encoding/json"
)
type Message struct {
	Title string
	Para string
	Para1 string
}
type Participant struct {
    Name string `json:"Name"`
    Email string `json:"Email"`
    RSVP string `json:"RSVP"`
}

var ParticipantDetails []Participant
func allParticipants(w http.ResponseWriter, r *http.Request){
	ParticipantDetails = []Participant{
        Participant{Name: "Rose", Email: "rose.@gmail.com", RSVP: "Maybe"},
        
    }
    fmt.Println("Endpoint Hit: allParticipants")
    json.NewEncoder(w).Encode(ParticipantDetails)
}
func homePage(w http.ResponseWriter, r *http.Request){
	v := Message{Title: "WELCOME TO MEETINGS", Para: "This is the homepage of meetings app.",Para1: " Use this link http://localhost:8080/participant"}
	p, _:= template.ParseFiles("Appointy.html")
	p.Execute(w,v)
	
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/participant", allParticipants)
    http.ListenAndServe(":8080", nil)
}

func main() {
    handleRequests()
}

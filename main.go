package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	s "nwHacks-backend/scraper"
	"os"
)

type Job struct {
	Company  string `json: "company"`
	Position string `json: "position"`
	Location string `json: "location"`
	Salary   int    `json: "salary"`
	Link     string `json: "link"`
	Visa     string `json: "visa"`
	Open     bool   `json: "open"`
}

type UserData struct {
	Password string `json:"password"`
	Applied  []int  `json:"applied"`
}

type Users map[string]UserData

func fetch(w http.ResponseWriter, req *http.Request) {

	jsonFile, err := os.Open("jobs.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	w.Write(byteValue)

}

func addUser(w http.ResponseWriter, req *http.Request) {

}

func login(w http.ResponseWriter, req *http.Request) {

}

func apply(w http.ResponseWriter, req *http.Request) {

}

func generateResume(w http.ResponseWriter, req *http.Request) {

}

func main() {
	s.ScrapeGithub()
	http.HandleFunc("/fetch", fetch)
	http.HandleFunc("/adduser", addUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/generateresume", generateResume)
}

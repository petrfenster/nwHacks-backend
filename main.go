package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	os.Chdir("resources")
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

func homePage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome!"))
}

func setUp() {

	s.ScrapeGithub()

	jsonFile, err := os.Open("../resources/githubJobs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	github := []s.GithubJobStructure{}
	json.Unmarshal(byteValue, &github)

	companies := []string{}

	for _, s := range github {
		companies = append(companies, s.Company)
	}

	s.ScrapeLevels(companies)

	fmt.Println(companies)

	//jsonFile, err = os.Open("resources/levels.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer jsonFile.Close()
	//byteValue, _ = ioutil.ReadAll(jsonFile)
	//var levels Levels
	//json.Unmarshal(byteValue, &levels)

}

func main() {
	setUp()

	http.HandleFunc("/fetch", fetch)
	http.HandleFunc("/adduser", addUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/generateresume", generateResume)
	http.HandleFunc("/homepage", homePage)

	fmt.Println("Up and running")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	s "nwHacks-backend/scraper"
	"os"
	"strconv"
)

type Job struct {
	Company  string      `json: "company"`
	Position string      `json: "position"`
	Location string      `json: "location"`
	Salary   int         `json: "salary"`
	Link     string      `json: "link"`
	Visa     bool        `json: "visa"`
	Open     bool        `json: "open"`
	LeetCode [][2]string `json: "leetcode"`
}

type UserData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Applied  []int  `json:"applied"`
	Saved    []int  `json:"saved"`
}

type Jobs map[string]Job

type Users map[string]UserData

func fetch(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	jsonFile, err := os.Open("resources/jobs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	w.Write(byteValue)
}

func addUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	userName := req.FormValue("username")
	password := req.FormValue("password")
	name := req.FormValue("name")

	jsonFile, err := os.Open("resources/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	userData := UserData{}
	userData.Name = name
	userData.Password = password
	users[userName] = userData

	file, _ := json.MarshalIndent(users, "", "  ")
	_ = ioutil.WriteFile("resources/users.json", file, 0644)

	w.Write([]byte("User " + userName + " has been succesfully added"))
}

func login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	userName := req.FormValue("username")
	password := req.FormValue("password")

	jsonFile, err := os.Open("resources/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	if users[userName].Password == password {
		w.Write([]byte("Login Successful"))
	} else {
		w.Write([]byte("Login Failed"))
	}
}

func apply(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	userName := req.FormValue("username")
	jobId, _ := strconv.Atoi(req.FormValue("jobid"))

	jsonFile, err := os.Open("resources/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	user := users[userName]
	user.Applied = append(users[userName].Applied, jobId)
	users[userName] = user

	file, _ := json.MarshalIndent(users, "", "  ")
	_ = ioutil.WriteFile("resources/users.json", file, 0644)

	w.Write([]byte("Successfuly marked the job as applied"))
}

func getApplied(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	userName := req.FormValue("username")

	jsonFile, err := os.Open("resources/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	jsonFile, err = os.Open("resources/jobs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	var jobs Jobs
	json.Unmarshal(byteValue, &jobs)

	appliedIds := users[userName].Applied
	appliedJobs := []Job{}

	for _, id := range appliedIds {
		appliedJobs = append(appliedJobs, jobs[strconv.FormatInt(int64(id), 10)])
	}

	file, _ := json.MarshalIndent(appliedJobs, "", "  ")
	w.Write(file)
}

func getSaved(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	userName := req.FormValue("username")

	jsonFile, err := os.Open("resources/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	jsonFile, err = os.Open("resources/jobs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	var jobs Jobs
	json.Unmarshal(byteValue, &jobs)

	savedIds := users[userName].Saved
	savedJobs := []Job{}

	for _, id := range savedIds {
		savedJobs = append(savedJobs, jobs[strconv.FormatInt(int64(id), 10)])
	}

	file, _ := json.MarshalIndent(savedJobs, "", "  ")
	w.Write(file)
}

func saveJob(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	userName := req.FormValue("username")
	jobId, _ := strconv.Atoi(req.FormValue("jobid"))

	jsonFile, err := os.Open("resources/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	user := users[userName]
	user.Saved = append(users[userName].Saved, jobId)
	users[userName] = user

	file, _ := json.MarshalIndent(users, "", "  ")
	_ = ioutil.WriteFile("resources/users.json", file, 0644)

	w.Write([]byte("Successfuly marked the job as saved"))
}

func uploadResume(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

}

func generateResume(w http.ResponseWriter, req *http.Request) {

}

func homePage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome!"))
}

func (jobs Jobs) parseGithub(github []s.Job) []string {

	companies := []string{}

	for i, s := range github {
		companies = append(companies, s.Company)
		key := strconv.FormatInt(int64(i), 10)
		job := Job{}

		job.Company = s.Company
		job.Position = s.JobRole
		job.Visa = true

		if s.Status == "Closed" {
			job.Open = false
		} else {
			job.Open = true
		}

		job.Link = s.Link
		job.Location = s.Location

		jobs[key] = job
	}

	return companies
}

func (jobs Jobs) getLeetcode() {
	for key, element := range jobs {
		csvFile, err := os.Open("resources/leetcode/" + element.Company + ".csv")

		if err != nil {
			fmt.Println(err)
		}
		defer csvFile.Close()

		fileReader := csv.NewReader(csvFile)
		records, err := fileReader.ReadAll()

		leetCode := [][2]string{}

		if len(records) > 0 {
			for i := 1; i < len(records); i++ {
				entry := [2]string{records[i][0], records[i][1]}
				leetCode = append(leetCode, entry)
			}
		}

		element.LeetCode = leetCode

		jobs[key] = element
	}
}

func setUp() {

	jobs := Jobs{}
	jobData := s.ScrapeGithub()
	companies := jobs.parseGithub(jobData)
	fmt.Println(len(companies))

	jobs.getLeetcode()
	file, _ := json.MarshalIndent(jobs, "", "  ")
	_ = ioutil.WriteFile("resources/jobs.json", file, 0644)
}

func main() {
	setUp()

	http.HandleFunc("/fetch", fetch)
	http.HandleFunc("/adduser", addUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/generateresume", generateResume)
	http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/getsaved", getSaved)
	http.HandleFunc("/getapplied", getApplied)
	http.HandleFunc("/savejob", saveJob)
	http.HandleFunc("/uploadresume", uploadResume)

	fmt.Println("Up and running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

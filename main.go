package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net"
	"net/http"
	"os"
	"time"
)

type Job struct {
	Company  string `json:"companyName"`
	JobRole  string `json:"jobRole"`
	Location string `json:"location"`
	Link     string `json:"link"`
	Status   string `json:"status"`
}

func main() {

	c := colly.NewCollector()

	c.WithTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   90 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	var jobData []Job

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})
	c.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			tableData := Job{
				Company:  el.ChildText("td:nth-child(1)"),
				Location: el.ChildText("td:nth-child(2)"),
				Status:   el.ChildText("td:nth-child(3)"),
				JobRole:  el.ChildText("td:nth-child(4)"),
				Link:     el.ChildText("td:nth-child(1):link"),
			}
			fmt.Println(el.ChildText("td:nth-child(1):link"))
			jobData = append(jobData, tableData)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//ignored error fix later
	_ = c.Visit("https://github.com/bsovs/Fall2023-Internships/blob/main/Fall2022/README.md")

	content, err := json.Marshal(jobData)
	if err != nil {
		fmt.Println(err.Error())
	}
	os.WriteFile("githubsJobs.json", content, 0644)
	fmt.Println("Total jobs: ", len(jobData))

}

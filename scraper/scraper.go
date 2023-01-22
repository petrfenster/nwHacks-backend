package scraper

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type GithubJobStructure struct {
	Company  string `json:"companyName"`
	JobRole  string `json:"jobRole"`
	Location string `json:"location"`
	Link     string `json:"link"`
	Status   string `json:"status"`
}

type LevelsJobStructure struct {
	Company string `json:"companyName"`
	Pay     string `json:"pay"`
}

func ScrapeGithub() {

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

	var jobData []GithubJobStructure

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})
	c.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var link string
			el.ForEach("a", func(_ int, el *colly.HTMLElement) {
				link = el.Attr("href")
			})
			tableData := GithubJobStructure{
				Company:  el.ChildText("td:nth-child(1)"),
				Location: el.ChildText("td:nth-child(2)"),
				Status:   el.ChildText("td:nth-child(3)"),
				JobRole:  el.ChildText("td:nth-child(4)"),
				Link:     link,
			}
			jobData = append(jobData, tableData)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	_ = c.Visit("https://github.com/bsovs/Fall2023-Internships/blob/main/README.md")

	content, err := json.Marshal(jobData)
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Chdir("resources")
	os.WriteFile("githubJobs.json", content, 0644)
	fmt.Println("Total jobs: ", len(jobData))

}

func ScrapeLevels(companies []string) {
	// Initialize a new collector
	c := colly.NewCollector()

	// Set the callback function
	c.OnHTML("title", func(e *colly.HTMLElement) {
		completeText := e.Text

		splitText := strings.Split(completeText, " ")

		companyName := splitText[0]

		length := len(splitText) - 5

		//Amazon Software Engineer Intern Salaries | $69.70 / hr | Levels.fyi
		pay := splitText[length]

		fmt.Println(companyName)
		fmt.Println(pay)

	})

	for index := range companies {
		companyNameHyphenated := companies[index]
		companyNameHyphenated = strings.Replace(companyNameHyphenated, " ", "_", -1)
		_ = c.Visit("https://www.levels.fyi/internships/" + companyNameHyphenated + "/Software-Engineer-Intern/")

	}

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

}

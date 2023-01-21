package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"strconv"
)

type Job struct {
	Company           string `json:"companyName"`
	JobRole           string `json:"jobRole"`
	Location          string `json:"location"`
	Link              string `json:"link"`
	ApplicationPeriod bool   `json:"open"`
}

func main() {
	allFacts := make([]Job, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("Fall2023", "https://github.com/bsovs/Fall2023-Internships"),
	)

	collector.OnHTML(".markdown-body entry-content container-lg tr", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("td"))
		if err != nil {
			log.Println("Could not get id")
		}
		factDesc := element.Text

		fmt.Println(factId)
		fmt.Println(factDesc)
		/*
			fact := Fact{
				ID:          factId,
				Description: factDesc,
			}

			allFacts = append(allFacts, fact)

		*/
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://github.com/bsovs/Fall2023-Internships")

	writeJSON(allFacts)
}

func writeJSON(data []Job) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("jobs.json", file, 0644)
}

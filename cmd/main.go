package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const jsonPath = "/Users/che-chia/Downloads/taipei-arena-iceland-63dc831f0e4c.json"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("api.metro.taipei"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("span", func(e *colly.HTMLElement) {
		ctx := context.Background()

		//srv, err := sheets.NewService(ctx, option.WithCredentialsFile(jsonPath))
		srv, err := sheets.NewService(ctx, option.WithCredentialsFile(jsonPath))
		if err != nil {
			log.Fatalf("Unable to retrieve Sheets client: %v", err)
		}

		// fetch data from web page
		id := e.Attr("id")
		value, err := strconv.Atoi(e.Text)
		if err != nil {
			log.Fatal(err)
		}
		switch id {
		case "LabelServiceWaitNumber":
			log.Printf("Wait number: %d\n", value)
		case "LabelWaitQueueGroups":
			log.Printf("Wait queue: %d\n", value)
		default:
			log.Printf("Unknown span id: %s value: %d\n", id, value)
		}

		// TODO: write to google sheet
		// date / epoch / data
		spreadsheetId := "1e3N79TcR68JB7iOUI97T3UUXjh45zjnomr2ksHwQugg"
		readRange := "Class Data!A2:E"
		resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}

		if len(resp.Values) == 0 {
			fmt.Println("No data found.")
		} else {
			fmt.Println("Name, Major:")
			for _, row := range resp.Values {
				// Print columns A and E, which correspond to indices 0 and 4.
				fmt.Printf("%s, %s\n", row[0], row[4])
			}
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://api.metro.taipei/taipeiarenainfoboard/queuenumber.aspx")
}

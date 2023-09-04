package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type JSONData struct {
	Data struct {
		EventCards map[string]EventCardsData
		EventInfo  EventInfoData
	}
}

type EventInfoData struct {
	MetaFromDE MetaDEData
}

type MetaDEData struct {
	Title string
}

type EventCardsData map[string]DateData

type DateData map[string]TimeData

type TimeData map[string]LayoutData

type LayoutData struct {
	AvailableSeats int
	TotalSeats     int
	PriceDesc      string
	Price          string
}

const apiURL = "https://in-cdn.bmscdn.com/le-synopsis/leapi-eventinfo-"

var matchNameMap = map[string]string{
	"ET00367559": "Ind vs Pak",
	"ET00367578": "Ind vs Eng",
	"ET00367551": "Ind vs Aus",
	"ET00367570": "Ind vs Nz",
	"ET00367587": "Ind vs Sa",
	"ET00367583": "Ind vs Sl",
	"ET00367556": "Ind vs Afg",
	"ET00367564": "Ind vs Ban",
	"ET00367595": "Ind vs Ned",
	"ET00367532": "Ind vs Ned Practice Match",
	// "ET00367548": "Eng vs NZ",
	// "ET00367569": "Eng vs SA",
}

func main() {
	// printAllMatchData()
	printPerMatchData()
}

func printPerMatchData() {
	matchIds := []string{}

	options := ""
	i := 1
	for id, name := range matchNameMap {
		options += fmt.Sprintf("%v - %v\n", i, name)
		i++
		matchIds = append(matchIds, id)
	}
	totalList := len(matchIds)

	for {
		fmt.Printf("\n%v\n", options)
		fmt.Printf("Select match (1-%v): ", totalList)
		data := ""
		fmt.Scanln(&data)
		fmt.Printf("\nSelected match (1-%v): %v\n", totalList, data)
		index, err := strconv.Atoi(data)
		if err != nil || index > len(matchIds) || index < 1 {
			fmt.Println("choose a valid number")
			continue
		}
		loadMatchData(matchIds[index-1])
	}
}

func printAllMatchData() {
	for {
		for id := range matchNameMap {
			loadMatchData(id)
		}
		fmt.Print("\nEnter to refresh\n")
		data := ""
		fmt.Scanln(&data)
	}
}

func loadMatchData(id string) {
	reset := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"

	apiResponse, err := http.Get(fmt.Sprintf("%v%v.json?%v", apiURL, id, rand.Float64()))
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer apiResponse.Body.Close()

	byteValue, _ := ioutil.ReadAll(apiResponse.Body)

	var jsonData JSONData
	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		fmt.Println("Error parsing API response:", err)
		return
	}

	fmt.Printf("\n%v\n\n", strings.Split(jsonData.Data.EventInfo.MetaFromDE.Title, "-")[0])

	eventInfo := jsonData.Data.EventCards

	for _, eventData := range eventInfo {
		for _, dateData := range eventData {
			for _, timeData := range dateData {
				total := 0
				for _, layoutData := range timeData {
					matchDetails := fmt.Sprintf("%s	| %s	|	AvailableSeats: %v/%v", layoutData.PriceDesc, layoutData.Price, layoutData.AvailableSeats, layoutData.TotalSeats)
					colorData := matchDetails
					total = total + layoutData.TotalSeats
					if layoutData.AvailableSeats > 0 {
						colorData = green + matchDetails + reset
					} else if layoutData.TotalSeats > 0 {
						colorData = red + matchDetails + reset
					} else {
						// continue
					}
					fmt.Println(colorData)
				}
				fmt.Printf("Total Seats: %v\n", total)
			}
		}
	}
}

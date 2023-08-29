package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	apiURL := "https://in-cdn.bmscdn.com/le-synopsis/leapi-eventinfo-"

	matchNameMap := map[string]string{
		"ET00367559": "Ind vs Pak",
		"ET00367578": "Ind vs Eng",
		"ET00367551": "Ind vs Aus",
		"ET00367570": "Ind vs Nz",
		"ET00367587": "Ind vs Sa",
		"ET00367583": "Ind vs Sl",
		"ET00367556": "Ind vs Afg",
		"ET00367564": "Ind vs Ban",
		"ET00367595": "Ind vs Ned",
	}

	matchIds := []string{}

	reset := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"

	options := ""
	i := 1
	for id, name := range matchNameMap {
		options += fmt.Sprintf("%v - %v\n", i, name)
		i++
		matchIds = append(matchIds, id)
	}

	for {
		fmt.Printf("\n%v\n", options)
		fmt.Print("Select match (1-9): ")
		data := ""
		fmt.Scanln(&data)
		fmt.Printf("\nSelected match (1-9): %v\n", data)
		index, err := strconv.Atoi(data)
		if err != nil || index > len(matchIds) || index < 1 {
			fmt.Println("choose a valid number")
			continue
		}

		apiResponse, err := http.Get(apiURL + matchIds[index-1] + ".json")
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
					for _, layoutData := range timeData {
						matchDetails := fmt.Sprintf("%s	| %s	|	AvailableSeats: %v/%v", layoutData.PriceDesc, layoutData.Price, layoutData.AvailableSeats, layoutData.TotalSeats)
						colorData := matchDetails
						if layoutData.AvailableSeats > 0 {
							colorData = green + matchDetails + reset
						} else if layoutData.TotalSeats > 0 {
							colorData = red + matchDetails + reset
						}
						fmt.Println(colorData)
					}
				}
			}
		}
	}
}

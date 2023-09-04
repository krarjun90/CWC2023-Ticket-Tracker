# World Cup 2023 Ticket Tracker

## How to configure new matches?

Append the map with match code and build a new executable.

https://in.bookmyshow.com/sports/england-vs-new-zealand-icc-mens-cwc-2023/ET00367548?groupEventCode=ET00367203

"ET00367548" is the match code from the book my show URL.

```
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
```

## How to build new executable after making changes?

- go build -o CWC2023-Ticket-Tracker
- GOOS=windows GOARCH=amd64 go build -o CWC2023-Ticket-Tracker.exe

## Where is the data?

https://in-cdn.bmscdn.com/le-synopsis/leapi-eventinfo-ET00367595.json

Replace the match code in the above url

## How to get multiple match data at once ?

```
func main() {
	// printAllMatchData()
	printPerMatchData()
}
```

Configure the matches you need in matchNameMap and uncomment "printAllMatchData" in main.

package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Event struct that contains the total event of the plateform
type Event struct {
	Status       string       `json:"status"`
	EndEvent     string       `json:"endAt"`
	Path         string       `json:"path"`
	Registration Registration `json:"registration"`
}

// Event struct that contains the total registration date of the event
type Registration struct {
	StartRegistration string `json:"startAt"`
	EndRegistration   string `json:"endAt"`
	StartEvent        string `json:"eventStartAt"`
}

// EventDate struct that contains the total of date event of the plateform
type EvenDate struct {
	Debut_reg   string
	Fin_reg     string
	Debut_event string
	Fin_event   string
	Path        string
}

var (
	events []Event

	listevents []string

	new_events []EvenDate
	new_event  *EvenDate
)

// display the date of event of ytrack
func Progress(c *gin.Context) {
	c.HTML(http.StatusOK, "progress.html", gin.H{"new_events": new_events, "title": "Evènements"})
}

// create list, split the path of event, format the date of event
func ListEvent() {
	new_event = new(EvenDate)
	last_str := ""

	for _, i := range events {
		listevents = append(listevents, i.Registration.StartRegistration, i.Registration.EndRegistration, i.Registration.StartEvent, i.EndEvent, i.Path)
	}

	for i := range listevents {

		if i%5 != 4 {
			dateString := listevents[i]
			date, error := time.Parse("2006-01-02T15:04:05Z07:00", dateString)

			dateformat := date.Format("2006-01-02 15:04:05")

			if error != nil {
				fmt.Println(error)
			}

			switch i % 5 {
			case 0:
				new_event.Debut_reg = dateformat
			case 1:
				new_event.Fin_reg = dateformat
			case 2:
				new_event.Debut_event = dateformat
			case 3:
				new_event.Fin_event = dateformat
			}

		} else {
			s := strings.Split(events[i/5].Path, "/")
			if 3 >= len(s) {
				last_str = s[2]
			} else {
				last_str = s[2] + "/" + s[3]
			}
			new_event.Path = last_str
			new_events = append(new_events, *new_event)
			sort.SliceStable(new_events, func(i, j int) bool {
				return new_events[i].Debut_event > new_events[j].Debut_event
			})
			new_event = new(EvenDate)
		}
	}
}

// read the JSON file
func ReadJsonEvent() {
	// Open our jsonFile
	jsonFile, err := os.Open("assets/json/events.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened events.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &events)
}

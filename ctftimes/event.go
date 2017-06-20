package ctftimes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tukejonny/g0tiu5a-bot/common"
)

// event
type Event struct {
	Organizers    []Organizer `json:"organizers"`
	OnSite        bool        `json:"onsite"`
	Finish        time.Time   `json:"finish"`
	Description   string      `json:"description"`
	Weight        float64     `json:"weight"`
	Title         string      `json:"title"`
	Url           string      `json:"url"`
	IsVotableNow  bool        `json:"is_votable_now"`
	Restrictions  string      `json:"restrictions"`
	Format        string      `json:"format"`
	Start         time.Time   `json:"start"`
	Participants  int         `json:"participants"`
	CtftimeUrl    string      `json:"ctftime_url"`
	Location      string      `json:"location"`
	LiveFeed      string      `json:"live_feed"`
	PublicVotable bool        `json:"public_votable"`
	Duration      Duration    `json:"duration"`
	Logo          string      `json:"logo"`
	FormatId      int         `json:"format_id"`
	Id            int         `json:"id"`
	CtfId         int         `json:"ctf_id"`
}

//event.organizers
type Organizer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//event.duration
type Duration struct {
	Hours int `json:"hours"`
	Days  int `json:"days"`
}

const (
	URL_PREFIX = "https://ctftime.org/api/v1"
	LIMIT      = "3"
)

func BuildUrl() string {
	now := time.Now().Unix()
	url := URL_PREFIX + "/events/?limit=" + LIMIT + "&start=" + fmt.Sprintf("%v", now)
	return url
}

func GetAPIData() []Event {
	var events []Event
	url := BuildUrl()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	common.Decode(response, &events)

	return events
}

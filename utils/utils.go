package utils

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/api/calendar/v3"
	"log"
	"time"
)

func PrintCalendarListEntry(calendar *calendar.CalendarListEntry) {
	fmt.Printf("Calendar{Id=%v, Summary=%v, BackgroundColor=%v, ColorId=%v}\n",
		calendar.Id, calendar.Summary, calendar.BackgroundColor, calendar.ColorId)
}

func PrintEvent(event *calendar.Event) {
	getDate := func(eventDataTime *calendar.EventDateTime) string {
		res := eventDataTime.DateTime
		if res == "" {
			res = eventDataTime.Date
		}
		return res
	}

	fmt.Printf("Event{Summary=%v, Start=%v, End=%v, ColorId=%v, Id=%v}\n",
		event.Summary,
		getDate(event.Start),
		getDate(event.End),
		event.ColorId,
		event.Id,
	)
}

func Time(eventDateTime *calendar.EventDateTime) time.Time {
	timeStr := eventDateTime.DateTime

	if timeStr == "" {
		timeStr = eventDateTime.Date
		if timeStr != "" {
			timeStr += "T00:00:00Z"
		} else {
			spew.Dump(eventDateTime)
		}
	}

	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

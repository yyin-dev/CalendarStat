package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/api/calendar/v3"
)

func PrintCalendarListEntry(calendar *calendar.CalendarListEntry) {
	fmt.Printf("Calendar{Summary=%v, Id=%v, BackgroundColor=%v, ColorId=%v}\n",
		calendar.Summary, calendar.Id, calendar.BackgroundColor, calendar.ColorId)
}

func PrintEvent(event *calendar.Event) {
	getDate := func(eventDataTime *calendar.EventDateTime) string {
		res := eventDataTime.DateTime
		if res == "" {
			res = eventDataTime.Date
		}
		return res
	}

	fmt.Printf("Event{Summary=%v, Start=%v, End=%v, ColorId=%v, Id=%v, Source=%v}\n",
		event.Summary,
		getDate(event.Start),
		getDate(event.End),
		event.ColorId,
		event.Id,
		event.Source,
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

func DurationRatio(d1, d2 time.Duration) float32 {
	nano1 := float32(d1 / time.Nanosecond)
	nano2 := float32(d2 / time.Nanosecond)
	return nano1 / nano2
}

func DurationDivide(d time.Duration, by float32) time.Duration {
	nano := float32(d / time.Nanosecond)
	nano /= by
	return time.Duration(nano) * time.Nanosecond
}

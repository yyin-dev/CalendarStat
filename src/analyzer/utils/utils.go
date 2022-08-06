package utils

import (
	"fmt"
	"google.golang.org/api/calendar/v3"
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

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

package analyze

import (
	"fmt"
	"google.golang.org/api/calendar/v3"
	"log"
	"time"
)

type ColorId = string

type Activity struct {
	colorId     string
	description string
	events      []*calendar.Event
	duration    time.Duration
}

func buildActivities(events []*calendar.Event) []Activity {
	// Job search : "1"
	// Life: "4"
	// Courses: ""/"2"
	color2Activity := make(map[ColorId]Activity)
	color2Activity["1"] = Activity{colorId: "1", description: "Job search"}
	color2Activity["2"] = Activity{colorId: "", description: "Courses"}
	color2Activity["4"] = Activity{colorId: "4", description: "Others"}

	for _, event := range events {
		if event.ColorId == "" {
			event.ColorId = "2"
		}
		activity := color2Activity[event.ColorId]
		activity.events = append(activity.events, event)
		{
			startTimeStr, endTimeStr := event.Start.DateTime, event.End.DateTime
			startTime, err := time.Parse(time.RFC3339, startTimeStr)
			if err != nil {
				log.Fatal(err)
			}
			endTime, err := time.Parse(time.RFC3339, endTimeStr)
			if err != nil {
				log.Fatal(err)
			}
			activity.duration += endTime.Sub(startTime)
		}
		color2Activity[event.ColorId] = activity
	}

	activities := make([]Activity, 0)
	for _, activity := range color2Activity {
		activities = append(activities, activity)
	}

	return activities
}

func Analyze(events []*calendar.Event) {
	activities := buildActivities(events)
	for _, activity := range activities {
		fmt.Printf("Activity: %v\n", activity.description)
		fmt.Printf("  Events: %v.\n", len(activity.events))
		fmt.Printf("  Total duration: %v.\n", activity.duration)
	}
}

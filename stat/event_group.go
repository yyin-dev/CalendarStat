package stat

import (
	"fmt"
	"github.com/yinfredyue/CalendarStat/utils"
	"google.golang.org/api/calendar/v3"
	"time"
)

type GroupEventBy int64

const (
	ByCalendar GroupEventBy = iota
	ByColorId
)

type ColorId = string

type EventGroup struct {
	calendar *calendar.CalendarListEntry
	events   []*calendar.Event
	start    time.Time
	end      time.Time
}

type EvenGroupStat struct {
	duration       time.Duration
	weeklyDuration time.Duration
}

func (eg *EventGroup) String() string {
	return fmt.Sprintf("EventGroup{calendar=%v, len(events)=%v, start=%v, end=%v}",
		eg.calendar.Summary, len(eg.events), eg.start, eg.end)
}

func (egs *EvenGroupStat) String() string {
	return fmt.Sprintf("EventGroupStat{duration=%v, weeklyDuration=%v}",
		egs.duration, egs.weeklyDuration)
}

//func buildActivity(events []*calendar.Event) []EventGroup {
//	// Job search : "1"
//	// Life: "4"
//	// Courses: ""/"2"
//	color2Activity := make(map[ColorId]Activity)
//	color2Activity["1"] = Activity{colorId: "1", description: "Job search"}
//	color2Activity["2"] = Activity{colorId: "2", description: "Courses"}
//	color2Activity["4"] = Activity{colorId: "4", description: "Others"}
//
//	for _, event := range events {
//		if event.ColorId == "" {
//			event.ColorId = "2"
//		}
//
//		activity := color2Activity[event.ColorId]
//		activity.events = append(activity.events, event)
//		{
//			startTimeStr, endTimeStr := event.Start.DateTime, event.End.DateTime
//			startTime, err := time.Parse(time.RFC3339, startTimeStr)
//			if err != nil {
//				log.Fatal(err)
//			}
//			endTime, err := time.Parse(time.RFC3339, endTimeStr)
//			if err != nil {
//				log.Fatal(err)
//			}
//			activity.duration += endTime.Sub(startTime)
//		}
//		color2Activity[event.ColorId] = activity
//	}
//
//	activities := make([]Activity, 0)
//	for _, activity := range color2Activity {
//		activities = append(activities, activity)
//	}
//
//	return activities
//}

//activities := buildActivity(events)
//
//timeRange := end.Sub(start)
//numWeeks := timeRange / (time.Hour * 7 * 24)
//
//for _, activity := range activities {
//	fmt.Printf("Activity: %v\n", activity.description)
//	fmt.Printf("  ColorID: %v\n", activity.colorId)
//	fmt.Printf("  Events: %v\n", len(activity.events))
//	fmt.Printf("  Total duration: %v\n", activity.duration)
//	if numWeeks != 1 {
//		fmt.Printf("  Duration per week: %v\n", activity.duration/numWeeks)
//	}
//}

func BuildEventGroup(events []*calendar.Event, calendar *calendar.CalendarListEntry, groupEventBy GroupEventBy,
	start time.Time, end time.Time) *EventGroup {

	return &EventGroup{
		calendar: calendar,
		events:   events,
		start:    start,
		end:      end,
	}
}

func (eg *EventGroup) Stat() *EvenGroupStat {
	var duration time.Duration
	for _, event := range eg.events {
		startTime := utils.Time(event.Start)
		endTime := utils.Time(event.End)
		duration += endTime.Sub(startTime)
	}

	numWeeks := eg.end.Sub(eg.start) / (time.Hour * 7 * 24)
	weeklyDuration := duration / numWeeks

	return &EvenGroupStat{duration: duration, weeklyDuration: weeklyDuration}
}

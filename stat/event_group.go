package stat

import (
	"fmt"
	"strings"
	"time"

	"github.com/yinfredyue/CalendarStat/utils"
	"google.golang.org/api/calendar/v3"
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
	byColor  map[ColorId]*EventGroup
}

type EventGroupStat struct {
	duration       time.Duration
	weeklyDuration time.Duration
	byColor        map[ColorId]*EventGroupStat
}

func (eg *EventGroup) String() string {
	// Manually format string to print nested byColor. A bit ugly.
	byColorStr := func(eg *EventGroup) string {
		sList := make([]string, 0)
		for colorId, subEg := range eg.byColor {
			s := fmt.Sprintf("colorId=%v, group=%v", colorId, subEg)
			sList = append(sList, s)
		}

		sep := "\n  "
		res := strings.Join(sList, sep)
		if len(res) > 0 {
			res = sep + res
		}
		return res
	}(eg)

	return fmt.Sprintf("EventGroup{calendar=%v, len(events)=%v, start=%v, end=%v, byColor={%v}}",
		eg.calendar.Summary, len(eg.events), eg.start, eg.end, byColorStr)
}

func (egs *EventGroupStat) String() string {
	// Manually format string to print nested byColor. A bit ugly.
	byColorStr := func(egs *EventGroupStat) string {
		sList := make([]string, 0)
		for colorId, subEg := range egs.byColor {
			s := fmt.Sprintf("colorId=%v, stat=%v", colorId, subEg)
			sList = append(sList, s)
		}

		sep := "\n    "
		res := strings.Join(sList, sep)
		if len(res) > 0 {
			res = sep + res
		}
		return res
	}(egs)

	return fmt.Sprintf("EventGroupStat{duration=%v, weeklyDuration=%v, byColor={%v}}",
		egs.duration, egs.weeklyDuration, byColorStr)
}

func BuildEventGroup(events []*calendar.Event, cal *calendar.CalendarListEntry, groupEventBy GroupEventBy,
	start time.Time, end time.Time) *EventGroup {

	var byColor map[ColorId]*EventGroup
	if groupEventBy == ByColorId {
		byColor = make(map[ColorId]*EventGroup)
		for _, event := range events {
			// Create if not exist
			if _, ok := byColor[event.ColorId]; !ok {
				byColor[event.ColorId] = &EventGroup{
					calendar: cal,
					events:   []*calendar.Event{},
					start:    start,
					end:      end,
					byColor:  nil,
				}
			}

			// Update group.events
			byColor[event.ColorId].events = append(byColor[event.ColorId].events, event)
		}
	}

	return &EventGroup{
		calendar: cal,
		events:   events,
		start:    start,
		end:      end,
		byColor:  byColor,
	}
}

func (eg *EventGroup) Stat() *EventGroupStat {
	var duration time.Duration
	for _, event := range eg.events {
		startTime := utils.Time(event.Start)
		endTime := utils.Time(event.End)
		duration += endTime.Sub(startTime)
	}

	numWeeks := utils.DurationRatio(eg.end.Sub(eg.start), time.Hour*24*7)
	weeklyDuration := utils.DurationDivide(duration, numWeeks)

	var byColor map[ColorId]*EventGroupStat
	if eg.byColor != nil {
		byColor = make(map[ColorId]*EventGroupStat)
		for colorId, subEg := range eg.byColor {
			byColor[colorId] = subEg.Stat()
		}
	}

	return &EventGroupStat{
		duration:       duration,
		weeklyDuration: weeklyDuration,
		byColor:        byColor,
	}
}

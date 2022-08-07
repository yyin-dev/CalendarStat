# CalendarStat

CalendarStat generates report and statistics from your Google Calendar.

Usage
```
$ cmd
NAME:
   CalendarStat - A tool for analyzing your Google Calendar

USAGE:
   CalendarStat [global options] command [command options] [arguments...]

COMMANDS:
   colors      Show all colorIds
   calendars   Show all calendars
   events      Show all events
   event-stat  Get statistics about events
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

Display calendars
```
$ cmd calendars
Calendar{Id=yyin5@andrew.cmu.edu, Summary=yyin5@andrew.cmu.edu, BackgroundColor=#d06b64, ColorId=2}
Calendar{Id=en.usa#holiday@group.v.calendar.google.com, Summary=Holidays in United States, BackgroundColor=#9fc6e7, ColorId=15}
Calendar{Id=legendmakeryy@gmail.com, Summary=legendmakeryy@gmail.com, BackgroundColor=#9fe1e7, ColorId=14}
```
Display events
``` 
$ cmd events -start-date 2021-09-19T00:00:00Z -end-date 2021-10-23T00:00:00Z -calendar-id yyin5@andrew.cmu.edu
Event{Summary=LeetCode Weekly Contest, Start=2021-09-18T22:30:00-04:00, End=2021-09-19T00:00:00-04:00, ColorId=, Id=30vo8ircrvmatr0k9tbu10o881_20210919T023000Z, Source=<nil>}
Event{Summary=notes, Start=2021-09-19T08:30:00-04:00, End=2021-09-19T10:30:00-04:00, ColorId=2, Id=6kbiv5akr39dsd8tc5cqibgj36, Source=<nil>}
Event{Summary=Q3, Start=2021-09-19T10:30:00-04:00, End=2021-09-19T13:00:00-04:00, ColorId=2, Id=5lgvcf310lin56983sit0evncq, Source=<nil>}
Event{Summary=HW2 Disussion, Start=2021-09-19T18:00:00-04:00, End=2021-09-19T20:00:00-04:00, ColorId=2, Id=7jpf74gpmtgg7b6bl5i0v4eofr, Source=<nil>}
Event{Summary=Leetcode questions, Start=2021-09-20T08:40:00-04:00, End=2021-09-20T13:40:00-04:00, ColorId=1, Id=2ubeoe2o4jvi098lfkqb8v7tmr, Source=<nil>}
Event{Summary=Online Applications, Start=2021-09-20T14:45:00-04:00, End=2021-09-20T15:35:00-04:00, ColorId=1, Id=790h4d1t6oujm0vt6ghfnh01d8, Source=<nil>}
Event{Summary=Storage Systems, Start=2021-09-20T16:40:00-04:00, End=2021-09-20T18:00:00-04:00, ColorId=, Id=5k9t0jjpeh9onib6cehbqsigpu_20210920T204000Z, Source=<nil>}
```
Display statistics about events (calendar-wise)
``` 
$ cmd event-stat -start-date 2021-09-19T00:00:00Z -end-date 2021-10-23T00:00:00Z -group-event-by calendar
EventGroup{calendar=Job Search, len(events)=0, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  EventGroupStat{duration=0s, weeklyDuration=0s, byColor={}}
EventGroup{calendar=Study, len(events)=0, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  EventGroupStat{duration=0s, weeklyDuration=0s, byColor={}}
EventGroup{calendar=Course Schedule, len(events)=0, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  EventGroupStat{duration=0s, weeklyDuration=0s, byColor={}}
EventGroup{calendar=Others, len(events)=250, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  EventGroupStat{duration=314h55m0s, weeklyDuration=78h43m45s, byColor={}}
```
Display statistics about events (calendar-wise AND color-wise)
```
$ cmd event-stat -start-date 2021-09-19T00:00:00Z -end-date 2021-10-23T00:00:00Z -group-event-by colorId
EventGroup{calendar=Job Search, len(events)=0, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  EventGroupStat{duration=0s, weeklyDuration=0s, byColor={}}
EventGroup{calendar=CMU Study, len(events)=0, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  EventGroupStat{duration=0s, weeklyDuration=0s, byColor={}}
EventGroup{calendar=Course Schedule, len(events)=0, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  EventGroupStat{duration=0s, weeklyDuration=0s, byColor={}}
EventGroup{calendar=Others, len(events)=250, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={
  colorId=5, group=EventGroup{calendar=Others, len(events)=1, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  colorId=2, group=EventGroup{calendar=Others, len(events)=77, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  colorId=, group=EventGroup{calendar=Others, len(events)=39, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  colorId=1, group=EventGroup{calendar=Others, len(events)=87, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}
  colorId=4, group=EventGroup{calendar=Others, len(events)=46, start=2021-09-19 00:00:00 +0000 UTC, end=2021-10-23 00:00:00 +0000 UTC, byColor={}}}}
  EventGroupStat{duration=314h55m0s, weeklyDuration=78h43m45s, byColor={
    colorId=2, stat=EventGroupStat{duration=118h25m0s, weeklyDuration=29h36m15s, byColor={}}
    colorId=, stat=EventGroupStat{duration=45h40m0s, weeklyDuration=11h25m0s, byColor={}}
    colorId=1, stat=EventGroupStat{duration=117h10m0s, weeklyDuration=29h17m30s, byColor={}}
    colorId=4, stat=EventGroupStat{duration=33h25m0s, weeklyDuration=8h21m15s, byColor={}}
    colorId=5, stat=EventGroupStat{duration=15m0s, weeklyDuration=3m45s, byColor={}}}}
```
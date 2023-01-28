package main

import (
	"strconv"
	"time"
)

const (
	TodayMsg    = "今日"
	DistanceMsg = "距离"
	UntilMsg    = "还有"
	DayDiffMsg  = "天"
	MonthMsg    = "月"
	DayMsg      = "日"
)

var weekDaysMap = map[time.Weekday]string{
	time.Sunday:    "星期日",
	time.Monday:    "星期一",
	time.Tuesday:   "星期二",
	time.Wednesday: "星期三",
	time.Thursday:  "星期四",
	time.Friday:    "星期五",
	time.Saturday:  "星期六",
}

type Text struct {
	FirstUpcomingPrefix string
	HighLightArea       string
	FirstUpcomingDiff   string
	Date                string
	SecondUpcoming      string
	ThirdUpcoming       string
}

func (t *Text) SetDate(date time.Time) {
	month := strconv.Itoa(int(date.Month())) + MonthMsg
	day := strconv.Itoa(date.Day()) + DayMsg
	t.Date = month + day + " " + weekDaysMap[date.Weekday()]
}

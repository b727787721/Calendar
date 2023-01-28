package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var MinCountYear = 2
var MinCountMemorialDays = 3

type Configs struct {
	MemorialDaysConfig []MemorialDaysConfig `json:"MemorialDaysConfig"`
}

type MemorialDaysConfig struct {
	Year         int                  // format 2022
	MemorialDays []MemorialDaysString `json:"MemorialDays"`
}

type MemorialDaysString struct {
	Date     string // "1.30"
	Memorial string // "Chinese Lunar New Year"
}

type MemorialDays struct {
	Date     time.Time
	Memorial string
}

type Calendar struct {
}

func (c *Calendar) LoadConfigFile(path string) (*Configs, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Configs
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Calendar) GetFirstAndLastDate(year int) (time.Time, time.Time) {
	firstDate := time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.Now().Location())
	lastDate := firstDate.AddDate(1, 0, -1)
	return firstDate, lastDate
}

func (c *Calendar) TransformMemorialDays(configs []MemorialDaysConfig) ([]MemorialDays, error) {
	var memorialDays []MemorialDays
	formatErrorMessage := fmt.Sprintf("The Memorial days format error, please input Date like format \"1.30\"")

	for _, config := range configs {
		for _, memorialDaysString := range config.MemorialDays {
			splitDates := strings.Split(memorialDaysString.Date, ".")
			month, err := strconv.Atoi(splitDates[0])
			if err != nil {
				return nil, errors.New(formatErrorMessage)
			}
			day, err := strconv.Atoi(splitDates[1])
			if err != nil {
				return nil, errors.New(formatErrorMessage)
			}

			memorialDay := MemorialDays{Date: time.Date(config.Year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location()), Memorial: memorialDaysString.Memorial}
			memorialDays = append(memorialDays, memorialDay)
		}
	}

	return memorialDays, nil
}

func (c *Calendar) Process(configs []MemorialDaysConfig) {
	memorialDays, err := c.TransformMemorialDays(configs)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}

	memorialDaysPos := 0
	texts := []Text{}
	firstUpcoming, secondUpcoming, thirdUpcoming := memorialDays[memorialDaysPos], memorialDays[memorialDaysPos+1], memorialDays[memorialDaysPos+2]
	firstDate, lastDate := c.GetFirstAndLastDate(configs[0].Year)
	for date := firstDate; !date.After(lastDate); date = date.AddDate(0, 0, 1) {
		text := Text{}
		text.SetDate(date)
		secondUpcomingMsg := DistanceMsg + secondUpcoming.Memorial + UntilMsg + strconv.Itoa(int(secondUpcoming.Date.Sub(date).Hours())/24) + DayDiffMsg
		thirdUpcomingMsg := DistanceMsg + thirdUpcoming.Memorial + UntilMsg + strconv.Itoa(int(thirdUpcoming.Date.Sub(date).Hours())/24) + DayDiffMsg
		text.SecondUpcoming = secondUpcomingMsg
		text.ThirdUpcoming = thirdUpcomingMsg

		if date.Equal(firstUpcoming.Date) {
			text.FirstUpcomingPrefix = TodayMsg
			text.HighLightArea = firstUpcoming.Memorial

			memorialDaysPos++
			firstUpcoming = memorialDays[memorialDaysPos]
			secondUpcoming = memorialDays[memorialDaysPos+1]
			thirdUpcoming = memorialDays[memorialDaysPos+2]
		} else {
			text.FirstUpcomingPrefix = DistanceMsg + firstUpcoming.Memorial + UntilMsg
			text.HighLightArea = strconv.Itoa(int(firstUpcoming.Date.Sub(date).Hours()) / 24)
			text.FirstUpcomingDiff = DayDiffMsg
		}

		texts = append(texts, text)
	}

	pdf := Pdf{}
	pdf.GeneratedCalendarPdf(texts)
}

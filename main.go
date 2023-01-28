package main

import (
	"flag"
	"fmt"
)

func main() {
	var path string
	flag.StringVar(&path, "file", "configs.json", "The file path of the calendar")
	flag.Parse()

	calendar := &Calendar{}
	configs, err := calendar.LoadConfigFile(path)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	memorialDaysConfig := configs.MemorialDaysConfig
	if len(memorialDaysConfig) < MinCountYear || len(memorialDaysConfig[0].MemorialDays) < MinCountMemorialDays || len(memorialDaysConfig[1].MemorialDays) < MinCountMemorialDays {
		panicMessage := fmt.Sprintf("You need to add at least current year and next year Memorial days, and at least %d Memorial days for each year", MinCountMemorialDays)
		fmt.Println("Got an error: ", panicMessage)
		return
	}

	calendar.Process(memorialDaysConfig)
}

package icalendar

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func LoadCalendar(icsCalendarUrl string) (iCalendar ICalendar, err error) {
	log.Println("Downloading calendar from:", icsCalendarUrl)

	response, err := http.Get(icsCalendarUrl)
	if err != nil {
		return
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	result := string(responseBody)
	iCalendar = parseCalendar(result)
	return
}

func parseCalendar(icsContent string) (iCalendar ICalendar) {
	log.Println("====== Start of parseCalendar ======")
	iCalendar = ICalendar{}
	icsContentLines := strings.Split(icsContent, "\n")

	dateLayout := "20060102 -0700"

	vEventList := []VEvent{}
	var vEvent VEvent

	for _, line := range icsContentLines {
		trimmedLine := strings.Replace(line, "\r", "", -1)
		if strings.HasPrefix(trimmedLine, VERSION) {
			log.Println(trimmedLine)
			iCalendar.VERSION = trimmedLine[len(VERSION):]
		} else if strings.HasPrefix(trimmedLine, PRODID) {
			log.Println(trimmedLine)
			iCalendar.PRODID = trimmedLine[len(PRODID):]
		} else if strings.HasPrefix(trimmedLine, CALSCALE) {
			log.Println(trimmedLine)
			iCalendar.CALSCALE = trimmedLine[len(CALSCALE):]
		} else if strings.HasPrefix(trimmedLine, VEVENT_BEGIN) {
			log.Println("------ Start of Event ------")
			log.Println(trimmedLine)
			vEvent = VEvent{}
		} else if strings.HasPrefix(trimmedLine, VEVENT_DTSTAMP) {
			log.Println(trimmedLine)
			vEvent.DTSTAMP = trimmedLine[len(VEVENT_DTSTAMP):]
		} else if strings.HasPrefix(trimmedLine, VEVENT_UID) {
			log.Println(trimmedLine)
			vEvent.UID = trimmedLine[len(VEVENT_UID):]
		} else if strings.HasPrefix(trimmedLine, VEVENT_DTSTART) {
			log.Println(trimmedLine)
			dateTime, err := time.Parse(dateLayout, fmt.Sprintf("%s +0800", trimmedLine[len(VEVENT_DTSTART):]))
			if err != nil {
				log.Fatal(err)
			}

			vEvent.DTSTART = dateTime
		} else if strings.HasPrefix(trimmedLine, VEVENT_DTEND) {
			log.Println(trimmedLine)
			dateTime, err := time.Parse(dateLayout, fmt.Sprintf("%s +0800", trimmedLine[len(VEVENT_DTEND):]))
			if err != nil {
				log.Fatal(err)
			}

			vEvent.DTEND = dateTime
		} else if strings.HasPrefix(trimmedLine, VEVENT_SUMMARY) {
			log.Println(trimmedLine)
			vEvent.SUMMARY = trimmedLine[len(VEVENT_SUMMARY):]
		} else if strings.HasPrefix(trimmedLine, VEVENT_END) {
			vEventList = append(vEventList, vEvent)
			log.Println("------ End of Event ------")
		}
	}

	iCalendar.VEVENT = vEventList
	log.Println("====== End of parseCalendar ======")

	return
}

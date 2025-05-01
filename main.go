package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sgpublicholiday/icalendar"
	"time"
)

const MOM_GOV_PUBLIC_HOLIDAY_ICS_URL = "https://www.mom.gov.sg/-/media/mom/documents/employment-practices/public-holidays/public-holidays-sg-%d.ics"

func GetMomPublicHolidayUrl() (momPublicHolidayIcsUrl string, currentYear int, err error) {
	sgt, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		return
	}

	currentDateTime := time.Now().In(sgt)
	currentYear = currentDateTime.Year()

	momPublicHolidayIcsUrl = fmt.Sprintf(MOM_GOV_PUBLIC_HOLIDAY_ICS_URL, currentYear)

	return
}

func main() {
	momPublicHolidayUrl, year, err := GetMomPublicHolidayUrl()
	if err != nil {
		log.Fatal(err)
	}

	ical, err := icalendar.LoadCalendar(momPublicHolidayUrl)
	if err != nil {
		log.Fatal(err)
	}

	icalJson, err := json.MarshalIndent(ical, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./data/latest.json", icalJson, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fmt.Sprintf("./data/public_holiday_%d.json", year), icalJson, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

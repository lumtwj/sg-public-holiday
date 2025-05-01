package icalendar

import "time"

const (
	VERSION        = "VERSION:"
	PRODID         = "PRODID:"
	CALSCALE       = "CALSCALE:"
	VEVENT_BEGIN   = "BEGIN:VEVENT"
	VEVENT_DTSTAMP = "DTSTAMP:"
	VEVENT_UID     = "UID:"
	VEVENT_DTSTART = "DTSTART;VALUE=DATE:"
	VEVENT_DTEND   = "DTEND;VALUE=DATE:"
	VEVENT_SUMMARY = "SUMMARY:"
	VEVENT_END     = "END:VEVENT"
)

type ICalendar struct {
	VERSION  string   `json:"version"`
	PRODID   string   `json:"prod_id"`
	CALSCALE string   `json:"cal_scale"`
	VEVENT   []VEvent `json:"vevent"`
}

type VEvent struct {
	DTSTAMP string    `json:"dt_stamp"`
	UID     string    `json:"uid"`
	DTSTART time.Time `json:"dt_start"`
	DTEND   time.Time `json:"dt_end"`
	SUMMARY string    `json:"summary"`
}

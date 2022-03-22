package icalendar

const Colon = ":"
const BEGIN = "BEGIN"
const PRODID = "PRODID"
const VERSION = "VERSION"
const DTSTART = "DTSTART"
const DTEND = "DTEND"
const LOCATION = "LOCATION"
const CATEGORIES = "CATEGORIES"
const DESCRIPTION = "DESCRIPTION"
const SUMMARY = "SUMMARY"
const PRIORITY = "PRIORITY"
const END = "END"

// BEGIN:VCALENDAR
// PRODID:-//Microsoft Corporation//Outlook MIMEDIR//EN
// VERSION:1.0
// BEGIN:VEVENT
// DTSTART:20210314T210000Z
// DTEND:20210314T230000Z
// LOCATION:行政楼221
// CATEGORIES:活动类型
// DESCRIPTION;ENCODING=QUOTED-PRINTABLE:这是测试的活动内容详情=0D=0A
// SUMMARY:测试的活动的SUMMARY
// PRIORITY:3
// END:VEVENT
// END:VCALENDAR


func SetBEGIN(value string) string {
	return BEGIN + Colon + value
}

func SetPRODID(value string) string {
	return PRODID + Colon + value
}

func SetVERSION(value string) string {
	return VERSION + Colon + value
}
func SetDTSTART(value string) string {
	return DTSTART + Colon + value
}
func SetDTEND(value string) string {
	return DTEND + Colon + value
}
func SetLOCATION(value string) string {
	return LOCATION + Colon + value
}
func SetCATEGORIES(value string) string {
	return CATEGORIES + Colon + value
}
func SetDESCRIPTION(value string) string {
	return DESCRIPTION + Colon + value
}
func SetSUMMARY(value string) string {
	return SUMMARY + Colon + value
}

func SetPRIORITY(value string) string {
	return PRIORITY + Colon + value
}
func SetEND(value string) string {
	return END + Colon + value
}

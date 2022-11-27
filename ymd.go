// package for handling YYYYMMDD string easily
package ymd

import (
	"fmt"
	"time"

	"github.com/Songmu/go-httpdate"
	"github.com/jinzhu/now"
)

const (
	//value for current date format
	YYYYMMDD = "20060102"

	//value for current datetime format
	YYYYMMDDHHMMSS = "20060102150405"

	//value for current date format
	YYYY_MM_DD = "2006-01-02"

	//value for current datetime format
	YYYY_MM_DD_HH_MM_SS = "2006-01-02-15-04-05"
)

// type for YYYY-MM-DD
type HyphenYmd string

// convert to string
func (h HyphenYmd) String() string {
	return (string)(h)
}

func (h HyphenYmd) getYear() string {
	s := (string)(h)
	return s[0:4]
}

func (h HyphenYmd) getMonth() string {
	s := (string)(h)
	return s[5:7]
}

func (h HyphenYmd) getDay() string {
	s := (string)(h)
	return s[8:10]
}

// Covert "YYYY-MM-DD" string to "YYYYMMDD" string
func (h HyphenYmd) ToYmd() Ymd {
	s := h.getYear() + h.getMonth() + h.getDay()
	return Ymd(s)
}

func ExampleToYmd() {
	fmt.Println(HyphenYmd("2022-11-23").ToYmd())
	// Output: "20221123"
}

// type for YYYYMMDD
type Ymd string

// convert to string
func (y Ymd) String() string {
	return (string)(y)
}

func (y Ymd) getYear() string {
	s := (string)(y)
	return s[0:4]
}

func (y Ymd) getMonth() string {
	s := (string)(y)
	return s[4:6]
}

func (y Ymd) getDay() string {
	s := (string)(y)
	return s[6:8]
}

// Get "YYYYMM" string from "YYYYMMDD" string
func (y Ymd) ToYm() string {
	return y.getYear() + y.getMonth()
}

// Covert "YYYYMMDD" string to "YYYY-MM-DD" string
func (y Ymd) ToHyphen() HyphenYmd {

	return HyphenYmd(y.getYear() + "-" + y.getMonth() + "-" + y.getDay())
}

func ExampleToHyphen() {
	fmt.Println(Ymd("20221123").ToHyphen())
	// Output: "2022-11-23"
}

// Add and Sub day with "YYYYMMDD" string
func (y Ymd) AddDay(c int) Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	d = d.AddDate(0, 0, c)
	return Ymd(d.Format(YYYYMMDD))
}

// Add and Sub month with "YYYYMMDD" string
func (y Ymd) AddMonth(c int) Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	expect_date := time.Date(d.Year(), d.Month(), 1, d.Hour(), d.Minute(), d.Second(), d.Nanosecond(), d.Location())
	d = d.AddDate(0, c, 0)

	expect_date = expect_date.AddDate(0, c, 0)
	// if expected month is different from actual month, use last date of last month
	if d.Month() != expect_date.Month() {
		last_date := now.With(expect_date).EndOfMonth()
		d = last_date
	}
	return Ymd(d.Format(YYYYMMDD))
}

// Add and Sub year with "YYYYMMDD" string
func (y Ymd) AddYear(c int) Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	expect_date := time.Date(d.Year(), d.Month(), 1, d.Hour(), d.Minute(), d.Second(), d.Nanosecond(), d.Location())
	d = d.AddDate(c, 0, 0)

	expect_date = expect_date.AddDate(c, 0, 0)
	// if expected month is different from actual month, use last date of last month
	if d.Month() != expect_date.Month() {
		last_date := now.With(expect_date).EndOfMonth()
		d = last_date
	}
	return Ymd(d.Format(YYYYMMDD))
}

// Get first "YYYYMMDD" date with the "YYYYMMDD"
func (y Ymd) FirstDateOfMonth() Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	d = now.With(d).BeginningOfMonth()
	return Ymd(d.Format(YYYYMMDD))
}

// Get end "YYYYMMDD" date with the "YYYYMMDD"
func (y Ymd) EndDateOfMonth() Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	d = now.With(d).EndOfMonth()
	return Ymd(d.Format(YYYYMMDD))
}

// Calculate the duration of two "YYYYMMDD" string
func (y Ymd) Between(y2 Ymd) int {
	d1, _ := httpdate.Str2Time(y.String(), nil)
	d2, _ := httpdate.Str2Time(y2.String(), nil)
	duration := d2.Sub(d1)
	hours0 := int(duration.Hours())
	days := hours0 / 24
	return days
}

// Explode "YYYYMMDD" string to "YYYY","MM","DD"
func (y Ymd) Explode() (string, string, string) {
	return y.getYear(), y.getMonth(), y.getDay()
}

// Get today's "YYYYMMDD" string
func TodayYmd() Ymd {
	now := time.Now().UTC()
	return Ymd(now.Format(YYYYMMDD))
}

// type for YYYYMMDDHHMMSS
type Ymdhms string

// convert to string
func (h Ymdhms) String() string {
	return (string)(h)
}

// Get today's "YYYYMMDDHHMMSS" string
func TodayYmdhms() Ymdhms {
	now := time.Now().UTC()
	return Ymdhms(now.Format(YYYYMMDDHHMMSS))
}

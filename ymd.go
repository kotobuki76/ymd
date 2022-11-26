package ymd

import (
	"time"

	"github.com/Songmu/go-httpdate"
	"github.com/jinzhu/now"
)

const (
	YYYYMMDD            = "20060102"
	YYYYMMDDHHMMSS      = "20060102150405"
	YYYY_MM_DD          = "2006-01-02"
	YYYY_MM_DD_HH_MM_SS = "2006-01-02-15-04-05"
)

type HyphenYmd string

func (h HyphenYmd) String() string {
	return (string)(h)
}

func (h HyphenYmd) GetYear() string {
	s := (string)(h)
	return s[0:4]
}

func (h HyphenYmd) GetMonth() string {
	s := (string)(h)
	return s[5:7]
}

func (h HyphenYmd) GetDay() string {
	s := (string)(h)
	return s[8:10]
}

func (h HyphenYmd) ToYmd() Ymd {
	s := h.GetYear() + h.GetMonth() + h.GetDay()
	return Ymd(s)

}

type Ymd string

func (y Ymd) String() string {
	return (string)(y)
}

func (y Ymd) GetYear() string {
	s := (string)(y)
	return s[0:4]
}

func (y Ymd) GetMonth() string {
	s := (string)(y)
	return s[4:6]
}

func (y Ymd) GetDay() string {
	s := (string)(y)
	return s[6:8]
}

func (y Ymd) ToHyphen() HyphenYmd {

	return HyphenYmd(y.GetYear() + "-" + y.GetMonth() + "-" + y.GetDay())
}

func (y Ymd) ToYm() string {
	return y.GetYear() + y.GetMonth()
}

func (y Ymd) AddDay(c int) Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	d = d.AddDate(0, 0, c)
	return Ymd(d.Format(YYYYMMDD))
}

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

func (y Ymd) FirstDateOfMonth() Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	d = now.With(d).BeginningOfMonth()
	return Ymd(d.Format(YYYYMMDD))
}

func (y Ymd) EndDateOfMonth() Ymd {
	d, _ := httpdate.Str2Time(y.String(), nil)
	d = now.With(d).EndOfMonth()
	return Ymd(d.Format(YYYYMMDD))
}

func (y Ymd) Between(y2 Ymd) int {
	d1, _ := httpdate.Str2Time(y.String(), nil)
	d2, _ := httpdate.Str2Time(y2.String(), nil)
	duration := d2.Sub(d1)
	hours0 := int(duration.Hours())
	days := hours0 / 24
	return days
}

func (y Ymd) Explode() (string, string, string) {
	return y.GetYear(), y.GetMonth(), y.GetDay()
}

func TodayYmd() Ymd {
	now := time.Now().UTC()
	return Ymd(now.Format(YYYYMMDD))
}

type Ymdhms string

func TodayYmdhms() Ymdhms {
	now := time.Now().UTC()
	return Ymdhms(now.Format(YYYYMMDDHHMMSS))
}

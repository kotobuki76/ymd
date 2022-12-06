package ymd

import "fmt"

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

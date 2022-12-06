package ymd

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

type Yyyymmdd struct {
	value string
}

func NewYyyymmdd(s string) (*Yyyymmdd, error) {

	err := ValidateYmd(s)

	if err != nil {
		return nil, err
	} else {
		return &Yyyymmdd{value: s}, nil
	}

}

func (item *Yyyymmdd) String() string {
	return item.value
}

func (item *Yyyymmdd) Set(s string) error {
	item.value = s
	if !(item.validate(s)) {
		return errors.New("invalid input for Yyyymmdd:" + s)
	} else {
		return nil
	}
}

func (item *Yyyymmdd) Scan(arg interface{}) error {
	if arg == nil {
		item.value = ""
	}

	if reflect.TypeOf(arg).String() == "string" {
		s := string((arg.([]byte)))
		if !(item.validate(s)) {
			return errors.New("invalid input for Yyyymmdd:" + s)
		} else {
			item.value = s
		}
	}
	return nil
}

func (item *Yyyymmdd) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if !(item.validate(s)) {
		return errors.New("invalid input for Yyyymmdd:" + s)
	} else {
		item.value = s
	}
	return nil
}

func (item *Yyyymmdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(item.value)
}

func (item *Yyyymmdd) validate(s string) bool {
	err := ValidateYmd(s)
	if err != nil {
		return false
	} else {
		return true
	}
}

func ValidateYmd(s string) error {
	if len(s) != 8 {
		return errors.New("YYYYMMDD length must be 8")
	}

	_, err := strconv.Atoi(s)
	if err != nil {
		return errors.New("YYYYMMDD must be numbers")
	}

	reg := regexp.MustCompile(`[-|/|:| |　]`)
	str := reg.ReplaceAllString(s, "")
	format := string([]rune("20060102150405")[:len(str)])
	_, errParse := time.Parse(format, str)

	if errParse != nil {
		return errors.New("YYYYMMDD must be valid comnination with year, month and day:" + errParse.Error())
	}

	return nil

}

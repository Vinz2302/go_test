package helpers

import (
	"fmt"
	"strings"
	"time"
)

type DateConverter struct {
	UnixTimestamp int64
	Time          time.Time
}

type IDateConverter interface {
	New(UnixTimestamp int)
	ConvertIntoFirstYear() error
	ConvertIntoFirstMonth() error
	ConvertIntoFirstDay() error
}

func (d *DateConverter) New(UnixTimestamp int64) {
	d.UnixTimestamp = UnixTimestamp
}

func (d *DateConverter) ConvertIntoFirstYear() error {
	d.Time = time.Unix(d.UnixTimestamp, 0)
	year, _, _ := d.Time.Date()
	parsingString := fmt.Sprintf("%v", year)
	result, err := time.Parse("2006", parsingString)
	if err != nil {
		return err
	} else {
		d.Time = result
		return nil
	}
}

func (d *DateConverter) ConvertIntoFirstDay() error {
	d.Time = time.Unix(d.UnixTimestamp, 0)
	year, month, day := d.Time.Date()
	parsingString := fmt.Sprintf("%v-%v-%v", year, month, day)
	fmt.Printf("parsingString : %v \n", parsingString)
	result, err := time.Parse("2006-January-2", parsingString)
	if err != nil {
		return err
	} else {
		d.Time = result
		return nil
	}
}

func ValidateDate(date string) (*time.Time, error) {
	dateFormat := strings.Replace(date, "/", "-", -1)

	result, errParseTime := time.Parse("2006-01-02", dateFormat)
	if errParseTime != nil {
		return nil, errParseTime
	}
	return &result, nil
}

func ConvertUnixToDate(date int) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	result := time.Unix(int64(date), 0).In(loc)

	return result
}

func ConvertUnix(date string) (*time.Time, error) {

	resultDate, err := ValidateDate(date)
	if err != nil {
		return nil, err
	}

	myTime := time.Unix(resultDate.Unix(), 0)

	return &myTime, nil
}

func ConvertDateToUnix(date time.Time) int64 {
	result := date.Unix()

	return result
}

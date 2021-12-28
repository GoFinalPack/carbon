package carbon

import (
	"fmt"
	"time"
)

type Carbon struct {
	Date time.Time
}
type FormatStr interface{}

// 获取当前的时间
func Now() string {
	dateTime := time.Now()
	return dateTime.Format("2006-01-02 15:04:05")
}

func Today() string {
	dateTime := time.Now()
	return dateTime.Format("2006-01-02") + " 00:00:00"
}

func Create(year int, month int, day int, hour int, minute int, second int) *Carbon {
	return &Carbon{time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)}
}

func (c *Carbon) Ymd(format ...FormatStr) string {
	if len(format) == 0 {
		return c.Date.Format("2006-01-02")
	} else {
		return c.Date.Format("2006" + format[0].(string) + "01" + format[0].(string) + "02")
	}

}

func Hello() {
	fmt.Println("Hello, World!")
}

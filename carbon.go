package carbon

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Carbon struct {
	Date time.Time
}
type FormatStr interface{}

// 获取当前的时间
func (c *Carbon) Now(options ...interface{}) string {
	if len(options) > 0 {
		loc, _ := time.LoadLocation(options[0].(string))
		if c.Date.IsZero() {
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
			c.Date = t.In(loc)
		} else {
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", c.Date.Format("2006-01-02 15:04:05"), time.Local)
			c.Date = t.In(loc)
		}
	} else {
		c.Date = time.Now()
	}

	return c.Date.Format("2006-01-02 15:04:05")
}

// 获取当前的年月日
func (c *Carbon) Today() string {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	return c.Ymd()
}

func Create(options ...interface{}) *Carbon {
	loc, _ := time.LoadLocation("UTC")
	format := "2006-01-02"
	if len(options) == 0 {
		return &Carbon{time.Now()}
	}
	if len(options) >= 2 {
		loc, _ = time.LoadLocation(options[1].(string))
	}
	t, _ := time.ParseInLocation(string(format), options[0].(string), loc)
	return &Carbon{
		Date: t,
	}
}

// 获取当前时间的开始时间
func (c *Carbon) StartOfDay() string {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	return c.Ymd() + " 00:00:00"
}

// 获取当前时间的结束时间
func (c *Carbon) EndOfDay() string {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	return c.Ymd() + " 23:59:59"
}

// 格式化时间

func (c *Carbon) Format(format string) string {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	return c.Date.Format("2006" + format + "01" + format + "02 15:04:05")
}

// 获取时间戳
func (c *Carbon) Timestamp() int64 {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	return c.Date.Unix()
}

// 时间戳转换时间

func (c *Carbon) TimestampToDate(timestamp int64, format ...interface{}) string {
	if len(format) == 0 {
		return time.Unix(timestamp, 0).String()
	}
	var t_format = ""
	switch format[0].(string) {
	case "Ymd":
		t_format = "2006-01-02"
	case "Ymd/":
		t_format = "2006/01/02"
	case "Ymdh":
		t_format = "2006-01-02 15"
	default:
		t_format = "2006-01-02 15:04:05"
	}
	return time.Unix(timestamp, 0).Format(t_format)
}

// 获取本周周一时间
func (c *Carbon) StartOfWeek() string {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	offset := int(time.Monday - c.Date.Weekday())
	if offset > 0 {
		offset = -6
	}
	return c.Date.AddDate(0, 0, offset).Format("2006-01-02 00:00:00")
}

// 获取本周周日时间
func (c *Carbon) EndOfWeek() string {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	offset := int(time.Saturday - c.Date.Weekday())
	if offset > 0 {
		offset = +6
	}
	return c.Date.AddDate(0, 0, offset).Format("2006-01-02 00:00:00")
}

// 获取年月日
func (c *Carbon) Ymd(format ...FormatStr) string {
	if len(format) == 0 {
		return c.Date.Format("2006-01-02")
	} else {
		return c.Date.Format("2006" + format[0].(string) + "01" + format[0].(string) + "02")
	}
}

func (c *Carbon) Parse(parse ...interface{}) *Carbon {
	if len(parse) == 0 {
		c.Date = time.Now()
		return c
	}
	switch {

	case parse[0].(string) == "today":
		c.Date = time.Now()
	case parse[0].(string) == "yesterday":
		c.Date = time.Now().AddDate(0, 0, -1)
	case parse[0].(string) == "tomorrow":
		c.Date = time.Now().AddDate(0, 0, 1)

	case strings.Contains(parse[0].(string), "+") || strings.Contains(parse[0].(string), "-"):
		actions := strings.Split(parse[0].(string), " ")
		if len(actions) > 1 {
			i, _ := strconv.Atoi(actions[0])
			switch {
			case actions[1] == "days":
				c.Date = time.Now().AddDate(0, 0, i)
			case actions[1] == "weeks":
				c.Date = time.Now().AddDate(0, 0, i*7)
			case actions[1] == "months":
				c.Date = time.Now().AddDate(0, i, 0)
			case actions[1] == "year":
				c.Date = time.Now().AddDate(i, 0, 0)
			}
		} else {
			c.Date = Create(actions[0], "PRC").Date
		}
	case strings.Contains(parse[0].(string), "next") || strings.Contains(parse[0].(string), "last"):
		actions := strings.Split(parse[0].(string), " ")
		offset := 0
		switch actions[1] {
		case "monday":
			offset = int(time.Monday - time.Now().Weekday())
		case "tuesday":
			offset = int(time.Tuesday - time.Now().Weekday())
		case "wednesday":
			offset = int(time.Wednesday - time.Now().Weekday())
		case "thursday":
			offset = int(time.Thursday - time.Now().Weekday())
		case "friday":
			offset = int(time.Friday - time.Now().Weekday())
		case "saturday":
			offset = int(time.Saturday - time.Now().Weekday())
		case "sunday":
			offset = int(time.Sunday-time.Now().Weekday()) + 7
		}
		if actions[0] == "next" {
			c.Date = time.Now().AddDate(0, 0, offset+7)
		} else {
			c.Date = time.Now().AddDate(0, 0, offset)
		}
	default:
		c.Date = Create(parse[0].(string), "PRC").Date
	}
	return c
}

func (c *Carbon) IsWeekday() bool {
	if c.Date.IsZero() {
		c.Date = time.Now()
	}
	var wednesday = c.Date.Weekday()
	if (wednesday == 0) || (wednesday == 5) {
		return false
	}
	return true
}

func Hello() {
	fmt.Println("Hello, World!")
}

package time

import (
	"fmt"
	"time"
)

type Time interface {
	Timestamp() (int64, error)
	Now() (string, error)
	FormatTime(timestamp int64) (string, error)
}

type TimeImpl struct {
}

func (obj *TimeImpl) Timestamp() (int64, error) {
	return time.Now().Unix(), nil
}

func (obj *TimeImpl) Now() (string, error) {
	return time.Now().Format("2006-01-02 15:04:05"), nil
}

func (obj *TimeImpl) FormatTime(timestamp int64) (string, error) {
	tm := time.Unix(timestamp, 0)
	year := tm.Year()
	month := tm.Month()
	day := tm.Day()
	hour := tm.Hour()
	minute := tm.Minute()
	second := tm.Second()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, int(month), day, hour, minute, second), nil
}

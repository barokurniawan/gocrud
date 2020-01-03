package util

import "time"

type Datetime time.Time

func (d Datetime) String() string {
	return time.Time(d).String()
}

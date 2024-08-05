package entity

import (
	"daily_matter/constant"
	"time"
)

type Matter struct {
	Title       string
	Desc        string
	TimeStart   time.Time
	TimeEnd     time.Time
	State       constant.State
	MatterClock *Clock
}

type Clock struct {
	FatherMatter *Matter
	Set          bool
	ClockTime    time.Time
}

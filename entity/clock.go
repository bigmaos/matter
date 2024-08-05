package entity

import (
	"time"
)

type Clock struct {
	FatherMatter *Matter
	Set          bool
	ClockTime    time.Time
}

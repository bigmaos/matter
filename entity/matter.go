package entity

import (
	"time"
)

type Matter struct {
	Title       string
	Desc        string
	TimeStart   time.Time
	TimeEnd     time.Time
	State       string
	MatterClock *Clock
}

func (m *Matter) Print() []string {
	return []string{m.Title, m.Desc, m.TimeStart.String(), m.TimeEnd.String(), m.State}
}

func (m *Matter) GetTitle() string {
	return m.Title
}

func (m *Matter) GetDesc() string {
	return m.Desc
}

func (m *Matter) SetDesc(s string) {
	m.Desc = s
}

func (m *Matter) GetStartTime() time.Time {
	return m.TimeStart
}

func (m *Matter) SetStartTime(t time.Time) {
	m.TimeStart = t
}

func (m *Matter) GetEndTime() time.Time {
	return m.TimeEnd
}

func (m *Matter) SetEndTime(t time.Time) {
	m.TimeEnd = t
}

func (m *Matter) GetState() string {
	return m.State
}

func (m *Matter) SetState(s string) {
	m.State = s
}

func (m *Matter) GetClock() *Clock {
	return m.MatterClock
}

func (m *Matter) SetClock(c *Clock) {
	m.MatterClock = c
}

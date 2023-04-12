package clock

import (
	"fmt"
	"math"
)

const MinutesInHour = 60
const HoursInDay = 24

type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	clock := Clock{}
	if math.Abs(float64(m)) >= MinutesInHour {
		clock.hours = m / MinutesInHour
	}
	clock.minutes = m % MinutesInHour
	clock.hours += h

	if math.Abs(float64(clock.hours)) >= HoursInDay {
		clock.hours = clock.hours % HoursInDay
	}

	if clock.minutes < 0 {
		clock.hours--
		clock.minutes = MinutesInHour + clock.minutes
	}
	if clock.hours < 0 {
		clock.hours = HoursInDay + clock.hours
	}

	return clock
}

func (c Clock) Add(m int) Clock {
	newClock := New(c.hours, c.minutes+m)
	return newClock
}

func (c Clock) Subtract(m int) Clock {
	newClock := New(c.hours, c.minutes-m)
	return newClock
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

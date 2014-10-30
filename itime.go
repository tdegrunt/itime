package itime

import "time"

type Itime int

func (a Itime) split() (int, int) {
	aHours := int(a / 100)
	aMinutes := int(a) - int(aHours*100)
	return aHours, aMinutes
}

func New(hours, minutes int) Itime {
	return Itime(hours*100 + minutes)
}

func (a Itime) Sub(b Itime) Itime {

	aHours, aMinutes := a.split()
	bHours, bMinutes := b.split()

	rHours := aHours - bHours
	rMinutes := aMinutes - bMinutes

	if rMinutes < 0 {
		rHours -= 1
		rMinutes = 60 + rMinutes
	}

	return Itime(rHours*100 + rMinutes)
}

// Translates for example 2230 into a duration.
func (a Itime) Duration() time.Duration {
	aHours, aMinutes := a.split()
	duration := time.Duration(aHours) * time.Hour
	duration += time.Duration(aMinutes) * time.Minute
	return duration
}

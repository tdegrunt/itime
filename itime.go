package itime

import "time"

type Itime int

// New returns an initialized Itime
func New(hours, minutes int) Itime {
	return Itime(hours*100 + minutes)
}

// Int returns integer value
func (a Itime) Int() int {
	return int(a)
}

// Split returns hours and minutes
func (a Itime) Split() (int, int) {
	aHours := int(a / 100)
	aMinutes := int(a) - int(aHours*100)
	return aHours, aMinutes
}

// Add returns the time a+b
func (a Itime) Add(b Itime) Itime {

	aHours, aMinutes := a.Split()
	bHours, bMinutes := b.Split()

	rHours := aHours + bHours
	rMinutes := aMinutes + bMinutes

	if rMinutes >= 60 {
		rHours += 1
		rMinutes -= 60
	}

	return Itime(rHours*100 + rMinutes)
}

// Sub returns the time a-b
func (a Itime) Sub(b Itime) Itime {

	aHours, aMinutes := a.Split()
	bHours, bMinutes := b.Split()

	rHours := aHours - bHours
	rMinutes := aMinutes - bMinutes

	if rMinutes < 0 {
		rHours -= 1
		rMinutes = 60 + rMinutes
	}

	return Itime(rHours*100 + rMinutes)
}

// Before reports whether the Itime instant a is before b.
func (a Itime) Before(b Itime) bool {
	return int(a) < int(b)
}

// After reports whether the Itime instant a is after b.
func (a Itime) After(b Itime) bool {
	return int(a) > int(b)
}

// Equal reports whether the Itime instant a is equal to b.
func (a Itime) Equal(b Itime) bool {
	return int(a) == int(b)
}

// Duration returns Itime in time.Duration
func (a Itime) Duration() time.Duration {
	aHours, aMinutes := a.Split()
	duration := time.Duration(aHours) * time.Hour
	duration += time.Duration(aMinutes) * time.Minute
	return duration
}

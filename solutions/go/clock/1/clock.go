package clock

import "fmt"

const (
	minutesPerHour = 60
	hoursPerDay    = 24
	minutesPerDay  = minutesPerHour * hoursPerDay
)

type Clock struct {
	minutes int // минуты от полуночи, 0..1439
}

func New(h, m int) Clock {
	total := h*minutesPerHour + m
	total %= minutesPerDay
	if total < 0 {
		total += minutesPerDay
	}
	return Clock{minutes: total}
}

func (c Clock) Add(m int) Clock {
	total := c.minutes + m
	total %= minutesPerDay
	if total < 0 {
		total += minutesPerDay
	}
	c.minutes = total
	return c
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	h := c.minutes / minutesPerHour
	m := c.minutes % minutesPerHour
	return fmt.Sprintf("%02d:%02d", h, m)
}

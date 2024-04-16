package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	totalMinutes := (h*60 + m) % (24 * 60)
	if totalMinutes < 0 {
		totalMinutes += 24 * 60
	}
	return Clock{
		hours:   totalMinutes / 60,
		minutes: totalMinutes % 60,
	}
}

func (c Clock) Add(m int) Clock {
	totalMinutes := (c.hours*60 + c.minutes + m) % (24 * 60)
	if totalMinutes < 0 {
		totalMinutes += 24 * 60
	}
	return Clock{
		hours:   totalMinutes / 60,
		minutes: totalMinutes % 60,
	}
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)

}

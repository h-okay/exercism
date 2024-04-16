package clock

import "fmt"

type Clock int

func New(h, m int) Clock {
	clock := Clock((h*60 + m) % (24 * 60))
	if clock < 0 {
		clock += 24 * 60
	}
	return clock

}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)

}

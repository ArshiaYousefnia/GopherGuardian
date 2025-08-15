package monitor

import "time"

type TimedChecker struct {
	Checker
}

func (t *TimedChecker) CheckWithDuration() (time.Duration, error) {
	start := time.Now()
	err := t.Check()

	return time.Since(start), err
}

type Checker interface {
	Check() error
	Interval() time.Duration
	Name() string
	Type() string
}

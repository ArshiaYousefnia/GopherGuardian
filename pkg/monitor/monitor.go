package monitor

import "time"

type Checker interface {
	Check() error
	Interval() time.Duration
	Name() string
	Type() string
}

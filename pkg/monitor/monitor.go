package monitor

import (
	"gopehrguardian/pkg/alert"
	"log"
	"time"
)
import "gopehrguardian/pkg/config"

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

func GetChecker(t *config.Target) Checker {
	switch t.Type {
	case config.TypeHTTP:
		return &HttpChecker{
			NameVal:     t.Name,
			AddressVal:  t.Address,
			IntervalVal: time.Second * time.Duration(t.Interval),
		}
	case config.TypeTCP:
		return &TCPChecker{
			NameVal:     t.Name,
			AddressVal:  t.Address,
			IntervalVal: time.Second * time.Duration(t.Interval),
		}
	}

	return nil
}

func Monitor(checker *TimedChecker, target *config.Target) {
	for {
		duration, err := checker.CheckWithDuration()
		if err != nil {
			alert.AlertTarget(target)
		} else {
			log.Printf("%s - %s - %s", checker.Name(), target.Address, duration)
		}

		time.Sleep(checker.Interval())
	}
}

package monitor

import (
	"gopehrguardian/pkg/config"
	"time"
)

type ICMPChecker struct {
	NameVal     string
	AddressVal  string
	IntervalVal time.Duration
}

func (c *ICMPChecker) Name() string {
	return c.NameVal
}

func (c *ICMPChecker) Interval() time.Duration {
	return c.IntervalVal
}

func (c *ICMPChecker) Type() string {
	return config.TypeICMP
}

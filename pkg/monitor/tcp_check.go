package monitor

import (
	"gopehrguardian/pkg/config"
	"net"
	"time"
)

type TCPChecker struct {
	NameVal     string
	AddressVal  string
	IntervalVal time.Duration
}

func (h *TCPChecker) Check() error {
	timeout := min(10*time.Second, h.IntervalVal)

	conn, err := net.DialTimeout(config.TypeTCP, h.AddressVal, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}

func (h *TCPChecker) Name() string {
	return h.NameVal
}

func (h *TCPChecker) Interval() time.Duration {
	return h.IntervalVal
}

func (h *TCPChecker) Type() string {
	return config.TypeTCP
}

package monitor

import (
	"fmt"
	"gopehrguardian/pkg/config"
	"net/http"
	"time"
)

type HttpChecker struct {
	NameVal     string
	AddressVal  string
	IntervalVal time.Duration
}

func (h *HttpChecker) Check() error {
	timeout := min(10*time.Second, h.IntervalVal)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Head(h.AddressVal)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("http status code %d", resp.StatusCode)
	}

	return nil
}

func (h *HttpChecker) Name() string {
	return h.NameVal
}

func (h *HttpChecker) Interval() time.Duration {
	return h.IntervalVal
}

func (h *HttpChecker) Type() string {
	return config.TypeHTTP
}

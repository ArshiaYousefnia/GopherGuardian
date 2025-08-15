package config

import "fmt"

const (
	TypeHTTP = "http"
	TypeTCP  = "tcp"
	TypeICMP = "icmp"
)

var validTypes = map[string]struct{}{
	TypeHTTP: {},
	TypeTCP:  {},
	TypeICMP: {},
}

func validateConfig(config *Config) error {
	for i := range config.Targets {
		err := validateTarget(&config.Targets[i])
		if err != nil {
			return fmt.Errorf("invalid target a index %d: %s", i, err.Error())
		}
	}

	return nil
}

func validateTarget(target *Target) error {
	if target.Name == "" {
		return fmt.Errorf("target name is required")
	}

	if _, ok := validTypes[target.Type]; !ok {
		return fmt.Errorf("invalid target type: %s", target.Type)
	}

	err := validateAddress(target.Address, target.Type)
	if err != nil {
		return err
	}

	if target.Interval <= 0 {
		return fmt.Errorf("invalid interval %d", target.Interval)
	}

	if target.Alert != nil {
		err := validateAlert(target.Alert)
		if err != nil {
			return fmt.Errorf("invalid alert: %s", err)
		}
	}

	return nil
}

func validateAddress(address string, targetType string) error {
	switch targetType {
	case TypeHTTP:
		if !httpRegex.MatchString(address) {
			return fmt.Errorf("invalid http address: %s", address)
		}
	case TypeTCP:
		if !tcpRegex.MatchString(address) {
			return fmt.Errorf("invalid tcp address: %s", address)
		}
	case TypeICMP:
		if !icmpRegex.MatchString(address) {
			return fmt.Errorf("invalid icmp address: %s", address)
		}
	}

	return nil
}

func validateAlert(alert *Alert) error {
	if alert.Email == "" && alert.Telegram == "" {
		return fmt.Errorf("alert email or telegram is required")
	}
	return nil
}

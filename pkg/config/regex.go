package config

import "regexp"

var (
	httpRegex = regexp.MustCompile(`^https?://\S+$`)
	tcpRegex  = regexp.MustCompile(`^[\w.-]+:\d+$`)
	icmpRegex = regexp.MustCompile(`^[\w.-]+$`)
)

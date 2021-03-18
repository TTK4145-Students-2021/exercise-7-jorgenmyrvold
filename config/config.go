package config

import "time"

var PeerPort int = 12345
var BcastPort int = 16569
var Timeout time.Duration = 5000 * time.Millisecond

type Counter struct {
	Counter   int
	Timestamp time.Time
}

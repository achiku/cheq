package cheq

import "time"

// QueStats que stats struct
type QueStats struct {
	Time              time.Time
	ScheduledJobCount int
	RunningJobCount   int
	ErrorCount        int
}

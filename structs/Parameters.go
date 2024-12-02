package structs

import "time"

type Parameters struct {
	ParametersID int
	MinimumCount int
	MaximumCount int
	Deadline     time.Time
}

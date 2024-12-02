package structs

import "time"

type Parameters struct {
	ParametersID int       `db:"parameter_id"`
	MinimumCount int       `db:"minimum_count"`
	MaximumCount int       `db:"maximum_count"`
	Deadline     time.Time `db:"deadline"`
}

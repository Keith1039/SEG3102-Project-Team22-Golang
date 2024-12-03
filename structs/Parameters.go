package structs

import "time"

type Parameters struct {
	ParametersID int       `db:"parameters_id"`
	MinimumCount int       `db:"min_count"`
	MaximumCount int       `db:"max_count"`
	Deadline     time.Time `db:"deadline"`
}

func (params Parameters) Validate() map[string]string {
	errors := make(map[string]string)
	if params.MinimumCount < 1 {
		errors["Minimum"] = "Invalid minimum, minimum must be a positive number greater than 1"
	}
	if params.MaximumCount < 1 || params.MaximumCount < params.MinimumCount {
		errors["Maximum"] = "Invalid maximum, maximum must be a positive number greater than 1 and greater than or equal to the minimum"
	}
	return errors
}

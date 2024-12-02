package structs

import "time"

type Team struct {
	TeamID       int
	ParametersID int
	TeamName     string
	CreationDate time.Time
	Liaison      int
	Status       string
}

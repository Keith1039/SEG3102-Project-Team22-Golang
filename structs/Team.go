package structs

type Team struct {
	TeamID       int    `db:"team_id"`
	ParametersID int    `db:"parameters_id"`
	TeamName     string `db:"team_name"`
	Liaison      int    `db:"liaison"`
	Status       string `db:"status"`
}

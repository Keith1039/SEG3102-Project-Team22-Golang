package repositories

import (
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var teamTable = "TEAMS"

func SaveTeam(team structs.Team, dbpool *pgxpool.Pool) {

}

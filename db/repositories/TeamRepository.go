package repositories

import (
	"context"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var teamTable = "TEAMS"

func SaveTeam(ctx context.Context, team structs.Team, dbpool *pgxpool.Pool) {

}

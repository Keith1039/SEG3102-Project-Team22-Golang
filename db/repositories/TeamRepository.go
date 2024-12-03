package repositories

import (
	"context"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

var teamTable = "TEAMS"

func SaveTeam(ctx context.Context, dbpool *pgxpool.Pool, team structs.Team) {

}

func GetAllTeams(ctx context.Context, dbpool *pgxpool.Pool) []structs.Team {
	var teams []structs.Team
	err := pgxscan.Get(ctx, dbpool, &teams, `SELECT * FROM teams`)
	if err != nil {
		return teams
	}
	return teams
}

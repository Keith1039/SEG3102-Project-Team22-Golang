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

func GetAllTeams(ctx context.Context, dbpool *pgxpool.Pool) []*structs.Team {
	var teams []*structs.Team
	rows, err := dbpool.Query(ctx, `SELECT team_id, team_name, liaison, parameters_id FROM teams`)
	err = pgxscan.ScanAll(&teams, rows)
	if err != nil {
		return nil
	}

	return teams
}

package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var teamTable = "TEAMS"

func SaveTeam(ctx context.Context, dbpool *pgxpool.Pool, team structs.Team) {

}

func GetTeam(ctx context.Context, dbpool *pgxpool.Pool, team *structs.Team, teamID int) bool {
	var err error
	flag := false

	err = pgxscan.Get(ctx, dbpool, team, `SELECT team_id, team_name, liaison, parameters_id FROM teams WHERE team_id=$1;`, teamID)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	} else {
		flag = true
	}
	return flag
}

func UpdateTeam(ctx context.Context, dbpool *pgxpool.Pool, team *structs.Team) bool {
	flag := false
	_, err := dbpool.Query(ctx, "UPDATE teams SET team_name=$1, liaison=$2 WHERE team_id=$3;", team.TeamName, team.Liaison, team.TeamID)
	if err != nil {
		log.Fatal(err)
	} else {
		flag = true
	}
	return flag
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

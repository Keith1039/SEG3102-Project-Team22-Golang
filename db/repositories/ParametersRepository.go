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

var paramTable = "PARAMETERS"

func SaveParameters(ctx context.Context, dbpool *pgxpool.Pool, param structs.Parameters) bool {
	flag := false
	_, err := dbpool.Query(ctx, `INSERT INTO parameters(min_count, max_count) VALUES ($1, $2);`, param.MinimumCount, param.MaximumCount)
	if err != nil {
		log.Fatal(err)

	} else {
		flag = true
	}
	return flag
}

func GetParameters(ctx context.Context, dbpool *pgxpool.Pool, params *structs.Parameters, paramID int) bool {
	var err error
	flag := false

	err = pgxscan.Get(ctx, dbpool, params, `SELECT parameters_id, min_count, max_count, deadline FROM parameters WHERE parameters_id=$1;`, paramID)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	} else {
		flag = true
	}
	return flag
}

func UpdateParameters(ctx context.Context, dbpool *pgxpool.Pool, param structs.Parameters, paramID int) bool {
	flag := false
	_, err := dbpool.Query(ctx, "UPDATE parameters SET min_count=$1, max_count=$2 WHERE parameters_id=$3;", param.MinimumCount, param.MaximumCount, paramID)
	if err != nil {
		log.Fatal(err)
	} else {
		flag = true
	}
	return flag
}

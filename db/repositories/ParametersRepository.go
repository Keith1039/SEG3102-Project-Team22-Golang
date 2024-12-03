package repositories

import (
	"context"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var paramTable = "PARAMETERS"

func SaveParameters(ctx context.Context, param structs.Parameters, dbpool *pgxpool.Pool) bool {
	flag := false
	_, err := dbpool.Query(ctx, `INSERT INTO parameters(min_count, max_count) VALUES ($1, $2);`, param.MinimumCount, param.MaximumCount)
	if err != nil {
		log.Fatal(err)

	} else {
		flag = true
	}
	return flag
}

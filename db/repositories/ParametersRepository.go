package repositories

import (
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var paramTable = "PARAMETERS"

func SaveParameters(param structs.Parameters, dbpool *pgxpool.Pool) {

}
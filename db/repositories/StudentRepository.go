package repositories

import (
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var studentTable = "STUDENTS"

func SaveStudent(student structs.Student, dbpool *pgxpool.Pool) {

}

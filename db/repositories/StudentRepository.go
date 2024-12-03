package repositories

import (
	"context"
	"errors"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var studentTable = "STUDENTS"

func SaveStudent(ctx context.Context, student structs.Student, dbpool *pgxpool.Pool) {

}

func CheckStudentNumber(ctx context.Context, dbpool *pgxpool.Pool, studentNumber int) bool {
	flag := false
	var studentNumber2 int
	err := dbpool.QueryRow(ctx, `SELECT student_id FROM students WHERE student_id=$1;`, studentNumber).Scan(&studentNumber2)
	if !errors.Is(err, pgx.ErrNoRows) {
		flag = true
	}
	return flag
}

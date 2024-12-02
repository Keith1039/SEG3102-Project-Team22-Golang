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

func SaveUser(ctx context.Context, dbpool *pgxpool.Pool, user structs.User) bool {
	var err error
	var userID int
	flag := false
	err = dbpool.QueryRow(ctx, `SELECT user_id FROM users WHERE email=$1;`, user.Email).Scan(&userID)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = dbpool.Query(ctx, `INSERT INTO users(first_name, last_name, email, role) VALUES ($1, $2, $3, $4);`, user.FirstName, user.LastName, user.Email, user.Role)
		if err != nil {
			log.Fatal(err)
		}
		flag = true
	}
	return flag
}

func GetUser(ctx context.Context, dbpool *pgxpool.Pool, user *structs.User, email string) bool {
	var err error
	flag := false

	err = pgxscan.Get(ctx, dbpool, user, `SELECT user_id, first_name, last_name, email, role FROM users WHERE email=$1;`, email)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	} else {
		flag = true
	}
	return flag
}

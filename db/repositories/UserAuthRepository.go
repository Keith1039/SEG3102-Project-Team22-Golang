package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func SaveCredentials(ctx context.Context, dbpool *pgxpool.Pool, auth structs.UserAuth) bool {
	var err error
	var userID int

	flag := false
	err = dbpool.QueryRow(ctx, `SELECT user_id FROM user_auth WHERE username=$1`, auth.Username).Scan(&userID)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = dbpool.Query(ctx, `INSERT INTO user_auth(username, password, user_id) VALUES ($1, $2, $3);`, auth.Username, auth.Password, auth.UserID)
		if err != nil {
			log.Fatal(err)
		}
		flag = true
	}
	return flag
}

func CheckUsername(ctx context.Context, dbpool *pgxpool.Pool, username string) bool {
	var err error
	var userID int

	flag := false
	err = dbpool.QueryRow(ctx, `SELECT user_id FROM user_auth WHERE username=$1`, username).Scan(&userID)
	if errors.Is(err, sql.ErrNoRows) {
		flag = true
	}
	return flag
}

package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/templates"
	"github.com/a-h/templ"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
)

var dbpool *pgxpool.Pool
var loggedIn bool
var user structs.User

var ctx = context.Background()

func init() {
	var err error
	err = os.Setenv("DATABASE_URL", "postgres://postgres:localDB12@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return
	}
	dbpool, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	var greeting string
	err = dbpool.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("The database is connected")
}

func main() {

	s := templates.Login()
	home := templates.Hello(&user)
	signUpPage := templates.SignUp()

	http.Handle("/", templ.Handler(s))
	http.Handle("/home", templ.Handler(home))
	http.Handle("/register", templ.Handler(signUpPage))
	http.HandleFunc("/login/", LoginRequest)
	http.HandleFunc("/signup/", SignUpRequest)

	fmt.Println("listening on port 8080")
	defer dbpool.Close()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	var userId int

	username := r.PostFormValue("username")
	userPassword := r.PostFormValue("password")

	err := dbpool.QueryRow(ctx, `SELECT user_id FROM user_auth WHERE username=$1 AND password=$2;`, username, userPassword).Scan(&userId)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("query is broken")
		log.Fatal(err)
	} else if err == nil {
		loggedIn = true
		fmt.Println("Logging in")
		err := pgxscan.Get(ctx, dbpool, &user, `SELECT first_name, last_name, email, role FROM users WHERE user_id=$1;`, userId)
		if err != nil {
			panic(err)
		}
		//user = structs.User{FirstName: firstName, LastName: lastName, Email: email, Role: role}
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		fmt.Println("Something broke")
		panic(err)
	}
}

func SignUpRequest(w http.ResponseWriter, r *http.Request) {
	var userID int
	username := r.PostFormValue("username")
	userPassword := r.PostFormValue("password")
	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	email := r.PostFormValue("email")
	role := r.PostFormValue("role")

	err := dbpool.QueryRow(ctx, `SELECT user_id FROM users WHERE email=$1;`, email).Scan(&userID)
	if errors.Is(err, sql.ErrNoRows) {
		//tempUser := structs.User{FirstName: firstName, LastName: lastName, Email: email, Role: role}
		_, err = dbpool.Query(ctx, `INSERT INTO users(first_name, last_name, email, role) VALUES ($1, $2, $3, $4);`, firstName, lastName, email, role)
		if err != nil {
			return
		}

		err = pgxscan.Get(ctx, dbpool, &user, `SELECT user_id, first_name, last_name, email, role FROM users WHERE email=$1;`, email)
		if err != nil {
			log.Println(err)
		}

		_, err := dbpool.Query(ctx, `INSERT INTO user_auth(username, password, user_id) VALUES ($1, $2, $3);`, username, userPassword, user.UserID)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/home", http.StatusFound)

	} else {
		fmt.Println("email already exists")
	}
}

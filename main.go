package main

import (
	"context"
	"fmt"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/db/repositories"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/structs"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/templates"
	"github.com/a-h/templ"
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
	username := r.PostFormValue("username")
	userPassword := r.PostFormValue("password")
	tempAuth := structs.UserAuth{Username: username, Password: userPassword}

	flag := repositories.GetCredentials(ctx, dbpool, &tempAuth)
	if flag {
		fmt.Println("Logging in")
		repositories.GetUserByID(ctx, dbpool, &user, tempAuth.UserID)
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		fmt.Println("User not found")
	}
}

func SignUpRequest(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	userPassword := r.PostFormValue("password")
	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	email := r.PostFormValue("email")
	role := r.PostFormValue("role")

	tempAuth := structs.UserAuth{UserID: 0, Username: username, Password: userPassword}
	if repositories.CheckUsername(ctx, dbpool, tempAuth.Username) {
		tempUser := structs.User{UserID: 0, FirstName: firstName, LastName: lastName, Email: email, Role: role}
		flag := repositories.SaveUser(ctx, dbpool, tempUser)
		if flag {
			flag = repositories.GetUser(ctx, dbpool, &user, tempUser.Email)
			if flag {
				tempAuth.UserID = user.UserID
				repositories.SaveCredentials(ctx, dbpool, tempAuth)
			}
			http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			fmt.Println("email already exists")
		}
	} else {
		fmt.Println("username already taken")
	}

}

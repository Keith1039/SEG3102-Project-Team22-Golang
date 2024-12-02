package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/templates"
	"github.com/a-h/templ"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	firstName string
	lastName  string
	email     string
	role      string
}

var db *sql.DB
var loggedIn bool
var userStruct User

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "localDB12"
	dbname   = "postgres"
)

func init() {
	var err error

	loggedIn = false
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}

func main() {

	s := templates.Login()
	http.Handle("/", templ.Handler(s))
	http.HandleFunc("/login/", LoginRequest)
	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer db.Close()
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	var userId int
	var firstName string
	var lastName string
	var email string
	var role string

	username := r.PostFormValue("username")
	userPassword := r.PostFormValue("password")

	err := db.QueryRow(`SELECT user_id FROM user_auth WHERE username=$1 AND password=$2;`, username, userPassword).Scan(&userId)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("query is broken")
		log.Fatal(err)
	} else if err == nil {
		loggedIn = true
		fmt.Println("Logging in")
		row := db.QueryRow(`SELECT first_name, last_name, email, role FROM users WHERE user_id=$1;`, userId)
		err = row.Scan(&firstName, &lastName, &email, &role)
		if err != nil {
			panic(err)
		}
		userStruct = User{firstName, lastName, email, role}
		fmt.Println(userStruct)
		http.Redirect(w, r, "/welcome", http.StatusFound)
	} else {
		fmt.Println("Something broke")
		panic(err)
	}

}

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
	"strconv"
)

var (
	dbpool *pgxpool.Pool
	user   structs.User
	team   structs.Team
	params structs.Parameters
	ctx    = context.Background()
)

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

	fmt.Println("The database is connected")
}

func main() {
	initialSignUp := templates.SignUpForm{"", "", "", "", "", ""}
	initialLogin := templates.LoginForm{Username: "", Password: ""}
	initialParamsCreate := templates.ParametersForm{Minimum: "", Maximum: ""}

	home := templates.Hello(&user)
	signUpPage := templates.SignUp(initialSignUp, map[string]string{"Username": "", "Email": ""}) // initially have nothing in the form
	login := templates.Login(initialLogin, "")                                                    // initially we have nothing in the form
	createParamsPage := templates.ParameterCreate(initialParamsCreate, map[string]string{"Minimum": "", "Maximum": ""})
	teams := templates.Teams(&user, repositories.GetAllTeams(ctx, dbpool))
	teamEdit := templates.TeamEdit(&team, &params, map[string]string{})

	http.Handle("/", templ.Handler(login))
	http.Handle("/home", templ.Handler(home))
	http.Handle("/signup", templ.Handler(signUpPage))
	http.Handle("/params-create", templ.Handler(createParamsPage))
	http.Handle("/teams", templ.Handler(teams))
	http.Handle("/team", templ.Handler(teamEdit))

	http.HandleFunc("/login/", HandleLoginRequest)
	http.HandleFunc("/register/", HandleSignUpRequest)
	http.HandleFunc("/create-parameter/", HandleParameterCreation)
	http.HandleFunc("/get-teams/", GetTeamsAndRedirect)
	http.HandleFunc("/team-edit", GetTeamAndRedirect)
	http.HandleFunc("/edit-team", ValidateAndSave)

	fmt.Println("listening on port 8080")
	defer dbpool.Close()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func HandleLoginRequest(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	userPassword := r.PostFormValue("password")
	tempAuth := structs.UserAuth{Username: username, Password: userPassword}

	flag := repositories.GetCredentials(ctx, dbpool, &tempAuth)
	if flag {
		fmt.Println("Logging in")
		repositories.GetUserByID(ctx, dbpool, &user, tempAuth.UserID)
		//http.Redirect(w, r, "/home", http.StatusFound)
		hxRedirect(w, r, "/home")
	} else {
		Render(w, r, templates.Login(templates.LoginForm{Username: username, Password: userPassword}, "username or password is incorrect"))
	}
}

func HandleSignUpRequest(w http.ResponseWriter, r *http.Request) {
	errors := map[string]string{}
	username := r.PostFormValue("username")
	userPassword := r.PostFormValue("password")
	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	email := r.PostFormValue("email")
	role := r.PostFormValue("role")

	errors["Username"] = ""
	errors["Email"] = ""

	values := templates.SignUpForm{Username: username, Password: userPassword, FirstName: firstName, LastName: lastName, Email: email, Role: role}

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
			//http.Redirect(w, r, "/home", http.StatusFound)
			hxRedirect(w, r, "/home")
		} else {
			errors["Email"] = "Email already exists"
			Render(w, r, templates.SignUp(values, errors))
		}
	} else {
		errors["Username"] = "Username already exists"
		Render(w, r, templates.SignUp(values, errors))
	}
}

func HandleParameterCreation(w http.ResponseWriter, r *http.Request) {
	var minCount int
	var maxCount int
	var err, err2 error
	paramForm := templates.ParametersForm{Minimum: r.PostFormValue("minimum"), Maximum: r.PostFormValue("maximum")}
	errors := paramForm.Validate()
	if len(errors) > 0 {
		Render(w, r, templates.ParameterCreate(paramForm, errors))
	} else {
		minCount, err = strconv.Atoi(r.PostFormValue("minimum"))
		maxCount, err2 = strconv.Atoi(r.PostFormValue("maximum"))
		if err == nil && err2 == nil {
			tempParams := structs.Parameters{MinimumCount: minCount, MaximumCount: maxCount}
			repositories.SaveParameters(ctx, dbpool, tempParams)
			Render(w, r, templates.ParameterCreate(templates.ParametersForm{Minimum: "", Maximum: ""}, errors))
		} else {
			errors["Conversion"] = "Conversion error"
			Render(w, r, templates.ParameterCreate(paramForm, errors))
		}
	}
}

func GetTeamsAndRedirect(w http.ResponseWriter, r *http.Request) {
	teams := repositories.GetAllTeams(ctx, dbpool)
	Render(w, r, templates.Teams(&user, teams))
}

func GetTeamAndRedirect(w http.ResponseWriter, r *http.Request) {
	teamID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	repositories.GetTeam(ctx, dbpool, &team, teamID)
	repositories.GetParameters(ctx, dbpool, &params, team.ParametersID)
	hxRedirect(w, r, "/team")
}

func ValidateAndSave(w http.ResponseWriter, r *http.Request) {
	var paramID, minimum, maximum, teamID, liaison int
	errors := map[string]string{}
	tempTeam := structs.Team{}
	tempParams := structs.Parameters{}
	var err error

	paramID = params.ParametersID
	tempParams.ParametersID = paramID
	minimum, err = strconv.Atoi(r.PostFormValue("minimum"))
	if err != nil {
		errors["Minimum"] = "Invalid value entered"
	} else {
		tempParams.MinimumCount = minimum
	}

	maximum, err = strconv.Atoi(r.PostFormValue("maximum"))
	if err != nil {
		errors["Maximum"] = "Invalid value entered"
	} else {
		tempParams.MaximumCount = maximum
	}
	teamID = team.TeamID
	tempTeam.TeamID = teamID
	teamName := r.PostFormValue("team_name")
	tempTeam.TeamName = teamName

	liaison, err = strconv.Atoi(r.PostFormValue("liaison"))
	if err != nil {
		errors["liaison"] = "Invalid value"
	} else {
		if repositories.CheckStudentNumber(ctx, dbpool, liaison) {
			tempTeam.Liaison = liaison
		} else {
			errors["liaison"] = "Student number doesn't exist"
		}
	}

	if len(errors) > 0 {
		fmt.Println("inputed data is bad")
		fmt.Println(errors)

	} else {
		errors = tempParams.Validate()
		if len(errors) > 0 {
			fmt.Println("failed validate")
			fmt.Println(errors)

		} else {
			repositories.UpdateTeam(ctx, dbpool, &tempTeam)
			repositories.UpdateParameters(ctx, dbpool, &tempParams)
			team = tempTeam
			params = tempParams
		}
	}
	Render(w, r, templates.TeamEdit(&tempTeam, &tempParams, errors))
}

func hxRedirect(w http.ResponseWriter, r *http.Request, url string) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusSeeOther)
		return nil
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
	return nil
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

package app

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/user/{id}").
		HandlerFunc(app.getFunction)

	app.Router.
		Methods("POST").
		Path("/user").
		HandlerFunc(app.postFunction)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	user := &User{}
	err := app.Database.QueryRow("SELECT * FROM `users` WHERE id = ?", id).Scan(&user)
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	log.Println("You fetched a thing!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	sqlStatement := `
			INSERT INTO users (age, email, first_name, last_name)
			VALUES ($1, $2, $3, $4)
			RETURNING id`
	id := 0
	err := app.Database.QueryRow(sqlStatement, user.age, user.firstname, user.lastName, user.email).Scan(&id)
	if err != nil {
		panic(err)
	}

	log.Println("new user got created id: ", id)
	user.id = id
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

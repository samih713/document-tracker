package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/document-tracker/db"
)

const (
	DBNAME = "demo.db"
)

var (
	database *sql.DB
)

func main() {
	// init databse
	var err error
	database, err = db.Open(DBNAME)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	if err := db.Init(database); err != nil {
		log.Fatal(err)
	}
	db.QueryAgents(database)
	// server
	mux := http.NewServeMux()
	// index page
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	mux.HandleFunc("POST /save_agent", SaveAgent)
	mux.HandleFunc("GET /get_agents", GetAgents)
	// server startup
	log.Println("server running on :7979")
	log.Fatal(http.ListenAndServe(":7979", mux))
}

func SaveAgent(w http.ResponseWriter, r *http.Request) {

	_, err := db.InsertAgent(database, map[string]string{
		"name":       r.FormValue("name"),
		"contact_no": r.FormValue("contact_no"),
		"email":      r.FormValue("email"),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func GetAgents(w http.ResponseWriter, r *http.Request) {
	agents, err := db.QueryAgents(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, agent := range agents {
		fmt.Println(agent)
	}
}

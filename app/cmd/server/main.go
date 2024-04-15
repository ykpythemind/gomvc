package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/controllers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("starting...")

	rawdb, err := sql.Open("sqlite3", gomvc.MustEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	app := gomvc.NewApp(rawdb)
	router := controllers.InitRouter(app)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

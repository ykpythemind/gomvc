package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/controllers"
)

func main() {
	fmt.Println("a")

	app := &gomvc.App{}
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

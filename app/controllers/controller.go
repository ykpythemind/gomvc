package controllers

import (
	"github.com/gorilla/mux"
	"github.com/ykpythemind/gomvc"
)

func InitRouter(app *gomvc.App) *mux.Router {
	r := mux.NewRouter()

	posts_controller := NewPostsController(app)
	r.HandleFunc("/", posts_controller.Index)
	// do something

	return r
}

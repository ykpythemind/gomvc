package controllers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/db"
)

func InitRouter(app *gomvc.App) *mux.Router {
	r := mux.NewRouter()
	middleware := NewEssentialMiddleware(app)
	r.Use(middleware.MiddlewareFunc)

	posts_controller := NewPostsController(app)
	r.HandleFunc("/", posts_controller.Index)
	r.HandleFunc("/p", posts_controller.Users)
	// do something

	return r
}

func NewEssentialMiddleware(app *gomvc.App) *EssentialMiddleware {
	return &EssentialMiddleware{App: app}
}

type EssentialMiddleware struct {
	App *gomvc.App
}

func (e *EssentialMiddleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqeustContextDb := e.App.DB()
		con := context.WithValue(r.Context(), db.DBContextKey, reqeustContextDb)
		// do something
		next.ServeHTTP(w, r.WithContext(con))
	})
}

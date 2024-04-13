package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/models"
)

// InitRouter is router initializer
func InitRouter(app *gomvc.App) *mux.Router {
	r := mux.NewRouter()
	middleware := NewEssentialMiddleware(app)
	r.Use(middleware.MiddlewareFunc)

	controller := NewController(app)
	r.HandleFunc("/", controller.IndexPosts)
	r.HandleFunc("/ping", controller.Ping)
	r.HandleFunc("/users", controller.IndexUsers)

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
		reqeustContextDb := e.App.UseDB()
		con := context.WithValue(r.Context(), models.DBContextKey, reqeustContextDb)
		// do something

		next.ServeHTTP(w, r.WithContext(con))
	})
}

func responseError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// Controller is controller
type Controller struct {
	App *gomvc.App
}

// NewController is constructor
func NewController(app *gomvc.App) *Controller {
	return &Controller{
		App: app,
	}
}

// IndexPosts is posts index handler
func (c *Controller) IndexPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post index"))
}

// Ping is handler
func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// IndexUsers is users index handler
func (c *Controller) IndexUsers(w http.ResponseWriter, r *http.Request) {
	db := c.App.UseDB()
	ctx := r.Context()
	users, err := models.ListUsers(ctx, db)
	if err != nil {
		responseError(w, err)
		return
	}
	responseStr := ""
	for _, user := range users {
		fullname, err := user.FullName(ctx)
		if err != nil {
			responseError(w, err)
			return
		}
		responseStr = fmt.Sprintf("%s\nid: %d, name: %s", responseStr, user.ID, fullname)
	}

	w.Write([]byte(responseStr))
}

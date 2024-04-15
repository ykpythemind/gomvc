package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/models"
	"github.com/ykpythemind/gomvc/views/jsonresponse"
	"golang.org/x/xerrors"
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
	r.Methods("POST").Path("/users/new/{name}").HandlerFunc(controller.NewUsers)
	r.HandleFunc("/error", controller.ErrorTest)

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
		// reqeustContextDb := e.App.UseDB()
		// con := context.WithValue(r.Context(), models.DBContextKey, reqeustContextDb)

		// FIXME: random request id
		con := context.WithValue(r.Context(), gomvc.RequestIDKey, "123456")
		// do something

		next.ServeHTTP(w, r.WithContext(con))
	})
}

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	slog.ErrorContext(r.Context(), fmt.Sprintf("error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func responseJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
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

// Error test is test error
func (c *Controller) ErrorTest(w http.ResponseWriter, r *http.Request) {
	err := newerror()

	err = xerrors.Errorf("wrapped error : %w", err)

	// test stacktrace FIXME: not working
	slog.ErrorContext(r.Context(), fmt.Sprintf("error: %v", err))

	w.Write([]byte("pong"))
}

func newerror() error {
	return xerrors.New("test")
}

func (c *Controller) NewUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := "its name"
	if v, ok := vars["name"]; ok {
		a = v
	}

	db := c.App.UseDB()
	ctx := r.Context()

	u := &models.User{Name: a}
	if err := u.Create(ctx, db); err != nil {
		responseError(w, r, err)
		return
	}
}

// IndexUsers is users index handler
func (c *Controller) IndexUsers(w http.ResponseWriter, r *http.Request) {
	db := c.App.UseDB()
	ctx := r.Context()
	users, err := models.ListUsers(ctx, db)
	if err != nil {
		responseError(w, r, err)
		return
	}

	/* レスポンス例
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
	*/

	if resp, err := jsonresponse.JSONUsers(users); err != nil {
		responseError(w, r, err)
		return
	} else {
		if err := responseJSON(w, resp); err != nil {
			responseError(w, r, err)
			return
		}
	}
}

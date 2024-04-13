package controllers

import (
	"net/http"

	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/db"
)

type PostsController struct {
	App *gomvc.App
}

func NewPostsController(app *gomvc.App) *PostsController {
	return &PostsController{
		App: app,
	}
}

// Index is handler func
func (c *PostsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostsController Index"))
}

// Index is handler func
func (c *PostsController) Users(w http.ResponseWriter, r *http.Request) {
	d := db.MustUseDB(r.Context())
	d.Query("SELECT * FROM users", nil)
	w.Write([]byte("PostsController Inex"))
}

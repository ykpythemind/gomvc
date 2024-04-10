package controllers

import (
	"net/http"

	"github.com/ykpythemind/gomvc"
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

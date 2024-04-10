package request_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/controllers"
)

func prepare() (*gomvc.App, func()) {
	app := &gomvc.App{}
	fn := func() {
	}
	return app, fn
}

func handlerTest(r *http.Request, handler http.HandlerFunc) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	handler(recorder, r)
	return recorder
}

func TestPostsIndex(t *testing.T) {
	app, cleanup := prepare()
	t.Cleanup(cleanup)
	con := controllers.NewPostsController(app)
	req := httptest.NewRequest("GET", "/", nil)
	resp := handlerTest(req, con.Index)
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "PostsController Index", string(body))
}

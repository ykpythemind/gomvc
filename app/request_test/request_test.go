package request_test

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ykpythemind/gomvc"
	"github.com/ykpythemind/gomvc/controllers"

	_ "github.com/mattn/go-sqlite3"
)

func prepare() (*gomvc.App, func()) {
	rawdb, err := sql.Open("sqlite3", gomvc.MustEnv("TEST_DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	app := gomvc.NewApp(rawdb)
	fn := func() {
		rawdb.Close()
	}
	return app, fn
}

func handlerTest(r *http.Request, handler http.HandlerFunc) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	handler(recorder, r)
	return recorder
}

func TestUserIndex(t *testing.T) {
	t.SkipNow()

	app, cleanup := prepare()
	t.Cleanup(cleanup)
	r := controllers.InitRouter(app)
	req := httptest.NewRequest("GET", "/users", nil)
	resp := handlerTest(req, r.ServeHTTP)
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "PostsController Index", string(body))
}

func TestPing(t *testing.T) {
	app, cleanup := prepare()
	t.Cleanup(cleanup)
	r := controllers.InitRouter(app)
	req := httptest.NewRequest("GET", "/ping", nil)
	resp := handlerTest(req, r.ServeHTTP)
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "pong", string(body))
}

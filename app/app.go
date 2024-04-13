package gomvc

import (
	"database/sql"

	"github.com/ykpythemind/gomvc/db"
)

// App is the application and config struct
type App struct {
	Revision string // Revision is the git revision of the build

	rawDB *sql.DB
}

func NewApp(rawDB *sql.DB) *App {
	return &App{rawDB: rawDB}
}

func (a *App) DB() *db.DB {
	return db.NewDB(a.rawDB)
}

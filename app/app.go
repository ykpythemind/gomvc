package gomvc

import (
	"database/sql"
	"os"

	"github.com/ykpythemind/gomvc/models"
)

// App is the application and config struct
type App struct {
	Revision string // Revision is the git revision of the build

	rawDB *sql.DB
}

func NewApp(rawDB *sql.DB) *App {
	return &App{rawDB: rawDB}
}

// useDB returns a new DB instance
func (a *App) UseDB() *models.DB {
	return models.NewDB(a.rawDB)
}

func MustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(key + " is not set")
	}
	return v
}

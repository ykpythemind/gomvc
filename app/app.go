package gomvc

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/ykpythemind/gomvc/implements"
	"github.com/ykpythemind/gomvc/interfaces"
	"github.com/ykpythemind/gomvc/models"
)

// App is the application and config struct
type App struct {
	Revision string // Revision is the git revision of the build

	rawDB *sql.DB

	// for dependency injection
	CoffeeList interfaces.CoffeeList
}

func NewApp(rawDB *sql.DB) *App {
	logger := slog.New(NewLogHandler(slog.NewJSONHandler(os.Stdout, nil)))
	slog.SetDefault(logger)

	var coffeeList interfaces.CoffeeList

	// 環境によって実装をさしかえる
	coffeeList = &implements.CoffeeListImpl{}

	return &App{rawDB: rawDB, CoffeeList: coffeeList}
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

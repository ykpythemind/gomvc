package interfaces

import (
	"context"

	"github.com/ykpythemind/gomvc/models"
)

type CoffeeList interface {
	Fetch(ctx context.Context) ([]models.Coffee, error)
}

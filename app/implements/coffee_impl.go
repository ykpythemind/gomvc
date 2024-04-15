package implements

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ykpythemind/gomvc/interfaces"
	"github.com/ykpythemind/gomvc/models"
	"golang.org/x/xerrors"
)

type CoffeeListImpl struct {
}

type coffeeResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *CoffeeListImpl) Fetch(ctx context.Context) ([]models.Coffee, error) {
	coffees := []models.Coffee{}

	url := "https://api.sampleapis.com/coffee/hot"

	client := new(http.Client)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to fetch coffee: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var coffeeResponses []coffeeResponse
	if err := json.Unmarshal(body, &coffeeResponses); err != nil {
		return nil, xerrors.Errorf("failed to parse response of coffee: %w", err)
	}

	for _, cr := range coffeeResponses {
		coffees = append(coffees, models.Coffee{Title: cr.Title, Description: cr.Description})
	}

	return coffees, nil
}

var _ interfaces.CoffeeList = (*CoffeeListImpl)(nil)

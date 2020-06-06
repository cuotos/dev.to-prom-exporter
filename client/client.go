package client

import (
	"dev.to-prom-exporter/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DevtoClient struct {
	ApiKey string
}

func (c DevtoClient) GetUserArticleData() ([]models.Article, error) {

	articles := []models.Article{}

	r, err := http.NewRequest("GET", "https://dev.to/api/articles/me", nil)
	if err != nil {
		return articles, err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Api-Key", c.ApiKey)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return articles, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return articles, err
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(body, &articles); err != nil {
		return articles, err
	}

	return articles, nil
}

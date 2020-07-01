package client

import (
	"dev.to-prom-exporter/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://dev.to/api"
)

type DevtoClient struct {
	ApiKey string
}

func (c *DevtoClient) GetUserArticleData() ([]models.Article, error) {

	articles := []models.Article{}

	resp, err := c.makeApiCall("/articles/me")
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

func (c *DevtoClient) GetUserDetails()(models.User, error){
	user := models.User{}

	resp, err := c.makeApiCall("/users/me")
	if err != nil {
		return user, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(body, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *DevtoClient) makeApiCall(endpoint string) (*http.Response, error){

	endpoint = strings.TrimPrefix(endpoint, "/")

	url := fmt.Sprintf("%s/%s", baseUrl, endpoint)

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Api-Key", c.ApiKey)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
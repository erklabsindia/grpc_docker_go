package news_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// These global variables makes it easy
// to mock these dependencies
// in unit tests.
var (
	jsonUnmarshal = json.Unmarshal
	ioUtilReadAll = ioutil.ReadAll
)

type NewsResponse struct {
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type NewsesResponse struct {
	Articles     []NewsResponse `json:"articles"`
	Status       string         `json:"status"`
	TotalResults int            `json:"totalResults"`
}

type NewsAPI interface {
	GetNews(query string, pageSize int) (NewsesResponse, error)
}

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type NewsApiClient struct {
	baseUrl    string
	apiKey     string
	httpClient HttpClient
}

func (p *NewsApiClient) GetNews(query string, pageSize int) (NewsesResponse, error) {
	var newsResponse NewsesResponse
	//newsResponse := make([]NewsResponse, 0)
	endpoint := fmt.Sprintf("%s?apiKey=%s&pageSize=%d&q=%s", p.baseUrl, p.apiKey, pageSize, query)
	response, err := p.httpClient.Get(endpoint)
	if err != nil {
		return newsResponse, errors.Wrapf(err, `performing request to "%s"`, endpoint)
	}
	defer response.Body.Close()
	statusCode := response.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		return newsResponse, errors.New(fmt.Sprintf("got status code %d", statusCode))
	}
	body, err := ioUtilReadAll(response.Body)
	if err != nil {
		return newsResponse, errors.Wrap(err, "reading response")
	}
	err = jsonUnmarshal(body, &newsResponse)
	if err != nil {
		return newsResponse, errors.Wrap(err, "parsing response")
	}
	return newsResponse, nil
}

func GetNewsApiClient(baseUrl string, apiKey string, timeoutInSeconds int) NewsAPI {
	httpClient := &http.Client{
		Timeout: time.Duration(timeoutInSeconds) * time.Second,
	}
	return &NewsApiClient{
		baseUrl:    baseUrl,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

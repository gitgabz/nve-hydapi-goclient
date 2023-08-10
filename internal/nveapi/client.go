package nveapi

import (
	"net/http"
	"net/url"
	"time"
)

type NveApiClient struct {
	apiKey string
	Client *http.Client
}

func NewClient(apikey string) (c *NveApiClient) {

	c = &NveApiClient{}

	tr := http.Transport{
		TLSHandshakeTimeout:   time.Second * 10,
		ResponseHeaderTimeout: time.Second * 10,
	}

	to := time.Second * 10

	c.Client = &http.Client{
		Transport: &tr,
		Timeout:   to,
	}

	c.apiKey = apikey

	return
}

func joinApiEndpointAndPath(contextPath string) (u string, err error) {

	up, err := url.Parse(apiBaseUri + contextPath)
	if err != nil {
		return
	}

	u = up.String()

	return
}

func appendQueryParams(req *http.Request, qm map[string]string) {

	q := req.URL.Query()

	for k, v := range qm {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

}

func (c *NveApiClient) createRequest(method string, endPointContextPath string) (req *http.Request, err error) {

	u, err := joinApiEndpointAndPath(endPointContextPath)
	if err != nil {
		return
	}

	req, err = http.NewRequest(method, u, nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("X-API-Key", c.apiKey)

	return
}

package nveapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *NveApiClient) GetObservations(rqo RequestQueryObservations) (jr ResponseQueryObservations, err error) {

	qm, err := generateQueryParameterMap(rqo)
	if err != nil {
		return
	}

	req, err := c.createRequest("GET", apiEndpointsV1[Observation])
	if err != nil {
		return
	}

	appendQueryParams(req, qm)

	resp, err := c.Client.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = checkForFailedRequests(resp.StatusCode, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &jr)
	if err != nil {
		return
	}

	return
}

// checkForFailedRequests checks and returns a standard error format for failed requests
func checkForFailedRequests(statusCode int, body []byte) (err error) {

	if statusCode == 400 {
		message := fmt.Sprintf("request failed, status code [%d] - %s", statusCode, body)
		err = errors.New(message)
		return
	}

	if statusCode == 401 {
		message := fmt.Sprintf("authorization failed, status code [%d]", statusCode)
		err = errors.New(message)
		return
	}

	return
}

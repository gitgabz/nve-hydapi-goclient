package nveapi

import "time"

type RequestQueryObservations struct {
	StationId       string `queryParameter:"StationId" required:"true"`
	Parameter       string `queryParameter:"Parameter" required:"true"`
	ResolutionTime  string `queryParameter:"ResolutionTime" required:"true"`
	VersionNumber   string `queryParameter:"VersionNumber"`
	ReferenceTime   string `queryParameter:"ReferenceTime"`
	QualityTypes    string `queryParameter:"QualityTypes"`
	Method          string `queryParameter:"Method"`
	TimeOffset      string `queryParameter:"TimeOffset"`
	CorrectionTypes string `queryParameter:"CorrectionTypes"`
}

type ResponseQueryObservations struct {
	CurrentLink string    `json:"currentLink"`
	ApiVersion  string    `json:"apiVersion"`
	License     string    `json:"license"`
	CreatedAt   time.Time `json:"createdAt"`
	QueryTime   string    `json:"queryTime"`
	ItemCount   int       `json:"itemCount"`
	Data        []struct {
		StationId        string `json:"stationId"`
		StationName      string `json:"stationName"`
		Parameter        int    `json:"parameter"`
		ParameterName    string `json:"parameterName"`
		ParameterNameEng string `json:"parameterNameEng"`
		SerieVersionNo   int    `json:"serieVersionNo"`
		Method           string `json:"method"`
		Unit             string `json:"unit"`
		ObservationCount int    `json:"observationCount"`
		Observations     []struct {
			Time       time.Time `json:"time"`
			Value      float64   `json:"value"`
			Correction int       `json:"correction"`
			Quality    int       `json:"quality"`
		} `json:"observations"`
	} `json:"data"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Status  int    `json:"status"`
	TraceId string `json:"traceId,omitempty"`
	Errors  struct {
		Parameter []string `json:"Parameter"`
	} `json:"errors,omitempty"`
}

package nveapi

/*
	http://api.nve.no/doc/hydrologiske-data/
	https://hydapi.nve.no/UserDocumentation/
*/

const (
	apiBaseUri = "https://hydapi.nve.no"
)

const (
	Observation = 1001
)

type apiEndpoints map[int]string

var apiEndpointsV1 = apiEndpoints{
	1001: "/api/v1/Observations",
}

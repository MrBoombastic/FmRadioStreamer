package dashboard

import "errors"

//Map with all endpoints
var endpointsList = map[string]EndpointData{
	"config":    configEndpoint,
	"music":     musicEndpoint,
	"loudstop":  loudstopEndpoint,
	"superstop": superstopEndpoint,
	"yt":        ytEndpoint,
	"play":      playEndpoint,
	"save":      saveEndpoint,
}

//Finds endpoint by name or alias
func findEndpoint(name string) (EndpointData, error) {
	if endpointsList[name].Endpoint != nil {
		return endpointsList[name], nil
	} else {
		return EndpointData{}, errors.New("404")
	}
}

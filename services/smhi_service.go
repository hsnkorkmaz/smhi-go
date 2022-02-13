package services

import (
	"app/structs"
	"app/utilities"
)

func GetCategoryResult() (result structs.CategoryResult, err error) {
	requestUri := ""
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}	

func GetVersionResult(version string) (result structs.VersionResult, err error) {
	requestUri := "/version/" + version
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}

func GetParameterResult(version string, parameter string) (result structs.ParameterResult, err error) {
	requestUri := "/version/" + version + "/parameter/" + parameter
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}

func GetStationResult(version string, parameter string, station string) (result structs.StationResult, err error) {
	requestUri := "/version/" + version + "/parameter/" + parameter + "/station/" + station
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}
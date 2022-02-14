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

func GetStationSetResult(version string, parameter string, stationSet string) (result structs.StationSetResult, err error) {
	requestUri := "/version/" + version + "/parameter/" + parameter + "/station-set/" + stationSet
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}

func GetPeriodStationResult(version string, parameter string, station string, period string) (result structs.PeriodResult, err error) {
	requestUri := "/version/" + version + "/parameter/" + parameter + "/station/" + station + "/period/" + period
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}

func GetPeriodStationSetResult(version string, parameter string, stationSet string, period string) (result structs.PeriodResult, err error) {
	requestUri := "/version/" + version + "/parameter/" + parameter + "/station-set/" + stationSet + "/period/" + period
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}

func GetDataStationResult(version string, parameter string, station string, period string) (result structs.DataStationResult, err error) {
	requestUri := "/version/" + version + "/parameter/" + parameter + "/station/" + station + "/period/" + period + "/data"
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}

func GetDataStationSetResult(version string, parameter string, stationSet string, period string) (result structs.DataStationSetResult, err error) {
	requestUri := "/version/" + version + "/parameter/" + parameter + "/station-set/" + stationSet + "/period/" + period + "/data"
	err = HttpRequest(utilities.GetEndpoint(requestUri), &result)
	return result, err
}
